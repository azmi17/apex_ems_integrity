package repository

import (
	"database/sql"
	"integrasi_apex_ems/model/domain"
)

// Abstraction
type NasabahRepository interface {
	CreateNasabah(tx *sql.Tx, nasabah domain.Nasabah) domain.Nasabah
}
