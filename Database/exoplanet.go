package Database

import (
	"NTTData/Models"
	"errors"
)

type ExoplanetDb interface {
	AddExoplanetToDb(exoplanet Models.Exoplanet) error
	GetAllExoplanetsDb() ([]Models.Exoplanet, error)
	GetExoplanetByIdDb(id string) (Models.Exoplanet, error)
	UpdateExoplanetByIdDb(id string, exoplanet Models.Exoplanet) error
	DeleteExoplanetByIdDb(id string) error
}

func (m *Db) AddExoplanetToDb(exoplanet Models.Exoplanet) error {
	err := m.DB.Create(&exoplanet).Error
	if err != nil {
		return errors.New("failed to add exoplanet")
	}
	return nil
}

func (m *Db) GetAllExoplanetsDb() ([]Models.Exoplanet, error) {
	var exoplanets []Models.Exoplanet
	err := m.DB.Find(&exoplanets).Error
	if err != nil {
		return exoplanets, errors.New("failed to get all exoplanets")
	}
	return exoplanets, nil
}

func (m *Db) GetExoplanetByIdDb(id string) (Models.Exoplanet, error) {
	var exoplanet Models.Exoplanet
	if err := m.DB.First(&exoplanet, "id = ?", id).Error; err != nil {
		return Models.Exoplanet{}, errors.New("failed to get exoplanet: " + err.Error())
	}
	return exoplanet, nil
}

func (m *Db) UpdateExoplanetByIdDb(id string, exoplanet Models.Exoplanet) error {
	err := m.DB.Where("id = ?", id).Updates(&exoplanet).Error
	if err != nil {
		return errors.New("failed to update exoplanet")
	}
	return nil
}

func (m *Db) DeleteExoplanetByIdDb(id string) error {
	err := m.DB.Where("id = ?", id).Delete(&Models.Exoplanet{}).Error
	if err != nil {
		return errors.New("failed to delete exoplanet")
	}
	return nil
}
