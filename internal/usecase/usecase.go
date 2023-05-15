package usecase

import (
	"udv-test-task-go/config"
)

type BackupInterface interface {
	BackupDatabase(cfg config.Config) error
}
type DeleteDBInterface interface {
	DeleteDatabase() error
}
type GetListDBInterface interface {
	GetListDatabases() ([]string, error)
}
type RestoreInterface interface {
	RestoreDatabase(cfg config.Config) error
}
