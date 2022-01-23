package github_provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aemloviji/golang-guide-microservices/src/api/clients/restclient"
	"github.com/aemloviji/golang-guide-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accesToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accesToken))

	// Test: #TestCreateRepoErrorRestClient
	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	// Test: #TestCreateRepoInvalidResponseBody
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		// Test: #TestCreateRepoInvalidErrorInterface
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		// return actual response status code. Because response body does not contain status code in it's content
		errResponse.StatusCode = response.StatusCode
		// Test: #TestCreateRepoUnauthorized
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	// Test: #TestCreateInvalidSuccessResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}

	// Test: #TestCreateRepoNoError
	return &result, nil
}
