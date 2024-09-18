package store_test

import (
	"os"
	"testing"
)
var (
	databaseURL string
)
func TestMain(m *testing.M) {
	databaseURL=os.Getenv("DATABASE_URL")
	if databaseURL==""{
		//databaseURL="host=localhost dbname=restapi_test sslmode=disable"
		databaseURL = "host=localhost port=5432 user=postgres password=postgres dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}