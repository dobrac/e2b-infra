package handlers

import (
	"fmt"
	"net/http"

	"github.com/devbookhq/orchestration-services/api/pkg/cockroach"
	"github.com/gin-gonic/gin"
)

func (a *APIStore) PostPrismaHubDb(
	c *gin.Context,
) {
	dbURL, err := cockroach.CreateDatabase(c)
	if err != nil {
		fmt.Printf("Error when creating new Cockroach database: %s", err)
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error creating DB: %s", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"dbURL": dbURL})
}