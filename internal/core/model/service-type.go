package model

import "github.com/google/uuid"

type ServiceTypeStatus string

const (
	Active   ServiceTypeStatus = "active"
	Inactive ServiceTypeStatus = "inactive"
)

type ServiceType struct {
	Id            uuid.UUID
	Name          string
	Code          string
	Status        ServiceTypeStatus
	InstitutionId string
}
