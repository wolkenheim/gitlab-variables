package gitlab

import (
	"fmt"
	"github.com/spf13/viper"
)

func getVariableDetailURL(key string) string {
	return fmt.Sprintf("%s/%s", getVariableURL(), key)
}

func getVariableURL() string {
	return fmt.Sprintf("%s/api/v4/projects/%d/variables", viper.Get("gitlabURL"), viper.Get("projectId"))
}
