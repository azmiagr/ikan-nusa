package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
)

type IProvinceService interface {
	GetAllProvinces() ([]*model.GetProvinceResponse, error)
}

type ProvinceService struct {
	ProvinceRepository repository.IProvinceRepository
}

func NewProvinceService(provinceRepository repository.IProvinceRepository) IProvinceService {
	return &ProvinceService{
		ProvinceRepository: provinceRepository,
	}
}

func (p *ProvinceService) GetAllProvinces() ([]*model.GetProvinceResponse, error) {
	var res []*model.GetProvinceResponse

	provinces, err := p.ProvinceRepository.GetAllProvince()
	if err != nil {
		return nil, err
	}

	for _, v := range provinces {
		res = append(res, &model.GetProvinceResponse{
			ProvinceID:   v.ProvinceID,
			ProvinceName: v.ProvinceName,
		})
	}

	return res, nil
}
