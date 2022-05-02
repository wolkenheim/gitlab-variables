package backup

import (
	"github.com/spf13/afero"
	"kafka-certificates/src/util"
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
