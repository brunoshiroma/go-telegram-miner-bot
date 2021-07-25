package internal

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"
)

func DoJSONPRC20(minerConfig MinerConfigJson, wg *sync.WaitGroup, results chan MinerResult) {
	defer wg.Done()

	url, err := url.Parse(minerConfig.Request.URL)
	if err != nil {
		results <- MinerResult{
			Success: false,
			Name:    minerConfig.Name,
			Result:  err.Error(),
		}
		return
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", url.Hostname(), url.Port()))
	if err != nil {
		results <- MinerResult{
			Success: false,
			Name:    minerConfig.Name,
			Result:  err.Error(),
		}
		return
	}
	conn.SetDeadline(time.Now().Add(time.Second * 10))

	_, err = fmt.Fprintln(conn, minerConfig.Request.Body)
	if err != nil {
		results <- MinerResult{
			Success: false,
			Name:    minerConfig.Name,
			Result:  err.Error(),
		}
		return
	}

	result, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		results <- MinerResult{
			Success: false,
			Name:    minerConfig.Name,
			Result:  err.Error(),
		}
	}
	results <- CreateResult(result, minerConfig)
}
