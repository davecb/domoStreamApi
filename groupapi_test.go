package domo

import (
	"reflect"
	"testing"
)

func TestGroupService_Retrieve(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		groupID int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantGroup Group
		wantErr   bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{  
						"id":180920085,
						"name":"Content",
						"active":true,
						"memberCount":0,
						"creatorId":"2",
						"default":false
					 }`,
				},
			},
			args: args{groupID: 180920085},
			wantGroup: Group{
				ID:          180920085,
				Name:        "Content",
				Active:      true,
				MemberCount: 0,
				CreatorID:   "2",
				Default:     false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateTestClient(tt.fields.myDoer)
			gotGroup, err := g.Group.Retrieve(tt.args.groupID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroup, tt.wantGroup) {
				t.Errorf("GroupService.Retrieve() = %v, want %v", gotGroup, tt.wantGroup)
			}
		})
	}
}

// func ExampleGroupService_Retrieve() {

// 	groupID := 180920085

// 	myDoer := testDoer{
// 		responseCode: 200,
// 		response: `{
// 			"id":180920085,
// 			"name":"Content",
// 			"active":true,
// 			"memberCount":0,
// 			"creatorId":"2",
// 			"default":false
// 		 }`,
// 	}
// 	g := CreateTestClient(myDoer)

// 	gotGroup, err := g.Group.Retrieve(groupID)

// 	if err != nil {
// 		fmt.Println(gotGroup)
// 	}

// 	// Output: .....
// }

func TestGroupService_List(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name           string
		fields         fields
		wantUserGroups Groups
		wantErr        bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `[{  
						"id":650589601,
						"name":"US East Division",
						"memberCount":2,
						"default":false
					 }]`,
				},
			},
			wantUserGroups: Groups{
				{ID: 650589601,
					Name:        "US East Division",
					MemberCount: 2,
					Default:     false},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateTestClient(tt.fields.myDoer)
			gotUserGroups, err := g.Group.List()

			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUserGroups, tt.wantUserGroups) {
				t.Errorf("GroupService.List() = %v, want %v", gotUserGroups, tt.wantUserGroups)
			}
		})
	}
}

func TestGroupService_Find(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantGroupID int
		wantErr     bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `[  
						{  
						   "id":180920085,
						   "name":"Content",
						   "memberCount":0,
						   "default":false
						},
						{  
						   "id":650589601,
						   "name":"US East Division",
						   "memberCount":2,
						   "default":false
						}
					 ]`,
				},
			},
			args:        args{name: "Content"},
			wantGroupID: 180920085,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateTestClient(tt.fields.myDoer)
			gotGroupID, err := g.Group.Find(tt.args.name)

			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotGroupID != tt.wantGroupID {
				t.Errorf("GroupService.Find() = %v, want %v", gotGroupID, tt.wantGroupID)
			}
		})
	}
}

func TestGroupService_AddUser(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		groupID int
		userID  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{
					responseCode: 204,
				},
			},
			args:    args{groupID: 180920085, userID: 388908335},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateTestClient(tt.fields.myDoer)

			if err := g.Group.AddUser(tt.args.groupID, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("GroupService.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGroupService_ListUsers(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		groupID int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantGroupUsers GroupUsers
		wantErr        bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{
					response:     `[27,1304116356]`,
					responseCode: 200,
				},
			},
			args:           args{groupID: 180920085},
			wantGroupUsers: GroupUsers{27, 1304116356},
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateTestClient(tt.fields.myDoer)
			gotGroupUsers, err := g.Group.ListUsers(tt.args.groupID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GroupService.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroupUsers, tt.wantGroupUsers) {
				t.Errorf("GroupService.ListUsers() = %v, want %v", gotGroupUsers, tt.wantGroupUsers)
			}
		})
	}
}

func TestGroupService_RemoveUser(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		groupID int
		userID  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Basic Test",
			fields: fields{
				testDoer{
					response:     ``,
					responseCode: 204,
				},
			},
			args:    args{groupID: 180920085, userID: 388908335},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateTestClient(tt.fields.myDoer)

			if err := g.Group.RemoveUser(tt.args.groupID, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("GroupService.RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
