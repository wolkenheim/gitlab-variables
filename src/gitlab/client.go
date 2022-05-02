package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
)

type ApiClient struct {
	client *http.Client
}

func NewApiClient(client *http.Client) *ApiClient {
	return &ApiClient{client: client}
}

func (api *ApiClient) call(url, method string, requestBody io.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("PRIVATE-TOKEN", viper.GetString("gitlabToken"))
	req.Header.Set("Content-Type", "application/json")
	response, err := api.client.Do(req)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func PrettyString(response []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, response, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
