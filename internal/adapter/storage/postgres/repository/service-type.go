package repository

import (
	"context"

	"github.com/davifrjose/My_Turn/internal/adapter/storage/postgres"
	"github.com/davifrjose/My_Turn/internal/core/model"
)

type ServiceTypeRepository struct {
	db *postgres.Db
}

func NewServiceTypeRepository(db *postgres.Db) *ServiceTypeRepository {
	return &ServiceTypeRepository{
		db,
	}
}

const createServiceTypeQuery = `INSERT INTO service_types 
(id, name, code, status, institution_id)
 VALUES ($1, $2, $3, $4, $5)
 RETURNING *
 `

func (serviceTypeRepository *ServiceTypeRepository) CreateServiceType(ctx context.Context, serviceType *model.ServiceType) (*model.ServiceType, error) {
	connection, error := serviceTypeRepository.db.Pool.Acquire(ctx)
	if error != nil {
		return nil, error
	}

	err := connection.QueryRow(ctx,
		createServiceTypeQuery,
		serviceType.Id,
		serviceType.Name,
		serviceType.Code,
		serviceType.Status,
		serviceType.InstitutionId).Scan(
		&serviceType.Id,
		&serviceType.Name,
		&serviceType.Code,
		&serviceType.Status,
		&serviceType.InstitutionId,
	)
	if err != nil {
		return nil, err
	}

	return serviceType, error
}
