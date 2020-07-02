package main
import(
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/tree/master/core/chaincode/shim/ext/cid"

)
type Person struct{
	Id string `json: "id"`
	Class string `json: "class"`
	Name string `json:"name"`
	Email string `json:"email"`

}
type Art struct{
	Id string `json:"id"`
	Class string `json:"name"`
	Description string `json:"description"`
	Artist string `json:"artist"`
	Owner string `json:"owner"`
	CreatedAt time.time `json:"created_at"`

}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response{
	return shim.Success(nil)

}

func (s *SmartContract) CreateUser(APIstub shim.ChaincodeStubInterface) sc.Response{
	Id := "user-"+utils.RandStringBytes(32)//utisl is a custom package.Ypu can 
	var user = Person{Class:"Person", Id: Id, Name: args[0], Email:args[1]}
	UserBytes, _ := json.Marshal(user)
	APIstub.PutState(Id, UserBytes)
	return shim.Success(nil)
}

func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string){
	if len(args) !=1{
		return shim.Error("Incorrect number of arguments.Expecting UserId")
    }

    if err!=nil{
    	return shim.Error(err.Error())
    }
    return shim.Success(UserBytes)
}