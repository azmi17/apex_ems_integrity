package web

type ApexRequest struct {
	KodeLkm      string `json:"kode_lkm" binding:"required"`
	Nama_Lembaga string `json:"nama_lembaga" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
	Telpon       string `json:"telpon" binding:"required"`
	User_Id      int    `json:"user_id" binding:"required"`
}
