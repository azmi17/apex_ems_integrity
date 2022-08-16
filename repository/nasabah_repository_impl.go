package repository

import (
	"database/sql"
	"integrasi_apex_ems/helper"
	"integrasi_apex_ems/model/domain"
)

type NasabahRepositoryImpl struct {
}

// Constructor returning Nasabah Repository
func NewNasabahRepository() NasabahRepository {
	return &NasabahRepositoryImpl{}
}

func (repository *NasabahRepositoryImpl) CreateNasabah(tx *sql.Tx, nasabah domain.Nasabah) domain.Nasabah {

	// SQL Statement
	SQL := `INSERT INTO nasabah(nasabah_id, nama_nasabah, alamat, telpon, jenis_kelamin, tempatlahir, tgllahir, jenis_id, no_id, kode_group1, kode_group2, kode_group3, kode_agama, desa, kecamatan, kota_kab, propinsi, verifikasi, hp, tgl_register, nama_ibu_kandung, kodepos, kode_kantor, userid, nama_alias, status_gelar, jenis_debitur, kode_area, negara_domisili, gol_debitur, langgar_bmpk, lampaui_bmpk, nama_nasabah_sid, alamat2, flag_masa_berlaku, status_marital, hp1, hp2, status_tempat_tinggal, masa_berlaku_ktp) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	// Exec
	_, err := tx.Exec(SQL,
		nasabah.Nasabah_Id,
		nasabah.Nama_Nasabah,
		nasabah.Alamat,
		nasabah.Telpon,
		nasabah.Jenis_Kelamin,
		nasabah.TempatLahir,
		nasabah.TglLahir,
		nasabah.Jenis_Id,
		nasabah.No_Id,
		nasabah.Kode_Group1,
		nasabah.Kode_Group2,
		nasabah.Kode_Group3,
		nasabah.Kode_Agama,
		nasabah.Desa,
		nasabah.Kecamatan,
		nasabah.Kota_Kab,
		nasabah.Provinsi,
		nasabah.Verifikasi,
		nasabah.Hp,
		nasabah.Tgl_Register,
		nasabah.Nama_Ibu_Kandung,
		nasabah.Kodepos,
		nasabah.Kode_Kantor,
		nasabah.UserId,
		nasabah.Nama_Alias,
		nasabah.Status_Gelar,
		nasabah.Jenis_Debitur,
		nasabah.Kode_Area,
		nasabah.Negara_Domisili,
		nasabah.Gol_Debitur,
		nasabah.Langgar_Bmpk,
		nasabah.Lampaui_Bmpk,
		nasabah.Nama_Nasabah_Sid,
		nasabah.Alamat2,
		nasabah.Flag_Masa_Berlaku,
		nasabah.Status_Marital,
		nasabah.Hp1,
		nasabah.Hp2,
		nasabah.Status_Tempat_Tinggal,
		nasabah.Masa_Berlaku_Ktp,
	)
	helper.FatalIfError(err)
	return nasabah
}

func (repository *NasabahRepositoryImpl) GetMaxNasabahID() {

}
