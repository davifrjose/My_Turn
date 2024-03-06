package http

import (
	"github.com/davifrjose/My_Turn/internal/core/model"
	"github.com/davifrjose/My_Turn/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceTypeHandler struct {
	serviceTypeService port.ServiceTypeService
}

func NewServiceTypeHandler(serviceTypeService port.ServiceTypeService) *ServiceTypeHandler {
	return &ServiceTypeHandler{
		serviceTypeService,
	}
}

type createServiceTypeRequest struct {
	Name          string                  `json:"name" binding:"required"`
	Code          string                  `json:"code" binding:"required"`
	Status        model.ServiceTypeStatus `json:"status"`
	InstitutionId string                  `json:"institutionId" binding:"required"`
}

type serviceTypeResponse struct {
	Id            uuid.UUID               `json:"id"`
	Name          string                  `json:"name"`
	Code          string                  `json:"code"`
	Status        model.ServiceTypeStatus `json:"status"`
	InstitutionId string                  `json:"institutionId"`
}

func newServiceTypeResponse(response *model.ServiceType) serviceTypeResponse {
	return serviceTypeResponse{
		Id:            response.Id,
		Name:          response.Name,
		Code:          response.Code,
		Status:        response.Status,
		InstitutionId: response.InstitutionId,
	}
}

func (serviceTypeHandler *ServiceTypeHandler) CreateServiceType(ctx *gin.Context) {
	var request createServiceTypeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}

	serviceType, err := serviceTypeHandler.serviceTypeService.CreateServiceType(ctx, &port.CreateServiceTypeParams{
		Name:          request.Name,
		Code:          request.Code,
		Status:        request.Status,
		InstitutionId: request.InstitutionId,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := newServiceTypeResponse(serviceType)

	handleSuccess(ctx, response)
}
