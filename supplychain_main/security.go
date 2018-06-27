package main

import (
	//"github.com/hyperledger/fabric/core/utils"
	//"github.com/hyperledger/fabric/common/tools/cryptogen/msp"

	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const ROLE_INTERNAL_SYSTEM = "INTERNAL_SYSTEM"

//==============================================================================================================================
//	 Security Functions
//==============================================================================================================================
//	 getCallerDetails - Retrieves details about a caller from the cert
//==============================================================================================================================

func GetCallerDetails(stub shim.ChaincodeStubInterface) (CallerDetails, error) {
	attributes := make(map[string][]byte)
	attributes["username"] = []byte("balee")
	attributes["role"] = []byte("INTERNAL_SYSTEM")

	user, _, err := cid.GetAttributeValue(stub, "username")
	fmt.Println("user is: " + user)
	if err != nil {
		user = "dummy-user"
	}

	role, err := getRole(stub)
	if err != nil {
		role = "dummy-role"
	}

	fmt.Printf("caller_data: %s %s\n", user, role)

	return NewCallerDetails(user, role), err
}

//==============================================================================================================================
//	 get_username - Retrieves the username of the user who invoked the chaincode from cert attributes.
//				    Returns the username as a string.
//==============================================================================================================================
func getUsername(stub shim.ChaincodeStubInterface) (string, error) {

	//username, _, err := cid.GetAttributeValue(stub, "attr1")
	id, err := cid.GetID(stub)
	if err != nil {
		fmt.Printf("error in function getUsername: Couldn't get attribute 'id'. Error: %s\n", id)
		return "", errors.New("Couldn't get attribute 'username'. Error: " + err.Error())
	}

	return string(id), nil
}

/*func getUsername(stub shim.ChaincodeStubInterface) (peer.Response, error) {

	serializedID, _ := stub.GetCreator()

	sId := &msp.SerializedIdentity{}

	v, ok, err = cid.GetAttributeValue(stub, "username")
	err := proto.Unmarshal(serializedID, sId)
	if err != nil {
		return shim.Error(fmt.Sprintf("Could not deserialize a SerializedIdentity, err %s", err)), err
	}

	bl, _ := pem.Decode(sId.IdBytes)
	if bl == nil {
		return shim.Error(fmt.Sprintf("Failed to decode PEM structure")), err
	}
	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to parse certificate %s", err)), err
	}

	return pb.Response.Payload, nil
}*/
//==============================================================================================================================
//	 getRole - Retrieves the role of the user who invoked the chaincode from cert attributes.
//				    Returns the role as a string.
//==============================================================================================================================
func getRole(stub shim.ChaincodeStubInterface) (string, error) {
	affiliation, _, err := cid.GetAttributeValue(stub, "role")
	if err != nil {
		return "", errors.New("Couldn't get attribute 'role'. Error: " + err.Error())
	}
	return string(affiliation), nil
}
