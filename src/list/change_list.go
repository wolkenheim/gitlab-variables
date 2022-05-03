package list

import (
	"fmt"
	"gitlab-variables/src/util"
)

func (c *Compound) initChangeList() []util.ChangeVariable {
	newList := c.backup.ReadNewVariablesFile()
	currentList := c.fetchAllAndBackupAndParse()
	return c.buildChangeList(newList, currentList)
}

func (c *Compound) buildChangeList(updateList []util.Variable, currentList []util.Variable) []util.ChangeVariable {

	curMap := make(map[string]util.Variable)
	for _, curVar := range currentList {
		curMap[curVar.Key] = curVar
	}

	updateMap := make(map[string]util.Variable)
	for _, upVar := range updateList {
		updateMap[upVar.Key] = upVar
	}

	// update: key exists in both lists and value (or other attributes) has changed
	// create: key does not exist in currList but only in updateList
	// delete: key does not exist in updateList but exists in currList. This will not work with the regular loop

	var changeList []util.ChangeVariable

	// create or update case
	for _, newVar := range updateList {

		old, ok := curMap[newVar.Key]
		if ok {
			if newVar.Value != old.Value {
				changeList = append(changeList, util.ChangeVariable{Variable: newVar, ChangeType: util.UPDATE})
			}
		} else {
			changeList = append(changeList, util.ChangeVariable{Variable: newVar, ChangeType: util.CREATE})
		}
	}

	// delete case
	for _, curVar := range currentList {
		_, ok := updateMap[curVar.Key]
		if !ok {
			changeList = append(changeList, util.ChangeVariable{Variable: curVar, ChangeType: util.DELETE})
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
