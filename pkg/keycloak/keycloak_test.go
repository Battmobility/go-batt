package keycloak

import (
	"fmt"
	"testing"
)

const (
	stagingHost   = "https://batt-staging.slamit.be/auth/realms/Battmobiel/protocol/openid-connect/certs"
	stagingHeader = ""
	bogus         = ""
)

func TestKeycloakValidator(t *testing.T) {
	kv, err := NewKeycloakValidator(stagingHost, Config{})
	if err != nil {
		t.Fatal(err)
	}
	sub, err := kv.ParseToken(stagingHeader)
	fmt.Println(sub)
	fmt.Println(err)
	_, err = kv.ParseToken(stagingHeader)
	fmt.Println(err)
	_, err = kv.ParseToken(bogus)
	fmt.Println(err)
}
