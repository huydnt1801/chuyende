package user

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

func (InvalidOTPError) Error() string {
	return "Mã OTP không hợp lệ"
}

type OTPIntervalError struct {
	http.BadRequestError
}

func (OTPIntervalError) Error() string {
	return "Chưa thể gửi lại OTP"
}

type UserNotFoundError struct {
	http.NotFoundError
}

func (UserNotFoundError) Error() string {
	return "Không tìm thấy người dùng"
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	notfoundErr := &UserNotFoundError{}
	return errors.As(err, notfoundErr) || errors.As(err, &notfoundErr)
}

type InvalidPasswordError struct {
	http.UnauthorizedError
}

func (InvalidPasswordError) Error() string {
	return "Mật khẩu không đúng"
}

type PasswordComplexityError struct {
	ErrDescription string
	http.BadRequestError
}

func (PasswordComplexityError) Error() string {
	return "Mật khẩu không hợp lệ"
}
