package integrationtest

import (
	"log"
	"net/http"
	"os"
	"testing"

	gohit "github.com/Eun/go-hit"
)

const (
	host     = "127.0.0.1:3000"
	basePath = "http://" + host + "/api/v1"
	attempts = 20
	requests = 10
)

func TestMain(m *testing.M) {
	log.Printf("Integration tests: host %s is available", host)
	code := m.Run()
	os.Exit(code)
}

func TestGetCampaigns(t *testing.T) {
	gohit.Test(t, gohit.Description("Get campaigns"),
		gohit.Get(basePath+"/campaigns"),
		gohit.Send().Headers("Content-Type").Add("application/json"),
		// gohit.Send().Body()
		gohit.Expect().Status().Equal(http.StatusOK),
	)
}

func TestGetCampaign(t *testing.T) {

}
