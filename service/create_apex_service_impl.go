package service

import (
	"database/sql"
	"integrasi_apex_ems/helper"
	"integrasi_apex_ems/model/domain"
	"integrasi_apex_ems/model/web"
	"integrasi_apex_ems/repository"
	"time"
)

type ApexServiceImpl struct {
	NasabahRepo       repository.NasabahRepository
	TabungRepo        repository.TabungRepository
	SysDaftarUserRepo repository.SysDaftarUserRepository
	DB1, DB2          *sql.DB
}

// Constructor returning Apex Service
func NewApexService(nasabahRepository repository.NasabahRepository, tabungRepository repository.TabungRepository, sysDaftarUserRepository repository.SysDaftarUserRepository, db1, db2 *sql.DB) ApexService {

	return &ApexServiceImpl{
		NasabahRepo:       nasabahRepository,
		TabungRepo:        tabungRepository,
		SysDaftarUserRepo: sysDaftarUserRepository,
		DB1:               db1,
		DB2:               db2,
	}
}

func (service *ApexServiceImpl) CreateApex(request web.ApexRequest) web.ApexResponse {

	// Begin Transaction to nasabah & tabung
	tx, err := service.DB1.Begin()
	helper.FatalIfError(err)
	defer helper.CommitOrRollback(tx) // Handle if err happens then Rollback

	nasabah := domain.Nasabah{}
	nasabah.Nasabah_Id = "000001107"
	nasabah.Nama_Nasabah = request.Nama_Lembaga
	nasabah.Alamat = request.Alamat
	nasabah.Telpon = request.Telpon
	nasabah.Jenis_Kelamin = "L"
	nasabah.TempatLahir = "Bandung"
	nasabah.TglLahir = "1997-09-17"
	nasabah.Jenis_Id = "1"
	nasabah.No_Id = "001928477766641"
	nasabah.Kode_Group1 = "1"
	nasabah.Kode_Group2 = "01"
	nasabah.Kode_Group3 = "001"
	nasabah.Kode_Agama = "Islam"
	nasabah.Desa = "Desa"
	nasabah.Kecamatan = "Kecamatan"
	nasabah.Kota_Kab = "0121"
	nasabah.Provinsi = "008"
	nasabah.Verifikasi = "1"
	nasabah.Hp = request.Telpon
	nasabah.Tgl_Register = time.Now()
	nasabah.Nama_Ibu_Kandung = "Ibu"
	nasabah.Kodepos = "12345"
	nasabah.Kode_Kantor = "001"
	nasabah.UserId = 118
	nasabah.Nama_Alias = request.Nama_Lembaga
	nasabah.Status_Gelar = "0100"
	nasabah.Jenis_Debitur = "0"
	nasabah.Kode_Area = "23"
	nasabah.Negara_Domisili = "ID"
	nasabah.Gol_Debitur = "907"
	nasabah.Langgar_Bmpk = "T"
	nasabah.Lampaui_Bmpk = "T"
	nasabah.Nama_Nasabah_Sid = request.Nama_Lembaga
	nasabah.Alamat2 = request.Alamat
	nasabah.Flag_Masa_Berlaku = "1"
	nasabah.Status_Marital = "Single"
	nasabah.Hp1 = request.Telpon
	nasabah.Hp2 = request.Telpon
	nasabah.Status_Tempat_Tinggal = "Milik Sendiri"
	nasabah.Masa_Berlaku_Ktp = time.Now().AddDate(7, 0, 0)
	nasabah = service.NasabahRepo.CreateNasabah(tx, nasabah)

	tabung := domain.Tabung{}
	tabung.No_Rekening = request.KodeLkm
	tabung.Nasabah_Id = "000001107"
	tabung.Kode_Bi_Pemilik = "874"
	tabung.Suku_Bunga = 0
	tabung.Persen_Pph = 0
	tabung.Tgl_Register = time.Now()
	tabung.Saldo_Akhir = 0
	tabung.Kode_Group1 = "001"
	tabung.Kode_Group2 = "01"
	tabung.Verifikasi = "1"
	tabung.Status = 1
	tabung.Kode_Kantor = "001"
	tabung.Kode_Integrasi = "01"
	tabung.Kode_Produk = "01"
	tabung.UserId = 118
	tabung.Kode_Group3 = ""
	tabung.Minimum = 0
	tabung.Setoran_Minimum = 0
	tabung.Jkw = 0
	tabung.Abp = 0
	tabung.Setoran_Wajib = 0
	tabung.Adm_Per_Bln = 0
	tabung.Target_Nominal = 0
	tabung.Saldo_Akhir_Titipan_bunga = 0
	tabung.Kode_Bi_Lokasi = "1"
	tabung.Saldo_Akhir_Titipan_bunga = 0
	tabung.Saldo_Titipan_Bunga_Ks = 0
	tabung.Saldo_Blokir = 0
	tabung.Premi = 0
	tabung.Kode_Keterkaitan = "2"
	tabung.Kode_Kantor_Kas = "01"
	tabung.No_Rekening_Virtual = request.KodeLkm
	tabung = service.TabungRepo.CreateTabung(tx, tabung)

	// Begin Transaction to sys_daftar_user
	tx, err = service.DB2.Begin()
	helper.FatalIfError(err)
	defer helper.CommitOrRollback(tx) // Handle if err happens then Rollback
	sysDaftarUser := domain.SysDaftarUser{}
	sysDaftarUser.User_Name = request.KodeLkm
	sysDaftarUser.User_Password = "TKkRamfizZc="
	sysDaftarUser.Nama_Lengkap = request.Nama_Lembaga
	sysDaftarUser.Unit_Kerja = "001"
	sysDaftarUser.Jabatan = "Echannel"
	sysDaftarUser.User_Code = "1"
	sysDaftarUser.Tgl_Expired = time.Now().AddDate(7, 0, 0)
	sysDaftarUser.User_Web_Password = helper.HashSha1Pass()
	sysDaftarUser.Flag = 0
	sysDaftarUser.Status_Aktif = 1
	sysDaftarUser.Penerimaan = 0
	sysDaftarUser.Pengeluaran = 0
	sysDaftarUser = service.SysDaftarUserRepo.CreateSysDaftarUser(tx, sysDaftarUser)

	// Converting data => 3 repo to ApexResponse by helper..
	return helper.ToApexResponse(nasabah, tabung, sysDaftarUser)
}
