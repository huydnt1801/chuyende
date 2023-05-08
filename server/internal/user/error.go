package user

import "github.com/huydnt1801/chuyende/internal/utils/http"

type InvalidPhoneError struct {
	http.NotFoundError
}

func (InvalidPhoneError) Error() string {
	return "Phone number is not valid"
}

type InvalidOTPError struct {
	http.BadRequestError
}

func (InvalidOTPError) Error() string {
	return "Invalid OTP"
}

type OTPIntervalError struct {
	http.BadRequestError
}

func (OTPIntervalError) Error() string {
	return "Can not send new OTP"
}

type UserNotFoundError struct {
	http.NotFoundError
}

func (UserNotFoundError) Error() string {
	return "User Not Found"
}
