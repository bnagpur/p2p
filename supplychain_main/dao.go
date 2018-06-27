package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CRUD_CHAINCODE_ID_KEY Stored the chaincodeId of the CRUD chaincode
const CRUD_CHAINCODE_ID_KEY = "CRUD_CHAINCODE_ID"

//SAVE_FUNCTION
const SAVE_FUNCTION = "save"
const RETRIEVE_FUNCTION = "retrieve"

type Dao struct{}

type IdsHolder struct {
	Ids []string `json:"ids"`
}

func InitDao(stub shim.ChaincodeStubInterface, crudChaincodeId string) {
	stub.PutState(CRUD_CHAINCODE_ID_KEY, []byte(crudChaincodeId))
}

func SavePurchaseOrderHistory(stub shim.ChaincodeStubInterface, id string, purchaseOrderHistory PurchaseOrderHistory) (PurchaseOrderHistory, error) {
	var err error
	err = saveObject(stub, id, purchaseOrderHistory)

	return purchaseOrderHistory, err
}

func RetrievePurchaseOrderHistory(stub shim.ChaincodeStubInterface, id string) (PurchaseOrderHistory, error) {
	var purchaseOrderHistory PurchaseOrderHistory
	var err error
	err = retrieveObject(stub, id, &purchaseOrderHistory)

	return purchaseOrderHistory, err
}

func SavePurchaseOrder(stub shim.ChaincodeStubInterface, purchaseOrder PurchaseOrder) (PurchaseOrder, error) {
	var err error
	err = saveObject(stub, purchaseOrder.Id, purchaseOrder)

	return purchaseOrder, err
}

func SavePOReceiptHistory(stub shim.ChaincodeStubInterface, id string, poReceiptHistory POReceiptHistory) (POReceiptHistory, error) {
	var err error
	err = saveObject(stub, id, poReceiptHistory)

	return poReceiptHistory, err
}

func RetrievePOReceiptHistory(stub shim.ChaincodeStubInterface, id string) (POReceiptHistory, error) {
	var poReceiptHistory POReceiptHistory
	var err error
	err = retrieveObject(stub, id, &poReceiptHistory)

	return poReceiptHistory, err
}

func SavePOReceipt(stub shim.ChaincodeStubInterface, poReceipt POReceipt) (POReceipt, error) {
	var err error
	err = saveObject(stub, poReceipt.Id, poReceipt)

	return poReceipt, err
}

func SaveGoodsReceiptNoteHistory(stub shim.ChaincodeStubInterface, id string, goodsReceiptNoteHistory GoodsReceiptNoteHistory) (GoodsReceiptNoteHistory, error) {
	var err error
	err = saveObject(stub, id, goodsReceiptNoteHistory)

	return goodsReceiptNoteHistory, err
}

func RetrieveGoodsReceiptNoteHistory(stub shim.ChaincodeStubInterface, id string) (GoodsReceiptNoteHistory, error) {
	var goodsReceiptNoteHistory GoodsReceiptNoteHistory
	var err error
	err = retrieveObject(stub, id, &goodsReceiptNoteHistory)

	return goodsReceiptNoteHistory, err
}

func SaveGoodsReceiptNote(stub shim.ChaincodeStubInterface, goodsReceiptNoteId string, goodsReceiptNote GoodsReceiptNote) (GoodsReceiptNote, error) {
	var err error
	err = saveObject(stub, goodsReceiptNoteId, goodsReceiptNote)

	return goodsReceiptNote, err
}

func SaveItem(stub shim.ChaincodeStubInterface, item Item, itemId string) (Item, error) {
	var err error
	err = saveObject(stub, itemId, item)
	//retrieveObject
	return item, err
}

// RetrievePurchaseOrder...
func RetrievePurchaseOrder(stub shim.ChaincodeStubInterface, id string) (PurchaseOrder, error) {
	var purchaseOrder PurchaseOrder

	//var err error
	err := retrieveObject(stub, id, &purchaseOrder)

	return purchaseOrder, err
}

func RetrievePOReceipt(stub shim.ChaincodeStubInterface, id string) (POReceipt, error) {
	var poReceipt POReceipt
	var err error
	err = retrieveObject(stub, id, &poReceipt)

	return poReceipt, err
}

func RetrieveGoodsReceiptNote(stub shim.ChaincodeStubInterface, id string) (GoodsReceiptNote, error) {
	var goodsReceiptNote GoodsReceiptNote
	var err error
	err = retrieveObject(stub, id, &goodsReceiptNote)

	return goodsReceiptNote, err
}

func RetrieveItem(stub shim.ChaincodeStubInterface, id string) (Item, error) {
	var item Item

	var err error
	err = retrieveObject(stub, id, &item)

	return item, err
}

/*Need to add the funtions for receipts and GRN

 */

func SaveIdsHolder(stub shim.ChaincodeStubInterface, id string, idsHolder IdsHolder) (IdsHolder, error) {
	var err error
	err = saveObject(stub, id, idsHolder)

	return idsHolder, err
}

func RetrieveIdsHolder(stub shim.ChaincodeStubInterface, id string) (IdsHolder, error) {
	var idsHolder IdsHolder

	var err error
	err = retrieveObject(stub, id, &idsHolder)

	return idsHolder, err
}

//Invoice Functions Start
func SaveInvoiceHistory(stub shim.ChaincodeStubInterface, id string, invoiceHistory PoInvoiceHistory) (PoInvoiceHistory, error) {
	var err error
	err = saveObject(stub, id, invoiceHistory)

	return invoiceHistory, err
}

