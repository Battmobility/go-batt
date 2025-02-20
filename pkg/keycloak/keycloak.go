package keycloak

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	BattAdminRole = "BattAdmin"
	AdminClientID = "admin-cli"
)

var (
	ErrNoRS256Key             = errors.New("no RS256 key found on keycloak url")
	ErrFailedToGetToken       = errors.New("failed to get token")
	ErrFailedToGetPublicKey   = errors.New("failed to get public key from keycloak")
	ErrFailedToParsePublicKey = errors.New("failed to parse public key from keycloak")
	ErrInvalidClaims          = errors.New("invalid claims")
)

type Claims struct {
	Sub   string   `json:"sub"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

// CertsResponse is the structure of the response from the Keycloak certs endpoint
type CertsResponse struct {
	Keys []struct {
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		Alg string `json:"alg"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

type Validator struct {
	pk  *rsa.PublicKey
	cfg Config
}

type Config struct {
	PassUnauthenticated bool
}

func NewKeycloakValidator(ctx context.Context, url string, cfg Config) (*Validator, error) {
	pk, err := getRSAPublicKeyFromKeycloak(ctx, url)
	if err != nil {
		return nil, err
	}
	return &Validator{
		pk:  pk,
		cfg: cfg,
	}, nil
}

// getRSAPublicKeyFromKeycloak fetches the RSA public key from Keycloak
func getRSAPublicKeyFromKeycloak(ctx context.Context, kcURL string) (*rsa.PublicKey, error) {
	parsedURL, err := url.Parse(kcURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToGetPublicKey, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrFailedToGetPublicKey
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParsePublicKey, err)
	}

	var certs CertsResponse
	if err := json.Unmarshal(body, &certs); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParsePublicKey, err)
	}
	return parseRSAPublicKey(certs)
}

func parseRSAPublicKey(certs CertsResponse) (*rsa.PublicKey, error) {
	keyIndex := -1
	for index, key := range certs.Keys {
		if key.Alg == "RS256" {
			keyIndex = index
			break
		}
	}
	if keyIndex == -1 {
		return nil, ErrNoRS256Key
	}
	key := certs.Keys[keyIndex]

	nBytes, err := jwt.DecodeSegment(key.N)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParsePublicKey, err)
	}
	eBytes, err := jwt.DecodeSegment(key.E)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParsePublicKey, err)
	}
	e := new(big.Int).SetBytes(eBytes).Int64()

	publicKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: int(e),
	}

	return publicKey, nil
}

func (kv *Validator) validateToken(_ *jwt.Token) (interface{}, error) {
	return kv.pk, nil
}

func (kv *Validator) ParseToken(header string) (*Claims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(header, claims, kv.validateToken)
	// extract the roles from realm_access
	if err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}
	realmClaims, ok := claims["realm_access"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("%w: realm_access not found in claims", ErrInvalidClaims)
	}
	roles, ok := realmClaims["roles"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("%w: roles not found in claims", ErrInvalidClaims)
	}
	rolesParsed := make([]string, len(roles))
	for i, v := range roles {
		parsedRole, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("%w: could not parse role", ErrInvalidClaims)
		}
		rolesParsed[i] = parsedRole
	}
	subClaim, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("%w: could not parse sub claim", ErrInvalidClaims)
	}
	emailClaim, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("%w: could not parse email claim", ErrInvalidClaims)
	}
	emailVerified, ok := claims["email_verified"].(bool)
	if !ok || !emailVerified {
		// if keycloak is configured correctly, this should never happen
		return nil, fmt.Errorf("%w: email not verified", ErrInvalidClaims)
	}
	return &Claims{
		Sub:   subClaim,
		Email: emailClaim,
		Roles: rolesParsed,
	}, nil
}

type contextKey string

const SubKey contextKey = "sub"
const EmailKey contextKey = "email"
const RolesKey contextKey = "roles"

// IsAdminRequest checks if the request is from an admin user
func IsAdminRequest(r *http.Request) bool {
	roles, ok := r.Context().Value(RolesKey).([]string)
	return ok && contains(roles, BattAdminRole)
}

// GetSubAndEmail returns the sub and email from the request context or empty strings if not present
func GetSubAndEmail(r *http.Request) (sub, email string) { //nolint: nonamedreturns
	sub, _ = r.Context().Value(SubKey).(string)
	email, _ = r.Context().Value(EmailKey).(string)
	return
}

func (kv *Validator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		header = strings.TrimPrefix(header, "Bearer ")
		if header == "" && kv.cfg.PassUnauthenticated {
			next.ServeHTTP(w, r)
			return
		}
		claims, err := kv.ParseToken(header)
		if err != nil {
			log.Println("Error parsing jwt token", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), SubKey, claims.Sub)
		ctx = context.WithValue(ctx, EmailKey, claims.Email)
		ctx = context.WithValue(ctx, RolesKey, claims.Roles)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (kv *Validator) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		header = strings.TrimPrefix(header, "Bearer ")
		parsed, err := kv.ParseToken(header)
		if err != nil || !contains(parsed.Roles, BattAdminRole) {
			log.Println("Error parsing jwt token", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func contains(s []string, search string) bool {
	for _, v := range s {
		if v == search {
			return true
		}
	}
	return false
}

type TokenProvider struct {
	url             string
	data            string
	token           string
	tokenExpiration time.Time
}

func NewTokenProvider(providerURL, realm, username, password, clientID string) (*TokenProvider, error) {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	data.Add("grant_type", "password")
	data.Add("client_id", clientID)
	tp := &TokenProvider{
		url:  fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", providerURL, realm),
		data: data.Encode(),
	}
	return tp, nil
}

// GetKeycloakToken fetches the token from Keycloak if the cached token is expired. Returns the cached token.
// Not thread-safe.
func (tp *TokenProvider) GetKeycloakToken(ctx context.Context) (string, error) {
	if time.Now().Before(tp.tokenExpiration) {
		return tp.token, nil
	}
	tp.tokenExpiration = time.Time{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tp.url, bytes.NewBufferString(tp.data))
	if err != nil {
		return "", fmt.Errorf("creating POST request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("POST call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("%w (%d): %s", ErrFailedToGetToken, resp.StatusCode, string(bodyBytes))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("json decoding failed: %w", err)
	}
	token, _, err := new(jwt.Parser).ParseUnverified(tokenResponse.AccessToken, &jwt.StandardClaims{})
	if err != nil {
		return "", fmt.Errorf("jwt parsing failed: %w", err)
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("jwt claims parsing failed: %w", err)
	}
	// ask for a new token a minute before the actual expiration
	tp.tokenExpiration = time.Unix(claims.ExpiresAt, 0).Add(-time.Minute)
	tp.token = tokenResponse.AccessToken
	return tp.token, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"` //nolint:tagliatelle
}
