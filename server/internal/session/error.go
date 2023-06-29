package session

import (
	"github.com/huydnt1801/chuyende/internal/utils/http"
)

type InvalidSessionError struct {
	http.UnauthorizedError
}

func (InvalidSessionError) Error() string {
	return "Đã hết phiên đăng nhập"
}

type NotFoundSessionError struct {
	http.NotFoundError
}

func (NotFoundSessionError) Error() string {
	return "Phiên đăng nhập không hợp lệ"
}
