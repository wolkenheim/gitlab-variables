package gitlab

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab-variables/src/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAll(t *testing.T) {
	t.Run("all success", func(t *testing.T) {

		// init gitlab mock api server
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[{\"variable_type\":\"env_var\",\"key\":\"AWS_REDIS_PASSWORD\",\"value\":\"uec89SdsddsQ\","+
				"\"protected\":false,\"masked\":false,\"environment_scope\":\"*\"}]")
		}))
		defer ts.Close()

		// load config and set dynamic url of mock api server
		app.ReadConfig("testing")
		viper.Set("gitlabURL", ts.URL)

		// init service
		gitlabService := NewGitlabService(NewApiClient(&http.Client{}))

		list := gitlabService.FetchAll()

		if list[0].Key != "AWS_REDIS_PASSWORD" {
			t.Error("service should return a list with one variable object")
		}
	})
}
