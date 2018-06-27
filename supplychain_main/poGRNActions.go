package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ALL_GRN_IDS_KEY = "GOODS_RECEIPT_NOTES"

//const STATUS_TYPE_SOURCE = "SOURCE"

//const STATUS_TYPE_TRANSPORT = "TRANSPORT"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================

func AddGoodsReceiptNote(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, goodsReceiptNote GoodsReceiptNote) error {
	//TODO Validate order

	if callerDetails.Role != ROLE_INTERNAL_SYSTEM {
		return LogAndError("Caller does not have permission to add an order")
	}

	//Add to state
	_, err := SaveGoodsReceiptNote(stub, goodsReceiptNote.GRNId, goodsReceiptNote)
	if err != nil {
		return err
	}

	err = addGoodsReceiptNoteIdToHolder(stub, goodsReceiptNote.GRNId)
	if err != nil {
		return LogAndError(err.Error())
	}

	//Add blank order history
	goodsReceiptNoteHistory := GoodsReceiptNoteHistory{[]GoodsReceiptNoteUpdate{}}
	_, err = SaveGoodsReceiptNoteHistory(stub, GRN_HISTORY_KEY_PREFIX+goodsReceiptNote.GRNId, goodsReceiptNoteHistory)

	return err
}

func UpdateGoodsReceiptNoteStatus(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, goodsReceiptNoteId string, statusType string, statusValue string, comment string) error {
	//TODO Validate update (define state progression)
	//TODO permissions
	goodsReceiptNote, err := RetrieveGoodsReceiptNote(stub, goodsReceiptNoteId)

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
	_, err = SaveGoodsReceiptNote(stub, goodsReceiptNote.GRNId, goodsReceiptNote)
	if err != nil {
		return LogAndError(err.Error())
	}

	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return LogAndError(err.Error())
	}

	//Update history
	goodsReceiptNoteUpdate := NewGoodsReceiptNoteUpdate(updateType, fromValue, statusValue, comment, callerDetails.Username, timestamp.String())

	return UpdateGoodsReceiptNoteHistory(stub, goodsReceiptNote.GRNId, goodsReceiptNoteUpdate)
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetGoodsReceiptNote(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, goodsReceiptNoteId string) (GoodsReceiptNote, bool, error) {

	//TODO Has permission to retrieve?

	goodsReceiptNote, err := RetrieveGoodsReceiptNote(stub, goodsReceiptNoteId)

	return goodsReceiptNote, false, err
}

func GetAllGoodsReceiptNotes(stub shim.ChaincodeStubInterface, callerDetails CallerDetails) (GoodsReceiptNotes, error) {

	//TODO Has permission to retrieve?
	goodsReceiptNotes := GoodsReceiptNotes{}

	goodsReceiptNoteIds, err := RetrieveIdsHolder(stub, ALL_GRN_IDS_KEY)

	if err != nil {
		return goodsReceiptNotes, LogAndError("Unable to retrieve order id holder")
	}

	for _, goodsReceiptNoteId := range goodsReceiptNoteIds.Ids {
		goodsReceiptNote, accessDenied, err := GetGoodsReceiptNote(stub, callerDetails, goodsReceiptNoteId)

		if accessDenied {
			fmt.Println("Access denied when reading order: " + goodsReceiptNoteId)
		} else if err != nil {
			return goodsReceiptNotes, LogAndError("There was an error when retrieving order: " + err.Error())
		} else {
			goodsReceiptNotes.GoodsReceiptNotes = append(goodsReceiptNotes.GoodsReceiptNotes, goodsReceiptNote)
		}
	}

	return goodsReceiptNotes, err
}

//==============================================================================================================================
//	 Internal
//==============================================================================================================================
func addGoodsReceiptNoteIdToHolder(stub shim.ChaincodeStubInterface, goodsReceiptNoteId string) error {
	idHolder, err := RetrieveIdsHolder(stub, ALL_GRN_IDS_KEY)

	if err != nil {
		fmt.Println("Unable to retrieve id holder so this is probably the first Purchase order...adding")

		idHolder = IdsHolder{}
	}

	idHolder.Ids = append(idHolder.Ids, goodsReceiptNoteId)

	_, err = SaveIdsHolder(stub, ALL_GRN_IDS_KEY, idHolder)

	return err
}
