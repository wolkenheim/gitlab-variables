package backup

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func getProjectPath() string {
	return fmt.Sprintf("./data/%s", viper.Get("projectName"))
}

func getFullPathForBackupFilePath(name string) string {
	return fmt.Sprintf("%s/%s", getProjectBackupPath(), name)
}

func getProjectBackupPath() string {
	return fmt.Sprintf("%s/backup", getProjectPath())
}

func getNewBackupFilePath() string {
	return fmt.Sprintf("%s/%s.json", getProjectBackupPath(), time.Now().Format(time.RFC3339))
}

func getUpdateVarsFilePath() string {
	return fmt.Sprintf("%s/%s", getProjectPath(), viper.Get("newVariablesFile"))
}

type youngestDate struct {
	name string
	time time.Time
}

func getLastBackupFileName(names []string) (string, error) {

	if len(names) == 0 {
		return "", errors.New("no backup backup found")
	}

	var y youngestDate
	for _, name := range names {
		nameParts := strings.Split(name, ".")
		if len(nameParts) != 2 {
			continue
		}
		date, err := time.Parse(time.RFC3339, nameParts[0])
		if err != nil {
			continue
		}

		if y.name == "" {
			y = youngestDate{name, date}
			continue
		}

		if date.After(y.time) {
			y = youngestDate{name, date}
		}
	}

	if y.name == "" {
		return "", errors.New("no valid backup file found")
	}

	return y.name, nil
}
