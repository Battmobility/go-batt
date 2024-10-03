package keycloak

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

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
	pk *rsa.PublicKey
}

func NewKeycloakValidator(url string) (res *KeycloakValidator, err error) {
	pk, err := getRSAPublicKeyFromKeycloak(url)
	if err != nil {
		return nil, err
	}
	return &KeycloakValidator{
		pk: pk,
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

	// Assuming there's only one key. You might need to select the correct one based on `kid`
	if len(certs.Keys) < 1 {
		return nil, fmt.Errorf("no keys found on keycloak url")
	}
	key := certs.Keys[0]

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

func (kv *KeycloakValidator) checkHeader(header string) (err error) {
	_, err = jwt.Parse(header, kv.validateToken)
	return err
}

func (kv *KeycloakValidator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		header = strings.TrimPrefix(header, "Bearer ")
		err := kv.checkHeader(header)
		if err != nil {
			fmt.Println("Error parsing jwt token", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
