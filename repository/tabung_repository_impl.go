package repository

import (
	"database/sql"
	"integrasi_apex_ems/helper"
	"integrasi_apex_ems/model/domain"
)

type TabungRepositoryImpl struct {
}

// Constructor returning Tabung Repository
func NewTabungRepository() TabungRepository {
	return &TabungRepositoryImpl{}
}
func (repository *TabungRepositoryImpl) CreateTabung(tx *sql.Tx, tabung domain.Tabung) domain.Tabung {

	// SQL Statement
	SQL := `INSERT INTO tabung(no_rekening, nasabah_id, kode_bi_pemilik, suku_bunga, persen_pph, tgl_register, saldo_akhir, kode_group1, kode_group2, verifikasi, status, kode_kantor, kode_integrasi, kode_produk, userid, kode_group3, minimum, setoran_minimum, jkw, abp, setoran_wajib, adm_per_bln, target_nominal, saldo_akhir_titipan_bunga, kode_bi_lokasi, saldo_titipan_pokok, saldo_titipan_bunga_ks, saldo_blokir, premi, kode_keterkaitan, kode_kantor_kas, no_rekening_virtual) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	// Exec
	_, err := tx.Exec(SQL,
		tabung.No_Rekening,
		tabung.Nasabah_Id,
		tabung.Kode_Bi_Pemilik,
		tabung.Suku_Bunga,
		tabung.Persen_Pph,
		tabung.Tgl_Register,
		tabung.Saldo_Akhir,
		tabung.Kode_Group1,
		tabung.Kode_Group2,
		tabung.Verifikasi,
		tabung.Status,
		tabung.Kode_Kantor,
		tabung.Kode_Integrasi,
		tabung.Kode_Produk,
		tabung.UserId,
		tabung.Kode_Group3,
		tabung.Minimum,
		tabung.Setoran_Minimum,
		tabung.Jkw,
		tabung.Abp,
		tabung.Setoran_Wajib,
		tabung.Adm_Per_Bln,
		tabung.Target_Nominal,
		tabung.Saldo_Akhir_Titipan_bunga,
		tabung.Kode_Bi_Lokasi,
		tabung.Saldo_Titipan_Pokok,
		tabung.Saldo_Titipan_Bunga_Ks,
		tabung.Saldo_Blokir,
		tabung.Premi,
		tabung.Kode_Keterkaitan,
		tabung.Kode_Kantor_Kas,
		tabung.No_Rekening_Virtual,
	)
	helper.FatalIfError(err)
	return tabung
}
