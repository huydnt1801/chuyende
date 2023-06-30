//go:build integration
// +build integration

package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/huydnt1801/chuyende/internal/ent"
	"github.com/huydnt1801/chuyende/internal/user"
	"github.com/huydnt1801/chuyende/test"
	"github.com/stretchr/testify/assert"
)

const (
	AccountEndpoint string = "/api/v1/account"
)

func mockUsers(client *ent.Client, len int) []*ent.User {
	var ret []*ent.User
	passwordHasher := &user.BcryptPasswordHasher{}
	for idx := 1; idx <= len; idx++ {
		user := &ent.User{
			PhoneNumber: "012345678" + fmt.Sprint(idx),
			Confirmed:   true,
			FullName:    "test user " + fmt.Sprint(idx),
			Password:    "123456",
		}
		ret = append(ret, user)
	}
	bulk := make([]*ent.UserCreate, 0)
	for _, user := range ret {
		hashPass, _ := passwordHasher.HashPassword(user.Password)
		q := client.User.Create().
			SetFullName(user.FullName).
			SetConfirmed(user.Confirmed).
			SetPhoneNumber(user.PhoneNumber).
			SetPassword(hashPass)
		bulk = append(bulk, q)
	}

	mockData, _ := client.User.CreateBulk(bulk...).
		Save(context.Background())
	return mockData
}

func mockDrivers(client *ent.Client, len int) []*ent.VehicleDriver {
	var ret []*ent.VehicleDriver
	for idx := 1; idx <= len; idx++ {
		user := &ent.VehicleDriver{
			PhoneNumber: "098765432" + fmt.Sprint(idx),
			FullName:    "test driver " + fmt.Sprint(idx),
			Password:    "testdriver",
		}
		ret = append(ret, user)
	}

	bulk := make([]*ent.VehicleDriverCreate, 0)
	for _, user := range ret {
		q := client.VehicleDriver.Create().
			SetFullName(user.FullName).
			SetPhoneNumber(user.PhoneNumber).
			SetPassword(user.Password)
		bulk = append(bulk, q)
	}

	mockData, _ := client.VehicleDriver.CreateBulk(bulk...).
		Save(context.Background())
	return mockData
}

type TestLoginInfo struct {
	body   map[string]interface{}
	output *LoginResponse
}

func TestLogin(t *testing.T) {
	// Create env
	env := test.NewTestEnv(t, container, "TestLogin")
	assert.NotNil(t, env)

	client := env.Client
	db := env.Database

	// Mock mysql data
	mockUsers := mockUsers(client, 1)

	tests := []struct {
		name string
		info *TestLoginInfo
	}{
		{
			name: "[TestLogin][Success] Return 200 - Return login success",
			info: &TestLoginInfo{
				body: map[string]interface{}{
					"phoneNumber": mockUsers[0].PhoneNumber,
					"password":    "123456",
				},
				output: &LoginResponse{
					Code: http.StatusOK,
					Data: &user.User{
						ID: 1,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doLogin(t, tc.info, db)
		})
	}
}

func doLogin(t *testing.T, info *TestLoginInfo, db *sql.DB) {
	req := test.BuildPostQuery(AccountEndpoint+"/login", info.body)
	rec := httptest.NewRecorder()
	c := NewTestEchoContext().NewContext(req, rec)

	accountSv := NewAccountServer(db)
	err := accountSv.Login(c)

	if info.output.Code == http.StatusOK {
		assert.Nil(t, err)
		var actual *LoginResponse
		err := json.NewDecoder(rec.Body).Decode(&actual)
		assert.Nil(t, err)
		assert.Equal(t, info.output.Code, actual.Code)
	} else {
		assert.Equal(t, info.output.Code, GetErrorCode(err))
	}
}
