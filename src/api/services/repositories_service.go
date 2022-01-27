package services

import (
	"net/http"
	"sync"

	"github.com/aemloviji/golang-guide-microservices/src/api/config"
	"github.com/aemloviji/golang-guide-microservices/src/api/domain/github"
	"github.com/aemloviji/golang-guide-microservices/src/api/domain/repositories"

	//"github.com/aemloviji/golang-guide-microservices/src/api/log/logger_logrus"
	"github.com/aemloviji/golang-guide-microservices/src/api/log/logger_zap"
	"github.com/aemloviji/golang-guide-microservices/src/api/providers/github_provider"
	"github.com/aemloviji/golang-guide-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	//log_logrus.Info("about to send request to external api", fmt.Sprintf("client_id:%s", clientId), "status:pending")
	logger_zap.Info("about to send request to external api",
		logger_zap.Field("client_id", clientId),
		logger_zap.Field("status", "pending"),
		logger_zap.Field("authenticated", clientId != ""))

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		//log_logrus.Error("response obtained from external api", err, fmt.Sprintf("client_id:%s", clientId), "status:error")
		logger_zap.Error("response obtained from external api", err,
			logger_zap.Field("client_id", clientId),
			logger_zap.Field("status", "error"),
			logger_zap.Field("authenticated", clientId != ""))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	//log_logrus.Info("response obtained from external api", fmt.Sprintf("client_id:%s", clientId), "status:success")
	logger_zap.Info("response obtained from external api",
		logger_zap.Field("client_id", clientId),
		logger_zap.Field("status", "success"),
		logger_zap.Field("authenticated", clientId != ""))

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResult(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *reposService) handleRepoResult(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingResult := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingResult.Response,
			Error:    incomingResult.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}

	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo("TODO_client_id", input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
