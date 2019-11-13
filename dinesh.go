package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	r "math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type timedevice struct {
	Timeinv     time.Time `json:",Timeinv"`
	Temperature int32     `json:"Temperature"`
}

// logisticstrans type
type logisticstrans struct {
	//product might be food,fish,phone,other itmes
	//Product id should be unique such as FISH123,Prawns456,ICECREAM789
	ProductID         string       `json:"ProductID"`
	ProductType       string       `json:"ProductType"`
	SellerID          string       `json:"SellerID"`
	SellerLocation    string       `json:"SellerLocation"`
	BuyerID           string       `json:"BuyerID"`
	BuyerLocation     string       `json:"BuyerLocation"`
	LogisticsID       string       `json:"LogisticsID"`
	LogisticsLocation string       `json:"LogisticsLocation"`
	JourneyStartTime  string       `json:",JourneyStartTime"`
	JourneyEndTime    string       `json:",JourneyEndTime"`
	Status            string       `json:"Status"`
	Timefromdevice    []timedevice `json:"timefromdevice"`
}

func main() {

	err := shim.Start(new(logisticstrans))
	if err != nil {
		fmt.Println("Error with chaincode")
	} else {
		fmt.Println("Chaincode installed successfully")
	}
}

//Init logisticstrans
func (t *logisticstrans) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initiate the chaincode")
	return shim.Success(nil)
}

//Invoke logisticstrans
func (t *logisticstrans) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fun, args := stub.GetFunctionAndParameters()
	fmt.Println("Arguements for function  ", fun)
	switch fun {
	case "RequestLogistic":
		return t.RequestLogistic(stub, args)
	case "TransitLogistics":
		return t.TransitLogistics(stub, args)
	case "DeliveryLogistics":
		return t.DeliveryLogistics(stub, args)
	case "GetAllProducts":
		return t.GetAllProducts(stub, args)
	case "QueryName":
		return t.QueryName(stub, args)
	case "GetTransactionHistoryForKey":
		return t.GetTransactionHistoryForKey(stub, args)
	}
	fmt.Println("Function not found!")
	return shim.Error("Recieved unknown function invocation!")
}

//Genlogistics for

func (t *logisticstrans) RequestLogistic(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	var ProductID string

	if len(args) < 1 {
		fmt.Println("Invalid number od arguements")
		return shim.Error(err.Error())
	}
	if err != nil {
		return shim.Error("Invalid Request Number")
	}

	var logobj = logisticstrans{ProductID: ProductID, ProductType: args[1], BuyerID: args[2], BuyerLocation: args[3], SellerID: args[4], SellerLocation: args[5]}
	logobj.Status = "Requested"

	logobjasBytes, _ := json.Marshal(logobj)
	stub.PutState(args[0], logobjasBytes)

	return shim.Success(nil)
}

