package backup

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func getProjectPath() string {
	return fmt.Sprintf("./data/%s", viper.Get("projectName"))
}

func getPathForBackupFileName(name string) string {
	return fmt.Sprintf("%s/%s", getProjectBackupPath(), name)
}

func getProjectBackupPath() string {
	return fmt.Sprintf("%s/backup", getProjectPath())
}

func getBackupFilePath() string {
	return fmt.Sprintf("%s/%s.json", getProjectBackupPath(), time.Now().Format(time.RFC3339))
}

func getNewVarsFilePath() string {
	return fmt.Sprintf("%s/%s", getProjectPath(), viper.Get("newVariablesFile"))
}
