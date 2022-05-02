package list

import (
	"gitlab-variables/src/util"
	"testing"
)

func TestBuildChangList(t *testing.T) {
	c := NewCompound(nil, nil)

	t.Run("update existing variable", func(t *testing.T) {

		oldList := make([]util.Variable, 1)
		oldList[0] = util.Variable{Key: "DB_PASSWORD", Value: "oldPassword"}

		newList := make([]util.Variable, 1)
		newList[0] = util.Variable{Key: "DB_PASSWORD", Value: "c3po"}

		changeList := c.buildChangeList(newList, oldList)

		if len(changeList) != 1 {
			t.Error("Result List Length should be 1")
		}

		if changeList[0].ChangeType != util.UPDATE {
			t.Error("Type should be update")
		}

		if changeList[0].Variable.Key != "DB_PASSWORD" {
			t.Error("Key not correct")
		}

		if changeList[0].Variable.Value != "c3po" {
			t.Error("Password should get updated")
		}
	})

	t.Run("create new variable", func(t *testing.T) {

		oldList := make([]util.Variable, 1)
		oldList[0] = util.Variable{Key: "DB_PASSWORD", Value: "c3po"}

		newList := make([]util.Variable, 2)
		newList[0] = util.Variable{Key: "DB_PASSWORD", Value: "c3po"}
		newList[1] = util.Variable{Key: "NEW_VAR", Value: "hello9"}

		changeList := c.buildChangeList(newList, oldList)

		if len(changeList) != 1 {
			t.Error("Result List Length should be 1")
		}

		if changeList[0].ChangeType != util.CREATE {
			t.Error("Type should be create")
		}

		if changeList[0].Variable.Key != "NEW_VAR" {
			t.Error("Key not correct")
		}

		if changeList[0].Variable.Value != "hello9" {
			t.Error("Value not correct")
		}

	})
}
