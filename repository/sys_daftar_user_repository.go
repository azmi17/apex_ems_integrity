package repository

import (
	"database/sql"
	"integrasi_apex_ems/model/domain"
)

// Abstraction
type SysDaftarUserRepository interface {
	CreateSysDaftarUser(tx *sql.Tx, tabung domain.SysDaftarUser) domain.SysDaftarUser
}
