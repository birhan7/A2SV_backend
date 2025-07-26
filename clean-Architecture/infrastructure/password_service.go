package infrastructure

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifiedPassword(hpassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(password))
	return err == nil
}
