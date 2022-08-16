package helper

import (
	"integrasi_apex_ems/model/domain"
	"integrasi_apex_ems/model/web"
)

func ToApexResponse(nasabah domain.Nasabah, tabung domain.Tabung, sysDaftarUser domain.SysDaftarUser) web.ApexResponse {

	return web.ApexResponse{
		KodeLkm:        nasabah.Nasabah_Id,
		Nama_Lembaga:   nasabah.Nama_Nasabah,
		Alamat:         nasabah.Alamat,
		Telpon:         nasabah.Telpon,
		No_rekening:    tabung.No_Rekening,
		Saldo_Akhir:    int(tabung.Saldo_Akhir),
		User_Name_Smec: sysDaftarUser.User_Name,
		Password_Smec:  sysDaftarUser.User_Web_Password_Hash,
		User_Id:        tabung.UserId,
	}
}
