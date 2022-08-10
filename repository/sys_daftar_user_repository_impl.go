package repository

import (
	"database/sql"
	"integrasi_apex_ems/helper"
	"integrasi_apex_ems/model/domain"
)

type SysDaftarUserRepositoryImpl struct {
}

// Constructor returning sys_daftar_user Repository
func NewSysDaftarUserRepository() SysDaftarUserRepository {
	return &SysDaftarUserRepositoryImpl{}
}
func (repository *SysDaftarUserRepositoryImpl) CreateSysDaftarUser(tx *sql.Tx, sdu domain.SysDaftarUser) domain.SysDaftarUser {

	// SQL Statement
	SQL := `INSERT INTO sys_daftar_user(user_name, user_password, nama_lengkap, unit_kerja, jabatan, user_code, tgl_expired, user_web_password, flag, status_aktif, penerimaan, pengeluaran) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	// Exec
	_, err := tx.Exec(SQL,
		sdu.User_Name,
		sdu.User_Password,
		sdu.Nama_Lengkap,
		sdu.Unit_Kerja,
		sdu.Jabatan,
		sdu.User_Code,
		sdu.Tgl_Expired,
		sdu.User_Web_Password,
		sdu.Flag,
		sdu.Status_Aktif,
		sdu.Penerimaan,
		sdu.Pengeluaran,
	)
	helper.FatalIfError(err)
	return sdu
}
