package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
)

type IDistrictService interface {
	GetAllDistricts() ([]*model.GetDistrictResponse, error)
}

type DistrictService struct {
	DistrictRepository repository.IDistrictRepository
}

func NewDistrictService(districtRepository repository.IDistrictRepository) IDistrictService {
	return &DistrictService{
		DistrictRepository: districtRepository,
	}
}

func (d *DistrictService) GetAllDistricts() ([]*model.GetDistrictResponse, error) {
	var res []*model.GetDistrictResponse

	districts, err := d.DistrictRepository.GetAllDistricts()
	if err != nil {
		return nil, err
	}

	for _, v := range districts {
		res = append(res, &model.GetDistrictResponse{
			DistrictID:   v.DistrictID,
			DistrictName: v.DistrictName,
		})
	}

	return res, nil
}
