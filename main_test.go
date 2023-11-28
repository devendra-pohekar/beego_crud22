package main_test

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestMianFIle(t *testing.T) {
	t.Run("godotenv", func(t *testing.T) {

		result := godotenv.Load()

		if result != nil {
			t.Fatalf("Can't Load %s but got", result)
		}
		t.Log("env load success")
	})

}
