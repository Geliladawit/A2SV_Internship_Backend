package Infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	CheckPasswordHash(password, hash string) bool
}

type PasswordServiceImpl struct{}

func NewPasswordService() PasswordService {
	return &PasswordServiceImpl{}
}
func (p *PasswordServiceImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}