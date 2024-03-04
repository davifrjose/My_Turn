package port

import (
	"context"

	"github.com/davifrjose/My_Turn/internal/core/model"
)

type ServiceTypeRepository interface {
	CreateServiceType(ctx context.Context, serviceType *model.ServiceType) (*model.ServiceType, error)
}

type ServiceTypeService interface {
	CreateServiceType(ctx context.Context, serviceType *model.ServiceType) (*model.ServiceType, error)
}
