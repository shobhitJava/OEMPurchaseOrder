package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type PO struct {
	Order_Id         string `json:"order_Id"`
	Asset_ID         string `json:"asset_ID"`
	Asset_Name       string `json:"asset_Name"`
	Order_Desc       string `json:"order_Desc"`
	Order_Quantity   string `json:"order_Quantity"`
	Supplier_Id      string `json:"supplier_Id"`
	Supplier_Name    string `json:"supplier_Name"`
	Supplier_Address string `json:"supplier_Address"`
	Supplier_Contact string `json:"supplier_Contact"`
	Requested_Date   string `json:"requested_Date"`
	Order_Status     string `json:"order_Status"`
}
type SUBO struct {
	SubOrderId        string `json:"subOrder_Id"`
	Order_Id          string `json:"order_Id"`
	Tier1_Name        string `json:"tier1_Name"`
	Asset_ID          string `json:"asset_Id"`
	Asset_Name        string `json:"asset_Name"`
	SubOrder_Desc     string `json:"subOrder_Desc"`
	SubOrder_Quantity string `json:"subOrder_Quantity"`
	Supplier2_Id      string `json:"supplier2_Id"`
	Supplier2_Name    string `json:"supplier2_Name"`
	Supplier2_Address string `json:"supplier2_Address"`
	Supplier2_Contact string `json:"supplier2_Contact"`
	Requested_Date    string `json:"requested_Date"`
	SubOrder_Status   string `json:"subOrder_Status"`
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
	this.Asset_ID = row.Columns[2].GetString_()
	this.Tier1_Name = row.Columns[3].GetString_()
	this.Asset_Name = row.Columns[4].GetString_()
	this.SubOrder_Desc = row.Columns[5].GetString_()
	this.SubOrder_Quantity = row.Columns[6].GetString_()
	this.Supplier2_Id = row.Columns[7].GetString_()
	this.Supplier2_Name = row.Columns[8].GetString_()
	this.Supplier2_Address = row.Columns[9].GetString_()
	this.Supplier2_Contact = row.Columns[10].GetString_()
	this.Requested_Date = row.Columns[11].GetString_()
	this.SubOrder_Status = row.Columns[12].GetString_()
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
		&shim.ColumnDefinition{"Order_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Tier1_Name", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Asset_ID", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Asset_Name", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"SubOrder_Desc", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"SubOrder_Quantity", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier2_Id", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier2_Name", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier2_Address", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Supplier2_Contact", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Requested_Date", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"SubOrder_Status", shim.ColumnDefinition_STRING, false}})

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
	if function == "fetchAllSubOrdersByTier1" {

		return fetchAllSubOrdersByTier1(stub, args)
	}
	if function == "fetchAllOrdersBySupplierName" {

		return fetchAllOrdersBySupplierName(stub, args)
	}
	if function == "fetchAllSubOrdersbyOrderId" {

		return fetchAllSubOrdersbyOrderId(stub, args)
	}
	if function == "fetchSubOrderBySubOrderId" {

		return fetchSubOrderBySubOrderId(stub, args)
	}
	if function == "fetchOrderByOrderId" {

		return fetchOrderByOrderId(stub, args)
	}
	if function == "changeOrderStatus" {

		return changeOrderStatus(stub, args)
	}
	if function == "fetchAllSubOrders" {

		return fetchAllSubOrders(stub, args)
	}

	return nil, nil

}

func changeOrderStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var OrderRowBytes []byte
	po := PO{}
	var err error
	var message string

	OrderRowBytes, err = fetchOrderByOrderId(stub, args)

	err = json.Unmarshal(OrderRowBytes, &po)

	orderStatus := po.Order_Status

	if orderStatus == "New" && args[1] == "Accept" {
		po.Order_Status = "InProgress"
	}
	if orderStatus == "InProgress" && args[1] == "Dispatched" {
		po.Order_Status = "Completed"
	}
	if orderStatus == "New" && args[1] == "Reject" {
		po.Order_Status = "Rejected byTier1"
	}

	col_Val := args[0]
	col1Val := po.Asset_ID
	col2Val := po.Asset_Name
	col3Val := po.Order_Desc
	col4Val := po.Order_Quantity
	col5Val := po.Supplier_Id
	col6Val := po.Supplier_Name
	col7Val := po.Supplier_Address
	col8Val := po.Supplier_Contact
	col9Val := po.Requested_Date
	col10Val := po.Order_Status

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
	ok, err := stub.ReplaceRow("OEM", row)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + args[0] + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}

	message = "Order Status Updated"
	return []byte(message), nil

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

		ordBytes, err = fetchOrderByOrderId(stub, args)

		fmt.Println("ordBytes ", string(ordBytes))

		err = json.Unmarshal(ordBytes, &po)

		if err == nil {
			fmt.Println("inside iF")
		}

		poList = append(poList, po)

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
