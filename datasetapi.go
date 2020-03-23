package domo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DataSetService Data API service
type DataSetService service

// Get a dataset by name
// Returns all DataSet objects that meet argument criteria from original request.
func (d *DataSetService) Get(name string) (list DatasetSummary, err error) {
	data, err := d.List()
	for _, thisrow := range data {
		if thisrow.Name == name {
			Data := DatasetSummary{
				ID:          thisrow.ID,
				Name:        thisrow.Name,
				Rows:        thisrow.Rows,
				Columns:     thisrow.Columns,
				Description: thisrow.Description,
			}
			return Data, err
		}
	}
	return
}

// List of all DataSets in your Domo instance.
// Definition
// GET https://api.domo.com/v1/datasets
// Returns
// Returns all DataSet objects that meet argument criteria from original request.
func (d *DataSetService) List() (list Datasets, err error) {
	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets", baseURL)
	bodyBytes, _, err := d.client.genericGET(url, nil)

	if err != nil {
		return list, fmt.Errorf("Failed to get List from Domo API %s", err)
	}

	err = json.Unmarshal(bodyBytes, &list)
	if err != nil {
		err = fmt.Errorf("Failed to get unmarshal list %s", err)
	}

	return
}

// FIXME this has never worked for me
// Retrieve the details of an existing DataSet.
// Definition
// GET https://api.domo.com/v1/datasets/{DATASET_ID}
// Returns
// Returns a DataSet object if valid DataSet ID was provided. When requesting,
// if the DataSet ID is related to a DataSet that has been deleted,
// a subset of the DataSet's information will be returned, including a deleted property, which will be true.
func (d *DataSetService) Retrieve(id string) (list Dataset, err error) {
	err =  errors.New("domo.Retrieve has never worked for me")
	return
	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets/%s", baseURL, id)
	bodyBytes, _, err := d.client.genericGET(url, nil)

	if err != nil {
		return list, fmt.Errorf("Failed to retrieve dataset from Domo API %s", err)
	}

	err = json.Unmarshal(bodyBytes, &list)
	if err != nil {
		err = fmt.Errorf("Failed to unmarshal dataset %s", err)
	}

	return
}

// Query runs an sql query on a dataset, as per
// https://developer.domo.com/docs/dataset-api-reference/dataset
//
//POST https://api.domo.com/v1/datasets/query/execute/ce79d23f-ef7d-4318-9787-ebde54a8c5b4
//Accept: application/json
//Authorization: bearer <your-valid-oauth-access-token>
//{"sql": "SELECT * FROM table"}
func (d *DataSetService) Query(datasetID, query string) (data string, err error) {
	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets", baseURL, datasetID)
	body := strings.NewReader(query)
	
	header := map[string]string{
		"Content-Type": "application/json",
	}

	bodyBytes, statusCode, err := d.client.genericPOST(url, body, header)
	if err != nil {
		return data, fmt.Errorf("Failed to query via the Domo API %s", err)
	}

	logger(fmt.Sprintf("query Status Code : %d", statusCode))
	err = json.Unmarshal(bodyBytes, &data)
	bodyString := string(bodyBytes)
	return bodyString, err
}

// Export data from a DataSet in your Domo instance.
// Returns a raw CSV in the response body or error for the outcome of data being exported into DataSet.
// NOTE: Sometimes sends  a 406 error, have tried setting various Accept-Language and Accept-Charset
// headers with no avail
// Definition
// GET https://api.domo.com/v1/datasets/{DATASET_ID}/data
// Returns a raw CSV in the response body or error for the outcome of data being exported into DataSet.
func (d *DataSetService) Export(datasetID string) (data string, err error) {

	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets/%s/data?includeHeader=true&fileName=data.csv", baseURL, datasetID)
	header := map[string]string{
		"Accept": "text/csv",
	}

	bodyBytes, _, err := d.client.genericGET(url, header)

	data = fmt.Sprintf("%s", bodyBytes)

	if err != nil {
		err = fmt.Errorf("Failed to export %s", err)
	}

	return
}

// Create a new DataSet in your Domo instance.
// Once the DataSet has been created, data can then be imported into the DataSet.
// Definition POST https://api.domo.com/v1/datasets
// Returns
// Returns a DataSet object when successful.
// The returned object will have DataSet attributes based on the information that was provided when DataSet was created.
func (d *DataSetService) Create(schema string) (data *Dataset, err error) {
	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets", baseURL)
	body := strings.NewReader(schema)

	header := map[string]string{
		"Content-Type": "application/json",
	}

	bodyBytes, statusCode, err := d.client.genericPOST(url, body, header)

	if err != nil {
		return data, fmt.Errorf("Failed to get create from Domo API %s", err)
	}

	logger(fmt.Sprintf("createStream Status Code : %d", statusCode))
	err = json.Unmarshal(bodyBytes, &data)

	bodyString := string(bodyBytes)

	logger("createStream" + bodyString)

	if err != nil {
		err = fmt.Errorf("Failed to create stream %s", err)
	}

	return

}

// Import data into a DataSet in your Domo instance. This request will replace the data currently in the DataSet.
// Definition
// PUT https://api.domo.com/v1/datasets/{DATASET_ID}/data
// Returns
// Returns a response of success or error for the outcome of data being imported into DataSet.
func (d *DataSetService) Import(datasetID string, payload string) (err error) {
	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets/%s/data", baseURL, datasetID)
	body := strings.NewReader(payload)

	header := map[string]string{
		"Content-Type": "text/csv",
	}

	bodyBytes, _, err := d.client.genericPUT(url, body, header)

	if logging {
		buf := new(bytes.Buffer)
		json.Indent(buf, []byte(bodyBytes), "", "  ")
		fmt.Println(buf)
	}

	return
}

// Delete Permanently deletes a DataSet from your Domo instance.
// This can be done for all DataSets, not just those created through the API.
// Definition
// DELETE https://api.domo.com/v1/datasets/{DATASET_ID}
// Returns
// Returns an empty response. HTTP/1.1 204 No Content
func (d *DataSetService) Delete(datasetID string) (err error) {
	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets/%s", baseURL, datasetID)
	statusCode, err := d.client.genericDELETE(url, nil)

	if statusCode != 204 {
		// it failed need to do something
		err = fmt.Errorf("Failed to delete dataset %s", err)
	}

	return
}

// Update the specified DataSet’s metadata by providing values to parameters passed.
// Definition
// PUT https://api.domo.com/v1/datasets/317970a1-6a6e-4f70-8e09-44cf5f34cf44
// Returns
// Returns a full DataSet object.
func (d *DataSetService) Update(datasetID string, schema string) (data *Dataset, err error) {

	d.client.getAccessToken("data")
	url := fmt.Sprintf("%s/v1/datasets/%s", baseURL, datasetID)
	body := strings.NewReader(schema)

	header := map[string]string{"Content-Type": "application/json"}
	bodyBytes, statusCode, err := d.client.genericPUT(url, body, header)

	if err != nil {
		return data, fmt.Errorf("Failed to update dataset from Domo API %s", err)
	}

	logger(fmt.Sprintf("createStream Status Code : %d", statusCode))
	err = json.Unmarshal(bodyBytes, &data)

	bodyString := string(bodyBytes)

	logger("updatedataset" + bodyString)

	if err != nil {
		err = fmt.Errorf("Can't unmarshal data set response %s", err)
	}

	return
}
