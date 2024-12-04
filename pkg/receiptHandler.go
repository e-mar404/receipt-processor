package handlers 

import (
	"net/http"

  "github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func ProcessReceipt (c *gin.Context) {
  var newReceipt Receipt

  if err := c.ShouldBindJSON(&newReceipt); err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "description": "The receipt is invalid",
    })

    return 
  }

  // TODO: save the receipt somewhere

  if newReceipt.ID == "" {
    newReceipt.ID = uuid.New().String()
  }

  c.JSON(http.StatusOK, gin.H {
    "id" : newReceipt.ID,
  })
}

func CountReceiptPoints (c *gin.Context) {
  c.JSON(http.StatusOK, gin.H {
    "points": 0,
  })
}
