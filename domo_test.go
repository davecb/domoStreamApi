package domo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type wonkyReader struct{}

func (wr wonkyReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

type testDoer struct {
	response     string
	responseCode int
	http.Header
}

func (nd testDoer) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(nd.response))),
		StatusCode: nd.responseCode,
		Header:     nd.Header,
	}, nil
}

func ExampleNew() {
	d := New("ClientID", "secret")

	fmt.Printf("%v", d.clientID)
	// Output: ClientID
}

func Test_logger(t *testing.T) {
	type args struct {
		logthis string
		logger  bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic Logger",
			args: args{logthis: "Hello logger", logger: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger(tt.args.logthis)
		})
	}
}

func TestSetLogging(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
	}
	type args struct {
		setlogging bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Set logging on",
			args: args{setlogging: true},
			want: true,
		},
		{
			name: "Set logging off",
			args: args{setlogging: false},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
			}
			d.SetLogging(tt.args.setlogging)
			assert.Equal(t, logging, tt.want)
		})
	}
}

func TestClient_genericRequest(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		url     string
		method  string
		body    io.Reader
		headers map[string]string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantBodyBytes  []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				"<clientID>",
				"<secret>",
				"<accessToken>",
				time.Now(),
				testDoer{responseCode: 200, response: "Hellow World"},
			},
			args: args{
				"/bla",
				"GET",
				wonkyReader{},
				nil,
			},
			wantBodyBytes:  []byte("Hellow World"),
			wantStatusCode: 200,
			wantErr:        false,
		},
		{
			name: "No access token test",
			fields: fields{
				"<clientID>",
				"<secret>",
				"",
				time.Now(),
				testDoer{},
			},
			args: args{
				"/bla",
				"GET",
				wonkyReader{},
				nil,
			},
			wantBodyBytes:  []byte(""),
			wantStatusCode: 0,
			wantErr:        false,
		},
		{
			name: "Header Test",
			fields: fields{
				"<clientID>",
				"<secret>",
				"",
				time.Now(),
				testDoer{},
			},
			args: args{
				"/bla",
				"GET",
				wonkyReader{},
				map[string]string{"Accept": "Stuff"},
			},
			wantBodyBytes:  []byte(""),
			wantStatusCode: 0,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			gotBodyBytes, gotStatusCode, err := d.genericRequest(tt.args.url, tt.args.method, tt.args.body, tt.args.headers)
			assert.Equal(t, gotBodyBytes, tt.wantBodyBytes, "Bad reply")
			assert.Equal(t, gotStatusCode, tt.wantStatusCode, "Incorrect status code")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}
func TestClient_genericGET(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		url     string
		headers map[string]string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantBodyBytes  []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				"<clientID>",
				"<secret>",
				"<accessToken>",
				time.Now(),
				testDoer{},
			},
			args: args{
				"/bla",
				nil,
			},
			wantBodyBytes:  []byte(""),
			wantStatusCode: 0,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			gotBodyBytes, gotStatusCode, err := d.genericGET(tt.args.url, tt.args.headers)
			assert.Equal(t, gotBodyBytes, tt.wantBodyBytes, "Bad reply")
			assert.Equal(t, gotStatusCode, tt.wantStatusCode, "Incorrect status code")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}
func TestClient_genericPOST(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		url     string
		body    io.Reader
		headers map[string]string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantBodyBytes  []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				"<clientID>",
				"<secret>",
				"<accessToken>",
				time.Now(),
				testDoer{},
			},
			args: args{
				"/bla",
				wonkyReader{},
				nil,
			},
			wantBodyBytes:  []byte(""),
			wantStatusCode: 0,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			gotBodyBytes, gotStatusCode, err := d.genericPOST(tt.args.url, tt.args.body, tt.args.headers)
			assert.Equal(t, gotBodyBytes, tt.wantBodyBytes, "Bad reply")
			assert.Equal(t, gotStatusCode, tt.wantStatusCode, "Incorrect status code")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}
func TestClient_genericPUT(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		url     string
		body    io.Reader
		headers map[string]string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantBodyBytes  []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				"<clientID>",
				"<secret>",
				"<accessToken>",
				time.Now(),
				testDoer{},
			},
			args: args{
				"/bla",
				wonkyReader{},
				nil,
			},
			wantBodyBytes:  []byte(""),
			wantStatusCode: 0,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			gotBodyBytes, gotStatusCode, err := d.genericPUT(tt.args.url, tt.args.body, tt.args.headers)
			assert.Equal(t, gotBodyBytes, tt.wantBodyBytes, "Bad reply")
			assert.Equal(t, gotStatusCode, tt.wantStatusCode, "Incorrect status code")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}

