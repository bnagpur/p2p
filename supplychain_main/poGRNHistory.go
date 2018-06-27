package main

//==============================================================================================================================
//	OrderHistory - Defines the structure for an order history object.
//==============================================================================================================================
type GoodsReceiptNoteHistory struct {
	GoodsReceiptNoteUpdates []GoodsReceiptNoteUpdate `json:"gRNUpdates"`
}

type GoodsReceiptNoteUpdate struct {
	UpdateType string `json:"updateType"`
	FromValue  string `json:"fromValue"`
	ToValue    string `json:"toValue"`
	Comment    string `json:"comment"`
	Updater    string `json:"updater"`
	Timestamp  string `json:"timestamp"`
}

const UPDATE_TYPE_SOURCE_STATUS = "SOURCE_STATUS"
const UPDATE_TYPE_TRANSPORT_STATUS = "TRANSPORT_STATUS"

func NewGoodsReceiptNoteUpdate(updateType string, fromValue string, toValue string, comment string, updater string, timestamp string) GoodsReceiptNoteUpdate {
	var goodsReceiptNoteUpdate GoodsReceiptNoteUpdate

	goodsReceiptNoteUpdate.UpdateType = updateType
	goodsReceiptNoteUpdate.FromValue = fromValue
	goodsReceiptNoteUpdate.ToValue = toValue
	goodsReceiptNoteUpdate.Comment = comment
	goodsReceiptNoteUpdate.Updater = updater
	goodsReceiptNoteUpdate.Timestamp = timestamp

	return goodsReceiptNoteUpdate
}
