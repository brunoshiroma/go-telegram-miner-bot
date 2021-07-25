package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/test_ok":
				rw.WriteHeader(http.StatusOK)
				_, _ = rw.Write([]byte(`{"result":{"value":1, "status":"ok"}}`))
			default:
				rw.WriteHeader(http.StatusInternalServerError)
				_, _ = rw.Write([]byte("NOK"))
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
					URL:    fmt.Sprintf("%s/test_jsonrpc_ok", testServer.URL),
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
					URL:    "0.0.0.0",
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.valuea",
				},
			},
			{
				Name: "JSONRPC_NOK",
				Request: MinerConfigRequestJson{
					Method: "JSONRPC20",
					URL:    "0.0.0.0;0",
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.statusa",
				},
			},
			{
				Name: "NOK",
				Request: MinerConfigRequestJson{
					Method: "not implemented",
					URL:    "0.0.0.0",
				},
				Response: MinerConfigResponseJson{
					JSONPath: "result.status",
				},
			},
		},
	}

	t.Run("Test check OK", func(t *testing.T) {
		result, err := DoMinersCheck(configOk)

		assert.NoError(t, err)

		for _, minerResult := range result.Miners {
			if minerResult.Name == "TEST_OK" {
				assert.Equal(t, "1", minerResult.Result)
			} else {
				assert.Equal(t, "", minerResult.Result)
			}
		}
	})

	t.Run("Test check nok", func(t *testing.T) {
		result, err := DoMinersCheck(configNok)

		assert.NoError(t, err)
		assert.False(t, result.Miners[0].Success)
		assert.False(t, result.Miners[1].Success)
		assert.False(t, result.Miners[2].Success)
	})
}
