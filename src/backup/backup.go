package backup

import (
	"github.com/spf13/afero"
	"gitlab-variables/src/util"
	"io"
	"log"
)

type Backup struct {
	appFs afero.Fs
}

func NewBackup(appFs afero.Fs) *Backup {
	return &Backup{appFs}
}

func (backup *Backup) GetLastBackupFileSize() (fileSize int64, backupPath string) {

	previousBackupFilePath, err := getLastBackupFileName(backup.GetBackupFileNames())
	if err != nil {
		return 0, ""
	}

	path := getFullPathForBackupFilePath(previousBackupFilePath)
	info, _ := backup.appFs.Stat(path)

	return info.Size(), previousBackupFilePath
}

// BackupGitlabVariables creates a new backup file of variables
// if the new variables differ from the last backup file
// it will return in any case a path to the current valid backup file
func (backup *Backup) BackupGitlabVariables(list []byte) (backupPath string) {
	backup.CreateBackupDirIfNotExists()

	prevSize, prevPath := backup.GetLastBackupFileSize()
	backupPath = getNewBackupFilePath()

	file, err := backup.appFs.Create(backupPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	prettyJSON, _ := util.PrettyString(list)
	_, err = file.WriteString(prettyJSON)
	if err != nil {
		log.Fatal(err)
	}

	// if last backup file has same size delete current file and keep previous
	curInfo, _ := backup.appFs.Stat(backupPath)
	if prevSize == curInfo.Size() {
		backup.appFs.Remove(backupPath)
		return getFullPathForBackupFilePath(prevPath)
	}
	return backupPath
}

func (backup *Backup) CreateBackupDirIfNotExists() {
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

func (backup *Backup) CreateNewVariablesFileFromBackupFile(backupPath string) {
	var err error
	_, err = backup.appFs.Stat(getUpdateVarsFilePath())
	if err == nil {
		log.Fatal("Update Variable File exists already. Will not overwrite it.")
	}

	backupFile, err := backup.appFs.Open(backupPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer backupFile.Close()

	varFile, err := backup.appFs.Create(getUpdateVarsFilePath())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer varFile.Close()

	_, err = io.Copy(varFile, backupFile)

	if err != nil {
		log.Fatal(err.Error())
	}
}
