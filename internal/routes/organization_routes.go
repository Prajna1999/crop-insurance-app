package routes

import (
	"net/http"
	"strconv"

	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type OrganizationRoutes struct {
	orgService *service.OrganizationService
}

func NewOrganizationRoutes(orgService *service.OrganizationService) *OrganizationRoutes {
	return &OrganizationRoutes{
		orgService: orgService,
	}
}

// imported to another package
func (or *OrganizationRoutes) Setup(router *gin.RouterGroup) {
	orgs := router.Group("/organizations")
	{
		orgs.POST("", or.createOrganization)
		orgs.GET("/:id", or.getOrganization)
		orgs.GET("", or.getAllOrganizations)
	}
}

func (or *OrganizationRoutes) createOrganization(c *gin.Context) {
	var input struct {
		OrgName       string `json:"org_name" binding:"required"`
		AdminEmail    string `json:"admin_email" binding:"required"`
		AdminUserName string `json:"admin_user_name"  binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	err := or.orgService.CreateOrganization(input.OrgName, input.AdminEmail, input.AdminUserName)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Organization created Successfully"})

}

func (or *OrganizationRoutes) getOrganization(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	org, err := or.orgService.GetOrganization(uint(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	c.IndentedJSON(http.StatusFound, org)

}

// get all organizations
func (or *OrganizationRoutes) getAllOrganizations(c *gin.Context) {
	orgs, err := or.orgService.GetAllOrganizations()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// jsonify the resulting arrays
	c.IndentedJSON(http.StatusFound, orgs)
}