func RetrieveInvoiceHistory(stub shim.ChaincodeStubInterface, id string) (PoInvoiceHistory, error) {
	var invoiceHistory PoInvoiceHistory
	var err error
	err = retrieveObject(stub, id, &invoiceHistory)

	return invoiceHistory, err
}

func SaveInvoice(stub shim.ChaincodeStubInterface, invoiceId string, invoice PoInvoice) (PoInvoice, error) {
	var err error
	err = saveObject(stub, invoiceId, invoice)

	return invoice, err
}

func RetrieveInvoice(stub shim.ChaincodeStubInterface, id string) (PoInvoice, error) {
	var invoice PoInvoice

	var err error
	err = retrieveObject(stub, id, &invoice)

	return invoice, err
}

/*func SaveInvoiceHistory(stub shim.ChaincodeStubInterface, id string, invoiceHistory PoInvoiceHistory) (PoInvoiceHistory, error) {
	var err error
	err = saveObject(stub, id, goodsReceiptNoteHistory)

	return goodsReceiptNoteHistory, err
}*/
//Invoice Functions End

//PoShipping Functions Start
func SavePoShippingHistory(stub shim.ChaincodeStubInterface, id string, poShippingHistory PoShippingHistory) (PoShippingHistory, error) {
	var err error
	err = saveObject(stub, id, poShippingHistory)

	return poShippingHistory, err
}

func RetrievePoShippingHistory(stub shim.ChaincodeStubInterface, id string) (PoShippingHistory, error) {
	var poShippingHistory PoShippingHistory
	var err error
	err = retrieveObject(stub, id, &poShippingHistory)

	return poShippingHistory, err
}

func SavePoShipping(stub shim.ChaincodeStubInterface, poShippingId string, poShipping PoShipping) (PoShipping, error) {
	var err error
	err = saveObject(stub, poShippingId, poShipping)

	return poShipping, err
}

func RetrievePoShipping(stub shim.ChaincodeStubInterface, id string) (PoShipping, error) {
	var poShipping PoShipping

	var err error
	err = retrieveObject(stub, id, &poShipping)

	return poShipping, err
}

/*func SavePoShippingHistory(stub shim.ChaincodeStubInterface, id string, poShippingHistory PoShippingHistory) (PoShippingHistory, error) {
	var err error
	err = saveObject(stub, id, poShippingHistory)

	return poShippingHistory, err
}*/
//PoShipping Functions End

func ClearData(stub shim.ChaincodeStubInterface) pb.Response {
	return invoke(stub, getCrudChaincodeId(stub), "init", []string{})
}

func retrieve(stub shim.ChaincodeStubInterface, id string) pb.Response {
	return query(stub, getCrudChaincodeId(stub), RETRIEVE_FUNCTION, []string{id})
}

func retrieveObject(stub shim.ChaincodeStubInterface, id string, toStoreObject interface{}) error {

	var err error
	//var bytes bytes
	response := retrieve(stub, id)
	st := response.GetStatus()
	//var er string

	if st == 0 {
		fmt.Printf("RetrieveObject: Cannot retrieve object with id: "+id+" : %s", st)
		return err
	}

	err = unmarshal(response.GetPayload(), toStoreObject)

	if err != nil {
		fmt.Printf("RetrieveObject: Cannot unmarshall object with id: "+id+" : %s", err)
		return err
	}

	return nil
}

func saveObject(stub shim.ChaincodeStubInterface, id string, object interface{}) error {
	//var err error
	//var bytes bytes
	bytes, err := marshall(object)

	if err != nil {
		fmt.Printf("\nUnable to marshall object with id: "+id+" : %s", err)
		return err
	}
	//pb.Response
	response := save(stub, id, bytes)
	var st int32
	st = response.GetStatus()

	if st == 0 {
		fmt.Printf("Unable to save policy: %d", st)
	}
	//err= response.GetStatus()
	//return response

	return err
}

func save(stub shim.ChaincodeStubInterface, id string, toSave []byte) pb.Response {
	return invoke(stub, getCrudChaincodeId(stub), SAVE_FUNCTION, []string{id, string(toSave)})
}

func invoke(stub shim.ChaincodeStubInterface, chaincodeId, functionName string, args []string) pb.Response {
	//response pb.Response
	//var bytes bytes
	response := stub.InvokeChaincode(chaincodeId, createArgs(functionName, stub.GetTxID(), args), "")
	return response
}

func query(stub shim.ChaincodeStubInterface, chaincodeId, functionName string, args []string) pb.Response {
	return stub.InvokeChaincode(chaincodeId, createArgs(functionName, stub.GetTxID(), args), "")

	//stub.GetQueryResult
}

//func (w http.ResponseWriter, r *http.Request) {

func createArgs(functionName string, txId string, args []string) [][]byte {
	var funcAndArgs []string
	funcAndArgs = append([]string{functionName}, txId)
	funcAndArgs = append(funcAndArgs, args...)
	return util.ArrayToChaincodeArgs(funcAndArgs)
}

func marshall(toMarshall interface{}) ([]byte, error) {
	return json.Marshal(toMarshall)
}

func unmarshal(data []byte, object interface{}) error {
	return json.Unmarshal(data, object)
}

func getCrudChaincodeId(stub shim.ChaincodeStubInterface) string {
	//var bytes bytes
	bytes, _ := stub.GetState(CRUD_CHAINCODE_ID_KEY)
	return string(bytes)
}
