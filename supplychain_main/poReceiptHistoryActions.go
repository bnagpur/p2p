package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const PURCHASE_RECEIPT_HISTORY_KEY_PREFIX = "HIST"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================
func UpdatePOReceiptHistory(stub shim.ChaincodeStubInterface, poReceiptId string, update POReceiptUpdate) error {
	receiptHistoryId := PURCHASE_RECEIPT_HISTORY_KEY_PREFIX + poReceiptId
	_, err := RetrievePOReceiptHistory(stub, receiptHistoryId)

	if err != nil {
		return LogAndError("Unable to retrieve order history: " + err.Error())
	}
	poReceiptHistory := POReceiptHistory{[]POReceiptUpdate{}}
	poReceiptHistory.POReceiptUpdates = append(poReceiptHistory.POReceiptUpdates, update)

	_, err = SavePOReceiptHistory(stub, receiptHistoryId, poReceiptHistory)
	return err
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPOReceiptHistory(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poReceiptId string) (POReceiptHistory, bool, error) {

	//TODO Has permission to retrieve?

	poReceiptHistory, err := RetrievePOReceiptHistory(stub, PURCHASE_RECEIPT_HISTORY_KEY_PREFIX+poReceiptId)

	return poReceiptHistory, false, err
}
