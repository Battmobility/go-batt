package keycloak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	stagingHost   = "https://keycloak-staging.battmobility.be/auth/realms/Battmobiel/protocol/openid-connect/certs"
	stagingHeader = ""
)

func TestKeycloakValidator(t *testing.T) {
	t.Parallel()
	if stagingHeader == "" {
		t.Skip("skipping test; stagingHeader is empty")
	}
	kv, err := NewKeycloakValidator(stagingHost, Config{})
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := kv.ParseToken(stagingHeader)
	require.NoError(t, err)
	fmt.Println(parsed.Roles) //nolint:forbidigo
}