func TestClient_genericDELETE(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		url     string
		headers map[string]string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "Basic Delete",
			fields: fields{
				"<clientID>",
				"<secret>",
				"<accessToken>",
				time.Now(),
				testDoer{},
			},
			args: args{
				"/bla",
				nil,
			},
			wantStatusCode: 0,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			gotStatusCode, err := d.genericDELETE(tt.args.url, tt.args.headers)
			assert.Equal(t, gotStatusCode, tt.wantStatusCode, "Incorrect status code")
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}

func TestClient_getAccessToken(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		scope string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Token",
			fields: fields{
				"Xd0ba0e7-4c6e-4ba0-9f4d-79d3668fbb5X",
				"X6354422a08ec4cfc1b41b6f0dd21ddfe897e083dbb4de0c384bd7c501d389cX",
				"",
				time.Now(),
				testDoer{responseCode: 200, response: "{\"access_token\": \"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiUHJpdmlsZWdlZCIsInNjb3BlIjpbImRhdGEiXSwiZG9tYWluIjoianVtYm9pbnRlcmFjdGl2ZS5kb21vLmNvbSIsImV4cCI6MTUxNzgzNzk0MywiZW52IjoicHJvZDMiLCJ1c2VySWQiOjkxNTA4Mzk5NCwianRpIjoiMjRjNzg5OTktY2E4ZS00ZWEwLWEyZGUtMjNiMTI5Nzg4YTA2IiwiY2xpZW50X2lkIjoiYWQwYmEwZTctNGM2ZS00YmEwLTlmNGQtNzlkMzY2OGZiYjUxIiwiY3VzdG9tZXIiOiJqdW1ib2ludGVyYWN0aXZlIn0.tzW9hWQAtZMmmVieCF6ha7gBvx6bFT0z8fAahA-fxmH1gB0yxAKk_LIk10gVARGBa-XgcZAq46X9HTKhVbkcKc1AxLn65ur1LLsmzeK52WHvpKM_k7oGxiot8DLofiNUM1Y00mrQpCwjU4CGHD7cMx_z6f64YNsD-acOvSiDUAZ7P55p-9ESqrLf-Z_Zx1SHCgDwg7Gw2DiWPJSGhvNmG47Sl20YpYKbLghawhrIhTqbHkYxDbJxbd50bumZXzXoCqJdlQOplAdZ0z4sPkXpIEFhVmw4nUnG2ul4z89aaWaEhMAfSoNwQaILLWnRph90XjTK-Pu5Y7ct5Qo4zqlx5Q\",\"token_type\": \"bearer\",\"expires_in\": 3599,\"scope\": \"data\",\"customer\": \"jumbointeractive\",\"env\": \"prod3\",\"userId\": 915083994,\"role\": \"Privileged\",\"domain\": \"jumbointeractive.domo.com\",\"jti\": \"24c78999-ca8e-4ea0-a2de-23b129788a06\"}"},
			},
			args:    args{scope: "data"},
			wantErr: false,
		},
		{
			name: "Invalid Token",
			fields: fields{
				"Xd0ba0e7-4c6e-4ba0-9f4d-79d3668fbb5F",
				"X6354422a08ec4cfc1b41b6f0dd21ddfe897e083dbb4de0c384bd7c501d389cX",
				"",
				time.Now(),
				testDoer{responseCode: 401, response: "{\"status\": 401,\"statusReason\": \"Unauthorized\",\"path\": \"/oauth/token\",\"message\": \"Bad credentials\",\"toe\": \"8UMPDBCRK3-P8IVL-E3RIU\"}"},
			},
			args:    args{scope: "data"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			err := d.getAccessToken(tt.args.scope)
			assert.NotEqual(t, err, tt.wantErr, "Bad error code")
		})
	}
}

