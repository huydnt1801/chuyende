package server

import (
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/huydnt1801/chuyende/internal/config"
	"github.com/huydnt1801/chuyende/internal/user"
	"github.com/huydnt1801/chuyende/pkg/log"
	"github.com/labstack/echo/v4"
)

type AccountServer struct {
	logger    logr.Logger
	secretKey string

	userSvc user.Service
}

func NewAccountServer(db *sql.DB) *AccountServer {
	cfg := config.MustParseConfig()
	svc := user.NewService(db)
	srv := &AccountServer{
		logger:    log.ZapLogger(),
		secretKey: cfg.SecretKey,
		userSvc:   svc,
	}
	return srv
}

func (s AccountServer) register(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &RegisterRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}
	if data.Password != data.Password2 {
		return echo.NewHTTPError(http.StatusBadRequest, "Mật khẩu không khớp")
	}
	_, nextOTPSend, err := s.userSvc.CreateUser(ctx, &user.User{
		PhoneNumber: data.PhoneNumber,
		FullName:    data.FullName,
	}, data.Password)
	if err != nil {
		return err
	}
	token := s.signConfirmInfo("reg", data.PhoneNumber, nextOTPSend)
	return c.JSON(http.StatusCreated, SuccessResponse{Code: http.StatusCreated, Data: token})
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Password2   string `json:"password2" validate:"required"`
}

func (s *AccountServer) registerConfirm(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := &RegisterConfirmRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}
	phoneNumber, nextOTPSend, err := s.parseConfirmInfo("reg", data.Token, 1*time.Hour)
	if err != nil {
		return err
	}
	if phoneNumber == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid information")
	}
	usr, err := s.userSvc.FindUser(ctx, &user.UserParams{PhoneNumber: phoneNumber})
	if err != nil {
		return err
	}
	switch data.Type {
	case "submit-otp":
		err = s.userSvc.Confirm(ctx, usr, data.OTP)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusNoContent, "")
	case "resend-otp":
		if time.Now().Before(nextOTPSend) {
			return user.OTPIntervalError{}
		}
		nextOTPSend, err := s.userSvc.SendConfirmationToken(ctx, usr)
		if err != nil {
			return err
		}

		token := s.signConfirmInfo("reg", phoneNumber, nextOTPSend)
		return c.JSON(http.StatusCreated, SuccessResponse{Code: http.StatusOK, Data: token})
	default:
		return c.Redirect(http.StatusSeeOther, r.RequestURI)
	}
}

type RegisterConfirmRequest struct {
	Type  string `json:"type" validate:"required"`
	Token string `json:"token" validate:"required"`
	OTP   string `json:"otp"`
}

func (s *AccountServer) signData(data []byte) []byte {
	mac := hmac.New(sha1.New, []byte(s.secretKey))
	mac.Write(data)
	return mac.Sum(nil)
}

func (s *AccountServer) signConfirmInfo(typ, phoneNumber string, nextOTPSend time.Time) string {
	regData := fmt.Sprintf("%s;%s;%d;%d", typ, phoneNumber, nextOTPSend.Unix(), time.Now().Unix())
	signature := s.signData([]byte(regData))
	encodeData := base64.StdEncoding.EncodeToString([]byte(regData))
	encodeSig := base64.StdEncoding.EncodeToString(signature)
	token := fmt.Sprintf("%s.%s", encodeData, encodeSig)
	return token
}

func (s *AccountServer) parseConfirmInfo(typ, hashed string, exp time.Duration) (string, time.Time, error) {
	ss := strings.Split(hashed, ".")
	if len(ss) != 2 {
		return "", time.Time{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
	encodedData, sig := ss[0], ss[1]
	raw, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", time.Time{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
	mac := hmac.New(sha1.New, []byte(s.secretKey))
	mac.Write(raw)
	expMac := mac.Sum(nil)
	tokenMac, _ := base64.StdEncoding.DecodeString(sig)
	if !hmac.Equal(tokenMac, expMac) {
		return "", time.Time{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
	ss = strings.Split(string(raw), ";")
	if len(ss) != 4 || ss[0] != typ {
		return "", time.Time{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
	}
	nextOTPSendUnix, _ := strconv.ParseInt(ss[2], 10, 0)
	nextOTPSend := time.Unix(nextOTPSendUnix, 0)
	createdAtUnix, _ := strconv.ParseInt(ss[3], 10, 0)
	createdAt := time.Unix(createdAtUnix, 0)
	if createdAt.Add(exp).Before(time.Now()) {
		return "", time.Time{}, echo.NewHTTPError(http.StatusBadRequest, "Token is expired")
	}
	return ss[1], nextOTPSend, nil
}
