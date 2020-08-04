package dynamoDB

import (
	"encoding/json"
	"fmt"
	"github.com/BhaviD/BootCamp_Team3_gRPC/pkg/dynamoDB/dbUtil"
	"github.com/BhaviD/BootCamp_Team3_gRPC/pkg/dynamoDB/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io/ioutil"
	"log"
)

var orderTable = "T3_Order"
var customerTable = "T3_Customer"
var restaurantTable = "T3_Restaurant"

func AddCustomerTable() {
	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	// create the api params
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(customerTable),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("Id"), KeyType: aws.String("HASH")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("Id"), AttributeType: aws.String("N")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	// create the table
	resp, err := db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the response data
	fmt.Println(resp)
}

func DeleteCustomerTable() {
	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	// create the api params
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String(customerTable),
	}

	// delete the table
	resp, err := db.DeleteTable(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}
	// print the response data
	fmt.Printf("Response = %+v\n", resp)
}

func LoadCustomerData() {
	// read the json data file
	f, err := ioutil.ReadFile("./assets/customerData.json")
	if err != nil {
		panic("Could not read Customer json data file")
	}

	// parse the json movie data
	var customers []types.Customer
	if err := json.Unmarshal(f, &customers); err != nil {
		panic("Could not parse json movie data")
	}

	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	// iterate through each movie
	for _, cust := range customers {
		// marshal the movie struct into an aws attribute value map
		custAVMap, err := dynamodbattribute.MarshalMap(cust)
		if err != nil {
			panic("Cannot marshal Customer into AttributeValue map")
		}

		// create the api params
		params := &dynamodb.PutItemInput{
			TableName: aws.String(customerTable),
			Item:      custAVMap,
		}

		// put the item
		resp, err := db.PutItem(params)
		if err != nil {
			fmt.Printf("Unable to add customer: %v\n", err.Error())
		} else {
			// print the response data
			fmt.Printf("Put item successful: '%s' (resp = '%+v')\n", cust.FullName, resp)
		}
	}
}

func AddOrderTable() {
	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	awsParams := &dynamodb.CreateTableInput{
		TableName: aws.String(orderTable),

		KeySchema: [] *dynamodb.KeySchemaElement {
			{AttributeName: aws.String("CustId"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("Id"), KeyType: aws.String("RANGE")},
		},

		AttributeDefinitions: [] *dynamodb.AttributeDefinition {
			{AttributeName: aws.String("CustId"), AttributeType: aws.String("N")},
			{AttributeName: aws.String("Id"), AttributeType: aws.String("N")},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput {
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}
	response, err := db.CreateTable(awsParams)
	if err != nil {
		log.Fatalf("Sorry error creating table : %v", err)
	}
	// print the response
	fmt.Println(response)
}


func LoadOrderData() {
	// read the json data file
	f, err := ioutil.ReadFile("./assets/orderData.json")
	if err != nil {
		panic("Could not read Order json data file")
	}

	// parse the json movie data
	var orders []types.Order
	if err := json.Unmarshal(f, &orders); err != nil {
		panic("Could not parse json Order data")
	}

	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	// iterate through each movie
	for _, order := range orders {
		// marshal the movie struct into an aws attribute value map
		orderMap, err := dynamodbattribute.MarshalMap(order)
		if err != nil {
			panic("Cannot marshal Order into AttributeValue map")
		}

		// create the api params
		params := &dynamodb.PutItemInput{
			TableName: aws.String(orderTable),
			Item:      orderMap,
		}

		// put the item
		resp, err := db.PutItem(params)
		if err != nil {
			fmt.Printf("Unable to add order: %v\n", err.Error())
		} else {
			// print the response data
			fmt.Printf("Put item successful: '%s' (resp = '%+v')\n", order.Id, resp)
		}
	}
}

func AddRestaurantTable() {
	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	params := &dynamodb.CreateTableInput{
		TableName: aws.String(restaurantTable),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("Id"), KeyType: aws.String("HASH")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("Id"), AttributeType: aws.String("N")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	// create the table
	resp, err := db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the response data
	fmt.Println(resp)
}

func LoadRestaurantData() {
	// read the json data file
	f, err := ioutil.ReadFile("./assets/restaurantData.json")
	if err != nil {
		panic("Could not read Restaurant json data file")
	}

	// parse the json movie data
	var restaurants []types.Restaurant
	if err := json.Unmarshal(f, &restaurants); err != nil {
		panic("Could not parse json Restaurant data")
	}

	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	// iterate through each movie
	for _, rest := range restaurants {
		// marshal the movie struct into an aws attribute value map
		custAVMap, err := dynamodbattribute.MarshalMap(rest)
		if err != nil {
			panic("Cannot marshal Restaurant into AttributeValue map")
		}

		// create the api params
		params := &dynamodb.PutItemInput{
			TableName: aws.String(restaurantTable),
			Item:      custAVMap,
		}

		// put the item
		resp, err := db.PutItem(params)
		if err != nil {
			fmt.Printf("Unable to add restaurant: %v\n", err.Error())
		} else {
			// print the response data
			fmt.Printf("Put item successful: '%s' (resp = '%+v')\n", rest.Name, resp)
		}
	}
}

func DeleteRestaurantTable() {
	// create a dynamodb instance
	db := dbUtil.GetDBInstance()

	// create the api params
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String(restaurantTable),
	}

	// delete the table
	resp, err := db.DeleteTable(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}
	// print the response data
	fmt.Printf("Response = %+v\n", resp)
}