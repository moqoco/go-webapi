package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct{}

func (h *SampleHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "sample"})
}
