package metadata

import (
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/cqroot/ternote/pkg/config"
	"github.com/cqroot/ternote/pkg/types"
)

var (
	db *gorm.DB
)

func init() {
	basePath, err := config.BasePath()
	if err != nil {
		panic(err)
	}

	dsn := path.Join(basePath, "ternote.db")
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.Note{})
}
