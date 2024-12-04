package handlers 

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt (c *gin.Context) {
  c.JSON(http.StatusOK, gin.H {
    "id" : 0,
  })
}

func CountReceiptPoints (c *gin.Context) {
  c.JSON(http.StatusOK, gin.H {
    "points": 0,
  })
}