//TransitLogistics at the same time measuring the temp details from logistics
func (t *logisticstrans) TransitLogistics(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 7")
	}

	if err != nil {
		return shim.Error("Invalid ")
	}
	logisticsAsBytes, _ := stub.GetState(args[0])

	var logisticobj logisticstrans
	json.Unmarshal(logisticsAsBytes, &logisticobj)
	logisticobj.ProductID = args[0]
	logisticobj.LogisticsID = args[1]
	logisticobj.LogisticsLocation = args[2]
	logisticobj.JourneyStartTime = args[3]
	logisticobj.JourneyEndTime = args[4]

	if logisticobj.Status != "Requested" {
		fmt.Println("we cannnot transit  the product which was not requested")
		return shim.Error("we cannnot transit  the product which was not requested")
	}

	logisticobj.Status = "In-Transit"
	//calling generateTimeTempRandom func
	logisticobj.Timefromdevice = generateTimeTempRandom(logisticobj.JourneyStartTime, logisticobj.JourneyEndTime)

	logisticsAsBytes, _ = json.Marshal(logisticobj)
	stub.PutState(args[0], logisticsAsBytes)

	return shim.Success(nil)
}
func (t *logisticstrans) DeliveryLogistics(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) > 2 {
		return shim.Error("Invalid   no of arg for delivery function ")

	}

	logisticsasbytes1, _ := stub.GetState(args[0])

	var logisticobj1 logisticstrans

	json.Unmarshal(logisticsasbytes1, &logisticobj1)

	if logisticobj1.Status != "In-Transit" {
		fmt.Println("we cannnot delivery the product which is not in In_Transit")
		return shim.Error("we cannnot delivery the product which is not in In_Transit")
	}
	fmt.Println("length of the logibj timefrrom device", len(logisticobj1.Timefromdevice))
	fmt.Println("length of the logibj journry in  device", logisticobj1.JourneyEndTime)
	fmt.Println("length of the logibj  journey out timefrrom device", logisticobj1.JourneyStartTime)

	count := 0
	for i := 0; i < len(logisticobj1.Timefromdevice); i++ {
		if logisticobj1.Timefromdevice[i].Temperature >= 20 {
			//fmt.Println("Temperature from array is :", logisticobj1.Timefromdevice[i].Temperature)
			//fmt.Println("status of temp from array is :", logisticobj1.Status)

			count++
		} else {
			count = 0
		}
		//fmt.Println("Count is  from for loop:", count)
		//fmt.Println("status of temp  is :", logisticobj1.Status)

		if count >= 3 {
			logisticobj1.Status = "Rejected from Buyer"
			break

		} else {
			logisticobj1.Status = "Accepted  from Buyer"
		}
		//logisticobj1.Status = "Rejected from Buyer"

	}
	//fmt.Println("Count is :", count)

	logisticsasbytes1, _ = json.Marshal(logisticobj1)
	stub.PutState(args[0], logisticsasbytes1)

	return shim.Success(nil)
}

//generation of time for 10 munutes  and temp randomly between the in and out values
func generateTimeTempRandom(JourneyStartTime string, JourneyEndTime string) []timedevice {

	var tmpdevArray []timedevice
	var tmpdev timedevice
	// temp range fixing in between 18 to 23
	min := 18
	max := 23
	EndTime, _ := strconv.Atoi(JourneyEndTime)
	StartTime, _ := strconv.Atoi(JourneyStartTime)

	//journey time should in railway time and
	journeyTime := (EndTime - StartTime)
	//	start := time.Date(2019, 1, 1, 12, 10, 0, 0, time.UTC)
	start := time.Now()
	for i := 0; i < (journeyTime * 6); i++ {
		rand.Seed(time.Now().UnixNano())
		var interval int
		interval = 10 * (i + 1)
		afterTenMinutes := start.Add(time.Minute * time.Duration(interval))
		a := r.Intn(max-min+1) + min
		tmpdev.Timeinv = afterTenMinutes
		tmpdev.Temperature = int32(a)
		tmpdevArray = append(tmpdevArray, tmpdev)
	}
	return tmpdevArray
}

//GetAllProducts func
func (t *logisticstrans) GetAllProducts(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	queryString := fmt.Sprintf("{\"selector\":{\"ProductID\":{\"$ne\": \"%s\"}}}", "null")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here
	if err != nil {
		fmt.Printf("Unable to get All logistics details: %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Printf("logistics Details : %v\n", queryResults)

	return shim.Success(queryResults)
}

// getQueryResultForQueryString
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("***getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)

	if err != nil {
		fmt.Println("Error from getQueryResultForQueryString:  ", err)
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("***getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//QueryName with product id
func (t *logisticstrans) QueryName(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return shim.Error(err.Error())
	}
	//fetch data from couch db starts here
	var ProductID = args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"ProductID\":{\"$eq\": \"%s\"}}}", ProductID)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here
	if err != nil {
		fmt.Printf("Unable to get product  details: %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Printf("Details for product    : %v\n", queryResults)

	return shim.Success(queryResults)
}
func getTransHistory(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("***GetTransactionHistory for Key :\n%s\n", queryString)

	resultsIterator, err := stub.GetHistoryForKey(queryString)

	if err != nil {
		fmt.Println("Error from GetHistoryForKey:  ", err)
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("*** GetTransactionHistory:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *logisticstrans) GetTransactionHistoryForKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	queryResults, err := getTransHistory(stub, args[0])
	if err != nil {
		fmt.Printf("Unable to get all Transactions : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Printf("Transaction History: %v\n", queryResults)

	return shim.Success(queryResults)
}
