package service

import (
	"ikan-nusa/entity"
	"ikan-nusa/internal/repository"
)

type IProvinceService interface {
	GetAllProvinces() ([]*entity.Province, error)
}

type ProvinceService struct {
	ProvinceRepository repository.IProvinceRepository
}

func NewProvinceService(provinceRepository repository.IProvinceRepository) IProvinceService {
	return &ProvinceService{
		ProvinceRepository: provinceRepository,
	}
}

func (p *ProvinceService) GetAllProvinces() ([]*entity.Province, error) {
	provinces, err := p.ProvinceRepository.GetAllProvince()
	if err != nil {
		return nil, err
	}

	return provinces, nil
}
