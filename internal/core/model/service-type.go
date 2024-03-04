package model

type ServiceTypeStatus string

const (
	Active   ServiceTypeStatus = "active"
	Inactive ServiceTypeStatus = "inactive"
)

type ServiceType struct {
	Id            string
	Name          string
	Code          string
	Status        ServiceTypeStatus
	InstitutionId string
}
