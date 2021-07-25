package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheck(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/test_ok" {
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`{"result":{"value":1, "status":"ok"}}`))
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte("NOK"))
			}
		}),
	)

	defer testServer.Close()

	configOk := MinersConfigJson{
		TelegramToken:    "TEST",
		TelegramUsername: "TEST",
		Miners: []MinerConfigJson{
			{
				Name: "TEST_OK",
				Request: MinerConfigRequestJson{
					Method: "GET",
					URL:    fmt.Sprintf("%s/test_ok", testServer.URL),
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.value",
				},
			},
			{
				Name: "JSONRPC_OK",
				Request: MinerConfigRequestJson{
					Method: "JSONRPC20",
					URL:    fmt.Sprintf("%s/test_ok", testServer.URL),
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.status",
				},
			},
		},
	}

	configNok := MinersConfigJson{
		TelegramToken:    "TEST",
		TelegramUsername: "TEST",
		Miners: []MinerConfigJson{
			{
				Name: "TEST_NOK",
				Request: MinerConfigRequestJson{
					Method: "GET",
					URL:    fmt.Sprintf("%s/test_nok", testServer.URL),
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.valuea",
				},
			},
			{
				Name: "JSONRPC_NOK",
				Request: MinerConfigRequestJson{
					Method: "JSONRPC20",
					URL:    fmt.Sprintf("%s/test_nok", testServer.URL),
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.statusa",
				},
			},
			{
				Name: "NOK",
				Request: MinerConfigRequestJson{
					Method: "not implemented",
					URL:    fmt.Sprintf("%s/test_nok", testServer.URL),
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.status",
				},
			},
		},
	}

	t.Run("Test check OK", func(t *testing.T) {
		DoMinersCheck(configOk)
	})

	t.Run("Test check nok", func(t *testing.T) {
		DoMinersCheck(configNok)
	})
}
