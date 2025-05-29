package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
)

type IProductTypeService interface {
	GetAllTypes() ([]*model.GetAllTypesResponse, error)
}

type ProductTypeService struct {
	ProductTypeRepository repository.IProductTypeRepository
}

func NewProductTypeService(productTypeRepository repository.IProductTypeRepository) IProductTypeService {
	return &ProductTypeService{
		ProductTypeRepository: productTypeRepository,
	}
}

func (p *ProductTypeService) GetAllTypes() ([]*model.GetAllTypesResponse, error) {
	var res []*model.GetAllTypesResponse

	types, err := p.ProductTypeRepository.GetAllTypes()
	if err != nil {
		return nil, err
	}

	for _, v := range types {
		res = append(res, &model.GetAllTypesResponse{
			ProductTypeID: v.ProductTypeID,
			Type:          v.Type,
		})
	}

	return res, nil
}
