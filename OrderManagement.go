package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type PO struct {
	Order_Id         string `json:"Order_Id"`
	Asset_ID         string `json:"Asset_ID"`
	Asset_Name       string `json:"Asset_Name"`
	Order_Desc       string `json:"Order_Desc"`
	Order_Quantity   string `json:"Order_Quantity"`
	Supplier_Id      string `json:"Supplier_Id"`
	Supplier_Name    string `json:"Supplier_Name"`
	Supplier_Address string `json:"Supplier_Address"`
	Supplier_Contact string `json:"Supplier_Contact"`
	Requested_Date   string `json:"Requested_Date"`
	Order_Status     string `json:"Order_Status"`
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
	OrderIds []string `json:"order_Ids"`
}

type SUB_ORDERS_LIST struct {
	SubOderId []string `json:"subOderId"`
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
	this.Asset_ID = row.Columns[1].GetString_()
	this.Asset_Name = row.Columns[2].GetString_()
	this.Order_Desc = row.Columns[3].GetString_()
	this.Order_Quantity = row.Columns[4].GetString_()
	this.Supplier_Id = row.Columns[5].GetString_()
	this.Supplier_Name = row.Columns[6].GetString_()
	this.Supplier_Address = row.Columns[7].GetString_()
	this.Supplier_Contact = row.Columns[8].GetString_()
	this.Requested_Date = row.Columns[9].GetString_()
	this.Order_Status = row.Columns[10].GetString_()

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
		&shim.ColumnDefinition{"Asset_ID", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Asset_Name", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Desc", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Quantity", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier_Name", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier_Address", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier_Contact", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Requested_Date", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Order_Status", shim.ColumnDefinition_STRING, false}})

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
	if function == "fetchAllOrdersBySupplierName" {

		return fetchAllOrdersBySupplierName(stub, args)
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

func fetchAllOrdersBySupplierName(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var ordBytes []byte
	var obytes []byte

	orderIdsBytes, err := stub.GetState(args[0])

	poList := []PO{}
	po := PO{}

	if err != nil {
		return nil, errors.New("some error in getting orders with Order Id ")
	}

	var orderIds ORDERS_LIST
	json.Unmarshal(orderIdsBytes, &orderIds)

	for _, ord := range orderIds.OrderIds {

		fmt.Println("Inside for loop for getting orders. orderId is  ", ord)

		args[0] = ord

		ordBytes, err = fetchOrderById(stub, args)

		fmt.Println("ordBytes ", string(ordBytes))

		err = json.Unmarshal(ordBytes, &po)

		if err == nil {
			fmt.Println("inside iF")

			poList = append(poList, po)
		}

	}

	obytes, err = json.Marshal(poList)

	return obytes, nil

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

	for _, sub := range subOrderIds.SubOderId {

		fmt.Println("Inside for loop for getting suborders. SUBORDER Id is  ", sub)
		args[0] = sub

		sbytes, err = fetchSubOrderBySubOrderId(stub, args)
		fmt.Println("sbytes ", string(sbytes))

		err = json.Unmarshal(sbytes, &s)

		if err == nil {
			fmt.Println("inside iF")

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

	subo = new(SUBO)
	subo.convertSub(&row)

	jsonRows, err = json.Marshal(subo)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}

func fetchOrderByOrderId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

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

	fmt.Println("order id  Row ", rowString1)

	var po *PO

	po = new(PO)
	po.convert(&row)

	jsonRows, err = json.Marshal(po)
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

	col_Val := strCurrentId
	col1Val := args[0]
	col2Val := args[1]
	col3Val := args[2]
	col4Val := args[3]
	col5Val := args[4]
	col6Val := args[5]
	col7Val := args[6]
	col8Val := args[7]
	col9Val := args[8]
	col10Val := "Created"

	var columns []*shim.Column

	col0 := shim.Column{Value: &shim.Column_String_{String_: col_Val}}
	col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
	col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
	col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
	col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}
	col7 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}
	col8 := shim.Column{Value: &shim.Column_String_{String_: col8Val}}
	col9 := shim.Column{Value: &shim.Column_String_{String_: col9Val}}
	col10 := shim.Column{Value: &shim.Column_String_{String_: col10Val}}

	columns = append(columns, &col0)
	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)
	columns = append(columns, &col5)
	columns = append(columns, &col6)
	columns = append(columns, &col7)
	columns = append(columns, &col8)
	columns = append(columns, &col9)
	columns = append(columns, &col10)

	row := shim.Row{Columns: columns}
	ok, err := stub.InsertRow("OEM", row)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + args[0] + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}

	//store the orders Ids of the orders assigned to tier1 with Tier1Name as key

	existingBytes, err = stub.GetState(col6Val)
	var newOrderId ORDERS_LIST
	json.Unmarshal(existingBytes, &newOrderId)

	if err != nil {
		return nil, errors.New("error unmarshalling new Property Address")
	}

	newOrderId.OrderIds = append(newOrderId.OrderIds, row.Columns[0].GetString_())
	bytes, err = json.Marshal(newOrderId)
	if err != nil {

		return nil, errors.New("error marshalling new Property Address")
	}

	err = stub.PutState(col6Val, bytes)

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

	rowString1 := fmt.Sprintf("%s", row)

	fmt.Println("SubOrderRowInserted ", rowString1)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + col1Val + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}

	var getBytes []byte
	getBytes, err = stub.GetState(col1Val)

	fmt.Println("getBytes " + string(getBytes))
	fmt.Println(err)

	newSubOrderId := SUB_ORDERS_LIST{}
	var subOrderIdBytes []byte

	json.Unmarshal(getBytes, &newSubOrderId)

	fmt.Println("newSubOrderId", newSubOrderId)

	newSubOrderId.SubOderId = append(newSubOrderId.SubOderId, strCurrentId)
	subOrderIdBytes, err = json.Marshal(newSubOrderId)
	stub.PutState(col1Val, subOrderIdBytes)

	if err != nil {
		return nil, errors.New("error marshalling new subOrderIDS")
	}

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

				po.convert(&row)

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