func TestClient_GetToken(t *testing.T) {
	type fields struct {
		clientID    string
		secret      string
		accessToken string
		expiresIn   time.Time
		myDoer      Doer
	}
	type args struct {
		scope string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Valid Token",
			fields: fields{
				"Xd0ba0e7-4c6e-4ba0-9f4d-79d3668fbb5X",
				"X6354422a08ec4cfc1b41b6f0dd21ddfe897e083dbb4de0c384bd7c501d389cX",
				"",
				time.Now(),
				testDoer{responseCode: 200, response: "{\"access_token\": \"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiUHJpdmlsZWdlZCIsInNjb3BlIjpbImRhdGEiXSwiZG9tYWluIjoianVtYm9pbnRlcmFjdGl2ZS5kb21vLmNvbSIsImV4cCI6MTUxNzgzNzk0MywiZW52IjoicHJvZDMiLCJ1c2VySWQiOjkxNTA4Mzk5NCwianRpIjoiMjRjNzg5OTktY2E4ZS00ZWEwLWEyZGUtMjNiMTI5Nzg4YTA2IiwiY2xpZW50X2lkIjoiYWQwYmEwZTctNGM2ZS00YmEwLTlmNGQtNzlkMzY2OGZiYjUxIiwiY3VzdG9tZXIiOiJqdW1ib2ludGVyYWN0aXZlIn0.tzW9hWQAtZMmmVieCF6ha7gBvx6bFT0z8fAahA-fxmH1gB0yxAKk_LIk10gVARGBa-XgcZAq46X9HTKhVbkcKc1AxLn65ur1LLsmzeK52WHvpKM_k7oGxiot8DLofiNUM1Y00mrQpCwjU4CGHD7cMx_z6f64YNsD-acOvSiDUAZ7P55p-9ESqrLf-Z_Zx1SHCgDwg7Gw2DiWPJSGhvNmG47Sl20YpYKbLghawhrIhTqbHkYxDbJxbd50bumZXzXoCqJdlQOplAdZ0z4sPkXpIEFhVmw4nUnG2ul4z89aaWaEhMAfSoNwQaILLWnRph90XjTK-Pu5Y7ct5Qo4zqlx5Q\",\"token_type\": \"bearer\",\"expires_in\": 3599,\"scope\": \"data\",\"customer\": \"jumbointeractive\",\"env\": \"prod3\",\"userId\": 915083994,\"role\": \"Privileged\",\"domain\": \"jumbointeractive.domo.com\",\"jti\": \"24c78999-ca8e-4ea0-a2de-23b129788a06\"}"},
			},
			args: args{scope: "data"},
			want: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiUHJpdmlsZWdlZCIsInNjb3BlIjpbImRhdGEiXSwiZG9tYWluIjoianVtYm9pbnRlcmFjdGl2ZS5kb21vLmNvbSIsImV4cCI6MTUxNzgzNzk0MywiZW52IjoicHJvZDMiLCJ1c2VySWQiOjkxNTA4Mzk5NCwianRpIjoiMjRjNzg5OTktY2E4ZS00ZWEwLWEyZGUtMjNiMTI5Nzg4YTA2IiwiY2xpZW50X2lkIjoiYWQwYmEwZTctNGM2ZS00YmEwLTlmNGQtNzlkMzY2OGZiYjUxIiwiY3VzdG9tZXIiOiJqdW1ib2ludGVyYWN0aXZlIn0.tzW9hWQAtZMmmVieCF6ha7gBvx6bFT0z8fAahA-fxmH1gB0yxAKk_LIk10gVARGBa-XgcZAq46X9HTKhVbkcKc1AxLn65ur1LLsmzeK52WHvpKM_k7oGxiot8DLofiNUM1Y00mrQpCwjU4CGHD7cMx_z6f64YNsD-acOvSiDUAZ7P55p-9ESqrLf-Z_Zx1SHCgDwg7Gw2DiWPJSGhvNmG47Sl20YpYKbLghawhrIhTqbHkYxDbJxbd50bumZXzXoCqJdlQOplAdZ0z4sPkXpIEFhVmw4nUnG2ul4z89aaWaEhMAfSoNwQaILLWnRph90XjTK-Pu5Y7ct5Qo4zqlx5Q",
		},
		{
			name: "Invalid Token",
			fields: fields{
				"Xd0ba0e7-4c6e-4ba0-9f4d-79d3668fbb5F",
				"X6354422a08ec4cfc1b41b6f0dd21ddfe897e083dbb4de0c384bd7c501d389cX",
				"",
				time.Now(),
				testDoer{responseCode: 401, response: "{\"status\": 401,\"statusReason\": \"Unauthorized\",\"path\": \"/oauth/token\",\"message\": \"Bad credentials\",\"toe\": \"8UMPDBCRK3-P8IVL-E3RIU\"}"},
			},
			args: args{scope: "data"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Client{
				clientID:    tt.fields.clientID,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
				expiresIn:   tt.fields.expiresIn,
				myDoer:      tt.fields.myDoer,
			}
			got := d.GetToken(tt.args.scope)
			assert.Equal(t, got, tt.want, "Token error")
		})
	}
}

