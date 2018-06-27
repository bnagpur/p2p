package main

//==============================================================================================================================
//	OrderHistory - Defines the structure for an order history object.
//==============================================================================================================================
type POReceiptHistory struct {
	POReceiptUpdates []POReceiptUpdate `json:"orderUpdates"`
}

type POReceiptUpdate struct {
	UpdateType string `json:"updateType"`
	FromValue  string `json:"fromValue"`
	ToValue    string `json:"toValue"`
	Comment    string `json:"comment"`
	Updater    string `json:"updater"`
	Timestamp  string `json:"timestamp"`
}

//const UPDATE_TYPE_SOURCE_STATUS = "SOURCE_STATUS"
//const UPDATE_TYPE_TRANSPORT_STATUS = "TRANSPORT_STATUS"

func NewPOReceiptUpdate(updateType string, fromValue string, toValue string, comment string, updater string, timestamp string) POReceiptUpdate {
	var poReceiptUpdate POReceiptUpdate

	poReceiptUpdate.UpdateType = updateType
	poReceiptUpdate.FromValue = fromValue
	poReceiptUpdate.ToValue = toValue
	poReceiptUpdate.Comment = comment
	poReceiptUpdate.Updater = updater
	poReceiptUpdate.Timestamp = timestamp

	return poReceiptUpdate
}
