package service

import (
	"context"

	"github.com/davifrjose/My_Turn/internal/core/model"
	"github.com/davifrjose/My_Turn/internal/core/port"
)

type ServiceType struct {
	serviceTypeRepo port.ServiceTypeRepository
}

func NewServiceType(repo port.ServiceTypeRepository) *ServiceType {
	return &ServiceType{
		serviceTypeRepo: repo,
	}
}

func (serviceType *ServiceType) CreateServiceType(ctx context.Context, serviceTypeParams *model.ServiceType) (*model.ServiceType, error) {
	serviceTypeResponse, err := serviceType.serviceTypeRepo.CreateServiceType(ctx, serviceTypeParams)
	if err != nil {
		if err == model.ErrorConflictingData {
			return nil, err
		}
		return nil, model.ErrorInternal
	}

	return serviceTypeResponse, nil
}
