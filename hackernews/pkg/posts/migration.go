package posts

import (
	"github.com/gocopper/copper/cerrors"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{
		db: db,
	}
}

type Migration struct {
	db *gorm.DB
}

func (m *Migration) Run() error {
	err := m.db.AutoMigrate(Post{}, Vote{})
	if err != nil {
		return cerrors.New(err, "failed to auto migrate posts models", nil)
	}

	return nil
}