func ExampleClient_GetToken() {

	// Normal use would be
	// d := New("<clientID>", "<Secret>")
	// setting up client to avoid hitting real API
	d := &Client{
		clientID:    "Xd0ba0e7-4c6e-4ba0-9f4d-79d3668fbb5X",
		secret:      "X6354422a08ec4cfc1b41b6f0dd21ddfe897e083dbb4de0c384bd7c501d389cX",
		accessToken: "",
		expiresIn:   time.Now(),
		myDoer:      testDoer{responseCode: 200, response: "{\"access_token\": \"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiUHJpdmlsZWdlZCIsInNjb3BlIjpbImRhdGEiXSwiZG9tYWluIjoianVtYm9pbnRlcmFjdGl2ZS5kb21vLmNvbSIsImV4cCI6MTUxNzgzNzk0MywiZW52IjoicHJvZDMiLCJ1c2VySWQiOjkxNTA4Mzk5NCwianRpIjoiMjRjNzg5OTktY2E4ZS00ZWEwLWEyZGUtMjNiMTI5Nzg4YTA2IiwiY2xpZW50X2lkIjoiYWQwYmEwZTctNGM2ZS00YmEwLTlmNGQtNzlkMzY2OGZiYjUxIiwiY3VzdG9tZXIiOiJqdW1ib2ludGVyYWN0aXZlIn0.tzW9hWQAtZMmmVieCF6ha7gBvx6bFT0z8fAahA-fxmH1gB0yxAKk_LIk10gVARGBa-XgcZAq46X9HTKhVbkcKc1AxLn65ur1LLsmzeK52WHvpKM_k7oGxiot8DLofiNUM1Y00mrQpCwjU4CGHD7cMx_z6f64YNsD-acOvSiDUAZ7P55p-9ESqrLf-Z_Zx1SHCgDwg7Gw2DiWPJSGhvNmG47Sl20YpYKbLghawhrIhTqbHkYxDbJxbd50bumZXzXoCqJdlQOplAdZ0z4sPkXpIEFhVmw4nUnG2ul4z89aaWaEhMAfSoNwQaILLWnRph90XjTK-Pu5Y7ct5Qo4zqlx5Q\",\"token_type\": \"bearer\",\"expires_in\": 3599,\"scope\": \"data\",\"customer\": \"jumbointeractive\",\"env\": \"prod3\",\"userId\": 915083994,\"role\": \"Privileged\",\"domain\": \"jumbointeractive.domo.com\",\"jti\": \"24c78999-ca8e-4ea0-a2de-23b129788a06\"}"},
	}
	t := d.GetToken("data")

	fmt.Printf("%v", t)
	// Output: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiUHJpdmlsZWdlZCIsInNjb3BlIjpbImRhdGEiXSwiZG9tYWluIjoianVtYm9pbnRlcmFjdGl2ZS5kb21vLmNvbSIsImV4cCI6MTUxNzgzNzk0MywiZW52IjoicHJvZDMiLCJ1c2VySWQiOjkxNTA4Mzk5NCwianRpIjoiMjRjNzg5OTktY2E4ZS00ZWEwLWEyZGUtMjNiMTI5Nzg4YTA2IiwiY2xpZW50X2lkIjoiYWQwYmEwZTctNGM2ZS00YmEwLTlmNGQtNzlkMzY2OGZiYjUxIiwiY3VzdG9tZXIiOiJqdW1ib2ludGVyYWN0aXZlIn0.tzW9hWQAtZMmmVieCF6ha7gBvx6bFT0z8fAahA-fxmH1gB0yxAKk_LIk10gVARGBa-XgcZAq46X9HTKhVbkcKc1AxLn65ur1LLsmzeK52WHvpKM_k7oGxiot8DLofiNUM1Y00mrQpCwjU4CGHD7cMx_z6f64YNsD-acOvSiDUAZ7P55p-9ESqrLf-Z_Zx1SHCgDwg7Gw2DiWPJSGhvNmG47Sl20YpYKbLghawhrIhTqbHkYxDbJxbd50bumZXzXoCqJdlQOplAdZ0z4sPkXpIEFhVmw4nUnG2ul4z89aaWaEhMAfSoNwQaILLWnRph90XjTK-Pu5Y7ct5Qo4zqlx5Q
}
