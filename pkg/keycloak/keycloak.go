package keycloak

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	BattAdminRole = "BattAdmin"
	AdminClientId = "admin-cli"
)

var (
	ErrNoRS256Key       = errors.New("no RS256 key found on keycloak url")
	ErrFailedToGetToken = errors.New("failed to get token")
)

type Claims struct {
	Sub   string   `json:"sub"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

// KeycloakCertsResponse is the structure of the response from the Keycloak certs endpoint
type KeycloakCertsResponse struct {
	Keys []struct {
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		Alg string `json:"alg"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

type KeycloakValidator struct {
	pk  *rsa.PublicKey
	cfg Config
}

type Config struct {
	PassUnauthenticated bool
}

func NewKeycloakValidator(url string, cfg Config) (res *KeycloakValidator, err error) {
	pk, err := getRSAPublicKeyFromKeycloak(url)
	if err != nil {
		return nil, err
	}
	return &KeycloakValidator{
		pk:  pk,
		cfg: cfg,
	}, nil
}

// getRSAPublicKeyFromKeycloak fetches the RSA public key from Keycloak
func getRSAPublicKeyFromKeycloak(url string) (*rsa.PublicKey, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get public key from keycloak")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var certs KeycloakCertsResponse
	if err := json.Unmarshal(body, &certs); err != nil {
		return nil, err
	}

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
		return nil, err
	}
	eBytes, err := jwt.DecodeSegment(key.E)
	if err != nil {
		return nil, err
	}
	e := new(big.Int).SetBytes(eBytes).Int64()

	publicKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: int(e),
	}

	return publicKey, nil
}

func (kv *KeycloakValidator) validateToken(token *jwt.Token) (res interface{}, err error) {
	return kv.pk, nil
}

func (kv *KeycloakValidator) ParseToken(header string) (result *Claims, err error) {
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(header, claims, kv.validateToken)
	//extract the roles from realm_access
	if err != nil {
		return nil, err
	}
	realmAccess, ok := claims["realm_access"].(map[string]interface{})
	if !ok {
		return nil, errors.New("realm_access claim not found")
	}
	roles, ok := realmAccess["roles"].([]interface{})
	if !ok {
		return nil, errors.New("roles claim not found")
	}
	rolesParsed := make([]string, len(roles))
	for i, v := range roles {
		rolesParsed[i], ok = v.(string)
		if !ok {
			return nil, errors.New("role not a string")
		}
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("sub claim not found")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("email claim not found")
	}
	return &Claims{
		Sub:   sub,
		Email: email,
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

func (kv *KeycloakValidator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		header = strings.TrimPrefix(header, "Bearer ")
		if header == "" && kv.cfg.PassUnauthenticated {
			next.ServeHTTP(w, r)
			return
		}
		claims, err := kv.ParseToken(header)
		if err != nil {
			fmt.Println("Error parsing jwt token", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), SubKey, claims.Sub)
		ctx = context.WithValue(ctx, EmailKey, claims.Email)
		ctx = context.WithValue(ctx, RolesKey, claims.Roles)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (kv *KeycloakValidator) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		header = strings.TrimPrefix(header, "Bearer ")
		parsed, err := kv.ParseToken(header)
		if err != nil || !contains(parsed.Roles, BattAdminRole) {
			fmt.Println("Error parsing jwt token", err)
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

func NewTokenProvider(url, realm, username, password, client_id string) (res *TokenProvider, err error) {
	return &TokenProvider{
		url:  fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", url, realm),
		data: fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=%s", username, password, client_id),
	}, nil
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

	var tokenResponse KeycloakTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("json decoding failed: %w", err)
	}
	token, _, err := new(jwt.Parser).ParseUnverified(tokenResponse.AccessToken, &jwt.StandardClaims{})
	if err != nil {
		return "", fmt.Errorf("jwt parsing failed: %w", err)
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); !ok {
		return "", fmt.Errorf("jwt claims parsing failed: %w", err)
	} else {
		// ask for a new token a minute before the actual expiration
		tp.tokenExpiration = time.Unix(claims.ExpiresAt, 0).Add(-time.Minute)
		tp.token = tokenResponse.AccessToken
	}
	return tp.token, nil
}

type KeycloakTokenResponse struct {
	AccessToken string `json:"access_token"` //nolint:tagliatelle
}
