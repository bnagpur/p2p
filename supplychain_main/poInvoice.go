package main

//==============================================================================================================================
//	Order - Defines the structure for a Purchase order object.
//==============================================================================================================================
type PoInvoice struct {
	Id string `json:"invoiceid"`
	//POId          PurchaseOrder.Id `json:"poid"`
	Customer         string  `json:"customer"`
	PoInvoiceAmount  float64 `json:"invoiceamount"`
	PoInvoiceStatus  string  `json:"invoicestatus"`
	PoInvoiceBalance string  `json:"invoicebalance"`
	Vendor           Vendor  `json:"vendor"`
}
type PoInvoices struct {
	PoInvoices []PoInvoice `json:"invoices"`
}

//const PICK_STATUS_PENDING = "PENDING"
//const PICK_STATUS_PICKED = "PICKED"
//const PICK_STATUS_PARTIALLY_PICKED = "PARTIALLY_PICKED"

const INVOICE_STATUS_PAID = "PAID"
const INVOICE_STATUS_UNPAID = "UNPAID"
const INVOICE_STATUS_HOLD = "HOLD"
const INVOICE_STATUS_PARTIALLY_RECEIVED = "PARTIALLY_RECEIVED"
const INVOICE_STATUS_FAILURE = "FAILURE"
const INVOICE_STATUS_REJECTED = "REJECTED"

func NewPoInvoice(id string, customer string, poInvoiceAmount float64, poInvoiceStatus string, poInvoiceBalance string, vendor Vendor) PoInvoice {
	var poInvoice PoInvoice

	/*poInvoice.Id = id
	//Hard code to warehouse source for now
	poInvoice.Source = Source{"WAREHOUSE", sourceLocation, PICK_STATUS_PENDING}
	poInvoice.Destination = Destination{recipient, address}
	//order.Transport = Transport {deliveryCompany, DELIVERY_STATUS_AWAITING_PICKUP}
	poInvoice.Items = items.Items
	poInvoice.Details = Details{client, owner, timestamp}
	*/
	poInvoice.Id = id
	poInvoice.Customer = customer
	poInvoice.PoInvoiceAmount = poInvoiceAmount
	poInvoice.PoInvoiceStatus = poInvoiceStatus
	poInvoice.PoInvoiceBalance = poInvoiceBalance
	poInvoice.Vendor = vendor

	return poInvoice
}
