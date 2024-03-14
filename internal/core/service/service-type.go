package service

import (
	"context"
	"log/slog"

	"github.com/davifrjose/My_Turn/internal/core/model"
	"github.com/davifrjose/My_Turn/internal/core/port"
	"github.com/google/uuid"
)

type ServiceType struct {
	serviceTypeRepo port.ServiceTypeRepository
}

func NewServiceType(repo port.ServiceTypeRepository) *ServiceType {
	return &ServiceType{
		serviceTypeRepo: repo,
	}
}

func (serviceType *ServiceType) CreateServiceType(ctx context.Context, serviceTypeParams *port.CreateServiceTypeParams) (*model.ServiceType, error) {
	serviceTypeResponse, err := serviceType.serviceTypeRepo.CreateServiceType(ctx, &model.ServiceType{
		Id:            uuid.New(),
		Name:          serviceTypeParams.Name,
		Code:          serviceTypeParams.Code,
		Status:        serviceTypeParams.Status,
		InstitutionId: serviceTypeParams.InstitutionId,
	})
	if err != nil {
		slog.Info(err.Error())
		if err == model.ErrorConflictingData {
			return nil, err
		}
		return nil, model.ErrorInternal
	}

	return serviceTypeResponse, nil
}
