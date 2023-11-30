package security

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (*string, error) {
	bytes := []byte(password)
	hashBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hash := string(hashBytes)

	return &hash, nil
}

func VerifyPassword(password, hashedPassword string) error {
	hashedPasswordBytes := []byte(hashedPassword)
	passwordBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)

	if err != nil {
		return err
	}

	return nil
}
