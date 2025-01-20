package keycloak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	stagingHost   = "https://keycloak-staging.battmobility.be/auth/realms/Battmobiel/protocol/openid-connect/certs"
	stagingHeader = ""
)

func TestKeycloakValidator(t *testing.T) {
	kv, err := NewKeycloakValidator(stagingHost, Config{})
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := kv.ParseToken(stagingHeader)
	assert.NoError(t, err)
	fmt.Println(parsed.Roles)
}
