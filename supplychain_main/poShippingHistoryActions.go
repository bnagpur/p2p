package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const PO_SHIPPING_HISTORY_KEY_PREFIX = "HIST"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================
func UpdatePoShippingHistory(stub shim.ChaincodeStubInterface, poShippingId string, update PoShippingUpdate) error {
	shippingHistoryId := PO_SHIPPING_HISTORY_KEY_PREFIX + poShippingId
	_, err := RetrievePoShippingHistory(stub, shippingHistoryId)

	if err != nil {
		return LogAndError("Unable to retrieve order history: " + err.Error())
	}

	poShippingHistory := PoShippingHistory{[]PoShippingUpdate{}}

	poShippingHistory.PoShippingUpdates = append(poShippingHistory.PoShippingUpdates, update)

	_, err = SavePoShippingHistory(stub, shippingHistoryId, poShippingHistory)
	return err
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPoShippingHistory(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poShippingId string) (PoShippingHistory, bool, error) {

	//TODO Has permission to retrieve?

	poShippingHistory, err := RetrievePoShippingHistory(stub, PO_SHIPPING_HISTORY_KEY_PREFIX+poShippingId)

	return poShippingHistory, false, err
}
