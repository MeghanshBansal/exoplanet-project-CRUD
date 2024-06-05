package Domain

import (
	"errors"
	"fmt"
	"strings"
)

type Fuel interface {
	CalculateFuel(id string, crewSize int64) (float64, error)
}

func (d *Domain) CalculateFuel(id string, crewSize int64) (float64, error) {
	//f = d / (g^2) * c units
	exoplanet, err := d.DB.GetExoplanetByIdDb(id)
	if err != nil {
		return 0, errors.New("failed to get the exoplanet data")
	}

	var g float64
	if strings.ToLower(exoplanet.TypeOfExoplanet)[0] == 'g' {
		g = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else {
		fmt.Println(exoplanet.Mass, exoplanet.Radius)
		g = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}

	f := float64(exoplanet.DistanceFromEarth) / (g * g) * float64(crewSize)
	return f, nil
}
