package main

//==============================================================================================================================
//	Item - Defines the structure for an inventory object.
//==============================================================================================================================
type Inventory struct {
	Code         string `json:"code"`
	Description  string `json:"description"`
	SubInventory string `json:"subInventory"`
	Items        []Item `json:"items"`
	//RefID		 PurchaseOrder.Id 	`json:"refid"`

}

type Inventories struct {
	Inventories []Inventory `json:"inventories"`
}

func MarshallInventory(itemsRepresentation string) (Inventories, error) {
	var inventories Inventories

	err := unmarshal([]byte(itemsRepresentation), &inventories)

	return inventories, err
}
