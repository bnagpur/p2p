package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"github.com/hyperledger/fabric/common/util"
)

//SupplychainChaincode...
var ACCESS_DENIED_RESPONSE = []byte("{\"failure\": \"ACCESS_DENIED\"}")
var NOT_FOUND_RESPONSE = []byte("{\"failure\": \"NOT_FOUND\"}")

type SupplychainChaincode struct{}

func (t *SupplychainChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("initialising")

	InitDao(stub, "SupplychainChaincode")

	return shim.Success(nil)
}

// Invoke is the entry point to invoke a chaincode function
func (t *SupplychainChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("invoke is running ")

	callerDetails, err := GetCallerDetails(stub)
	if err != nil {
		fmt.Println("An error occured whilst obtaining the caller details: " + callerDetails.Role)
		return shim.Error("An error occured whilst obtaining the caller details")
	}
	fmt.Println("User is " + callerDetails.Username)
	fmt.Println("Role is " + callerDetails.Role)
	function, args := stub.GetFunctionAndParameters()
	// Handle different functions
	if function == "init" {
		b := t.Init(stub)
		return b
	} else if function == "clearData" {
		b, _ := t.processClearData(stub, callerDetails, args)
		return b
	} else if function == "addOrder" {
		_, err := t.processAddPurchaseOrder(stub, callerDetails, args)
		if err != nil {
			return shim.Error("Error Occured in processAddPurchaseOrder ")
		}

	} else if function == "updatePurchaseOrderStatus" {
		_, err := t.processUpdatePurchaseOrderStatus(stub, callerDetails, args)
		if err != nil {
			return shim.Error("Error Occured in processUpdatePurchaseOrderStatus ")
		}
	}
	fmt.Println("invoke did not find func: " + function)

	return shim.Success(nil) /*, errors.New("Received unknown function invocation: " + function)*/
}

func (t *SupplychainChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	callerDetails, err := GetCallerDetails(stub)
	if err != nil {
		fmt.Println("An error occured whilst obtaining the caller details")
		return nil, err
	}

	if function == "getPurchaseOrder" {
		return t.processGetPurchaseOrder(stub, callerDetails, args)
	} else if function == "getAllPurchaseOrders" {
		return t.processGetAllPurchaseOrders(stub, callerDetails, args)
	} else if function == "getPurchaseOrderHistory" {
		return t.processGetPurchaseOrderHistory(stub, callerDetails, args)
	}

	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

//=================================================================================================================================
//	 processAddOrder - Processes an addOrganistion request.
//          args -  Recipient,
//                  Address,
//                  SourceWarehouse,
//                  DeliveryCompany,
//                  Items,
//                  Client,
//                  Owner,
//                  Timestamp
//=================================================================================================================================
func (t *SupplychainChaincode) processAddPurchaseOrder(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, args []string) ([]byte, error) {

	fmt.Println("running processAddPurchaseOrder)")

	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting (Recipient, Address, SourceWarehouse, DeliveryCompany, Items, Client, Owner)")
	}

	/*items, err := MarshallItems(args[4])

	if err != nil {
		return nil, LogAndError("Invalid items: " + args[4] + ", error: " + err.Error())
	}*/

	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, LogAndError(err.Error())
	}
	quantity, _ := strconv.Atoi(args[2])
	fmt.Println("calling NewPurchaseOrder")

	purchaseOrder := NewPurchaseOrder(stub.GetTxID(), args[0], args[1], quantity, args[3], args[4], timestamp.String())
	callerDetails.Role = ROLE_INTERNAL_SYSTEM
	callerDetails.Username = "user01"

	return nil, AddPurchaseOrder(stub, callerDetails, purchaseOrder)
}

//=================================================================================================================================
//	 processAddOrder - Processes an addOrganistion request.
//          args -  orderId,
//                  statusType,
//                  statusValue,
//                  comment
//=================================================================================================================================
func (t *SupplychainChaincode) processUpdatePurchaseOrderStatus(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, args []string) ([]byte, error) {

	fmt.Println("running processUpdatePurchaseOrderStatus)")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting (OrderId, StatusType, StatusValue, Comment)")
	}

	err := UpdatePurchaseOrderStatus(stub, callerDetails, args[0])

	return nil, err

}

//=================================================================================================================================
//	 processClearData - Processes an clearData request.
//          args -  NONE
//=================================================================================================================================
func (t *SupplychainChaincode) processClearData(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, args []string) (pb.Response, error) {

	fmt.Println("running processClearData)")

	if len(args) != 0 {
		err := errors.New("Incorrect number of arguments. Expecting (NONE)")
		return shim.Success(nil), err
	}

	return shim.Success(nil), nil

}

//=================================================================================================================================
//	 processGetOrder - Processes a getOrder request.
//          args -  orderId
//=================================================================================================================================
func (t *SupplychainChaincode) processGetPurchaseOrder(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, args []string) ([]byte, error) {

	fmt.Println("running processGetPurchaseOrder()")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting (OrderId)")
	}

	purchaseOrder, err := GetPurchaseOrder(stub, args[0])

	/*if accessDenied {
		return ACCESS_DENIED_RESPONSE, nil
	}*/

	if err != nil {
		return NOT_FOUND_RESPONSE, nil
	}

	return marshall(purchaseOrder)
}

//=================================================================================================================================
//	 processGetAllOrders - Processes a getAllOrders request.
//
//=================================================================================================================================
func (t *SupplychainChaincode) processGetAllPurchaseOrders(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, args []string) ([]byte, error) {

	fmt.Println("running processGetAllOrders()")

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting NONE")
	}

	purchaseOrders, err := GetAllPurchaseOrders(stub, callerDetails)

	//Probably no orders, return empty orders
	if err != nil {
		purchaseOrders = PurchaseOrders{[]PurchaseOrder{}}
	}

	return marshall(purchaseOrders)
}

//=================================================================================================================================
//	 processGetOrderHistory - Processes a getOrderHistory request.
//          args -  orderId
//=================================================================================================================================
func (t *SupplychainChaincode) processGetPurchaseOrderHistory(stub shim.ChaincodeStubInterface, callerDetails CallerDetails, args []string) ([]byte, error) {

	fmt.Println("running processGetPurchaseOrderHistory()")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting (OrderId)")
	}

	purchaseOrderHistory, accessDenied, err := GetPurchaseOrderHistory(stub, callerDetails, args[0])

	if accessDenied {
		return ACCESS_DENIED_RESPONSE, nil
	}

	if err != nil {
		return NOT_FOUND_RESPONSE, nil
	}

	return marshall(purchaseOrderHistory)
}

func main() {
	err := shim.Start(new(SupplychainChaincode))
	if err != nil {
		fmt.Printf("Error starting Org Registrar chaincode: %s", err)
	}
}
