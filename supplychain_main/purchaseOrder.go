package main

//==============================================================================================================================
//	Order - Defines the structure for a Purchase order object.
//==============================================================================================================================
type PurchaseOrder struct {
	/*Id          string           `json:"id"`
	Vendors     Vendor           `json:"vendor"`
	Destination Destination      `json:"destination"`
	Logistic    Logistic         `json:"logistic"`
	Items       []Item           `json:"items"`
	Quantity    int              `json:"quantity"`
	Details     Details          `json:"details"`
	ContractNo  string           `json:"contractNo"`
	RefNo       string           `json:"RefNo"`
	Amount      string           `json:"amount"`
	Currency    string           `json:"currency"`
	Weight      string           `json:"Weight"`
	POReceipts  POReceipts       `json:"poreceipts"`
	PoTerms     PoTerms          `json:"poterms"`
	PoGRN       GoodsReceiptNote `json:"pogrn"`
	POStatus    string           `json:"postatus"`*/
	Id     string `json:"id"`
	Vendor string `json:"vendor"`
	//Destination Destination      `json:"destination"`
	//Logistic    Logistic         `json:"logistic"`
	Item     string `json:"items"`
	Quantity int    `json:"quantity"`
	//Details     Details          `json:"details"`
	//ContractNo  string           `json:"contractNo"`
	//RefNo       string           `json:"RefNo"`
	Amount string `json:"amount"`
	//Currency    string           `json:"currency"`
	//Weight      string           `json:"Weight"`
	//POReceipts  POReceipts       `json:"poreceipts"`
	//PoTerms     PoTerms          `json:"poterms"`
	//PoGRN       GoodsReceiptNote `json:"pogrn"`
	POStatus      string `json:"postatus"`
	POCreatedTime string `json:"poCreatedTime"`
}
type PoTerms struct {
	TermsOfTrade        string `json:"TermsOfTrade"`
	TermsOfInsurance    string `json:"TermsOfInsurance"`
	TermsOfPayment      string `json:"TermsOfPayment"`
	PackingMethod       string `json:"PackingMethod"`
	WayOfTransportation string `json:"WayOfTransportation"`
	TimeOfShipment      string `json:"TimeOfShipment"`
	PortOfShipment      string `json:"PortOfShipment"`
	PortOfDischarge     string `json:"PortOfDischarge"`
}

/*type POReceipt struct {
	Id				string		`json:"id"`
	Items			item[]		`json:"item"`
	QuantityRcvd	int			`json:"quantityrcvd"`
	QuantityBal		string		`json:"quantitybal"`
	QuantityRej		string		`json:"quantityrej"`
}*/

/*type GoodsReceiptNote struct {
	ReceiptID 		POReceipt.Id	`json:"receiptid"`
	GoodAccepted	bool			`json:"goodsaccepted"`
	GoodsRetured	bool			`json:"goodsreturned"`
	QuantityRcvd	int				`json:"quantityrcvd"`
	QuantityBal		string			`json:"quantitybal"`
	QuantityRej		string			`json:"quantityrej"`
	Details			Details			`json:"details"`
	ReceiptStaus	string			`json:"receiptstatus"`
}*/

type POProcesTrack struct {
	ProcessStatus       string `json:"ProcessStatus"`
	POInitialCreateTime string `json:"POInitialCreateTime"`
	UpdateTime          string `json:"UpdateTime "`
	POCreatedTime       string `json:"POCreatedTime"`
	POSubmittedTime     string `json:"POSubmittedTime"`
	CompanyIdOfExporter string `json:"CompanyIdOfExporter"`
	CompanyIdOfImporter string `json:"CompanyIdOfImporter"`
	PORejectReason      string `json:"PORejectReason"`
}

type PurchaseOrders struct {
	PurchaseOrders []PurchaseOrder `json:"purchaseorders"`
}

/*type POReceipt struct {
	POReceipts []POReceipt `json:"poreceipts"`
}*/
type Source struct {
	Type     string `json:"type"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

type Vendor struct {
	VendorID string `json:"vendorid"`
	VName    string `json:"vname"`
	VAddress string `json:"vaddress"`
}

type Vendors struct {
	Vendors []Vendor `json:"vendors"`
}

type Destination struct {
	Recipient string `json:"recipient"`
	Address   string `json:"address"`
}

type Logistic struct {
	Company string `json:"company"`
	Status  string `json:"status"`
}

type Details struct {
	Client    string `json:"client"`
	Owner     string `json:"owner"`
	Timestamp string `json:"timestamp"`
}

type Invoice struct {
	Id string `json:"invoiceid"`
	//POId          PurchaseOrder.Id `json:"poid"`
	Customer      string  `json:"customer"`
	InvoiceAmount float64 `json:"invoiceamount"`
	InvoiceStatus string  `json:"invoicestatus"`
}

//const PICK_STATUS_PENDING = "PENDING"
//const PICK_STATUS_PICKED = "PICKED"
//const PICK_STATUS_PARTIALLY_PICKED = "PARTIALLY_PICKED"
const PO_STATUS_CREATED = "CREATED"
const PO_STATUS_OPEN = "OPEN"
const PO_STATUS_CANCELLED = "CANCELLED"
const PO_STATUS_OPEN_FOR_RECEIVING = "OPEN_FOR_RECEIVING"
const PO_STATUS_RECEIVED = "RECEIVED"
const PO_STATUS_CLOSED_FOR_RECEIVING = "CLOSED_FOR_RECEIVING"
const PO_STATUS_OPEN_FOR_INVOICE = "OPEN_FOR_INVOICE"
const PO_STATUS_CLOSED_FOR_INVOICE = "CLOSED_FOR_INVOICE"
const PO_STATUS_INVOICE_COMPLETED = "INVOICE_COMPLETED"
const PO_STATUS_ACCPETED = "ACCEPTED"
const PO_STATUS_CLOSED = "CLOSED"

func NewPurchaseOrder(id string, vendor string, item string, quantity int, amount string, postatus string, poCreatedTime string) PurchaseOrder {
	var purchaseOrder PurchaseOrder

	purchaseOrder.Id = id
	purchaseOrder.Vendor = vendor
	purchaseOrder.Item = item
	purchaseOrder.Quantity = quantity
	purchaseOrder.Amount = amount
	purchaseOrder.POStatus = postatus
	purchaseOrder.POCreatedTime = poCreatedTime

	return purchaseOrder
}
