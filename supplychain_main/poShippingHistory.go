package main

//==============================================================================================================================
//	OrderHistory - Defines the structure for an order history object.
//==============================================================================================================================
type PoShippingHistory struct {
	PoShippingUpdates []PoShippingUpdate `json:"orderUpdates"`
}

type PoShippingUpdate struct {
	UpdateType string `json:"updateType"`
	FromValue  string `json:"fromValue"`
	ToValue    string `json:"toValue"`
	Comment    string `json:"comment"`
	Updater    string `json:"updater"`
	Timestamp  string `json:"timestamp"`
}

func NewPoShippingUpdate(updateType string, fromValue string, toValue string, comment string, updater string, timestamp string) PoShippingUpdate {
	var poShippingUpdate PoShippingUpdate

	poShippingUpdate.UpdateType = updateType
	poShippingUpdate.FromValue = fromValue
	poShippingUpdate.ToValue = toValue
	poShippingUpdate.Comment = comment
	poShippingUpdate.Updater = updater
	poShippingUpdate.Timestamp = timestamp

	return poShippingUpdate
}
