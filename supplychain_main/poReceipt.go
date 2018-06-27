package main

//==============================================================================================================================
//	Order - Defines the structure for a Purchase order object.
//==============================================================================================================================
type POReceipt struct {
	Id              string `json:"id"`
	QuantityRcvd    int    `json:"quantityrcvd"`
	QuantityBal     string `json:"quantitybal"`
	QuantityRej     string `json:"quantityrej"`
	QuantityOrdered string `json: "quantityordered"`
	//InvCode			Inventory.Code			`json:"invcode"`
	//SubInv			Inventory.SubInventory	`json:"subinv"`
	ReceiptStatus string `json:"receiptstatus"`
	Rejected      bool   `json:"rejected"`
	POReceiptDate string `json:"poreceiptdat"`
	Items         Items  `json:"items"`
}

/*type GoodReceiptNote struct {
	GRNId			string			`json:"grnid"`
	ReceiptID 		POReceipt.Id	`json:receiptid`
	GoodAccepted	bool			`json:goodsaccepted`
	GoodsRetured	bool			`json:goodsreturned`
	QuantityRcvd	int				`json:"quantityrcvd"`
	QuantityBal		string			`json:"quantitybal"`
	QuantityRej		string			`json:"quantityrej"`
	Details			Details			`json:"details"`
	GRNDate			string			`json:"grndate"`
}
*/
/*type PurchaseOrders struct {
	PurchaseOrders []PurchaseOrder `json:"purchaseorders"`
}*/

type POReceipts struct {
	POReceipts []POReceipt `json:"poreceipts"`
}

/*type Source struct {
	Type     string `json:"type"`
	Location string `json:"location"`
	Status   string `json:"status"`
}*/

//const PICK_STATUS_PENDING = "PENDING"
//const PICK_STATUS_PICKED = "PICKED"
//const PICK_STATUS_PARTIALLY_PICKED = "PARTIALLY_PICKED"

const RECEIPT_STATUS_AWAITING_RECEIPT = "AWAITING_RECEIPT"
const RECEIPT_STATUS_ENROUTE = "ENROUTE"
const RECEIPT_STATUS_RECEIVED = "RECEIVED"
const RECEIPT_STATUS_PARTIALLY_RECEIVED = "PARTIALLY_RECEIVED"
const RECEIPT_STATUS_FAILURE = "FAILURE"
const RECEIPT_STATUS_REJECTED = "REJECTED"

func NewPOReceipt(id string, quantityRcvd int, quantityBal string, quantityRej string, quantityOrdered string, receiptStatus string, rejected bool, poReceiptDate string, items Items) POReceipt {
	var poReceipt POReceipt

	poReceipt.Id = id
	poReceipt.QuantityRcvd = quantityRcvd
	poReceipt.QuantityBal = quantityBal
	poReceipt.QuantityRej = quantityRej
	poReceipt.QuantityOrdered = quantityOrdered
	poReceipt.ReceiptStatus = receiptStatus
	poReceipt.Rejected = rejected
	poReceipt.POReceiptDate = poReceiptDate
	poReceipt.Items = items

	return poReceipt
}
