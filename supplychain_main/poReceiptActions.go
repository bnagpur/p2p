package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ALL_PURCHASE_RECEIPT_IDS_KEY = "PO_RECEIPTS"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================

func AddPOReceipt(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poreceipt POReceipt, purchaseOrderid string) error {
	//TODO Validate order

	if callerDetails.Role != ROLE_INTERNAL_SYSTEM {
		return LogAndError("Caller does not have permission to add an order")
	}
	//Add to state
	purchaseorder, err := GetPurchaseOrder(stub, purchaseOrderid)
	if err == nil && purchaseorder.POStatus == PO_STATUS_ACCPETED {
		//poreceipt.POId = purchaseOrderid
		poreceipt.ReceiptStatus = RECEIPT_STATUS_AWAITING_RECEIPT
		_, err := SavePOReceipt(stub, poreceipt)
		if err != nil {
			return err
		}

		err = addPOReceiptIdToHolder(stub, poreceipt.Id)
		if err != nil {
			return LogAndError(err.Error())
		}

		//Add blank order history
		poReceiptHistory := POReceiptHistory{[]POReceiptUpdate{}}
		_, err = SavePOReceiptHistory(stub, PURCHASE_ORDER_HISTORY_KEY_PREFIX+poreceipt.Id, poReceiptHistory)
	}
	return err
}

func UpdatePOReceiptStatus(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poReceiptId string, statusType string, statusValue string, comment string) error {
	//TODO Validate update (define state progression)
	//TODO permissions
	poReceipt, err := RetrievePOReceipt(stub, poReceiptId)

	if err != nil {
		return LogAndError(err.Error())
	}

	//TODO make this more extendable
	fromValue := ""
	updateType := ""

	var newStatus string

	//PurchaseOrder.POStatus

	if poReceipt.ReceiptStatus == RECEIPT_STATUS_AWAITING_RECEIPT {

		rcv := poReceipt.QuantityRcvd
		ord, err := strconv.Atoi(poReceipt.QuantityOrdered)
		if err != nil {
			return LogAndError(err.Error())
		}

		//!= 0 && strconv.Atoi(poReceipt.QuantityRcvd) < strconv.Atoi(poReceipt.QuantityOrdered) {
		if rcv != 0 && rcv < ord {
			newStatus = RECEIPT_STATUS_PARTIALLY_RECEIVED
		} else if rcv == ord {
			newStatus = RECEIPT_STATUS_RECEIVED
		}

	} else if poReceipt.ReceiptStatus == RECEIPT_STATUS_PARTIALLY_RECEIVED {
		rcv := poReceipt.QuantityRcvd
		ord, err := strconv.Atoi(poReceipt.QuantityOrdered)
		if err != nil {
			return LogAndError(err.Error())
		}

		if rcv == ord {
			newStatus = RECEIPT_STATUS_RECEIVED
		} else if rcv != ord {

			newStatus = RECEIPT_STATUS_PARTIALLY_RECEIVED
		}

	} else if poReceipt.ReceiptStatus == RECEIPT_STATUS_AWAITING_RECEIPT && poReceipt.Rejected {

		newStatus = RECEIPT_STATUS_REJECTED

	}

	stateTransitionAllowed := false

	/*
		if PurchaseOrder.ReceiptStatus == "P/O Created" && newStatus == "P/O Submitted" {
		stateTransitionAllowed = true
		} else */

	if poReceipt.ReceiptStatus == RECEIPT_STATUS_AWAITING_RECEIPT && newStatus == RECEIPT_STATUS_PARTIALLY_RECEIVED {
		stateTransitionAllowed = true
	} else if poReceipt.ReceiptStatus == RECEIPT_STATUS_AWAITING_RECEIPT && newStatus == RECEIPT_STATUS_RECEIVED {
		stateTransitionAllowed = true
	} else if poReceipt.ReceiptStatus == RECEIPT_STATUS_AWAITING_RECEIPT && newStatus == RECEIPT_STATUS_RECEIVED {
		stateTransitionAllowed = true
	} else if poReceipt.ReceiptStatus == RECEIPT_STATUS_AWAITING_RECEIPT && newStatus == RECEIPT_STATUS_REJECTED {
		stateTransitionAllowed = true
	}
	if stateTransitionAllowed == false {
		return errors.New(" This state transition is not allowed ")
	}
	//}
	//Add to state
	_, err = SavePOReceipt(stub, poReceipt)
	if err != nil {
		return LogAndError(err.Error())
	}

	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return LogAndError(err.Error())
	}

	//Update history
	/*poReceiptUpdate := NewPOReceiptUpdate(updateType, fromValue, statusValue, comment, callerDetails.Username, timestamp.String())

	return UpdatePOReceiptHistory(stub, poReceiptId, poReceiptUpdate)*/
	updateType = UPDATE_TYPE_SOURCE_STATUS
	fromValue = poReceipt.ReceiptStatus
	statusValue = newStatus
	//Update history
	poReceiptUpdate := NewPOReceiptUpdate(updateType, fromValue, statusValue, "", callerDetails.Username, timestamp.String())

	return UpdatePOReceiptHistory(stub, poReceiptId, poReceiptUpdate)
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPOReceipt(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poReceiptId string) (POReceipt, bool, error) {

	//TODO Has permission to retrieve?

	poReceipt, err := RetrievePOReceipt(stub, poReceiptId)

	return poReceipt, false, err
}

func GetAllPOReceipts(stub shim.ChaincodeStubInterface, callerDetails CallerDetails) (POReceipts, error) {

	//TODO Has permission to retrieve?
	poReceipts := POReceipts{}

	poReceiptIds, err := RetrieveIdsHolder(stub, ALL_PURCHASE_ORDER_IDS_KEY)

	if err != nil {
		return poReceipts, LogAndError("Unable to retrieve order id holder")
	}

	for _, poReceiptId := range poReceiptIds.Ids {
		poReceipt, accessDenied, err := GetPOReceipt(stub, callerDetails, poReceiptId)

		if accessDenied {
			fmt.Println("Access denied when reading order: " + poReceiptId)
		} else if err != nil {
			return poReceipts, LogAndError("There was an error when retrieving order: " + err.Error())
		} else {
			poReceipts.POReceipts = append(poReceipts.POReceipts, poReceipt)
		}
	}

	return poReceipts, err
}

//==============================================================================================================================
//	 Internal
//==============================================================================================================================

/*func addOrderIdToHolder(stub shim.ChaincodeStubInterface, orderId string) (error) {
    idHolder, err := RetrieveIdsHolder(stub, ALL_ORDER_IDS_KEY)

    if err != nil {
        fmt.Println("Unable to retrieve id holder so this is probably the first order...adding")

        idHolder = IdsHolder{}
    }

    idHolder.Ids = append(idHolder.Ids, orderId)

    _, err = SaveIdsHolder(stub, ALL_ORDER_IDS_KEY, idHolder)

    return err
}*/

func addPOReceiptIdToHolder(stub shim.ChaincodeStubInterface, poReceiptId string) error {
	idHolder, err := RetrieveIdsHolder(stub, ALL_PURCHASE_RECEIPT_IDS_KEY)

	if err != nil {
		fmt.Println("Unable to retrieve id holder so this is probably the first Purchase order...adding")

		idHolder = IdsHolder{}
	}

	idHolder.Ids = append(idHolder.Ids, poReceiptId)

	_, err = SaveIdsHolder(stub, ALL_PURCHASE_RECEIPT_IDS_KEY, idHolder)

	return err
}
