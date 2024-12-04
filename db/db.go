package db

import (
  "errors"
	"e-mar404/receipt-processor/models"
)


var receiptList []models.Receipt

func Save (receipt *models.Receipt) {
  receiptList = append(receiptList, *receipt)
}

func GetReceipt (id string) (models.Receipt, error) {
  for _, receipt := range receiptList {
    if receipt.ID == id {
      return receipt, nil
    }
  }

  return models.Receipt{}, errors.New("No receipt found for that id")
}
