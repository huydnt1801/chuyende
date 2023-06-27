//go:build integration
// +build integration

package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/huydnt1801/chuyende/internal/ent"
	"github.com/huydnt1801/chuyende/internal/trip"
	"github.com/huydnt1801/chuyende/test"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

const (
	TripEndpoint string = "/api/v1/trips"
)

var container testcontainers.Container

func TestMain(m *testing.M) {
	// Setup
	var cleanUp func()
	container, cleanUp = test.MysqlContainer(context.Background())

	// Run test case
	exitCode := m.Run()

	// After test
	cleanUp()

	os.Exit(exitCode)
}

type TestListTripsInfo struct {
	queryParams map[string]string
	output      *ListTripResponse
}

func TestListTrips(t *testing.T) {
	// Create env
	env := test.NewTestEnv(t, container, "TestListTrips")
	assert.NotNil(t, env)

	client := env.Client
	db := env.Database

	// Mock mysql data
	mockUsers := mockUsers(client, 1)
	mockDrivers := mockDrivers(client, 1)
	mockTrips := mockTrips(client, mockUsers[0].ID, mockDrivers[0].ID, 3)
	respData, _ := trip.DecodeManyTrip(mockTrips)

	tests := []struct {
		name string
		info *TestListTripsInfo
	}{
		{
			name: "[ListTrips][Success] Return 200 - Return trips",
			info: &TestListTripsInfo{
				queryParams: map[string]string{},
				output: &ListTripResponse{
					Code: http.StatusOK,
					Data: respData,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doListTrips(t, tc.info, db)
		})
	}
}

func doListTrips(t *testing.T, info *TestListTripsInfo, db *sql.DB) {
	req := test.BuildGetQuery(TripEndpoint, info.queryParams)
	rec := httptest.NewRecorder()
	c := NewTestEchoContext().NewContext(req, rec)

	tripSrv := NewTripServer(db)
	err := tripSrv.ListTrip(c)

	if info.output.Code == http.StatusOK {
		assert.Nil(t, err)
		var actual *ListTripResponse
		err := json.NewDecoder(rec.Body).Decode(&actual)
		assert.Nil(t, err)
		assert.Equal(t, info.output.Code, actual.Code)
		assert.Equal(t, info.output.Data, actual.Data)
	} else {
		assert.Equal(t, info.output.Code, GetErrorCode(err))
	}
}

func mockTrips(client *ent.Client, userID, driverID, len int) []*ent.Trip {
	var ret []*ent.Trip
	for idx := 1; idx <= len; idx++ {
		trip := &ent.Trip{
			UserID:        userID,
			DriverID:      driverID,
			StartX:        float64(idx),
			StartY:        float64(idx),
			StartLocation: "di",
			EndX:          float64(idx),
			EndY:          float64(idx),
			EndLocation:   "den",
			Price:         20,
			Distance:      float64(idx),
		}
		ret = append(ret, trip)
	}

	bulk := make([]*ent.TripCreate, 0)
	for _, trip := range ret {
		q := client.Trip.Create().
			SetUserID(trip.UserID).
			SetDriverID(trip.DriverID).
			SetStartX(trip.StartX).
			SetStartY(trip.StartY).
			SetStartLocation(trip.StartLocation).
			SetEndX(trip.EndX).
			SetEndY(trip.EndY).
			SetEndLocation(trip.EndLocation).
			SetPrice(trip.Price).
			SetDistance(trip.Distance)
		bulk = append(bulk, q)
	}

	mockData, _ := client.Trip.CreateBulk(bulk...).
		Save(context.Background())
	return mockData
}
