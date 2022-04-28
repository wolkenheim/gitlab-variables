package files

import (
	"fmt"
	"github.com/spf13/afero"
	"gitlab-variables/src/app"
	"testing"
)

func TestLastBackupFileName(t *testing.T) {
	names := []string{"2022-04-28T12:57:03+02:00.json", ".DS_Store", "2022-04-28T13:03:19+02:00.json",
		"2021-12-23T18:02:19+01:00.json"}
	got := getLastBackupFileName(names)

	if got != "2022-04-28T13:03:19+02:00.json" {
		t.Error("Youngest Date not found")
	}
}

func TestGetStatForLastBackupFile(t *testing.T) {
	app.ReadConfig("testing")

	fmt.Println(getProjectBackupPath())

	osMock := afero.NewMemMapFs()

	osMock.Create(getProjectBackupPath() + "/sdsdsd.txt")
	//open, _ := osMock.Open(getProjectBackupPath())
	//open.WriteString("sdsdd.txt")

	backup := NewBackup(osMock)
	backup.createBackupDirIfNotExists()
	list := backup.GetBackupFileNames()

	fmt.Printf("size:%d\n", backup.GetLastBackupFileSize())

	fmt.Printf("%v", list)

}
