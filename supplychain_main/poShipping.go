package main

//==============================================================================================================================
//	Order - Defines the structure for a Purchase order object.
//==============================================================================================================================
type PoShipping struct {
	Id             string `json:"id"`
	Items          Items  `json:"items"`
	QuantityRcvd   int    `json:"quantityrcvd"`
	QuantityBal    string `json:"quantitybal"`
	QuantityRej    string `json:"quantityrej"`
	InvCode        string `json:"invcode"`
	SubInv         string `json:"subinv"`
	ShippingStatus string `json:"shippingstatus"`
	PoShippingDate string `json:"poreceiptdate"`
}

/*type PurchaseOrders struct {
	PurchaseOrders []PurchaseOrder `json:"purchaseorders"`
}*/

type PoShippings struct {
	PoShippings []PoShipping `json:"poreceipts"`
}

//const PICK_STATUS_PENDING = "PENDING"
//const PICK_STATUS_PICKED = "PICKED"
//const PICK_STATUS_PARTIALLY_PICKED = "PARTIALLY_PICKED"

const SHIPPING_STATUS_AWAITING_SHIPPING = "AWAITING_SHIPPING"
const SHIPPING_STATUS_ENROUTE = "ENROUTE"
const SHIPPING_STATUS_RECEIVED = "RECEIVED"
const SHIPPING_STATUS_PARTIALLY_RECEIVED = "PARTIALLY_RECEIVED"
const SHIPPING_STATUS_FAILURE = "FAILURE"
const SHIPPING_STATUS_REJECTED = "REJECTED"

func NewPoShipping(id string, quantityRcvd int, quantityBal string, quantityRej string, invCode string, subInv string,
	shippingStatus string, poShippingDate string, recipient string, address string, sourceLocation string, deliveryCompany string, items Items, client string, owner string, timestamp string) PoShipping {
	var poShipping PoShipping

	poShipping.Id = id
	poShipping.Items = items
	poShipping.QuantityRcvd = quantityRcvd
	poShipping.QuantityBal = quantityBal
	poShipping.QuantityRej = quantityRej
	poShipping.InvCode = invCode
	poShipping.SubInv = subInv
	poShipping.ShippingStatus = shippingStatus
	poShipping.PoShippingDate = poShippingDate
	return poShipping
}
