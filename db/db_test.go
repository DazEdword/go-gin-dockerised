package db

import (
	"fmt"
	"os"
	"testing"
)

func setup(t *testing.T) {
	os.Setenv("POSTGRES_PORT", "1234")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_USER", "dave")
	os.Setenv("POSTGRES_PASSWORD", "notverysecure")
	os.Setenv("POSTGRES_DB", "gogin")

	t.Cleanup(func() {
		os.Unsetenv("POSTGRES_PORT")
		os.Unsetenv("POSTGRES_HOST")
		os.Unsetenv("POSTGRES_USER")
		os.Unsetenv("POSTGRES_PASSWORD")
		os.Unsetenv("POSTGRES_DB")
	})
}

func TestBuildConnectionString(t *testing.T) {
	// Arrange
	expected := "host=localhost port=1234 user=dave password=notverysecure dbname=gogin sslmode=disable"

	setup(t)

	// Act
	actual := BuildConnectionString()
	fmt.Println(actual)

	// Assert
	if actual != expected {
		t.Error("hmmm")
	}
}
