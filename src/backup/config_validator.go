package backup

import (
	"strings"
)

func (backup *Backup) IsValidConfigName(configName string) (valid bool) {
	valid = false
	fileNames := backup.GetFileNamesForDir("./config")
	for _, name := range fileNames {
		nameSlice := strings.Split(name, ".")
		if configName == nameSlice[0] {
			valid = true
		}
	}
	return valid
}
