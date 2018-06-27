package main

//==============================================================================================================================
//	OrderHistory - Defines the structure for an order history object.
//==============================================================================================================================
type PurchaseOrderHistory struct {
	PurchaseOrderUpdates []PurchaseOrderUpdate `json:"purchaseorderUpdates"`
}

type PurchaseOrderUpdate struct {
	UpdateType string `json:"updateType"`
	FromValue  string `json:"fromValue"`
	ToValue    string `json:"toValue"`
	Comment    string `json:"comment"`
	Updater    string `json:"updater"`
	Timestamp  string `json:"timestamp"`
}

//const UPDATE_TYPE_SOURCE_STATUS = "SOURCE_STATUS"
//const UPDATE_TYPE_TRANSPORT_STATUS = "TRANSPORT_STATUS"

func NewPurchaseOrderUpdate(updateType string, fromValue string, toValue string, comment string, updater string, timestamp string) PurchaseOrderUpdate {
	var purchaseOrderUpdate PurchaseOrderUpdate

	purchaseOrderUpdate.UpdateType = updateType
	purchaseOrderUpdate.FromValue = fromValue
	purchaseOrderUpdate.ToValue = toValue
	purchaseOrderUpdate.Comment = comment
	purchaseOrderUpdate.Updater = updater
	purchaseOrderUpdate.Timestamp = timestamp

	return purchaseOrderUpdate
}
