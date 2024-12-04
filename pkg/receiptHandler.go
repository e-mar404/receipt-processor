package handlers

import (
	"net/http"
  "e-mar404/receipt-processor/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt (c *gin.Context) {
  var newReceipt db.Receipt

  if err := c.ShouldBindJSON(&newReceipt); err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "description": "The receipt is invalid",
    })

    return 
  }

  newReceipt.ID = uuid.New().String()

  // in a real case there would have to be error handling here but since it is just needing to save in memory it is fine for now
  db.Save(&newReceipt)

  c.JSON(http.StatusOK, gin.H {
    "id" : newReceipt.ID,
  })
}

func CountReceiptPoints (c *gin.Context) {
  c.IndentedJSON(http.StatusOK, db.GetList())
}
