package trip

import (
	"errors"

	"github.com/huydnt1801/chuyende/internal/utils/http"
)

type TripNotFoundError struct {
	http.NotFoundError
}

func (TripNotFoundError) Error() string {
	return "Không tìm thấy chuyến đi"
}

func IsTripNotFound(err error) bool {
	if err == nil {
		return false
	}
	notfoundErr := &TripNotFoundError{}
	return errors.As(err, notfoundErr) || errors.As(err, &notfoundErr)
}
