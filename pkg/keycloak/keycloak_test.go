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
	kv, err := NewKeycloakValidator(stagingHost)
	if err != nil {
		t.Fatal(err)
	}
	err = kv.checkHeader(stagingHeader, false)
	fmt.Println(err)
	err = kv.checkHeader(stagingHeader, true)
	fmt.Println(err)
	err = kv.checkHeader(bogus, false)
	fmt.Println(err)
}
