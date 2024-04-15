package database

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
)

//
// user related access functions
//

func (d *Database) NewIncident(incident models.Incident) error {
	res := d.Pool.Omit("resolverId").Create(&incident)
	return res.Error
}