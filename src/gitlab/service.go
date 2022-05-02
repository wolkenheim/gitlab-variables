package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kafka-certificates/src/util"
	"log"
	"net/http"
)

type Service struct {
	apiClient ApiClientInterface
}

type ApiClientInterface interface {
	call(url, method string, requestBody io.Reader) ([]byte, error)
}

func NewGitlabService(apiClient ApiClientInterface) *Service {
	return &Service{apiClient: apiClient}
}

// FetchAllRaw
// https://docs.gitlab.com/ee/api/project_level_variables.html
// curl --header "PRIVATE-TOKEN: <your_access_token>" "https://gitlab.example.com/api/v4/projects/1/variables"
func (service *Service) FetchAllRaw() ([]byte, error) {
	body, err := service.apiClient.call(getVariableURL(), http.MethodGet, nil)
	return body, err
}

func (service *Service) FetchAll() []util.Variable {
	body, err := service.FetchAllRaw()
	if err != nil {
		log.Fatal(err)
	}
	return util.ParseVariableJson(body)
}

func (service *Service) PrintListVariables() {
	body, err := service.FetchAllRaw()
	if err != nil {
		log.Fatal(err)
	}
	res, err := PrettyString(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}

func (service *Service) Show(key string) {
	service.apiClient.call(getVariableDetailURL(key), http.MethodGet, nil)
}

func (service *Service) Create(key string, value string) {
	uv := util.NewVariable(key, value)
	jsonData, _ := json.Marshal(uv)

	service.apiClient.call(getVariableURL(), http.MethodPost, bytes.NewBuffer(jsonData))
}

func (service *Service) Update(key string, value string) {
	uv := util.NewVariable(key, value)
	jsonData, _ := json.Marshal(uv)

	service.apiClient.call(getVariableDetailURL(key), http.MethodPut, bytes.NewBuffer(jsonData))
}

func (service *Service) Delete(key string) {
	service.apiClient.call(getVariableDetailURL(key), http.MethodDelete, nil)
}
