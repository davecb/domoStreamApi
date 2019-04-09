package domo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// StreamService Stream API service
type StreamService service

// Retrieve Retrieves the details of an existing stream.
// Returns a Stream object if valid Stream ID was provided. When requesting,
// if the Stream ID is related to a DataSet that has been deleted,
// a subset of the Stream's information will be returned, including a deleted property, which will be true.
//
// Returns a Stream object if valid Stream ID was provided.
func (s *StreamService) Retrieve(streamID int) (stream Stream, err error) {
	s.client.getAccessToken("data")
	stream, err = s.retrieve(streamID)
	return
}

// retrieveStream Retrieves the details of an existing stream.
// Definition
// https://api.domo.com/v1/streams/{SREAM_ID}
// Returns
// Returns a Stream object if valid Stream ID was provided. When requesting,
// if the Stream ID is related to a DataSet that has been deleted,
// a subset of the Stream's information will be returned, including a deleted property, which will be true.
func (s *StreamService) retrieve(streamID int) (data Stream, err error) {
	url := fmt.Sprintf("%s/v1/streams/%d", baseURL, streamID)
	bodyBytes, _, err := s.client.genericGET(url, nil)

	if err != nil {
		return data, fmt.Errorf(
			"Unable to retrieve stream from Domo API %s",
			err,
		)
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	return
}

// Create When creating a stream, specify the DataSet properties (name and description) and as a convenience,
// the create stream API will create a DataSet for you.
// In addition, you can only have one stream open at a time. If you need to add additional data,
// we recommended adding more parts to the currently open stream or executing a commit of the open stream before creating a new stream.
// KNOWN LIMITATION :
// The StreamAPI currently only allows you to import data to a DataSet created via the Stream API.
// For example, it is currently not supported to import data to a DataSet created by a Domo Connector.
//
// Returns a DataSet object when successful. The returned object will have DataSet attributes based on
// the information that was provided when DataSet was created from the Stream created.
func (s *StreamService) Create(dataset string) (data *Stream, err error) {
	s.client.getAccessToken("data")
	return s.create(dataset)
}

// create When creating a stream, specify the DataSet properties (name and description) and as a convenience,
// the create stream API will create a DataSet for you.
// In addition, you can only have one stream open at a time. If you need to add additional data,
// we recommended adding more parts to the currently open stream or executing a commit of the open stream before creating a new stream.
// KNOWN LIMITATION
// The StreamAPI currently only allows you to import data to a DataSet created via the Stream API.
// For example, it is currently not supported to import data to a DataSet created by a Domo Connector.
// Definition
// POST https://api.domo.com/v1/streams
// Returns
// Returns a DataSet object when successful. The returned object will have DataSet attributes based on
// the information that was provided when DataSet was created from the Stream created.
func (s *StreamService) create(dataSet string) (data *Stream, err error) {

	url := fmt.Sprintf("%s/v1/streams", baseURL)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	body := strings.NewReader(dataSet)
	bodyBytes, statusCode, err := s.client.genericPOST(url, body, header)

	if err != nil {
		return data, fmt.Errorf(
			"Unable to create stream from Domo API %s",
			err,
		)
	}

	logger(fmt.Sprintf("createStream Status Code : %d", statusCode))
	err = json.Unmarshal(bodyBytes, &data)

	bodyString := string(bodyBytes)

	logger("createStream" + bodyString)

	if err != nil {
		err = fmt.Errorf("Unable to unmarshal create response %s", err)
	}

	return
}

/* ------------  TODO ----------------*/
// Delete
// Updates the specified Stream’s metadata by providing values to parameters passed.
// Definition
// PATCH https://api.domo.com/v1/streams/{STREAM_ID}
// Returns
// Returns a full DataSet object of the Stream.

// Delete Deletes a Stream from your Domo instance. This does not a delete the associated DataSet.
//
// Returns a Stream object and parameter of success or error based on whether the Stream ID being valid.
func (s *StreamService) Delete(streamID int) error {

	s.client.getAccessToken("data")
	return s.delete(streamID)
}

// deleteStream Deletes a Stream from your Domo instance. This does not a delete the associated DataSet.
// Definition
// DELETE https://api.domo.com/v1/streams/{STREAM_ID}
// Returns
// Returns a Stream object and parameter of success or error based on whether the Stream ID being valid.
func (s *StreamService) delete(streamID int) error {
	var err error
	url := fmt.Sprintf("%s/v1/streams/%d", baseURL, streamID)
	statusCode, err := s.client.genericDELETE(url, nil)

	if err != nil {
		err = fmt.Errorf(
			"Unable to delete stream from Domo API %s",
			err,
		)
	}

	if statusCode != 204 {
		// it failed need to do something
		return fmt.Errorf("Failed to delete stream with ID %d", streamID)
	}
	return err
}

// List Get a list of all Streams for a specific DataSet.
//
// Returns all Stream objects that meet argument criteria from original request.
func (s *StreamService) List(ownerID int) (streamlist string, err error) {
	s.client.getAccessToken("data")
	streamlist, err = s.list(ownerID)
	return
}

// list Get a list of all Streams for a specific DataSet.
// Definition
// GET https://api.domo.com/v1/streams
// Returns
// Returns all Stream objects that meet argument criteria from original request.
func (s *StreamService) list(ownerID int) (streamlist string, err error) {
	url := fmt.Sprintf("%s/v1/streams/search?q=dataSource.owner.id:%d&fields=all", baseURL, ownerID)
	bodyBytes, _, err := s.client.genericGET(url, nil)

	if err != nil {
		return streamlist, fmt.Errorf(
			"Unable to list stream from Domo API %s",
			err,
		)
	}

	buf := new(bytes.Buffer)
	streamlist = fmt.Sprintf("%s", json.Indent(buf, []byte(bodyBytes), "", "  "))

	if err != nil {
		err = fmt.Errorf("Unable to get list %s", err)
	}

	return
}

/* ------------  TODO ----------------*/
// RetrieveStreamExecution Import data into a DataSet in your Domo instance. This request will replace the data currently in the DataSet.
// Definition
// GET https://api.domo.com/v1/streams/{STREAM_ID}/executions/{EXECUTION_ID}
// Returns
// Returns a subset fields of a Stream's object.

// createStreamExecution When you’re ready to upload data to your DataSet via a Stream,
// you first tell Domo that you’re ready to start sending data by creating an Execution.
// Definition
// POST https://api.domo.com/v1/streams/{STREAM_ID}/executions
// Returns
// Returns a subset of the stream object.
func (s *StreamService) createStreamExecution(streamID int) (data Execution, err error) {

	url := fmt.Sprintf("%s/v1/streams/%d/executions", baseURL, streamID)
	bodyBytes, _, err := s.client.genericPOST(url, nil, nil)

	if err != nil {
		return data, fmt.Errorf(
			"Unable to create stream execution from Domo API %s",
			err,
		)
	}

	err = json.Unmarshal(bodyBytes, &data)

	if err != nil {
		err = fmt.Errorf("Unable to post stream execution %s", err)
	}

	return

}

/* ------------  TODO ----------------*/
// ListStreamExecutions Retrieve a policy from a DataSet within Domo. A DataSet is required for a PDP policy to exist.
// Definition
// GET https://api.domo.com/v1/streams/{STREAM_ID}/executions
// Returns
// Returns a subset of the Stream execution object from the specified Stream.

// uploadDataPart Creates a data part within the Stream execution to upload chunks of rows to the DataSet.
// The calling client should keep track of parts and order them accordingly in an increasing sequence.
// If a part upload fails, retry the upload as all parts must be present before committing the stream execution.
// Definition
// PUT https://api.domo.com/v1/streams/{STREAM_ID}/executions/{EXECUTION_ID}/part/{PART_ID}
// Returns
// Returns a subset of a stream object and a parameter of success or error based on whether the data part within
// the stream execution being successful.
func (s *StreamService) uploadDataPart(streamID int, payload string, executionID int) error {
	var err error
	url := fmt.Sprintf("%s/v1/streams/%d/executions/%d/part/1", baseURL, streamID, executionID)

	header := make(map[string]string)
	header["Content-Type"] = "text/csv"

	body := strings.NewReader(payload)
	bodyBytes, _, err := s.client.genericPUT(url, body, header)

	if err != nil {
		return fmt.Errorf("Unable to put uploaddatapart %s", err)
	}

	if logging {
		buf := new(bytes.Buffer)
		json.Indent(buf, []byte(bodyBytes), "", "  ")
		fmt.Println(buf)
	}

	return err
}

// commitStreamExecution Commits stream execution to import combined set of data parts that have been successfully uploaded.
// KNOWN LIMITATION
// The Stream API only supports the ability to execute a “commit” every 15 minutes.
// Definition
// PUT https://api.domo.com/v1/streams/{STREAM_ID}/executions/{EXECUTION_ID}/commit
// Returns
// Returns a subset of a stream object and a parameter of success or error based on
// whether the stream execution successfully committed to Domo.
func (s *StreamService) commitStreamExecution(streamID int, executionID int) error {
	var err error
	url := fmt.Sprintf("%s/v1/streams/%d/executions/%d/commit", baseURL, streamID, executionID)
	_, _, err = s.client.genericPUT(url, nil, nil)

	if err != nil {
		err = fmt.Errorf("Unable to put commitStreamExecution %s", err)
	}

	return err
}

// abortStreamExecution If needed during an execution, aborts an entire Stream execution.
// BEST PRACTICES
// To abort the current stream execution within a Stream, simply identify the Stream’s ID within request
// Definition
// PUT https://api.domo.com/v1/streams/{STREAM_ID}/executions/{EXECUTION_ID}/abort
// Returns
// Returns a parameter of success or error based on whether the Stream ID being valid.
func (s *StreamService) abortStreamExecution(streamID int, executionID int) error {
	var err error
	url := fmt.Sprintf("%s/v1/streams/%d/executions/%d/abort", baseURL, streamID, executionID)
	_, statusCode, err := s.client.genericPUT(url, nil, nil)

	if statusCode != 204 {
		// it failed need to do something
		err = fmt.Errorf("Failed to abort stream")
	}

	return err
}

//UploadStringToStream sends data to Domo streamAPI
func (s *StreamService) UploadStringToStream(streamID int, payload string) error {
	var err error

	s.client.getAccessToken("data")

	thisExecutionID, err := s.createStreamExecution(streamID)

	if err != nil {
		return err
	}

	err = s.uploadDataPart(streamID, payload, thisExecutionID.ID)

	if err != nil {
		return fmt.Errorf("Failed to upload file")
	}

	return s.commitStreamExecution(streamID, thisExecutionID.ID)

}

// Get get a stream by name
// Get a list of all Streams for a specific DataSet.
func (s *StreamService) Get(streamname string) (list *StreamList, err error) {
	s.client.getAccessToken("data")
	return s.get(streamname)
}

//get List streams
// Get a list of all Streams for a specific DataSet.
func (s *StreamService) get(streamname string) (data *StreamList, err error) {

	url := fmt.Sprintf("%s/v1/streams/search?q=dataSource.name:%s", baseURL, streamname)
	bodyBytes, status, err := s.client.genericGET(url, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to get from domo %s", err)
	} else if status == http.StatusUnauthorized {
		return nil, errors.New("Unable to get from domo due to authorization failure")
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshal '%s' response %s", string(bodyBytes), err)
	}

	return
}
