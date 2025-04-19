package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("failed to hash password: %v", err)
	}

	if hash == "" {
		t.Error("hashed password is empty")
	}

	if hash == "password" {
		t.Error("hashed password should not be the same as the original password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("failed to hash password: %v", err)
	}

	err = CheckPasswordHash("password", hash)
	if err != nil {
		t.Errorf("failed to compare password: %v", err)
	}

	err = CheckPasswordHash("wrongpassword", hash)
	if err == nil {
		t.Error("expected error for wrong password, got nil")
	}
}
