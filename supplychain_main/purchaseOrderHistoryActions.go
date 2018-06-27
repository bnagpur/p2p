package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const PURCHASE_ORDER_HISTORY_KEY_PREFIX = "HIST"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================
func UpdatePurchaseOrderHistory(stub shim.ChaincodeStubInterface, purchaseOrderId string, update PurchaseOrderUpdate) error {
	orderHistoryId := PURCHASE_ORDER_HISTORY_KEY_PREFIX + purchaseOrderId
	_, err := RetrievePurchaseOrderHistory(stub, orderHistoryId)

	if err != nil {
		return LogAndError("Unable to retrieve order history: " + err.Error())
	}
	purchaseOrderHistory := PurchaseOrderHistory{[]PurchaseOrderUpdate{}}
	purchaseOrderHistory.PurchaseOrderUpdates = append(purchaseOrderHistory.PurchaseOrderUpdates, update)

	_, err = SavePurchaseOrderHistory(stub, orderHistoryId, purchaseOrderHistory)
	return err
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPurchaseOrderHistory(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, purchaseOrderId string) (PurchaseOrderHistory, bool, error) {

	//TODO Has permission to retrieve?

	purchaseOrderHistory, err := RetrievePurchaseOrderHistory(stub, PURCHASE_ORDER_HISTORY_KEY_PREFIX+purchaseOrderId)

	return purchaseOrderHistory, false, err
}
