package routes

import (
	"net/http"
	"strconv"

	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type EndPointRoutes struct {
	endpointService *service.EndpointService
}

func NewEndpointRoutes(endpointService *service.EndpointService) *EndPointRoutes {
	return &EndPointRoutes{endpointService: endpointService}
}

func (endpoint *EndPointRoutes) Setup(router *gin.RouterGroup) {
	endpoints := router.Group("/endpoints")
	{
		endpoints.GET("/:id", endpoint.getEndpointByID)
		endpoints.GET("/:id/parameters", endpoint.getAllParametersByEndpointID)
		endpoints.POST("/:id/parameters", endpoint.addParamToApi)
	}
}

func (h *EndPointRoutes) getEndpointByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endpoint, err := h.endpointService.GetEndpointByID(uint(id))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	if endpoint == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Endpoint does not exist for this id"})
	}
	c.IndentedJSON(http.StatusOK, endpoint)

}

func (h *EndPointRoutes) addParamToApi(c *gin.Context) {
	// parse endpoint id fields
	endpointID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parameter fields
	var parameter *models.Parameter
	if err := c.ShouldBindBodyWithJSON(&parameter); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// make the acall to add endppoint
	if err := h.endpointService.AddParameterToEndpoint(uint(endpointID), parameter); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully added parameters to the endpoint"})
}

func (h *EndPointRoutes) getAllParametersByEndpointID(c *gin.Context) {
	// parse the endpointif
	endpointID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parameters, err := h.endpointService.GetAllParametersForEndpoint(uint(endpointID))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(parameters) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Parameters are yet to be created for this endpoint"})
		return
	}
	c.IndentedJSON(http.StatusOK, parameters)
}
