package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

func DoMinersCheck(config MinersConfigJson) (MinersResult, error) {
	var (
		wg          *sync.WaitGroup
		resultsChan chan MinerResult
	)

	resultsChan = make(chan MinerResult, len(config.Miners))
	wg = &sync.WaitGroup{}

	result := MinersResult{
		Miners: make([]MinerResult, 0),
	}
	for _, minerConfig := range config.Miners {

		switch minerConfig.Request.Method {
		case "GET":
			wg.Add(1)
			go doCheckHttpGet(minerConfig, wg, resultsChan)
		case "JSONRPC20":
			wg.Add(1)
			go DoJSONPRC20(minerConfig, wg, resultsChan)
		default:
			result.Miners = append(result.Miners, struct {
				Name    string
				Success bool
				Result  string
			}{
				Name:    minerConfig.Name,
				Success: false,
				Result:  fmt.Sprintf("%s uses method %s, not implemented", minerConfig.Name, minerConfig.Request.Method),
			})
		}

	}
	wg.Wait()
	close(resultsChan)

	for minerResult := range resultsChan {
		result.Miners = append(result.Miners, minerResult)
	}

	return result, nil
}

func doCheckHttpGet(config MinerConfigJson, wg *sync.WaitGroup, results chan MinerResult) {
	defer wg.Done()
	http.DefaultClient.Timeout = time.Second * 10
	resp, err := http.Get(config.Request.URL)
	if err != nil {
		results <- MinerResult{
			Name:    config.Name,
			Success: false,
			Result:  err.Error(),
		}
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- MinerResult{
			Name:    config.Name,
			Success: false,
			Result:  err.Error(),
		}
		return
	}
	results <- CreateResult(string(body), config)
}

func CreateResult(body string, config MinerConfigJson) MinerResult {
	var resultBody string

	// exec the jsonPath and get the value
	if len(strings.TrimSpace(config.Response.JSONPath)) > 0 {
		jqResult := gjson.Get(body, config.Response.JSONPath)
		resultBody = jqResult.String()
	}

	return MinerResult{
		Name:    config.Name,
		Success: true,
		Result:  resultBody,
	}
}
