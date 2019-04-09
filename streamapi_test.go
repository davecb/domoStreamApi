package domo

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CreateTestClient(d Doer) *Client {
	client := New(
		"<clientID>",
		"<secret>",
	)
	client.myDoer = d
	return client
}

func TestClient_RetrieveStream(t *testing.T) {
	t1, _ := time.Parse("2006-01-02T15:04:05Z", "2018-02-06T09:37:11Z")

	s := Stream{
		ID: 73,
		DataSet: StreamDataSet{
			ID:          "aa4d2db5-8504-49d8-9816-93b4861ce8af",
			Name:        "kafka_jl_purchase",
			Description: "Purchase checkout events from Kafka, there is no checking for loss or duplication",
			Rows:        99267,
			Columns:     7,
			Owner: Owner{
				ID:   915083994,
				Name: "Chris Joyce",
			},
			DataCurrentAt: t1,
			CreatedAt:     t1,
			UpdatedAt:     t1,
		},
		UpdateMethod: "APPEND",
		CreatedAt:    t1,
		ModifiedAt:   t1,
	}

	type fields struct {
		myDoer Doer
	}
	type args struct {
		streamID int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStream Stream
		wantErr    bool
	}{
		{
			name: "stream by id",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"id": 73,
						"dataSet": {
						  "id": "aa4d2db5-8504-49d8-9816-93b4861ce8af",
						  "name": "kafka_jl_purchase",
						  "description": "Purchase checkout events from Kafka, there is no checking for loss or duplication",
						  "rows": 99267,
						  "columns": 7,
						  "owner": {
							"id": 915083994,
							"name": "Chris Joyce"
						  },
						  "dataCurrentAt": "2018-02-06T09:37:11Z",
						  "createdAt": "2018-02-06T09:37:11Z",
						  "updatedAt": "2018-02-06T09:37:11Z",
						  "pdpEnabled": false
						},
						"updateMethod": "APPEND",
						"lastExecution": {
						  "id": 498,
						  "startedAt": "2018-02-06T09:37:11Z",
						  "endedAt": "2018-02-06T09:37:11Z",
						  "currentState": "SUCCESS",
						  "createdAt": "2018-02-06T09:37:11Z",
						  "modifiedAt": "2018-02-06T09:37:11Z",
						  "updateMethod": "APPEND"
						},
						"lastSuccessfulExecution": {
						  "id": 498,
						  "startedAt": "2018-02-06T09:37:11Z",
						  "endedAt": "2018-02-06T09:37:11Z",
						  "currentState": "SUCCESS",
						  "createdAt": "2018-02-06T09:37:11Z",
						  "modifiedAt": "2018-02-06T09:37:11Z",
						  "updateMethod": "APPEND"
						},
						"createdAt": "2018-02-06T09:37:11Z",
						"modifiedAt": "2018-02-06T09:37:11Z"
					  }`,
				},
			},
			args:       args{streamID: 0},
			wantStream: s,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			gotStream, err := d.Stream.Retrieve(tt.args.streamID)
			assert.Equal(t, gotStream, tt.wantStream, "Bad reply")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}

func TestClient_retrieveStream(t *testing.T) {
	t1, _ := time.Parse("2006-01-02T15:04:05Z", "2018-02-06T09:37:11Z")

	s := Stream{
		ID: 73,
		DataSet: StreamDataSet{
			ID:          "aa4d2db5-8504-49d8-9816-93b4861ce8af",
			Name:        "kafka_jl_purchase",
			Description: "Purchase checkout events from Kafka, there is no checking for loss or duplication",
			Rows:        99267,
			Columns:     7,
			Owner: Owner{
				ID:   915083994,
				Name: "Chris Joyce",
			},
			DataCurrentAt: t1,
			CreatedAt:     t1,
			UpdatedAt:     t1,
		},
		UpdateMethod: "APPEND",
		CreatedAt:    t1,
		ModifiedAt:   t1,
	}

	type fields struct {
		myDoer Doer
	}
	type args struct {
		streamID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Stream
		wantErr bool
	}{
		{
			name: "stream by id",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"id": 73,
						"dataSet": {
						  "id": "aa4d2db5-8504-49d8-9816-93b4861ce8af",
						  "name": "kafka_jl_purchase",
						  "description": "Purchase checkout events from Kafka, there is no checking for loss or duplication",
						  "rows": 99267,
						  "columns": 7,
						  "owner": {
							"id": 915083994,
							"name": "Chris Joyce"
						  },
						  "dataCurrentAt": "2018-02-06T09:37:11Z",
						  "createdAt": "2018-02-06T09:37:11Z",
						  "updatedAt": "2018-02-06T09:37:11Z",
						  "pdpEnabled": false
						},
						"updateMethod": "APPEND",
						"lastExecution": {
						  "id": 498,
						  "startedAt": "2018-02-06T09:37:11Z",
						  "endedAt": "2018-02-06T09:37:11Z",
						  "currentState": "SUCCESS",
						  "createdAt": "2018-02-06T09:37:11Z",
						  "modifiedAt": "2018-02-06T09:37:11Z",
						  "updateMethod": "APPEND"
						},
						"lastSuccessfulExecution": {
						  "id": 498,
						  "startedAt": "2018-02-06T09:37:11Z",
						  "endedAt": "2018-02-06T09:37:11Z",
						  "currentState": "SUCCESS",
						  "createdAt": "2018-02-06T09:37:11Z",
						  "modifiedAt": "2018-02-06T09:37:11Z",
						  "updateMethod": "APPEND"
						},
						"createdAt": "2018-02-06T09:37:11Z",
						"modifiedAt": "2018-02-06T09:37:11Z"
					  }`,
				},
			},
			args:    args{streamID: 0},
			want:    s,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			got, err := d.Stream.retrieve(tt.args.streamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.retrieveStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.retrieveStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_createStreamExecution(t *testing.T) {

	t1, _ := time.Parse("2006-01-02T15:04:05Z", "2018-02-06T09:37:11Z")

	type fields struct {
		myDoer Doer
	}
	type args struct {
		streamID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Execution
		wantErr bool
	}{
		{
			name: "executions commit",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"id": 4,
						"startedAt": "2018-02-06T09:37:11Z",
						"currentState": "ACTIVE",
						"updateMethod": "REPLACE"
					  }`,
				},
			},
			args:    args{streamID: 0},
			want:    Execution{ID: 4, StartedAt: t1, CurrentState: "ACTIVE", UpdateMethod: "REPLACE"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			got, err := d.Stream.createStreamExecution(tt.args.streamID)
			assert.Equal(t, got, tt.want, "Bad reply")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}

func TestClient_listStreams(t *testing.T) {
	t.Skip("Test not yet working")
	type fields struct {
		myDoer Doer
	}
	type args struct {
		ownerID int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStreamlist string
		wantErr        bool
	}{
		{
			name: "list by owner",
			fields: fields{
				testDoer{
					responseCode: 200,
					response:     "[{\nid\n: 72,\ndataSet\n: {\nid\n: \ned8d30cb-1ad5-4eb1-9c5c-b13324f6c5db\n,\nname\n: \nkafka_jl_customer\n,\ndescription\n: \nCustomer events from Kafka, there is no checking for loss or duplication\n,\nrows\n: 18877,\ncolumns\n: 7,\nowner\n: {\nid\n: 915083994,\nname\n: \nChris Joyce\n},\ndataCurrentAt\n: \n2018-02-06T09:37:08Z\n,\ncreatedAt\n: \n2018-02-01T01:46:08Z\n,\nupdatedAt\n: \n2018-02-06T09:37:08Z\n,\npdpEnabled\n: false},\nupdateMethod\n: \nAPPEND\n,\nlastExecution\n: {\nid\n: 498,\nstartedAt\n: \n2018-02-06T09:37:04Z\n,\nendedAt\n: \n2018-02-06T09:37:08Z\n,\ncurrentState\n: \nSUCCESS\n,\ncreatedAt\n: \n2018-02-06T09:37:04Z\n,\nmodifiedAt\n: \n2018-02-06T09:37:08Z\n,\nupdateMethod\n: \nAPPEND\n},\nlastSuccessfulExecution\n: {\nid\n: 498,\nstartedAt\n: \n2018-02-06T09:37:04Z\n,\nendedAt\n: \n2018-02-06T09:37:08Z\n,\ncurrentState\n: \nSUCCESS\n,\ncreatedAt\n: \n2018-02-06T09:37:04Z\n,\nmodifiedAt\n: \n2018-02-06T09:37:08Z\n,\nupdateMethod\n: \nAPPEND\n},\ncreatedAt\n: \n2018-02-01T01:46:10Z\n,\nmodifiedAt\n: \n2018-02-06T09:37:08Z\n},{\nid\n: 73,\ndataSet\n: {\nid\n: \naa4d2db5-8504-49d8-9816-93b4861ce8af\n,\nname\n: \nkafka_jl_purchase\n,\ndescription\n: \nPurchase checkout events from Kafka, there is no checking for loss or duplication\n,\nrows\n: 99267,\ncolumns\n: 7,\nowner\n: {\nid\n: 915083994,\nname\n: \nChris Joyce\n},\ndataCurrentAt\n: \n2018-02-06T09:37:11Z\n,\ncreatedAt\n: \n2018-02-01T01:46:12Z\n,\nupdatedAt\n: \n2018-02-06T09:37:12Z\n,\npdpEnabled\n: false},\nupdateMethod\n: \nAPPEND\n,\nlastExecution\n: {\nid\n: 498,\nstartedAt\n: \n2018-02-06T09:37:07Z\n,\nendedAt\n: \n2018-02-06T09:37:12Z\n,\ncurrentState\n: \nSUCCESS\n,\ncreatedAt\n: \n2018-02-06T09:37:07Z\n,\nmodifiedAt\n: \n2018-02-06T09:37:12Z\n,\nupdateMethod\n: \nAPPEND\n},\nlastSuccessfulExecution\n: {\nid\n: 498,\nstartedAt\n: \n2018-02-06T09:37:07Z\n,\nendedAt\n: \n2018-02-06T09:37:12Z\n,\ncurrentState\n: \nSUCCESS\n,\ncreatedAt\n: \n2018-02-06T09:37:07Z\n,\nmodifiedAt\n: \n2018-02-06T09:37:12Z\n,\nupdateMethod\n: \nAPPEND\n},\ncreatedAt\n: \n2018-02-01T01:46:13Z\n,\nmodifiedAt\n: \n2018-02-06T09:37:12Z\n}]",
				},
			},
			args:    args{ownerID: 12345},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			gotStreamlist, err := d.Stream.list(tt.args.ownerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stream.list() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStreamlist != tt.wantStreamlist {
				t.Errorf("Stream.list() = %v, want %v", gotStreamlist, tt.wantStreamlist)
			}
		})
	}
}

func TestClient_commitStreamExecution(t *testing.T) {

	type fields struct {
		myDoer Doer
	}
	type args struct {
		streamID    int
		executionID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "executions commit",
			fields: fields{
				testDoer{
					responseCode: 200,
					response:     "{\"id\": 4,\"startedAt\": \"2018-02-06T09:47:33Z\"currentState\": \"ACTIVE\"rows\": 0,\"bytes\": 0,\"createdAt\": \"2018-02-06T09:47:33Z\"modifiedAt\": \"2018-02-06T09:47:33Z\"}",
				},
			},
			args:    args{streamID: 0, executionID: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			err := d.Stream.commitStreamExecution(tt.args.streamID, tt.args.executionID)
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}

func TestClient_abortStreamExecution(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		streamID    int
		executionID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "executions abort",
			fields: fields{
				testDoer{
					responseCode: 200,
					response:     "{\"id\": 5,\"startedAt\": \"2018-02-06T09:48:47Z\"endedAt\": \"2018-02-06T09:49:21Z\"currentState\": \"ERROR\"createdAt\": \"2018-02-06T09:48:47Z\"modifiedAt\": \"2018-02-06T09:48:47Z\"updateMethod\": \"REPLACE\"}",
				},
			},
			args:    args{streamID: 0, executionID: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			err := d.Stream.abortStreamExecution(tt.args.streamID, tt.args.executionID)
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")

		})
	}
}
