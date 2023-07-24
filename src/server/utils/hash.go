package utils

import (
	"golang.org/x/crypto/bcrypt"

)

func HashPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	if err != nil {
		return "", nil;
	}
	resHash := string(hash);
	return resHash, err;
}

func CmpHash(HashedPassword string, password string) (bool) {
	var err = bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password));
	return err == nil;
}
