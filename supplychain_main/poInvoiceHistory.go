package main

//==============================================================================================================================
//	OrderHistory - Defines the structure for an order history object.
//==============================================================================================================================
type PoInvoiceHistory struct {
	PoInvoiceUpdates []PoInvoiceUpdate `json:"orderUpdates"`
}

type PoInvoiceUpdate struct {
	UpdateType string `json:"updateType"`
	FromValue  string `json:"fromValue"`
	ToValue    string `json:"toValue"`
	Comment    string `json:"comment"`
	Updater    string `json:"updater"`
	Timestamp  string `json:"timestamp"`
}

//const UPDATE_TYPE_SOURCE_STATUS = "SOURCE_STATUS"
//const UPDATE_TYPE_TRANSPORT_STATUS = "TRANSPORT_STATUS"

func NewPoInvoiceUpdate(updateType string, fromValue string, toValue string, comment string, updater string, timestamp string) PoInvoiceUpdate {
	var poInvoiceUpdate PoInvoiceUpdate

	poInvoiceUpdate.UpdateType = updateType
	poInvoiceUpdate.FromValue = fromValue
	poInvoiceUpdate.ToValue = toValue
	poInvoiceUpdate.Comment = comment
	poInvoiceUpdate.Updater = updater
	poInvoiceUpdate.Timestamp = timestamp

	return poInvoiceUpdate
}
