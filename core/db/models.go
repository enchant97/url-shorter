package db

import (
	"github.com/enchant97/url-shorter/core"
	"gorm.io/gorm"
)

type Short struct {
	gorm.Model
	ShortID    string `gorm:"unique"`
	TargetURL  string
	VisitCount int `gorm:"default:0"`
}

func (s *Short) IntoCoreShort() core.Short {
	return core.Short{
		ShortID:   s.ShortID,
		TargetURL: s.TargetURL,
	}
}
