package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const GRN_HISTORY_KEY_PREFIX = "HIST"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================
func UpdateGoodsReceiptNoteHistory(stub shim.ChaincodeStubInterface, goodsReceiptNoteId string, update GoodsReceiptNoteUpdate) error {
	grnHistoryId := GRN_HISTORY_KEY_PREFIX + goodsReceiptNoteId
	_, err := RetrieveGoodsReceiptNoteHistory(stub, grnHistoryId)

	if err != nil {
		return LogAndError("Unable to retrieve order history: " + err.Error())
	}
	goodsReceiptNoteHistory := GoodsReceiptNoteHistory{[]GoodsReceiptNoteUpdate{}}

	goodsReceiptNoteHistory.GoodsReceiptNoteUpdates = append(goodsReceiptNoteHistory.GoodsReceiptNoteUpdates, update)

	_, err = SaveGoodsReceiptNoteHistory(stub, grnHistoryId, goodsReceiptNoteHistory)
	return err
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetGoodsReceiptNoteHistory(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, goodsReceiptNoteId string) (GoodsReceiptNoteHistory, bool, error) {

	//TODO Has permission to retrieve?

	goodsReceiptNoteHistory, err := RetrieveGoodsReceiptNoteHistory(stub, GRN_HISTORY_KEY_PREFIX+goodsReceiptNoteId)

	return goodsReceiptNoteHistory, false, err
}
