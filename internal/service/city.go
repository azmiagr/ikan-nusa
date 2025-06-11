package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
)

type ICityService interface {
	GetAllCities() ([]*model.GetCityResponse, error)
	GetCitiesByProvinceID(provinceID int) ([]*model.GetCityResponse, error)
}

type CityService struct {
	CityRepository repository.ICityRepository
}

func NewCityService(cityRepository repository.ICityRepository) ICityService {
	return &CityService{
		CityRepository: cityRepository,
	}
}

func (c *CityService) GetAllCities() ([]*model.GetCityResponse, error) {
	var res []*model.GetCityResponse

	cities, err := c.CityRepository.GetAllCities()
	if err != nil {
		return nil, err
	}

	for _, v := range cities {
		res = append(res, &model.GetCityResponse{
			CityID:   v.CityID,
			CityName: v.CityName,
		})
	}

	return res, nil
}

func (c *CityService) GetCitiesByProvinceID(provinceID int) ([]*model.GetCityResponse, error) {
	var res []*model.GetCityResponse

	cities, err := c.CityRepository.GetCitiesByProvinceID(provinceID)
	if err != nil {
		return nil, err
	}

	for _, v := range cities {
		res = append(res, &model.GetCityResponse{
			CityID:   v.CityID,
			CityName: v.CityName,
		})
	}

	return res, nil
}
