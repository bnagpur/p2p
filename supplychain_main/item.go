package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ALL_ITEM_IDS_KEY = "ITEM"

//==============================================================================================================================
//	Item - Defines the structure for an items object.
//==============================================================================================================================
type Item struct {
	ItemId      string `json:"itemid"`
	Code        string `json:"code"`
	Description string `json:"description"`
	UnitPrice   string `json:"UnitPrice"`
}

type Items struct {
	Items []Item `json:"items"`
}

func MarshallItems(itemsRepresentation string) (Items, error) {
	var items Items
	var err error
	err = unmarshal([]byte(itemsRepresentation), &items)

	return items, err
}

func AddItem(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, item Item, itemId string) error {
	//TODO Validate order

	if callerDetails.Role != ROLE_INTERNAL_SYSTEM {
		return LogAndError("Caller does not have permission to add an order")
	}
	var err error
	//Add to state
	_, err = SaveItem(stub, item, itemId)
	if err != nil {
		return err
	}

	err = addItemIdToHolder(stub, itemId)
	if err != nil {
		return LogAndError(err.Error())
	}

	//Add blank order history
	//  orderHistory := OrderHistory{[]OrderUpdate{}}
	// _, err = SaveOrderHistory(stub, ORDER_HISTORY_KEY_PREFIX + order.Id, orderHistory)

	return err
}

func addItemIdToHolder(stub shim.ChaincodeStubInterface, itemId string) error {
	idHolder, err := RetrieveIdsHolder(stub, ALL_ITEM_IDS_KEY)

	if err != nil {
		fmt.Println("Unable to retrieve id holder so this is probably the first Invoice...adding")

		idHolder = IdsHolder{}
	}

	idHolder.Ids = append(idHolder.Ids, itemId)

	_, err = SaveIdsHolder(stub, ALL_ITEM_IDS_KEY, idHolder)

	return err
}
