package integration

import (
	"net/http"
	"testing"
)

func TestIntegration(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/players", nil)

}
