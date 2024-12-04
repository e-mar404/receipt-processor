package handlers

import (
  "net/http"
	"e-mar404/receipt-processor/db"
	"e-mar404/receipt-processor/models"
	"e-mar404/receipt-processor/rules"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt (c *gin.Context) {
  var newReceipt models.Receipt

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
  receiptId := c.Param("id")

  receipt, err := db.GetReceipt(receiptId)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H {
      "description": err.Error(),
    })

    return
  }

  c.JSON(http.StatusOK, gin.H {
    "price": rules.CountPoints(&receipt),
  })
}
