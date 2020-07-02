package main
import ("fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"encoding/json"
	"net/http"
)
type Car struct{
	ModelName string
	Color string
	SerialNo string
	Manufacturer string
	owner Owner

}
type Owner struct{
	name string
	nationality string
	address string
	gender string

}
type CarChainCode struct{

}

func (t *Car) ChangeOwner(newOwner owner){
	newOwner=owner

}

func(c *CarChainCode) Init(stub shim.chaincodeStubInterface) peer.Response{
	tom:= Owner{name:"Tom", nationality:"American", address:"12-202 Park Street,New York", gender:"M"}
	Bob:= Owner{name:"Bob", nationality:"French", address:"A-202 L'Street,Paris", gender:"M"}
	car:= Car{ModelName:"XUV", Color:"white", SerialNo:"A2344fe2", Manufacturer:"Mahindra Motors", owner:"tom"}

	tomAsJSONbytes,_ := json.Marshal(tom)
	err := stub.PutState(tom.nationality, tomAsJSONbytes)

	if err != nil{
		return fmt.Println("Failed to create Asset"+tom.name)

	}
	bobAsJSONbytes,_ := json.Marshal(bob)
    err := stub.PutState(bob.nationality, bobAsJSONbytes)

    if err !=nil{
    	return fmt.Println("Failed to Create Asset"+bob.name)
    }

    carAsJSONbytes,_ := json.Marshal(car)
    err := stub.PutState(car.SerialNo,carAsJSONbytes)

    if err !=nil{
    	return fmt.Println("Failed to Create Asset"+car.ModelName)
    }        
}

func(c *CarChainCode) TransferOwnership(stub shim.ChaincodeStubInterface,args []string) peer.Response{
	carAsbytes,_ := stub.GetState(args[0])

	if carAsbytes==nil{
		return shim.Error("Car Asset not found")
	}
	car:= Car{}
	_ := json.Unmarshal(carAsbytes, &car)

	OwnerAsbytes,_ := stub.GetState(args[1])

	if OwnerAsbytes==nil{
		return shim.Error("Owner can't be found")
	}
	newOwner:= newOwner{}
	_ := json.Unmarshal(OwnerAsbytes, &newOwner)

	car.ChangeOwner(newOwner)
	carAsJSONbytes,_ := json.Marshal(car)

	err := stub.PutState(car.SerialNo, carAsJSONbytes)

	if err != nil{
		return shim.Error("Not able to create asset")
	}

	return shim.Success([] byte("Car Asset Modified"))
}
func(c *CarChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fn, args:= stub.GetFunctionAndParameters()
	if fn == "TransferOwnership"{
		return fn.TransferOwnership(stub,args)
	}
	return shim.Error("This function does not exists in chaincode")
}
func main(){
	logger.SetLevel(shim.info)
	err:= shim.Start(CarChainCode)
	if err!=nil{
		return shim.Error("Unable to start chaincode")
	}
}
