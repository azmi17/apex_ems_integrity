package helper

import (
	"integrasi_apex_ems/model/domain"
	"integrasi_apex_ems/model/web"
)

func ToApexResponse(nasabah domain.Nasabah, tabung domain.Tabung, sysDaftarUser domain.SysDaftarUser) web.ApexResponse {

	return web.ApexResponse{
		KodeLkm:      tabung.No_Rekening,
		Nama_Lembaga: nasabah.Nama_Nasabah,
		NasabahId:    nasabah.Nasabah_Id,
		No_rekening:  tabung.No_Rekening,
		Saldo_Akhir:  int(tabung.Saldo_Akhir),
		Nama_Nasabah: nasabah.Nama_Nasabah,
		Alamat:       nasabah.Alamat,
		Telpon:       nasabah.Telpon,
	}
}
