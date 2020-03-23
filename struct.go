package domo

import "time"

//Groups strut
type Groups []struct {
	Default     bool   `json:"default"`
	ID          int    `json:"id"`
	MemberCount int    `json:"memberCount"`
	Name        string `json:"name"`
}

//Group group details
type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	MemberCount int    `json:"memberCount"`
	CreatorID   string `json:"creatorId"`
	Default     bool   `json:"default"`
}

//Datasets is improper superset of datasets
type Datasets []struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Rows          int       `json:"rows"`
	Columns       int       `json:"columns"`
	Owner         Owner     `json:"owner"`
	DataCurrentAt time.Time `json:"dataCurrentAt"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	PdpEnabled    bool      `json:"pdpEnabled"`
	Description   string    `json:"description,omitempty"`
}

//DatasetSummary summary only
type DatasetSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Rows        int    `json:"rows"`
	Columns     int    `json:"columns"`
	Description string `json:"description,omitempty"`
}

//StreamDataSet ..
type StreamDataSet struct {
	ID            string    `json:"id"`            // The id of the DataSet associated to the Stream
	Name          string    `json:"name"`          // The name of the DataSet associated to the Stream
	Description   string    `json:"description"`   // The description of the DataSet associated to the Stream
	Rows          int       `json:"rows"`          // The number of rows in the DataSet
	Columns       int       `json:"columns"`       // The number of columns in the DataSet's schema
	Owner         Owner     `json:"owner"`         // The owner of the stream's underlying DataSet
	DataCurrentAt time.Time `json:"dataCurrentAt"` // An ISO-8601 representation of the create date of the Stream
	CreatedAt     time.Time `json:"createdAt"`     // An ISO-8601 representation of the create date of the Stream
	UpdatedAt     time.Time `json:"updatedAt"`     // An ISO-8601 representation of the time the Stream was last updated
	PdpEnabled    bool      `json:"pdpEnabled"`
}

//Dataset information on a Dataset
type Dataset struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rows        int       `json:"rows"`
	Columns     int       `json:"columns,omitempty"`
	Schema      Schema    `json:"schema"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	PdpEnabled  bool      `json:"pdpEnabled,omitempty"`
	Policies    Policies  `json:"policies,omitempty"`
}

// Schema ...
type Schema struct {
	Columns Columns `json:"columns"`
}

// Columns ..
type Columns []struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// Policies ...
type Policies []struct {
	ID      int           `json:"id"`
	Type    string        `json:"type"`
	Name    string        `json:"name"`
	Filters Filters       `json:"filters"`
	Users   []int         `json:"users"`
	Groups  []interface{} `json:"groups"`
}

// Filters ...
type Filters []struct {
	Column   string   `json:"column"`
	Values   []string `json:"values"`
	Operator string   `json:"operator"`
	Not      bool     `json:"not"`
}

//Access ...
type Access struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	Customer    string `json:"customer"`
	Env         string `json:"env"`
	UserID      int    `json:"userId"`
	Role        string `json:"role"`
	Jti         string `json:"jti"`
}

//Page The page object
type Page struct {
	Name          string       `json:"name"`          // The name of the page
	ID            int          `json:"id"`            // The ID of the page
	ParentID      int          `json:"parentId"`      // The ID of the page that is higher in organizational hierarchy
	OwnerID       int          `json:"ownerId"`       // The ID of the page owner
	Locked        bool         `json:"locked"`        // Determines whether users (besides the page owner) can make updates to page or its content - the default value is false
	CollectionIDs int          `json:"collectionIds"` // The IDs of collections within a page
	CardIds       []int        `json:"cardIds"`       // The ID of all cards contained within the page
	Children      ChildrenPage `json:"children"`      // All pages that are considered "sub pages" in organizational hierarchy
	Visibility    Visibility   `json:"visibility"`    // Determines the access given to both individual users or groups within Domo
	UserIds       int          `json:"userIds"`       // The ID of the page
	GroupIDs      int          `json:"groupIds"`      // The ID of the page

}

