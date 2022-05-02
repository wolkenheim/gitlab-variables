package backup

import (
	"errors"
	"github.com/spf13/afero"
	"log"
	"strings"
	"time"
)

type Backup struct {
	appFs afero.Fs
}

type youngestDate struct {
	name string
	time time.Time
}

func NewBackup(appFs afero.Fs) *Backup {
	return &Backup{appFs}
}

func (backup *Backup) GetLastBackupFileSize() int64 {

	previousBackupFileName, err := getLastBackupFileName(backup.GetBackupFileNames())
	if err != nil {
		return 0
	}

	path := getPathForBackupFileName(previousBackupFileName)
	info, _ := backup.appFs.Stat(path)

	return info.Size()
}

func (backup *Backup) BackupGitlabVariables(list []byte) {
	backup.createBackupDirIfNotExists()

	prevSize := backup.GetLastBackupFileSize()
	curPath := getBackupFilePath()

	file, err := backup.appFs.Create(curPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(list))
	if err != nil {
		log.Fatal(err)
	}

	// if last backup file has same size  - do delete
	curInfo, _ := backup.appFs.Stat(curPath)
	if prevSize == curInfo.Size() {
		backup.appFs.Remove(curPath)
	}
}

func (backup *Backup) createBackupDirIfNotExists() {
	_, err := backup.appFs.Stat(getProjectBackupPath())
	if err == nil {
		return
	}

	err = backup.appFs.MkdirAll(getProjectBackupPath(), 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (backup *Backup) GetBackupFileNames() (names []string) {
	open, _ := backup.appFs.Open(getProjectBackupPath())

	files, err := open.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return names
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
