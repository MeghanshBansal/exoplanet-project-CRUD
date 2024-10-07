package Domain

import (
	"Exoplanet/Models"
	"errors"
	"log"
	"strings"

	u "github.com/google/uuid"
)

type Exoplanet interface {
	AddExoplanet(exoplanet Models.Exoplanet) error
	GetAllExoplanet() ([]Models.Exoplanet, error)
	GetExoplanetById(id string) (Models.Exoplanet, error)
	UpdateExoplanetById(id string, updatePlanet Models.UpdateMapExoplanet) error
	DeleteExoplanetById(id string) error
}

func (d *Domain) AddExoplanet(exoplanet Models.Exoplanet) error {
	uniqueId := u.New().String()
	exoplanet.Id = uniqueId

	if err := d.constraintChecks(&exoplanet); err != nil {
		log.Println("constraint check failed with err: ", err.Error())
		return err
	}

	err := d.DB.AddExoplanetToDb(exoplanet)
	if err != nil {
		log.Println("failed to insert exoplanet with err: ", err.Error())
		return errors.New("failed to add exoplanet")
	}
	return nil
}

func (d *Domain) GetAllExoplanet() ([]Models.Exoplanet, error) {
	exoplanets, err := d.DB.GetAllExoplanetsDb()
	if err != nil {
		log.Println("failed to get all exoplanets with err: ", err)
		return []Models.Exoplanet{}, errors.New("failed to get all exoplanets")
	}
	return exoplanets, nil
}

func (d *Domain) GetExoplanetById(id string) (Models.Exoplanet, error) {
	exoplanet, err := d.DB.GetExoplanetByIdDb(id)
	if err != nil {
		log.Println("failed to get exoplanet with err: ", err)
		return Models.Exoplanet{}, errors.New("failed to get exoplanet by id: " + id)
	}
	return exoplanet, nil
}

func (d *Domain) UpdateExoplanetById(id string, updatePlanet Models.UpdateMapExoplanet) error {
	exoplanet, err := d.DB.GetExoplanetByIdDb(id)
	if err != nil {
		log.Println("failed to get exoplanet with err: ", err)
		return errors.New("failed to get exoplent to update")
	}

	d.createUpdateMapExoplanet(&exoplanet, updatePlanet)
	err = d.DB.UpdateExoplanetByIdDb(id, exoplanet)
	if err != nil {
		log.Println("failed to update exoplanet with err: ", err)
		return errors.New("failed to update the exoplanet")
	}
	return nil
}

func (d *Domain) DeleteExoplanetById(id string) error {
	err := d.DB.DeleteExoplanetByIdDb(id)
	if err != nil {
		log.Println("failed to get delete with err: ", err)
		return errors.New("failed to delete exoplanet by id: " + id)
	}

	return nil
}

func (d *Domain) createUpdateMapExoplanet(exoplanet *Models.Exoplanet, updatePlanet Models.UpdateMapExoplanet) {
	for idx := range updatePlanet.InputFields {
		if updatePlanet.InputFields[idx].Key == "name" {
			if name, ok := updatePlanet.InputFields[idx].Value.(string); ok {
				exoplanet.Name = name
			}
		}
		if updatePlanet.InputFields[idx].Key == "description" {
			if description, ok := updatePlanet.InputFields[idx].Value.(string); ok {
				exoplanet.Description = description
			}
		}
		if updatePlanet.InputFields[idx].Key == "distanceFromEarth" {
			if distanceFromEarth, ok := updatePlanet.InputFields[idx].Value.(int); ok {
				exoplanet.DistanceFromEarth = distanceFromEarth
			}
		}
		if updatePlanet.InputFields[idx].Key == "radius" {
			if radius, ok := updatePlanet.InputFields[idx].Value.(float64); ok {
				exoplanet.Radius = radius
			}
		}
		if updatePlanet.InputFields[idx].Key == "mass" {
			if mass, ok := updatePlanet.InputFields[idx].Value.(float64); ok {
				exoplanet.Mass = mass
			}
		}
		if updatePlanet.InputFields[idx].Key == "typeOfExoplanet" {
			if typeOfExoplanet, ok := updatePlanet.InputFields[idx].Value.(string); ok {
				if strings.ToLower(typeOfExoplanet)[0] == 't' {
					exoplanet.TypeOfExoplanet = Models.TypeTerrestrial
				} else {
					exoplanet.TypeOfExoplanet = Models.TypeGasGiant
				}
			}
		}
	}
}

func (d *Domain) constraintChecks(exoplanet *Models.Exoplanet) error {
	if strings.ToLower(exoplanet.TypeOfExoplanet) == strings.ToLower(Models.TypeTerrestrial) {
		exoplanet.TypeOfExoplanet = Models.TypeTerrestrial
	} else if strings.ToLower(exoplanet.TypeOfExoplanet) == strings.ToLower(Models.TypeGasGiant) {
		exoplanet.TypeOfExoplanet = Models.TypeGasGiant
	} else {
		return errors.New("wrong value for type of exoplanet")
	}
	if exoplanet.DistanceFromEarth <= 10 || exoplanet.DistanceFromEarth > 1000 {
		return errors.New("wrong value for distance from the earth")
	}
	if exoplanet.TypeOfExoplanet != Models.TypeGasGiant && exoplanet.Mass <= 0.1 || exoplanet.Mass > 10.0 {
		return errors.New("wrong value for mass")
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius > 10.0 {
		return errors.New("wrong value for radius")
	}
	return nil
}
