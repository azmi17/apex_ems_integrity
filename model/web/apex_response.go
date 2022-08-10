package web

type ApexResponse struct {
	KodeLkm        string `json:"kode_lkm"`
	Nama_Lembaga   string `json:"nama_lembaga"`
	Alamat         string `json:"alamat"`
	Telpon         string `json:"telpon"`
	NasabahId      string `json:"nasabah_id"`
	Nama_Nasabah   string `json:"nama_nasabah"`
	No_rekening    string `json:"no_rekening"`
	Saldo_Akhir    int    `json:"saldo_akhir"`
	User_Name_Smec string `json:"user_name_smec"`
	Password_Smec  string `json:"password_smec"`
}