// Pages List of pages
type Pages []struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Children []ChildrenPage `json:"children"`
}

// Visibility Determines the access given to both individual users or groups within Domo
type Visibility struct {
	UserIds  []int `json:"userIds"`  // The IDs of the users
	GroupIds []int `json:"groupIds"` // The IDs of the groups
}

// ChildrenPage All pages that are considered "sub pages" in organizational hierarchy
type ChildrenPage struct {
	Name string `json:"name"` // The name of the page
	ID   int    `json:"id"`   // The ID of the page
}

// Execution ...
type Execution struct {
	ID           int       `json:"id"` // ID of the Stream
	StartedAt    time.Time `json:"startedAt"`
	EndedAt      time.Time `json:"endedAt"`
	CurrentState string    `json:"currentState"`
	CreatedAt    time.Time `json:"createdAt"`  // An ISO-8601 representation of the create date of the Stream
	ModifiedAt   time.Time `json:"modifiedAt"` // An ISO-8601 representation of the time the Stream was last updated
	UpdateMethod string    `json:"updateMethod"`
}

//Stream ....
type Stream struct {
	ID           int           `json:"id"` //ID of the Stream
	DataSet      StreamDataSet `json:"dataSet"`
	UpdateMethod string        `json:"updateMethod"` // The data import behavior
	CreatedAt    time.Time     `json:"createdAt"`
	ModifiedAt   time.Time     `json:"modifiedAt"`
}

//Owner The owner of the underlying DataSet
type Owner struct {
	ID   int    `json:"id"`   // The ID of the owner of the stream's underlying DataSet
	Name string `json:"name"` // The name of the owner of the stream's underlying DataSet
}

//StreamList ....
type StreamList []struct {
	ID                      int           `json:"id"`
	DataSet                 StreamDataSet `json:"dataSet"`
	UpdateMethod            string        `json:"updateMethod"`
	LastExecution           Execution     `json:"lastExecution"`
	LastSuccessfulExecution Execution     `json:"lastSuccessfulExecution"`
	CreatedAt               time.Time     `json:"createdAt"`
	ModifiedAt              time.Time     `json:"modifiedAt"`
}

//User domo user object
type User struct {
	ID             int       `json:"id"`
	Title          string    `json:"title,omitempty"` //User's job title
	Email          string    `json:"email"`           //User's primary email used in profile
	Role           string    `json:"role"`            //The role of the user created (available roles are: 'Admin', 'Privileged', 'Participant')
	Name           string    `json:"name"`            //User's full name
	RoleID         int       `json:"roleId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Timezone       string    `json:"timezone,omitempty"` //Time zone used to display to user the system times throughout Domo application
	Phone          string    `json:"phone,omitempty"`    //Primary phone number of user
	Location       string    `json:"location,omitempty"` //Free text that can be used to define office location (e.g. City, State, Country)
	AlternateEmail string    `json:"alternateEmail"`     //User's secondary email in profile
	EmployeeNumber int       `json:"employeeNumber"`     //Employee number within company
	Image          string    `json:"image"`
	Groups         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
}

//Users domo users object
type Users []struct {
	ID        int       `json:"id"`
	Title     string    `json:"title,omitempty"` //User's job title
	Email     string    `json:"email"`           //User's primary email used in profile
	Role      string    `json:"role"`            //The role of the user created (available roles are: 'Admin', 'Privileged', 'Participant')
	Name      string    `json:"name"`            //User's full name
	RoleID    int       `json:"roleId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Timezone  string    `json:"timezone,omitempty"` //Time zone used to display to user the system times throughout Domo application
	Phone     string    `json:"phone,omitempty"`    //Primary phone number of user
	Location  string    `json:"location,omitempty"` //Free text that can be used to define office location (e.g. City, State, Country)
}

// ErrorMessage domo error reply
type ErrorMessage struct {
	Status       int    `json:"status"`
	StatusReason string `json:"statusReason"`
	Toe          string `json:"toe"`
}
