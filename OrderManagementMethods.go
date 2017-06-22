package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

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

	var revisedDate string
	var notification string

	revisedDate = " "
	notification = " "

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
	col10Val := args[9]
	col11Val := revisedDate
	col12Val := notification

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
	col11 := shim.Column{Value: &shim.Column_String_{String_: col11Val}}
	col12 := shim.Column{Value: &shim.Column_String_{String_: col12Val}}

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
	columns = append(columns, &col11)
	columns = append(columns, &col12)

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

	byteOrderId, err := stub.GetState("subOrderIdNumber")
	strOrderId := string(byteOrderId)
	intOrderId, _ := strconv.Atoi(strOrderId)

	currentId := intOrderId + 1
	fmt.Println("currentId : " + strconv.Itoa(currentId))
	str := strconv.Itoa(currentId)

	strCurrentId := "SUB" + strconv.Itoa(currentId)
	stub.PutState("subOrderIdNumber", []byte(str))

	var revisedDate string
	var notification string

	revisedDate = " "
	notification = " "

	col_Val := strCurrentId
	col0Val := args[0]
	col1Val := args[1]
	col2Val := args[2]
	col3Val := args[3]
	col4Val := args[4]
	col5Val := args[5]
	col6Val := args[6]
	col7Val := args[7]
	col8Val := args[8]
	col9Val := args[9]
	col10Val := args[10]
	col11Val := args[11]
	col12Val := revisedDate
	col13Val := notification

	var columns []*shim.Column

	col := shim.Column{Value: &shim.Column_String_{String_: col_Val}}
	col0 := shim.Column{Value: &shim.Column_String_{String_: col0Val}}
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
	col11 := shim.Column{Value: &shim.Column_String_{String_: col11Val}}
	col12 := shim.Column{Value: &shim.Column_String_{String_: col12Val}}
	col13 := shim.Column{Value: &shim.Column_String_{String_: col13Val}}

	columns = append(columns, &col)
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
	columns = append(columns, &col11)
	columns = append(columns, &col12)
	columns = append(columns, &col13)

	row := shim.Row{Columns: columns}

	ok, err := stub.InsertRow("TIER1", row)

	rowString1 := fmt.Sprintf("%s", row)

	fmt.Println("SubOrderRowInserted ", rowString1)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + col_Val + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}

	// Store the subOrder Ids assigned to a particular tier 2 supplier  with Tier2 Supplier name as key and SubOrder Ids as values

	var getsubOrderIdBytes []byte
	getsubOrderIdBytes, err = stub.GetState(args[7])

	fmt.Println("getBytes " + string(getsubOrderIdBytes))
	fmt.Println(err)

	newSubOrderIds := SUB_ORDERS_LIST{}
	var newSubOrderIdBytes []byte

	json.Unmarshal(getsubOrderIdBytes, &newSubOrderIds)

	fmt.Println("newSubOrderIds", newSubOrderIds)

	newSubOrderIds.SubOderId = append(newSubOrderIds.SubOderId, strCurrentId)
	newSubOrderIdBytes, err = json.Marshal(newSubOrderIds)
	stub.PutState(args[7], newSubOrderIdBytes)

	if err != nil {
		return nil, errors.New("error marshalling new subOrderIDS")
	}

	//store all suborders created for one particular Order in a list with Order Id as key and SuborderIds as values :

	var getBytes []byte
	getBytes, err = stub.GetState(col0Val)

	fmt.Println("getBytes " + string(getBytes))
	fmt.Println(err)

	newSubOrderId := SUB_ORDERS_LIST{}
	var subOrderIdBytes []byte

	json.Unmarshal(getBytes, &newSubOrderId)

	fmt.Println("newSubOrderId", newSubOrderId)

	newSubOrderId.SubOderId = append(newSubOrderId.SubOderId, strCurrentId)
	subOrderIdBytes, err = json.Marshal(newSubOrderId)
	stub.PutState(col0Val, subOrderIdBytes)

	fmt.Println("subOrders List for an oRder id " + string(subOrderIdBytes))

	if err != nil {
		return nil, errors.New("error marshalling new subOrderIDS")
	}

	return nil, nil
}

func fetchAllOrders(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	rowChannel, err := stub.GetRows("OEM", columns)

	var jsonRows []byte

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

	jsonRows, err = json.Marshal(orderArray)

	if err != nil {
		return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}
func fetchAllSubOrdersAssignedToTier2(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	rowChannel, err := stub.GetRows("TIER1", columns)

	var jsonRows []byte

	subOrderArray := []SUBO{}

	fmt.Println("args[0]  = ", args[0])

	for {
		select {

		case row, ok := <-rowChannel:

			fmt.Println("OK = ", ok)

			if !ok {
				rowChannel = nil
			} else {

				fmt.Println("Inside Else of for loop in query")
				subo := SUBO{}

				rowString1 := fmt.Sprintf("%s", row)

				fmt.Println("Suborer id  Row ", rowString1)

				subo.convertSub(&row)

				if subo.Supplier2_Name == args[0] {
					subOrderArray = append(subOrderArray, subo)
				}
			}

		}
		if rowChannel == nil {
			break
		}
	}

	jsonRows, err = json.Marshal(subOrderArray)

	if err != nil {
		return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}

func fetchAllSubOrdersByTier1(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	rowChannel, err := stub.GetRows("TIER1", columns)
	var jsonRows []byte

	subOrderArray := []SUBO{}

	fmt.Println("args[0]  = ", args[0])

	for {
		select {

		case row, ok := <-rowChannel:

			fmt.Println("OK = ", ok)

			if !ok {
				rowChannel = nil
			} else {

				fmt.Println("Inside Else of for loop in query")
				subo := SUBO{}

				rowString1 := fmt.Sprintf("%s", row)

				fmt.Println("Suborer id  Row ", rowString1)

				subo.convertSub(&row)

				if subo.Tier1_Name == args[0] {
					subOrderArray = append(subOrderArray, subo)
				}
			}

		}
		if rowChannel == nil {
			break
		}
	}

	jsonRows, err = json.Marshal(subOrderArray)

	if err != nil {
		return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}
