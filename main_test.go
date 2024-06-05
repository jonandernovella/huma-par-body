package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
			pingByte, _ := json.Marshal(ping)
			resp := api.Post("/ping", ping)
			t.Logf("Response: %s", resp.Body)
			if resp.Code != 200 {
				t.Fatalf("Expected status 200, got %d", resp.Code)
			}
			var data PingResponse
			body, _ := io.ReadAll(resp.Body)
			if err := json.Unmarshal(body, &data); err != nil {
				t.Fatalf("Error unmarshalling response: %s", err)
			}
			if len(pingByte) != len(data.PingRaw) {
				log.Printf("input: %s\n", string(pingByte))
				log.Printf("output: %s\n", string(data.PingRaw))
				log.Fatal("input byte is not same length as output byte")
			}
		})
	}
}
