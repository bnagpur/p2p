package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ALL_PURCHASE_ORDER_IDS_KEY = "PURCHASE_ORDERS"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================
func AddPurchaseOrder(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, purchaseorder PurchaseOrder) error {
	//TODO Validate order

	if callerDetails.Role != ROLE_INTERNAL_SYSTEM {
		return LogAndError("Caller does not have permission to add an order")
	}

	//Add to state

	_, err := SavePurchaseOrder(stub, purchaseorder)
	if err != nil {
		return err
	}

	err = addPurchaseOrderIdToHolder(stub, purchaseorder.Id)
	if err != nil {
		return LogAndError(err.Error())
	}

	//Add blank order history
	purchaseOrderHistory := PurchaseOrderHistory{[]PurchaseOrderUpdate{}}
	_, err = SavePurchaseOrderHistory(stub, PURCHASE_ORDER_HISTORY_KEY_PREFIX+purchaseorder.Id, purchaseOrderHistory)

	return err
}

func UpdatePurchaseOrderStatus(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, purchaseOrderId string) error {
	//TODO Validate update (define state progression)
	//TODO permissions
	purchaseOrder, err := RetrievePurchaseOrder(stub, purchaseOrderId)

	if err != nil {
		return LogAndError(err.Error())
	}

	var newStatus string

	//PurchaseOrder.POStatus

	if purchaseOrder.POStatus == PO_STATUS_CREATED {

		newStatus = PO_STATUS_ACCPETED

	} else if purchaseOrder.POStatus == PO_STATUS_ACCPETED {

		newStatus = PO_STATUS_OPEN_FOR_RECEIVING

	} else if purchaseOrder.POStatus == PO_STATUS_OPEN_FOR_RECEIVING {

		newStatus = PO_STATUS_RECEIVED

	} else if purchaseOrder.POStatus == PO_STATUS_RECEIVED {

		newStatus = PO_STATUS_CLOSED_FOR_RECEIVING

	} else if purchaseOrder.POStatus == PO_STATUS_CLOSED_FOR_RECEIVING {

		newStatus = PO_STATUS_OPEN_FOR_INVOICE

	} else if purchaseOrder.POStatus == PO_STATUS_OPEN_FOR_INVOICE {

		newStatus = PO_STATUS_INVOICE_COMPLETED

	} else if purchaseOrder.POStatus == PO_STATUS_INVOICE_COMPLETED {

		newStatus = PO_STATUS_CLOSED_FOR_INVOICE

	} else if purchaseOrder.POStatus == PO_STATUS_CLOSED_FOR_INVOICE {

		newStatus = PO_STATUS_CLOSED

		/*}else if PurchaseOrder.POStatus == "SubmitBL"{

		      newStatus = "B/L Submitted"

		  }else if PurchaseOrder.POStatus == "SubmitAN"{

		      newStatus = "A/N Submitted"

		  }else if PurchaseOrder.POStatus == "SubmitCRR"{

		      newStatus = "Request FWD"

		  } else if PurchaseOrder.POStatus == "NotifyImShip"{

		      newStatus = "Notify to Im Ship"

		  } else if PurchaseOrder.POStatus == "CargoReceived"{

		      newStatus = "Cargo Received"
		  }*/

		//Start- Check that the currentStatus to newStatus transition is accurate

		stateTransitionAllowed := false

		/*
			if PurchaseOrder.POStatus == "P/O Created" && newStatus == "P/O Submitted" {
			stateTransitionAllowed = true
			} else */

		if purchaseOrder.POStatus == PO_STATUS_CREATED && newStatus == PO_STATUS_ACCPETED {
			stateTransitionAllowed = true
		} else if purchaseOrder.POStatus == PO_STATUS_CREATED && newStatus == PO_STATUS_OPEN {
			stateTransitionAllowed = true
		} else if purchaseOrder.POStatus == PO_STATUS_OPEN && newStatus == PO_STATUS_RECEIVED {
			stateTransitionAllowed = true
		} else if purchaseOrder.POStatus == PO_STATUS_RECEIVED && newStatus == PO_STATUS_CLOSED_FOR_RECEIVING {
			stateTransitionAllowed = true
		} else if purchaseOrder.POStatus == PO_STATUS_CLOSED_FOR_RECEIVING && newStatus == PO_STATUS_OPEN_FOR_INVOICE {
			stateTransitionAllowed = true
		} else if purchaseOrder.POStatus == PO_STATUS_OPEN_FOR_INVOICE && newStatus == PO_STATUS_CLOSED_FOR_INVOICE {
			stateTransitionAllowed = true
		} else if purchaseOrder.POStatus == PO_STATUS_CLOSED_FOR_INVOICE && newStatus == PO_STATUS_CLOSED {
			stateTransitionAllowed = true
		}

		if stateTransitionAllowed == false {
			return errors.New(" This state transition is not allowed ")
		}

		//Add to state
		_, err = SavePurchaseOrder(stub, purchaseOrder)
		if err != nil {
			return LogAndError(err.Error())
		}

		timestamp, err := stub.GetTxTimestamp()
		if err != nil {
			return LogAndError(err.Error())
		}

		updateType := UPDATE_TYPE_SOURCE_STATUS
		fromValue := purchaseOrder.POStatus
		statusValue := newStatus
		//Update history
		purchaseOrderUpdate := NewPurchaseOrderUpdate(updateType, fromValue, statusValue, "", callerDetails.Username, timestamp.String())

		return UpdatePurchaseOrderHistory(stub, purchaseOrderId, purchaseOrderUpdate)
	}
	return err
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPurchaseOrder(stub shim.ChaincodeStubInterface, purchaseOrderId string) (PurchaseOrder, error) {

	//TODO Has permission to retrieve?

	purchaseOrder, err := RetrievePurchaseOrder(stub, purchaseOrderId)

	return purchaseOrder, err
}

func GetAllPurchaseOrders(stub shim.ChaincodeStubInterface, callerDetails CallerDetails) (PurchaseOrders, error) {

	//TODO Has permission to retrieve?
	purchaseOrders := PurchaseOrders{}

	purchaseOrderIds, err := RetrieveIdsHolder(stub, ALL_PURCHASE_ORDER_IDS_KEY)

	if err != nil {
		return purchaseOrders, LogAndError("Unable to retrieve order id holder")
	}

	for _, purchaseOrderId := range purchaseOrderIds.Ids {
		purchaseOrder, err := GetPurchaseOrder(stub, purchaseOrderId)

		/*if accessDenied {
			fmt.Println("Access denied when reading order: " + purchaseOrderId)
		} else*/if err != nil {
			return purchaseOrders, LogAndError("There was an error when retrieving order: " + err.Error())
		} else {
			purchaseOrders.PurchaseOrders = append(purchaseOrders.PurchaseOrders, purchaseOrder)
		}
	}

	return purchaseOrders, err
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

func addPurchaseOrderIdToHolder(stub shim.ChaincodeStubInterface, purchaseOrderId string) error {
	idHolder, err := RetrieveIdsHolder(stub, ALL_PURCHASE_ORDER_IDS_KEY)

	if err != nil {
		fmt.Println("Unable to retrieve id holder so this is probably the first Purchase order...adding")

		idHolder = IdsHolder{}
	}

	idHolder.Ids = append(idHolder.Ids, purchaseOrderId)

	_, err = SaveIdsHolder(stub, ALL_PURCHASE_ORDER_IDS_KEY, idHolder)

	return err
}
