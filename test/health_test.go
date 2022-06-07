package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"gotest.tools/v3/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode())
}
