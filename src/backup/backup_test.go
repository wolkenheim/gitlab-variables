package backup

import (
	"github.com/spf13/afero"
	"gitlab-variables/src/app"
	"testing"
)

func TestLastBackupFileName(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		var names []string
		_, err := getLastBackupFileName(names)

		if err == nil {
			t.Error("Empty slice should return error")
		}
	})

	t.Run("valid list", func(t *testing.T) {
		names := []string{"2022-04-28T12:57:03+02:00.json", ".DS_Store", "2022-04-28T13:03:19+02:00.json",
			"2021-12-23T18:02:19+01:00.json"}
		got, _ := getLastBackupFileName(names)

		if got != "2022-04-28T13:03:19+02:00.json" {
			t.Error("Youngest Date not found")
		}
	})

	t.Run("invalid list", func(t *testing.T) {
		names := []string{"ssds.txt", ".DS_Store"}
		_, err := getLastBackupFileName(names)

		if err == nil {
			t.Error("Invalid backup file names should return error")
		}
	})
}

func TestGetStatForLastBackupFile(t *testing.T) {
	app.ReadConfig("testing")

	t.Run("missing backup file should return 0", func(t *testing.T) {
		osMock := afero.NewMemMapFs()

		backup := NewBackup(osMock)
		backup.createBackupDirIfNotExists()

		got := backup.GetLastBackupFileSize()
		if got != 0 {
			t.Error("missing file should return 0")
		}
	})

	t.Run("invalid backup file name should return 0", func(t *testing.T) {
		osMock := afero.NewMemMapFs()
		osMock.Create(getProjectBackupPath() + "/sdsdsd.txt")

		backup := NewBackup(osMock)
		backup.createBackupDirIfNotExists()

		got := backup.GetLastBackupFileSize()
		if got != 0 {
			t.Error("Invalid file name should return 0")
		}
	})

	t.Run("valid backup file success", func(t *testing.T) {
		osMock := afero.NewMemMapFs()
		path := getProjectBackupPath() + "/2022-04-28T12:57:03+02:00.json"
		file, _ := osMock.Create(path)
		defer file.Close()

		file.WriteString("asdassda asdad")

		backup := NewBackup(osMock)
		backup.createBackupDirIfNotExists()

		got := backup.GetLastBackupFileSize()

		if got != 14 {
			t.Error("valid file should return correct size of 14")
		}
	})
}
