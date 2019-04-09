package domo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageService_RetrievePage(t *testing.T) {

	type fields struct {
		myDoer Doer
	}
	type args struct {
		pageID int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantPage Page
		wantErr  bool
		err      error
	}{
		{
			name: "Page found",
			fields: fields{
				testDoer{
					responseCode: 204,
					response: `{
						"id":333576354,
						"ownerId":804729568,
						"name":"Test Page One",
						"locked":false
						}`,
				},
			},
			args: args{pageID: 333576354},
			wantPage: Page{
				Name:    "Test Page One",
				ID:      333576354,
				OwnerID: 804729568,
				Locked:  false,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Page not found",
			fields: fields{
				testDoer{
					responseCode: 404,
					response: `{
						"status": 404,
						"statusReason": "Not Found",
						"toe": "8VGVLIN8CI-EODB4-G3CUS"
					  }`,
				},
			},
			args:     args{pageID: 0},
			wantPage: Page{},
			wantErr:  true,
			err:      fmt.Errorf("404 Not Found PageID: 0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateTestClient(tt.fields.myDoer)
			got, err := p.Page.Retrieve(tt.args.pageID)
			assert.Equal(t, tt.wantPage, got, "Bad reply")
			assert.Equal(t, tt.wantPage.ID, got.ID, "Incorrect Page")

			if tt.wantErr {
				assert.Equal(t, err, tt.err, "Incorrect error")
			}
		})
	}
}

// This test needs work
func TestPageService_ListPages(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name      string
		fields    fields
		wantPages Pages
		wantErr   bool
	}{
		{
			name: "Page list",
			fields: fields{
				testDoer{
					responseCode: 204,
					response: `[
						{
							"id":968425465,
							"name":"Developer Home",
						},
						{
							"id":842042130,
							"name":"Sample DataSets + Cards",
						},
						{
							"id":333576354,
							"name":"Test Page One",
						}
						]`,
				},
			},
			wantPages: nil,

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := CreateTestClient(tt.fields.myDoer)
			got, _ := p.Page.List()
			assert.Equal(t, tt.wantPages, got, "Bad reply")

			if tt.wantErr {
				//assert.Equal(t, err, tt.err, "Incorrect error")
			}

			// p := &PageService{
			// 	client: tt.fields.client,
			// }
			// gotPages, err := p.ListPages()
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("PageService.ListPages() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(gotPages, tt.wantPages) {
			// 	t.Errorf("PageService.ListPages() = %v, want %v", gotPages, tt.wantPages)
			// }
		})
	}
}
