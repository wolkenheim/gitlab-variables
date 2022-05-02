package backup

import (
	"github.com/spf13/afero"
	"gitlab-variables/src/util"
	"log"
)

func (backup *Backup) ReadNewVariablesFile() []util.Variable {
	return backup.readVariablesFile(getNewVarsFilePath())
}

func (backup *Backup) readVariablesFile(path string) []util.Variable {
	content, err := afero.ReadFile(backup.appFs, path)

	if err != nil {
		log.Fatal(err.Error())
	}

	return util.ParseVariableJson(content)
}
