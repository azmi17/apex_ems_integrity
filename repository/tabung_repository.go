package repository

import (
	"database/sql"
	"integrasi_apex_ems/model/domain"
)

// Abstraction
type TabungRepository interface {
	CreateTabung(tx *sql.Tx, tabung domain.Tabung) domain.Tabung
}
