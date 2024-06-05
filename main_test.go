package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func generateRandomPing() Ping {
	names := []string{"world", "golang", "developer", "tester", "user"}
	name := names[rand.Intn(len(names))]
	count := rand.Intn(10) + 1
	return Ping{Name: name, Count: count}
}

func TestPing(t *testing.T) {
	_, api := humatest.New(t)
	addRoute(api)

	for i := 0; i < 100; i++ {
		i := i
		t.Run(fmt.Sprintf("tc-%d", i), func(t *testing.T) {
			t.Parallel()
			ping := generateRandomPing()
			resp := api.Post("/ping", ping)
			t.Logf("Response: %s", resp.Body)
			if resp.Code != 200 {
				t.Fatalf("Expected status 200, got %d", resp.Code)
			}
		})
	}
}
