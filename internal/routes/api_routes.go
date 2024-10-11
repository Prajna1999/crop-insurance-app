package routes

import (
	"net/http"
	"strconv"

	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
	apiService *service.ApiService
}

func NewApiRoutes(apiService *service.ApiService) *ApiRoutes {
	return &ApiRoutes{
		apiService: apiService,
	}
}

func (api *ApiRoutes) Setup(router *gin.RouterGroup) {
	apis := router.Group("/apis")
	{
		apis.GET("", api.getAllApis)
		apis.POST("", api.createApi)
		apis.POST("/:id/endpoints", api.addEndpointToApi)
	}
}

// create api with details

func (api *ApiRoutes) createApi(c *gin.Context) {
	var input *models.Api
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.apiService.CreateApi(input); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully created API"})
}

// get apis by queryparams
func (api *ApiRoutes) getAllApis(c *gin.Context) {
	filters := make(map[string]interface{})

	// add to filters based on queryparams
	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}
	if version := c.Query("version"); version != "" {
		filters["version"] = version
	}
	if isPublic := c.Query("is_public"); isPublic != "" {
		publicBool, _ := strconv.ParseBool(isPublic)
		filters["is_public"] = publicBool
	}
	if categoryID := c.Query("category_id"); categoryID != "" {
		categoryIDUint, _ := strconv.ParseUint(categoryID, 10, 32)
		filters["category_id"] = uint(categoryIDUint)
	}
	if orgID := c.Query("organization_id"); orgID != "" {
		orgIDUint, _ := strconv.ParseUint(orgID, 10, 32)
		filters["organization_id"] = uint(orgIDUint)
	}
	apis, err := api.apiService.GetAllApis(filters)
	if len(apis) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "No APIs found for this filter combination"})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apis)
}

// add endpoints to existing api
func (api *ApiRoutes) addEndpointToApi(c *gin.Context) {
	// parse apiId fields
	apiID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// endpoint fields
	var endpoint *models.EndPoint
	if err := c.ShouldBindBodyWithJSON(&endpoint); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// make the call to the endpoint service
	if err := api.apiService.AddEndpointToApi(uint(apiID), endpoint); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully created endpoints"})

}
