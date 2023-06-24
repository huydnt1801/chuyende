package user

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type BcryptPasswordHasher struct{}

func (ph *BcryptPasswordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (ph *BcryptPasswordHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type PasswordValidator interface {
	ValidatePassword(password string) error
}

type PasswordComplexity struct {
	MinLength    int
	MaxLength    int
	NumUppercase int
	NumSpecial   int
}

var DefaultPasswordComplexity = PasswordComplexity{
	MinLength:    8,
	MaxLength:    255,
	NumUppercase: 0,
	NumSpecial:   0,
}

func (validator *PasswordComplexity) ValidatePassword(password string) error {
	var errDes string

	if len(password) < validator.MinLength {
		errDes = errDes + fmt.Sprintf("Độ dài tối thiểu là %v.", validator.MinLength)
	}
	if len(password) > validator.MaxLength {
		errDes = errDes + fmt.Sprintf("Độ dài tối đa là %v.", validator.MaxLength)
	}
	var numUppercase, numSpecial int
	for _, char := range password {
		if unicode.IsUpper(char) {
			numUppercase++
		}
		if unicode.IsSymbol(char) {
			numSpecial++
		}
	}
	if numUppercase < validator.NumUppercase {
		errDes = errDes + fmt.Sprintf("Mật khẩu phải chứa ít nhất %v từ viết hoa.", validator.NumUppercase)
	}
	if numSpecial < validator.NumSpecial {
		errDes = errDes + fmt.Sprintf("Mật khẩu phải chứa ít nhất %v từ viết thường.", validator.NumSpecial)
	}
	if errDes != "" {
		return &PasswordComplexityError{ErrDescription: errDes}
	}
	return nil
}
