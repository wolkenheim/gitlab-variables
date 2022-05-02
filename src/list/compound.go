package list

import (
	"fmt"
	"gitlab-variables/src/backup"
	"gitlab-variables/src/gitlab"
	"gitlab-variables/src/util"
	"log"
)

type Compound struct {
	gitlabService *gitlab.Service
	backup        *backup.Backup
}

func NewCompound(g *gitlab.Service, b *backup.Backup) *Compound {
	return &Compound{g, b}
}

func (c *Compound) Execute() {
	list := c.buildChangeList()
	c.processChangeList(list)
}

func (c *Compound) fetchAllAndBackup() []util.Variable {
	content, err := c.gitlabService.FetchAllRaw()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.backup.BackupGitlabVariables(content)
	return util.ParseVariableJson(content)
}

func (c *Compound) buildChangeList() []util.ChangeVariable {
	newList := c.backup.ReadNewVariablesFile()
	currentList := c.fetchAllAndBackup()

	if len(newList) != len(currentList) {
		log.Fatal("Have not yet implemented this case")
	}

	// update: key exists in both lists and value (or other attributes) has changed
	// create: key does not exist in currList but only in newList
	// delete: key does not exist in newList but exists in currList. This will not work with the regular loop

	var changeList []util.ChangeVariable
	for _, newVar := range newList {
		for _, curVar := range currentList {

			// key exists. either do nothing or update
			if newVar.Key == curVar.Key {
				if newVar.Value != curVar.Value {
					fmt.Printf("UPDATE: change detected for: %s // old: %s / new: %s\n", newVar.Key,
						curVar.Value, newVar.Value)
					changeList = append(changeList, util.ChangeVariable{Variable: newVar, ChangeType: util.UPDATE})
				}
			}
		}
	}
	return changeList
}

func (c *Compound) processChangeList(list []util.ChangeVariable) {
	if len(list) == 0 {
		fmt.Println("No changes detected. Nothing to do.")
		return
	}

	for _, changeVariable := range list {
		switch changeVariable.ChangeType {
		case util.CREATE:
			c.gitlabService.Create(changeVariable.Variable.Key, changeVariable.Variable.Value)
		case util.UPDATE:
			c.gitlabService.Update(changeVariable.Variable.Key, changeVariable.Variable.Value)
		case util.DELETE:
			c.gitlabService.Delete(changeVariable.Variable.Key)
		}
	}
}
