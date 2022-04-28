package files

import (
	"gitlab-variables/src/util"
	"io/ioutil"
	"log"
)

func ReadNewVariablesFile() []util.Variable {
	return readVariablesFile(getNewVarsFilePath())
}

func readVariablesFile(path string) []util.Variable {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	return util.ParseVariableJson(content)
}
