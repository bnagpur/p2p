package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const PO_INVOICE_HISTORY_KEY_PREFIX = "HIST"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================
func UpdatePoInvoiceHistory(stub shim.ChaincodeStubInterface, poInvoiceId string, update PoInvoiceUpdate) error {
	invoiceHistoryId := PO_INVOICE_HISTORY_KEY_PREFIX + poInvoiceId
	_, err := RetrieveInvoiceHistory(stub, invoiceHistoryId)

	if err != nil {
		return LogAndError("Unable to retrieve invoice history: " + err.Error())
	}
	poInvoiceHistory := PoInvoiceHistory{[]PoInvoiceUpdate{}}

	poInvoiceHistory.PoInvoiceUpdates = append(poInvoiceHistory.PoInvoiceUpdates, update)

	_, err = SaveInvoiceHistory(stub, invoiceHistoryId, poInvoiceHistory)
	return err
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPoInvoiceHistory(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poInvoiceId string) (PoInvoiceHistory, bool, error) {

	//TODO Has permission to retrieve?

	poInvoiceHistory, err := RetrieveInvoiceHistory(stub, PO_INVOICE_HISTORY_KEY_PREFIX+poInvoiceId)

	return poInvoiceHistory, false, err
}
