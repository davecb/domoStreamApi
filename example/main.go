package main

import (
	"flag"
	"fmt"
	domo "github.com/davecb/domoStreamApi"
	// formerly "github.com/JumboInteractiveLimited/domostreamapi"
)

func main() {

	d := domo.New("<clientID>", "<secret>")
	d.SetLogging(true)
	d.SetDebugLogging(true)

	flag.Parse()

	//Dataset API
	payload := `"Jacob Bernoulli","yes"
"Jakob Hermann","no"
"Christian Goldbach","yes"`

	schema := `
{
	"name": "Leonhard Euler Party",
	"description": "Mathematician Guest List",
	"rows": 0,
	"schema": {
	"columns": [ {
		"type": "STRING",
		"name": "Friend"
		}, {
		"type": "STRING",
		"name": "Attending"
		} ]
	}
}`

	schemaupdate := `
{
	"name": "Leonhard Euler Birthday Bash",
	"description": "VIP Guest List",
	"pdpEnabled": true
}`

	fmt.Println("\n-----------  d.DataSet.Create ----------")
	// create a new schema to use with example
	myDataID, err := d.DataSet.Create(schema)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myDataID)

	fmt.Println("\n-----------  d.DataSet.List ----------")
	// return a list of datasets
	fmt.Print(d.DataSet.List())

	fmt.Println("\n-----------  d.DataSet.Get ----------")
	// get a summary finding the next dataset by name
	datasetsummary, err := d.DataSet.Get("Leonhard Euler Party")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", datasetsummary)

	fmt.Println("\n-----------  d.DataSet.Retrieve ----------")
	// get the all details on the datsset we just created
	dataset, err := d.DataSet.Retrieve(myDataID.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", dataset)

	fmt.Println("\n-----------  d.DataSet.Import ----------")
	// let upload some data now
	err = d.DataSet.Import(myDataID.ID, payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n-----------  d.DataSet.Export ----------")
	// now our dataset has somke data lets get it
	fmt.Print(d.DataSet.Export(myDataID.ID))

	fmt.Println("\n-----------  d.DataSet.Update ----------")
	// now make a schema change
	fmt.Print(d.DataSet.Update(myDataID.ID, schemaupdate))

	fmt.Println("\n-----------  d.DataSet.Delete ----------")
	// thats all for now so delete the dataset
	err = d.DataSet.Delete(myDataID.ID)
	if err != nil {
		fmt.Println(err)
	}
}
