// stress_test.go

package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestStressRegisterEndpoint(t *testing.T) {
	runStressTest(t, "http://localhost:8080/register", makeRequestWithPayload)
}

func TestStressGetUserEndpoint(t *testing.T) {
	runStressTest(t, "http://localhost:8080/get-user", makeRquest)
}

func TestStressAdminListUsersEndpoint(t *testing.T) {
	runStressTest(t, "http://localhost:8080/admin/users", makeRquest)
}

func TestStressAdminSearchUserEndpoint(t *testing.T) {
	runStressTest(t, "http://localhost:8080/admin/users/username", makeRquest)
}

func runStressTest(t *testing.T, apiEndpoint string, requestFunc func(string, string)) {
	var wg sync.WaitGroup
	// concurrency := 10
	requests := 100

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestFunc(apiEndpoint, "POST")
		}()
	}

	wg.Wait()
}

func makeRequestWithPayload(apiEndpoint, method string) {
	payload := `{"username": "example", "password": "secret"}`
	makeRequest(apiEndpoint, method, strings.NewReader(payload))
}

func makeRquest(apiEndpoint, method string) {
	makeRequest(apiEndpoint, method, nil)
}

func makeRequest(apiEndpoint, method string, payload *strings.Reader) {
	resp, err := http.Post(apiEndpoint, "application/json", payload)
	handleResponse(apiEndpoint, resp, err)
}

func handleResponse(apiEndpoint string, resp *http.Response, err error) {
	if err != nil {
		fmt.Printf("Error making request to %s: %v\n", apiEndpoint, err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %v\n", apiEndpoint, err)
		return
	}

	fmt.Printf("Response Body from %s: %s\n", apiEndpoint, responseBody)
	fmt.Printf("Status Code from %s: %d\n", apiEndpoint, resp.StatusCode)
}
