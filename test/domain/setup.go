package domain

import (
	"NTTData/Database"
	"NTTData/Domain"
)

func InitializeDomain(dbService Database.DbService) Domain.Service {
	return Domain.NewDomainService(dbService)
}
