package driver

import (
	"errors"

	"github.com/huydnt1801/chuyende/internal/utils/http"
)

type InvalidPhoneError struct {
	http.NotFoundError
}

func (InvalidPhoneError) Error() string {
	return "Số điện thoại không hợp lệ"
}

type InvalidOTPError struct {
	http.BadRequestError
}

type DriverNotFoundError struct {
	http.NotFoundError
}

func (DriverNotFoundError) Error() string {
	return "Không tìm thấy người dùng"
}

func IsDriverNotFound(err error) bool {
	if err == nil {
		return false
	}
	notfoundErr := &DriverNotFoundError{}
	return errors.As(err, notfoundErr) || errors.As(err, &notfoundErr)
}

type InvalidPasswordError struct {
	http.UnauthorizedError
}

func (InvalidPasswordError) Error() string {
	return "Mật khẩu không đúng"
}
