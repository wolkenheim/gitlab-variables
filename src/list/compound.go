package list

import (
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

func (c *Compound) Init() {
	backupPath := c.fetchAllAndBackup()
	c.backup.CreateNewVariablesFileFromBackupFile(backupPath)
}

func (c *Compound) Update() {
	list := c.initChangeList()
	c.processChangeList(list)
}

func (c *Compound) fetchAll() []byte {
	content, err := c.gitlabService.FetchAllRaw()
	if err != nil {
		log.Fatal(err.Error())
	}
	return content
}

func (c *Compound) fetchAllAndBackup() (backupPath string) {
	content := c.fetchAll()
	return c.backup.BackupGitlabVariables(content)
}

func (c *Compound) fetchAllAndBackupAndParse() []util.Variable {
	content := c.fetchAll()
	c.backup.BackupGitlabVariables(content)
	return util.ParseVariableJson(content)
}
