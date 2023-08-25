package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

func DoPasswordsMatch(hashedPassword, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(requestPassword))
	return err == nil
}
