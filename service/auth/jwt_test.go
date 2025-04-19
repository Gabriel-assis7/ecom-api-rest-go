package auth

import "testing"

func TestCreateJwt(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJwt(secret, 1)
	if err != nil {
		t.Errorf("failed to create JWT: %v", err)
	}

	if token == "" {
		t.Error("JWT token is empty")
	}
}
