package domain

import "time"

type SysDaftarUser struct {
	User_Name               string
	User_Password           string
	Nama_Lengkap            string
	Penerimaan, Pengeluaran float32
	Unit_Kerja              string
	Jabatan                 string
	User_Code               string
	Tgl_Expired             time.Time
	Flag                    int
	Status_Aktif            int
	User_Web_Password       string
}
