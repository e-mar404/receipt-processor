package db


var receiptList []Receipt

func Save (receipt *Receipt) {
  receiptList = append(receiptList, *receipt)
}

func GetList () []Receipt {
  return receiptList
}
