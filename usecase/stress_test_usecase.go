package usecase

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type StressTestInterface interface {
	Execute() (StressTestOutput, error)
}

type StressTestOutput struct {
	AvgResponseTime   float64
	TotalResponseTime float64
	TotalRequest      int
	StatusCounts      map[int]int
	SuccessCount      int
}

type StressTest struct {
	URL         string
	Requests    int
	Concurrency int
}

func NewStressTest(url string, requests int, concurrency int) *StressTest {
	return &StressTest{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}

func (s *StressTest) Execute() (StressTestOutput, error) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, s.Concurrency)

	var totalResponseTime float64
	var totalRequests int
	var successCount int
	statusCounts := make(map[int]int)
	var mu sync.Mutex

	for i := 0; i < s.Requests; i++ {
		semaphore <- struct{}{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			responseTime, code, err := sendRequest(s.URL)

			mu.Lock()
			defer mu.Unlock()

			totalRequests++

			if err != nil {
				if code == http.StatusGatewayTimeout {
					statusCounts[http.StatusGatewayTimeout]++
				} else {
					statusCounts[http.StatusInternalServerError]++
				}
				<-semaphore
				return
			}

			totalResponseTime += responseTime
			statusCounts[code]++
			if code == http.StatusOK {
				successCount++
			}
			<-semaphore
		}()
	}

	wg.Wait()
	close(semaphore)

	avgResponseTime := 0.0
	if totalRequests > 0 {
		avgResponseTime = totalResponseTime / float64(totalRequests)
	}

	return StressTestOutput{
		AvgResponseTime:   avgResponseTime,
		TotalRequest:      totalRequests,
		StatusCounts:      statusCounts,
		TotalResponseTime: totalResponseTime,
		SuccessCount:      successCount,
	}, nil
}

func sendRequest(url string) (float64, int, error) {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return float64(time.Since(start).Milliseconds()), http.StatusInternalServerError, err
	}

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return float64(time.Since(start).Milliseconds()), http.StatusGatewayTimeout, nil
		}
		return float64(time.Since(start).Milliseconds()), http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	return float64(time.Since(start).Milliseconds()), res.StatusCode, nil
}
