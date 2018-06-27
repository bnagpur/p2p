package main

//==============================================================================================================================
//	Order - Defines the structure for a Purchase order object.
//==============================================================================================================================

type GoodsReceiptNote struct {
	GRNId string `json:"grnid"`
	//ReceiptID 		POReceipt.Id	`json:receiptid`
	GoodsAccepted bool    `json:goodsaccepted`
	GoodsRetured  bool    `json:goodsreturned`
	QuantityRcvd  int     `json:"quantityrcvd"`
	QuantityBal   string  `json:"quantitybal"`
	QuantityRej   string  `json:"quantityrej"`
	Details       Details `json:"details"`
	GRNDate       string  `json:"grndate"`
}

/*type PurchaseOrders struct {
	PurchaseOrders []PurchaseOrder `json:"purchaseorders"`
}*/

type GoodsReceiptNotes struct {
	GoodsReceiptNotes []GoodsReceiptNote `json:"goodsreceiptnotes"`
}

/*type Source struct {
	Type     string `json:"type"`
	Location string `json:"location"`
	Status   string `json:"status"`

}*/
//const PICK_STATUS_PENDING = "PENDING"
//const PICK_STATUS_PICKED = "PICKED"
//const PICK_STATUS_PARTIALLY_PICKED = "PARTIALLY_PICKED"

const GRN_STATUS_AWAITING_GRN = "AWAITING_GRN"
const GRN_STATUS_ENROUTE = "ENROUTE"
const GRN_STATUS_RECEIVED = "RECEIVED"
const GRN_STATUS_PARTIALLY_RECEIVED = "PARTIALLY_RECEIVED"
const GRN_STATUS_FAILURE = "FAILURE"
const GRN_STATUS_REJECTED = "REJECTED"

func NewGoodReceiptNote(gRNId string, goodsAccepted bool, goodsRetured bool, quantityRcvd int, quantityBal string, quantityRej string, details Details, gRNDate string) GoodsReceiptNote {
	var poGRN GoodsReceiptNote

	//poGRN = GoodsReceiptNote(gRNId, goodsAccepted, goodsRetured, quantityRcvd, quantityBal, quantityRej, details, gRNDate)
	poGRN.GRNId = gRNId
	//Hard code to warehouse source for now
	poGRN.GoodsAccepted = goodsAccepted
	poGRN.GoodsRetured = goodsRetured
	poGRN.QuantityRcvd = quantityRcvd
	poGRN.QuantityBal = quantityBal
	poGRN.QuantityRej = quantityRej
	poGRN.Details = details
	poGRN.GRNDate = gRNDate
	/*order.Transport = Transport {deliveryCompany, DELIVERY_STATUS_AWAITING_PICKUP}
	poGRN.Items = items.Items
	poGRN.Details = Details{client, owner, timestamp}
	*/
	return poGRN
}
