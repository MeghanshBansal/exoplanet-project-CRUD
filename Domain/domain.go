package Domain

import "NTTData/Database"

type Service interface {
	Exoplanet
	Fuel
}

type Domain struct {
	DB Database.DbService
}

func NewDomainService(db Database.DbService) Service {
	return &Domain{
		DB: db,
	}
}
