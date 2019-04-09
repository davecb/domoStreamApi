//Package domo access libary for Domo.com API's
//
// Domo’s does not officially supported this golang API libraries.
// Thay provide you a quick way to begin developing on Domo’s platform.
//
// The Dataset API allows you to create, import, export and manage DataSets and manage data permissions for DataSets within Domo.
// The DataSet API should be used to create and update small DataSets that occasionally need their data updated.
// For creating and updating massive, constantly changing, or rapidly growing DataSets, the Stream API is recommended.
//
// The Group API manages group objects allow you to manage a group and users associated to a group.
// Groups allow you to set access rights, send Buzz messages,
// or share content that stays consistent even when the group members may change.
// The API allows you to create, delete, retrieve a user or a list of users, and update user information.
//
// The Page API manages the page object is a screen where you can view a “collection” of data,
// which is typically displayed in cards. You use a page to organize, manage,
// and share content to other users in Domo. Pages allow you to send external reports,
// create holistic filters across all metrics within the page, or have conversations
// in Domo’s Buzz tool about the data associated to the entire page.
// The Page API allows you to create, delete,  retrieve a page or a list of pages,
// and update page information and content within a page.
//
// The Stream API allows you to automate the creation of new DataSets in your Domo Warehouse,
// featuring an accelerated upload Stream. A Domo Stream expedites uploads by dividing your data into parts,
// and uploading all of these parts simultaneously.
// BEST PRACTICES
// This API should be used to create and update massive, constantly changing, or rapidly growing DataSets.
// For creating and updating smaller DataSets that occasionally need data updated, leverage the DataSet API.
//
// The User API manages the User objects allow you to manage a user and the user’s attributes
// such as a department, phone number, employee number, email, and username.
// The API allows you to create, delete, retrieve a user or a list of users, and update user information.
package domo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Client domo client object
//
// Only holds key information that is needed to user Domo's API's
type Client struct {
	clientID    string
	secret      string
	accessToken string
	expiresIn   time.Time
	myDoer      Doer

	service

	//Services
	DataSet *DataSetService
	Stream  *StreamService
	Page    *PageService
	User    *UserService
	Group   *GroupService
}

type service struct {
	client *Client
}

var logging bool
var debuglog bool
var baseURL string

//New Domo Client
func New(clientID string, secret string) *Client {
	d := Client{}
	d.clientID = clientID
	d.secret = secret
	baseURL = "https://api.domo.com"
	d.myDoer = http.DefaultClient
	d.service.client = &d

	d.DataSet = (*DataSetService)(&d.service)
	d.Stream = (*StreamService)(&d.service)
	d.Page = (*PageService)(&d.service)
	d.User = (*UserService)(&d.service)
	d.Group = (*GroupService)(&d.service)

	return &d
}

// SetLogging set the logger on or off
//
// Logging is off by default
func (d *Client) SetLogging(setlogging bool) {

	if setlogging {
		logging = true
	} else {
		logging = false
	}
}

// SetDebugLogging set the logger to include debug
//
// Debug is off by default
func (d *Client) SetDebugLogging(setlogging bool) {

	if setlogging {
		debuglog = true
	} else {
		debuglog = false
	}
}

// debuglogger method to handle debug logging
//
// All log records with be prefixed with [DomoClient]
func logdebug(logthis string) {
	if debuglog {
		fmt.Printf("[DomoClient Debug] %s\n", logthis)
	}
}

// logger method to handle logging
//
// All log records with be prefixed with [DomoClient]
func logger(logthis string) {
	if logging {
		fmt.Printf("[DomoClient] %s\n", logthis)
	}
}

//GetToken To interact with Domo’s APIs through OAuth security,
// you will need to obtain authorization and authentication. In order to access Domo APIs,
// a user must first be authenticated (prove that they are whom they say they are) through a client ID and client secret.
// Once a user has been authenticated, they can then create an access token to authorize what the scope
// of functionality will be available for each API for that specific access token.
//
// Returns a oAuth token
func (d *Client) GetToken(scope string) string {
	d.getAccessToken(scope)
	return d.accessToken
}

//getAccessToken To interact with Domo’s APIs through OAuth security,
// you will need to obtain authorization and authentication. In order to access Domo APIs,
// a user must first be authenticated (prove that they are whom they say they are) through a client ID and client secret.
// Once a user has been authenticated, they can then create an access token to authorize what the scope
// of functionality will be available for each API for that specific access token.
// Definition
// GET https://api.domo.com/oauth/token
// Returns
// Returns a oAuth token
func (d *Client) getAccessToken(scope string) error {
	var err error

	now := time.Now()
	diff := d.expiresIn.Sub(now).Seconds()
	remaining := int64(diff)

	if remaining < 60 {

		d.accessToken = ""
		url := fmt.Sprintf("%s/oauth/token?grant_type=client_credentials&scope=%s", baseURL, scope)
		bodyBytes, statusCode, genericErr := d.genericGET(url, nil)
		if genericErr != nil {
			return genericErr
		}

		if statusCode != 200 {

			return fmt.Errorf("Unauthorized")
		}
		data := new(Access)
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			return err
		}

		now := time.Now()
		d.accessToken = data.AccessToken
		d.expiresIn = now.Add(time.Second * time.Duration(data.ExpiresIn))

	} else {
		logdebug("Token is valid for > 60 seconds")
	}
	return err
}

// Doer to make testing easer !
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

func (d *Client) genericRequest(url string, method string, body io.Reader, headers map[string]string) (bodyBytes []byte, statusCode int, err error) {

	logdebug(fmt.Sprintf("URL : %s", url))
	logdebug(fmt.Sprintf("Reduest Body : %v", body))

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	var noAccept = true
	if len(headers) > 0 {
		for key, value := range headers {
			if key == "Accept" {
				noAccept = false
			}
			req.Header.Set(key, value)
			logdebug(fmt.Sprintf("head : %s %s", key, value))
		}
	}

	if noAccept {
		req.Header.Set("Accept", "application/json")
		logdebug(fmt.Sprintf("head : %s %s", "Accept", "application/json"))
	}

	if d.accessToken == "" {
		req.SetBasicAuth(d.clientID, d.secret)
	} else {
		bearer := fmt.Sprintf("bearer %s", d.accessToken)
		req.Header.Set("Authorization", bearer)
		logdebug(fmt.Sprintf("head : %s %s", "Authorization", bearer))
	}

	resp, err := d.myDoer.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	logdebug(fmt.Sprintf("reply body : %s", string(bodyBytes)))

	statusCode = resp.StatusCode
	logdebug(fmt.Sprintf("status code : %d", statusCode))
	return
}

func (d *Client) genericGET(url string, headers map[string]string) (bodyBytes []byte, statusCode int, err error) {
	return d.genericRequest(url, "GET", nil, headers)
}

func (d *Client) genericPOST(url string, body io.Reader, headers map[string]string) (bodyBytes []byte, statusCode int, err error) {
	return d.genericRequest(url, "POST", body, headers)
}

func (d *Client) genericPUT(url string, body io.Reader, headers map[string]string) (bodyBytes []byte, statusCode int, err error) {
	return d.genericRequest(url, "PUT", body, headers)
}

func (d *Client) genericDELETE(url string, headers map[string]string) (statusCode int, err error) {
	_, statusCode, err = d.genericRequest(url, "DELETE", nil, headers)
	return
}

func bytesToErrorMessage(bodyBytes []byte) (errorMessage ErrorMessage, err error) {
	err = json.Unmarshal(bodyBytes, &errorMessage)
	return errorMessage, err
}
