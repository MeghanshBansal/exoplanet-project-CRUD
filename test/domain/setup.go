package domain

import (
	"Exoplanet/Database"
	"Exoplanet/Domain"
)

func InitializeDomain(dbService Database.DbService) Domain.Service {
	return Domain.NewDomainService(dbService)
}
