package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type PO struct {
	Order_Id       string `json:"order_Id"`
	Order_Desc     string `json:"order_desc"`
	Order_Quantity string `json:"order_quantity"`
	Assigned_To_Id string `json:"assigned_to_id"`
	Created_By_Id  string `json:"created_by_id"`
	Order_Status   string `json:"order_status"`
	Asset_ID       string `json:"asset_ID"`
}
type SUBO struct {
	SubOrderId     string `json:"subOrder_Id"`
	Order_Id       string `json:"order_Id"`
	Order_Desc     string `json:"order_desc"`
	Order_Quantity string `json:"order_quantity"`
	Assigned_To_Id string `json:"assigned_to_id"`
	Created_By_Id  string `json:"created_by_id"`
	Order_Status   string `json:"order_status"`
	Asset_ID       string `json:"asset_ID"`
}

type ORDERS_LIST struct {
	orderIds []string `json:"order_Ids"`
}

type SUB_ORDERS_LIST struct {
	suboOderId []string `json:"order_Ids"`
}

type SimpleChaincode struct {
}

func main() {

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)

	}
}

func (this *PO) convert(row *shim.Row) {

	this.Order_Id = row.Columns[0].GetString_()
	this.Order_Desc = row.Columns[1].GetString_()
	this.Order_Quantity = row.Columns[2].GetString_()
	this.Assigned_To_Id = row.Columns[3].GetString_()
	this.Created_By_Id = row.Columns[4].GetString_()
	this.Order_Status = row.Columns[5].GetString_()
	this.Asset_ID = row.Columns[6].GetString_()

}
func (this *SUBO) convertSub(row *shim.Row) {

	this.SubOrderId = row.Columns[0].GetString_()
	this.Order_Id = row.Columns[1].GetString_()
	this.Order_Desc = row.Columns[2].GetString_()
	this.Order_Quantity = row.Columns[3].GetString_()
	this.Assigned_To_Id = row.Columns[4].GetString_()
	this.Created_By_Id = row.Columns[5].GetString_()
	this.Order_Status = row.Columns[6].GetString_()
	this.Asset_ID = row.Columns[7].GetString_()

}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var orderIDsBytes []byte

	var err error

	var orderIds ORDERS_LIST

	orderIDsBytes, err = json.Marshal(orderIds)

	stub.PutState("tier1_orders", orderIDsBytes)

	//above array will save the orders assigned to Tier 1

	err = stub.CreateTable("OEM", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"OrderId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"Order_Desc", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Quantity", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Assigned_To_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Created_By_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Status", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Asset_ID", shim.ColumnDefinition_STRING, false}})

	if err != nil {
		return nil, errors.New("OEM table not created")
	}

	err = stub.CreateTable("TIER1", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"SubOrderId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"OrderId", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Desc", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Quantity", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Assigned_To_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Created_By_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Status", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Asset_ID", shim.ColumnDefinition_STRING, false}})

	if err != nil {
		return nil, errors.New("TIER1 table not created")
	}

	orderId := "10"
	subOrderId := "100"

	stub.PutState("orderIdNumber", []byte(orderId))
	stub.PutState("subOrderIdNumber", []byte(subOrderId))

	return nil, nil

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "createOrder" {

		return createOrder(stub, args)
	}

	if function == "createSubOrder" {

		return createSubOrder(stub, args)
	}

	return nil, nil
}
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "fetchAllOrders" {

		return fetchAllOrders(stub, args)
	}
	if function == "fetchOrderById" {

		return fetchOrderById(stub, args)
	}
	if function == "fetchAllSubOrdersbyOrderId" {

		return fetchAllSubOrdersbyOrderId(stub, args)
	}
	if function == "fetchSubOrderBySubOrderId" {

		return fetchSubOrderBySubOrderId(stub, args)
	}
	return nil, nil

}

func fetchAllSubOrdersbyOrderId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var suboBytes []byte
	var sbytes []byte

	orderId := args[0]
	subOrderBytes, err := stub.GetState(orderId)

	subo := []SUBO{}
	s := SUBO{}

	if err != nil {
		return nil, errors.New("some error in getting sub orders with Order Id ")
	}

	var subOrderIds SUB_ORDERS_LIST
	json.Unmarshal(subOrderBytes, &subOrderIds)

	for _, sub := range subOrderIds.suboOderId {

		fmt.Println("Inside for loop for getting suborders. SUBORDER Id is  ", sub)
		args[0] = sub

		sbytes, err = fetchSubOrderBySubOrderId(stub, args)

		err = json.Unmarshal(sbytes, &s)

		if err == nil {

			subo = append(subo, s)
		}

	}

	suboBytes, err = json.Marshal(subo)

	return suboBytes, nil

}

func fetchSubOrderBySubOrderId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	var err error
	var row shim.Row
	var jsonRows []byte

	col0 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
	columns = append(columns, col0)

	row, err = stub.GetRow("TIER1", columns)

	if err != nil {
		return nil, fmt.Errorf("getRow operation failed. %s", err)
	}

	rowString1 := fmt.Sprintf("%s", row)

	fmt.Println("Suborer id  Row ", rowString1)

	var subo *SUBO
	var suboList []*SUBO

	subo = new(SUBO)
	subo.convertSub(&row)

	suboList = append(suboList, subo)

	jsonRows, err = json.Marshal(suboList)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}

func createOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var existingBytes []byte
	var bytes []byte

	//OrderId
	byteOrderId, err := stub.GetState("orderIdNumber")
	strOrderId := string(byteOrderId)
	intOrderId, _ := strconv.Atoi(strOrderId)

	currentId := intOrderId + 1
	str := strconv.Itoa(currentId)
	strCurrentId := "PO" + strconv.Itoa(currentId)
	stub.PutState("orderIdNumber", []byte(str))

	col1Val := strCurrentId
	col2Val := args[0]
	col3Val := args[1]
	col4Val := "Tier 1"
	col5Val := "OEM"
	col6Val := "Order created. Pending with Tier1"
	col7Val := args[2]

	var columns []*shim.Column

	col0 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	col1 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
	col4 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
	col5 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}
	col6 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}

	columns = append(columns, &col0)
	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)
	columns = append(columns, &col5)
	columns = append(columns, &col6)

	row := shim.Row{Columns: columns}
	ok, err := stub.InsertRow("OEM", row)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + args[0] + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}

	existingBytes, err = stub.GetState("tier1_orders")
	var newOrderId ORDERS_LIST
	json.Unmarshal(existingBytes, &newOrderId)

	if err != nil {
		return nil, errors.New("error unmarshalling new Property Address")
	}

	newOrderId.orderIds = append(newOrderId.orderIds, row.Columns[0].GetString_())
	bytes, err = json.Marshal(newOrderId)
	if err != nil {

		return nil, errors.New("error marshalling new Property Address")
	}

	err = stub.PutState("tier1_orders", bytes)

	return nil, nil
}

func createSubOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	bytesub, err := stub.GetState("subOrderIdNumber")
	strSub := string(bytesub)
	intSub, _ := strconv.Atoi(strSub)

	currentSub := intSub + 1
	str := strconv.Itoa(currentSub)
	strCurrentId := "SUB" + strconv.Itoa(currentSub)
	stub.PutState("subOrderIdNumber", []byte(str))

	col0Val := strCurrentId
	col1Val := args[0]
	col2Val := args[1]
	col3Val := args[2]
	col4Val := "Tier 2"
	col5Val := "Tier1"
	col6Val := "Sub Order created. Pending with Tier2"
	col7Val := args[3]

	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	var columns []*shim.Column

	col0 := shim.Column{Value: &shim.Column_String_{String_: col0Val}}
	col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
	col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
	col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
	col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}
	col7 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}

	columns = append(columns, &col0)
	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)
	columns = append(columns, &col5)
	columns = append(columns, &col6)
	columns = append(columns, &col7)

	row := shim.Row{Columns: columns}

	ok, err := stub.InsertRow("TIER1", row)

	fmt.Println("err.Error is ", err.Error())
	fmt.Println("OK ", ok)

	rowString1 := fmt.Sprintf("%s", row)

	fmt.Println("SubOrderRowInserted ", rowString1)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + col1Val + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}

	newSubOrderId := SUB_ORDERS_LIST{}
	var getBytes []byte
	var bytes []byte

	getBytes, err = stub.GetState(col1Val)

	fmt.Println("getBytes " + string(getBytes))
	fmt.Println(err.Error())

	if err != nil {
		return nil, errors.New("error  in get state")
	}

	if getBytes != nil {
		err = json.Unmarshal(getBytes, &newSubOrderId)
	}

	fmt.Println("newSubOrderId.suboOderId ", newSubOrderId.suboOderId)
	fmt.Println("getBytes " + string(getBytes))
	fmt.Println("err " + err.Error())

	newSubOrderId.suboOderId = append(newSubOrderId.suboOderId, strCurrentId)
	bytes, err = json.Marshal(newSubOrderId)
	if err != nil {
		return nil, errors.New("error marshalling new subOrderIDS")
	}

	err = stub.PutState(args[0], bytes)

	return nil, nil
}

func fetchAllOrders(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	rowChannel, err := stub.GetRows("OEM", columns)

	orderArray := []PO{}

	for {
		select {

		case row, ok := <-rowChannel:

			if !ok {
				rowChannel = nil
			} else {

				fmt.Println("Inside Else of for loop in query")
				po := PO{}

				po.Order_Id = row.Columns[0].GetString_()
				po.Order_Desc = row.Columns[1].GetString_()
				po.Order_Quantity = row.Columns[2].GetString_()
				po.Assigned_To_Id = row.Columns[3].GetString_()
				po.Created_By_Id = row.Columns[4].GetString_()
				po.Order_Status = row.Columns[5].GetString_()
				po.Asset_ID = row.Columns[6].GetString_()

				orderArray = append(orderArray, po)
			}

		}
		if rowChannel == nil {
			break
		}
	}

	jsonRows, err := json.Marshal(orderArray)

	if err != nil {
		return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}
func fetchOrderById(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	var err error
	var row shim.Row
	var jsonRows []byte

	col0 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
	columns = append(columns, col0)

	row, err = stub.GetRow("OEM", columns)

	if err != nil {
		return nil, fmt.Errorf("getRow operation failed. %s", err)
	}

	rowString1 := fmt.Sprintf("%s", row)

	fmt.Println("OrderId Row ", rowString1)

	var po *PO
	var poList []*PO

	po = new(PO)
	po.convert(&row)

	poList = append(poList, po)

	jsonRows, err = json.Marshal(poList)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}
