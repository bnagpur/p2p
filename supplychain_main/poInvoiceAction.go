package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ALL_INVOICE_IDS_KEY = "INVOICE"

//const STATUS_TYPE_SOURCE = "SOURCE"

//const STATUS_TYPE_TRANSPORT = "TRANSPORT"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================

func AddInvoice(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, invoice PoInvoice) error {
	//TODO Validate order

	if callerDetails.Role != ROLE_INTERNAL_SYSTEM {
		return LogAndError("Caller does not have permission to add an order")
	}

	//Add to state
	_, err := SaveInvoice(stub, invoice.Id, invoice)
	if err != nil {
		return err
	}

	err = addInvoiceIdToHolder(stub, invoice.Id)
	if err != nil {
		return LogAndError(err.Error())
	}

	//Add blank order history
	poInvoiceHistory := PoInvoiceHistory{[]PoInvoiceUpdate{}}
	_, err = SaveInvoiceHistory(stub, GRN_HISTORY_KEY_PREFIX+invoice.Id, poInvoiceHistory)

	return err
}

func UpdateInvoiceStatus(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, invoiceId string, statusType string, statusValue string, comment string) error {
	//TODO Validate update (define state progression)
	//TODO permissions
	invoice, err := RetrieveInvoice(stub, invoiceId)

	if err != nil {
		return LogAndError(err.Error())
	}

	//TODO make this more extendable
	fromValue := ""
	updateType := ""
	/*  if statusType == STATUS_TYPE_SOURCE {
	        updateType = UPDATE_TYPE_SOURCE_STATUS
	        fromValue = order.Source.Status
	        order.Source.Status = statusValue
	    } else if statusType == STATUS_TYPE_TRANSPORT {
	        updateType = UPDATE_TYPE_TRANSPORT_STATUS
	        fromValue = order.Transport.Status
	        order.Transport.Status = statusValue
	    }
	*/
	//Add to state
	_, err = SaveInvoice(stub, invoice.Id, invoice)
	if err != nil {
		return LogAndError(err.Error())
	}

	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return LogAndError(err.Error())
	}

	//Update history
	invoiceUpdate := NewPoInvoiceUpdate(updateType, fromValue, statusValue, comment, callerDetails.Username, timestamp.String())

	return UpdatePoInvoiceHistory(stub, invoice.Id, invoiceUpdate)
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetInvoice(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, invoiceId string) (PoInvoice, bool, error) {

	//TODO Has permission to retrieve?

	invoice, err := RetrieveInvoice(stub, invoiceId)

	return invoice, false, err
}

func GetAllInvoices(stub shim.ChaincodeStubInterface, callerDetails CallerDetails) (PoInvoices, error) {

	//TODO Has permission to retrieve?
	invoices := PoInvoices{}

	invoiceIds, err := RetrieveIdsHolder(stub, ALL_GRN_IDS_KEY)

	if err != nil {
		return invoices, LogAndError("Unable to retrieve order id holder")
	}

	for _, invoiceId := range invoiceIds.Ids {
		invoice, accessDenied, err := GetInvoice(stub, callerDetails, invoiceId)

		if accessDenied {
			fmt.Println("Access denied when reading order: " + invoiceId)
		} else if err != nil {
			return invoices, LogAndError("There was an error when retrieving order: " + err.Error())
		} else {
			invoices.PoInvoices = append(invoices.PoInvoices, invoice)
		}
	}

	return invoices, err
}

//==============================================================================================================================
//	 Internal
//==============================================================================================================================
func addInvoiceIdToHolder(stub shim.ChaincodeStubInterface, invoiceId string) error {
	idHolder, err := RetrieveIdsHolder(stub, ALL_GRN_IDS_KEY)

	if err != nil {
		fmt.Println("Unable to retrieve id holder so this is probably the first Purchase order...adding")

		idHolder = IdsHolder{}
	}

	idHolder.Ids = append(idHolder.Ids, invoiceId)

	_, err = SaveIdsHolder(stub, ALL_GRN_IDS_KEY, idHolder)

	return err
}
