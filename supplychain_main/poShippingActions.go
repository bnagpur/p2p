package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ALL_PURCHASE_SHIPPING_IDS_KEY = "PO_SHIPPINGS"

//const STATUS_TYPE_SOURCE = "SOURCE"

//const STATUS_TYPE_TRANSPORT = "TRANSPORT"

//==============================================================================================================================
//	 Invocations
//==============================================================================================================================

func AddPoShipping(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poShipping PoShipping, purchaseOrderid string) error {
	//TODO Validate order

	if callerDetails.Role != ROLE_INTERNAL_SYSTEM {
		return LogAndError("Caller does not have permission to add an order")
	}

	//Add to state
	purchaseOrder, err := GetPurchaseOrder(stub, purchaseOrderid)
	if err == nil && purchaseOrder.POStatus == PO_STATUS_ACCPETED {
		//poreceipt.POId = purchaseOrderid
		poShipping.ShippingStatus = SHIPPING_STATUS_AWAITING_SHIPPING
		_, err := SavePoShipping(stub, poShipping.Id, poShipping)
		if err != nil {
			return err
		}

		err = addPoShippingIdToHolder(stub, poShipping.Id)
		if err != nil {
			return LogAndError(err.Error())
		}

		//Add blank order history
		poShippingHistory := PoShippingHistory{[]PoShippingUpdate{}}
		_, err = SavePoShippingHistory(stub, PURCHASE_ORDER_HISTORY_KEY_PREFIX+poShipping.Id, poShippingHistory)
	}
	return err
}

func UpdatePoShippingStatus(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poShippingId string, statusType string, statusValue string, comment string) error {
	//TODO Validate update (define state progression)
	//TODO permissions
	poShipping, err := RetrievePoShipping(stub, poShippingId)

	if err != nil {
		return LogAndError(err.Error())
	}

	//TODO make this more extendable
	fromValue := ""
	updateType := ""
	/*if statusType == STATUS_TYPE_SOURCE {
	  updateType = UPDATE_TYPE_SOURCE_STATUS
	          fromValue = order.Source.Status
	          order.Source.Status = statusValue
	      }else if statusType == STATUS_TYPE_TRANSPORT {
	  updateType = UPDATE_TYPE_TRANSPORT_STATUS
	          fromValue = order.Transport.Status
	          order.Transport.Status = statusValue
	      }
	*/
	//Add to state
	_, err = SavePoShipping(stub, poShipping.Id, poShipping)
	if err != nil {
		return LogAndError(err.Error())
	}

	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return LogAndError(err.Error())
	}

	//Update history
	poShippingUpdate := NewPoShippingUpdate(updateType, fromValue, statusValue, comment, callerDetails.Username, timestamp.String())

	return UpdatePoShippingHistory(stub, poShipping.Id, poShippingUpdate)
}

//==============================================================================================================================
//	 Queries
//==============================================================================================================================

func GetPoShipping(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, poShippingId string) (PoShipping, bool, error) {

	//TODO Has permission to retrieve?

	poShipping, err := RetrievePoShipping(stub, poShippingId)

	return poShipping, false, err
}

func GetAllPoShippings(stub shim.ChaincodeStubInterface, callerDetails CallerDetails) (PoShippings, error) {

	//TODO Has permission to retrieve?
	poShippings := PoShippings{}

	poShippingIds, err := RetrieveIdsHolder(stub, ALL_PURCHASE_ORDER_IDS_KEY)

	if err != nil {
		return poShippings, LogAndError("Unable to retrieve order id holder")
	}

	for _, poShippingId := range poShippingIds.Ids {
		poShipping, accessDenied, err := GetPoShipping(stub, callerDetails, poShippingId)

		if accessDenied {
			fmt.Println("Access denied when reading order: " + poShippingId)
		} else if err != nil {
			return poShippings, LogAndError("There was an error when retrieving order: " + err.Error())
		} else {
			poShippings.PoShippings = append(poShippings.PoShippings, poShipping)
		}
	}

	return poShippings, err
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

func addPoShippingIdToHolder(stub shim.ChaincodeStubInterface, poShippingId string) error {
	idHolder, err := RetrieveIdsHolder(stub, ALL_PURCHASE_SHIPPING_IDS_KEY)

	if err != nil {
		fmt.Println("Unable to retrieve id holder so this is probably the first Purchase order...adding")

		idHolder = IdsHolder{}
	}

	idHolder.Ids = append(idHolder.Ids, poShippingId)

	_, err = SaveIdsHolder(stub, ALL_PURCHASE_SHIPPING_IDS_KEY, idHolder)

	return err
}
