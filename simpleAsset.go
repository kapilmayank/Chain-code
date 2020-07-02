package main
import ("fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"html/template"
	"net/http"
)

type SimpleAsset struct{

}
//shim
//Chain code interface-Init, Invoke
//Chain code stub interface

//Init is to instantiate the chain code and to initalize any data
func(t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response{
	//1)tranx proposal
	args:= stub.GetStringArgs()
	//Tranx will have key and value
	if len(args)!=2{
		return shim.Error("Expecting two value i.e. key and value")
	}
    
    //store the key and value to the ledger
	err:= stub.PutState(args[0], []byte(args[1]))
	if err!=nil{
		return shim.Error(fmt.Sprintf("Falied to create the asset %s",args[0]))
	}
	return shim.Success(nil)
}
//Invoke the chain code
func(t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	
	var result string
	var err error
	fn, args:= stub.GetFunctionAndParameters()

	
	if fn == "set"{
		result,err = set(stub,args)
		
	} else {
		result,err = get(stub,args)
	}
	if err!=nil{
		return shim.Error( err.Error())
	}
	return shim.Success([]byte(result))

}
//implementing our chain code
func set(stub shim.ChaincodeStubInterface, args []string) (string, error){
	if len(args) !=2 {
		return "",fmt.Errorf("Expecting key and value for the asset")
	}
	err := stub.PutState(args[0],[]byte(args[1]))
	if err!=nil{
		return "",fmt.Errorf("Falied to set the asset %s",args[0])
	}

	return args[1],nil

}
func get(stub shim.ChaincodeStubInterface,args []string) (string,error){
	if len(args)!=2{
		return "",fmt.Errorf("Expecting key and value for the asset")
	}
	value,err:=stub.GetState(args[0])
	if err!=nil{
		return "",fmt.Errorf("Failed to get the asset %s",args[0])
	}
	if value==nil{
		return "",fmt.Errorf("No value was found.Please enter the correct key")
	}

	return string(value),nil
}

func main(){
	if err:= shim.Start(new(SimpleAsset)); err!=nil{
		fmt.Printf("Not able to run chaincode SimpleAsset")
	}
}