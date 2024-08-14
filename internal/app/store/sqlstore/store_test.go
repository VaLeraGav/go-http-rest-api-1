package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "sslmode=disable user=postgres password=yourpassword dbname=restapi-dev host=localhost port=5432"
	}
	os.Exit(m.Run())
}
