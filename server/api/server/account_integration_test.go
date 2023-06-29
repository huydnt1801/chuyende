//go:build integration
// +build integration

package server

import (
	"context"
	"fmt"

	"github.com/huydnt1801/chuyende/internal/ent"
)

const (
	AccountEndpoint string = "/api/v1/account"
)

func mockUsers(client *ent.Client, len int) []*ent.User {
	var ret []*ent.User
	for idx := 1; idx <= len; idx++ {
		user := &ent.User{
			PhoneNumber: "012345678" + fmt.Sprint(idx),
			Confirmed:   true,
			FullName:    "test user " + fmt.Sprint(idx),
			Password:    "testuser",
		}
		ret = append(ret, user)
	}

	bulk := make([]*ent.UserCreate, 0)
	for _, user := range ret {
		q := client.User.Create().
			SetFullName(user.FullName).
			SetConfirmed(user.Confirmed).
			SetPhoneNumber(user.PhoneNumber).
			SetPassword(user.Password)
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
