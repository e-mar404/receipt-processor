package main

import (
  "e-mar404/receipt-processor/pkg"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  receipts := router.Group("/receipts") 
  {
    receipts.POST("/process", handlers.ProcessReceipt)
    receipts.GET("/:id/points", handlers.CountReceiptPoints)
  }

  router.Run(":8080")
}
