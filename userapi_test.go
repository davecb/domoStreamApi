package domo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CheckRole(t *testing.T) {
	type args struct {
		role string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test cases.
		{
			name: "Admin",
			args: args{role: "Admin"},
			want: true,
		},
		{
			name: "Privileged",
			args: args{role: "Privileged"},
			want: true,
		},
		{
			name: "Participant",
			args: args{role: "Participant"},
			want: true,
		},
		{
			name: "Not A Role",
			args: args{role: "no role"},
			want: false,
		},
		{
			name: "Blank",
			args: args{role: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckRole(tt.args.role)
			assert.Equal(t, got, tt.want, "Incorret role")
		})
	}
}

func ExampleCheckRole() {
	isRoleOK := CheckRole("Participant")
	if isRoleOK {
		fmt.Println("Participant is a valid role")
	}
	// Output: Participant is a valid role
}

func TestClient_List(t *testing.T) {

	var users Users
	_ = json.Unmarshal([]byte("[{\"id\": 61014150,\"title\": \"Marketing Team\",\"email\": \"marketing@ozlotteries.com\",\"role\": \"Editor\",\"name\": \"Marketing\",\"location\": \"Brisbane\",\"roleId\": 3,\"createdAt\": \"1970-01-18T11:45:43.925Z\",\"updatedAt\": \"2018-02-01T23:42:52.843Z\"}]"), &users)

	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name      string
		fields    fields
		wantUsers Users
		wantErr   bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{responseCode: 200, response: "[{\"id\": 61014150,\"title\": \"Marketing Team\",\"email\": \"marketing@ozlotteries.com\",\"role\": \"Editor\",\"name\": \"Marketing\",\"location\": \"Brisbane\",\"roleId\": 3,\"createdAt\": \"1970-01-18T11:45:43.925Z\",\"updatedAt\": \"2018-02-01T23:42:52.843Z\"}]"},
			},
			wantUsers: users,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			gotUsers, err := d.User.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("Client.ListUsers() = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
}

func TestClient_Retrieve(t *testing.T) {
	var user User
	_ = json.Unmarshal([]byte("{\"id\": 27,\"email\": \"support@domo.com\",\"role\": \"Admin\",\"name\": \"DomoSupport\",\"timezone\": \"America/Denver\",\"roleId\": 1,\"createdAt\": \"1970-01-17T14:46:50.566Z\",\"updatedAt\": \"2016-11-01T04:35:37.243Z\",\"image\": \"https://jumbointeractive.domo.com/avatar/thumb/domo/27\",\"groups\": [{\"id\": 180920085,\"name\": \"Content\"},{\"id\": 1324037627,\"name\": \"Default\"}]}"), &user)

	type fields struct {
		myDoer Doer
	}
	type args struct {
		userid int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser User
		wantErr  bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{responseCode: 200, response: "{\"id\": 27,\"email\": \"support@domo.com\",\"role\": \"Admin\",\"name\": \"DomoSupport\",\"timezone\": \"America/Denver\",\"roleId\": 1,\"createdAt\": \"1970-01-17T14:46:50.566Z\",\"updatedAt\": \"2016-11-01T04:35:37.243Z\",\"image\": \"https://jumbointeractive.domo.com/avatar/thumb/domo/27\",\"groups\": [{\"id\": 180920085,\"name\": \"Content\"},{\"id\": 1324037627,\"name\": \"Default\"}]}"}},
			args:     args{userid: 27},
			wantUser: user,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CreateTestClient(tt.fields.myDoer)
			gotUser, err := d.User.Retrieve(tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RetrieveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Client.RetrieveUser() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserService_Find(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantUserID int
		wantErr    bool
	}{
		{
			name: "Paul Particpant",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `[
						{  
						   "id":388908335,
						   "title":"Chief Marketing Officer",
						   "email":"paul.participant@domosoftware.net",
						   "alternateEmail":"paul@email.com",
						   "role":"Participant",
						   "phone":"431.654.6548",
						   "name":"Paul Particpant",
						   "roleId":4,
						   "employeeNumber":1354,
						   "createdAt":"1970-01-18T13:19:16.479Z",
						   "updatedAt":"2018-01-24T01:34:25.531Z"
						},
						{  
						   "id":804729568,
						   "email":"tech_support@ozlotteries.com",
						   "role":"Admin",
						   "name":"Tech Support",
						   "roleId":1,
						   "createdAt":"1970-01-18T13:28:10.683Z",
						   "updatedAt":"2018-02-08T05:46:23.169Z"
						}
					 ]`,
				},
			},
			args:       args{name: "Paul Particpant"},
			wantUserID: 388908335,
			wantErr:    false,
		},
		{
			name: "Missing Mike",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `[
						{  
						   "id":804729568,
						   "email":"tech_support@ozlotteries.com",
						   "role":"Admin",
						   "name":"Tech Support",
						   "roleId":1,
						   "createdAt":"1970-01-18T13:28:10.683Z",
						   "updatedAt":"2018-02-08T05:46:23.169Z"
						}
					 ]`,
				},
			},
			args:       args{name: "Missing Mike"},
			wantUserID: 0,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := CreateTestClient(tt.fields.myDoer)
			gotUserID, err := u.User.Find(tt.args.name)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("UserService.Find() = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}
