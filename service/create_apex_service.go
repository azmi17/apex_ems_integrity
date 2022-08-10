package service

import "integrasi_apex_ems/model/web"

type ApexService interface {
	CreateApex(request web.ApexRequest) web.ApexResponse
}
