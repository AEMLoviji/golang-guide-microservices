package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/aemloviji/golang-guide-microservices/src/api/domain/repositories"
	"github.com/aemloviji/golang-guide-microservices/src/api/services"
	"github.com/aemloviji/golang-guide-microservices/src/api/utils/errors"
)

const requestsFile = "./requests_mini.txt"
const codeShouldGenerateRepoName = true

// to test with 360_000 requests
// const requestsFile = "./requests.txt"
// const codeShouldGenerateRepoName = false

var (
	success = make(map[string]string)
	failed  = make(map[string]errors.ApiError)
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open(requestsFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	idendifier := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		repoName := line
		// we need to make sure that request name is unique, as we store the success/failed results in map, key referenced to request name
		if codeShouldGenerateRepoName {
			repoName = fmt.Sprintf("%s---%d", line, idendifier)
		}

		request := repositories.CreateRepoRequest{Name: repoName}
		result = append(result, request)

		idendifier++
	}
	return result
}

func main() {
	requests := getRequests()

	fmt.Printf("about to process %d requests\n\n", len(requests))

	createRepoResultChan := make(chan createRepoResult)
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(&wg, createRepoResultChan)

	for _, request := range requests {
		//fmt.Printf("try to create a go routine for: %q if channel has space\n", request.Name)
		buffer <- true
		wg.Add(1)
		//fmt.Printf("creating a go routine for: %q\n", request.Name)
		go createRepo(&wg, buffer, createRepoResultChan, request)
	}

	wg.Wait()
	close(createRepoResultChan)

	i := 1
	for k, v := range success {
		fmt.Printf("success(%d) request: %q, response: %s\n", i, k, v)
		i++
	}

	i = 1
	for k, v := range failed {
		fmt.Printf("failed(%d) request: %q, response: %s\n", i, k, v.Message())
		i++
	}
}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		// time.Sleep(250 * time.Millisecond)
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}

		// Done on WaitGorup should exactly be called here, not in createRepo() method below
		// if we do otherwise, then that go routine can fail to complete if main go routine completes before
		// To test it, you can uncomment time.Sleep() call above
		wg.Done()
	}
}

func createRepo(wg *sync.WaitGroup, buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	fmt.Printf("request to Github sent for %q\n", request.Name)
	result, err := services.RepositoryService.CreateRepo("your_client_id", request)
	if err != nil {
		fmt.Printf("error response from Github for %q and error is %d\n", request.Name, err.Status())
	} else {
		fmt.Printf("success response from Github for %q and repo name is %q\n", request.Name, result.Name)
	}

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	//fmt.Println("release buffer")
	<-buffer
}
