package database

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"gorm.io/gorm"
)

func (d *Database) MarkResolved(incident models.Incident, ngoId int) error {
	return d.Pool.Transaction(func(tx *gorm.DB) error {
		return d.Pool.Model(&incident).Updates(&models.Incident{
			Resolved:   true,
			ResolverId: ngoId,
		}).Error
	})
}
