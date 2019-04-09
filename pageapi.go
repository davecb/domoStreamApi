package domo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PageService Page API service
type PageService service

// Retrieve Retrieves the details of an existing page.
// Definition
// https://api.domo.com/v1/pages/{PAGE_ID}
// Returns
// Returns a page object if valid page ID was provided.
func (p *PageService) Retrieve(pageID int) (page Page, err error) {
	p.client.getAccessToken("dashboard")
	url := fmt.Sprintf("%s/v1/pages/%d", baseURL, pageID)
	bodyBytes, statuscode, err := p.client.genericGET(url, nil)

	if err != nil {
		return Page{}, fmt.Errorf(
			"Unable to Retrieve pages response from DOMO API %s",
			err,
		)
	}

	if statuscode == 404 {
		message, _ := bytesToErrorMessage(bodyBytes)
		err = fmt.Errorf("%d %s PageID: %d", message.Status, message.StatusReason, pageID)
	} else {
		page, err = bytesToPage(bodyBytes)
	}
	return

	// Status :  404 (error)
	// {"status":404,"statusReason":"Not Found","toe":"8VGVLIN8CI-EODB4-G3CUS"}
	// Status :  204 OK
	// {"id":333576354,"ownerId":804729568,"name":"Test Page One","locked":false,"collectionIds":[],"cardIds":[],"visibility":{"userIds":[804729568],"groupIds":[]}}

}

// Create Creates a new page in your Domo instance.
// Definition
// POST https://api.domo.com/v1/pages
// Returns
// Returns a page object when successful.
func (p *PageService) Create() (err error) {
	p.client.getAccessToken("dashboard")

	url := fmt.Sprintf("%s/v1/pages", baseURL)
	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	// {
	// 	"name": "API test page",
	// 	 "parentId": 431196438,
	// 	 "locked": "TRUE",
	// 	 "cardIds": [12,2535,233,694],
	// 	 "visibility": {
	// 		  "userIds": [793,20,993,19234],
	// 		  "groupIds": [32,25,17,74]
	// 	  }
	//   }

	body := strings.NewReader("{\"name\":\"API test page\",\"id\":0,\"parentId\":431196438,\"ownerId\":0,\"locked\":true,\"collectionIds\":0,\"cardIds\":[12,2535,233,694],\"visibility\":{\"userIds\":[12,2535,233,694],\"groupIds\":[12,2535,233,694]},\"userIds\":0,\"groupIds\":0}")

	bodyBytes, statuscode, err := p.client.genericPOST(url, body, header)

	if err != nil {
		return fmt.Errorf(
			"Unable to Create response from DOMO API %s",
			err,
		)
	}

	str := fmt.Sprintf("%s", bodyBytes)

	fmt.Println(fmt.Sprintf("%v", body))

	fmt.Println(statuscode)
	fmt.Println("GET : " + url)
	fmt.Println(str)
	return
}

// UpdatePage Updates the specified page by providing values to parameters passed.
// Any parameter left out of the request will cause the specific page’s attribute to remain unchanged.
// Also, collections cannot be added or removed via this endpoint, only reordered.
// Giving access to a user or group will also cause that user or group to have access to the parent page (if the page is a subpage).
// Moving a page by updating the parentId will also cause everyone with access to the page to have access to the new parent page.
// Definition
// PUT https://api.domo.com/v1/pages/{PAGE_ID}
// Returns
// Returns the parameter of success or error based on the page ID being valid.

// DeletePage Permanently deletes a page from your Domo instance.
// WARNING: This is destructive and cannot be reversed.
// Definition
// DELETE https://api.domo.com/v1/pages/{PAGE_ID}
// Returns
// Returns the parameter of success or error based on the page ID being valid.
func (p *PageService) DeletePage(pageID int) (err error) {
	p.client.getAccessToken("dashboard")

	url := fmt.Sprintf("%s/v1/pages/%d", baseURL, pageID)
	statuscode, err := p.client.genericDELETE(url, nil)

	fmt.Println(statuscode)
	fmt.Println("DELETE : " + url)

	return
}

// List Get a list of all pages in your Domo instance.
// Definition
// GET https://api.domo.com/v1/pages
// Returns
// Returns all page objects that meet argument criteria from original request.
func (p *PageService) List() (pages Pages, err error) {
	p.client.getAccessToken("dashboard")

	url := fmt.Sprintf("%s/v1/pages", baseURL)
	bodyBytes, _, err := p.client.genericGET(url, nil)

	if err != nil {
		return Pages{}, fmt.Errorf(
			"Unable to ListPages response from DOMO API %s",
			err,
		)
	}

	pages, err = bytesToPages(bodyBytes)

	str := fmt.Sprintf("%s", bodyBytes)
	fmt.Println(str)

	return

	// Status :  404 (error)
	// {"status":404,"statusReason":"Not Found","toe":"8VGVLIN8CI-EODB4-G3CUS"}
	// Status :  204 OK
	// [{"id":968425465,"name":"Developer Home","children":[]},{"id":842042130,"name":"Sample DataSets + Cards","children":[]},{"id":333576354,"name":"Test Page One","children":[]}]

}

// RetrieveCollection Retrieve a page collection ...
// Definition
// GET https://api.domo.com/v1/pages/{PAGE_ID}/collections
// Returns
// Returns the parameter of success or error based on the page ID being valid.
func (p *PageService) RetrieveCollection(pageID int) (err error) {
	p.client.getAccessToken("dashboard")

	url := fmt.Sprintf("%s/v1/pages/%d/collections", baseURL, pageID)
	bodyBytes, _, err := p.client.genericGET(url, nil)

	if err != nil {
		return fmt.Errorf(
			"Unable to RetrieveCollection pages response from DOMO API %s",
			err,
		)
	}

	str := fmt.Sprintf("%s", bodyBytes)

	fmt.Println("GET : " + url)
	fmt.Println(str)
	return
}

// CreatePageCollection ...
// Definition
// POST https://api.domo.com/v1/pages/{PAGE_ID}/collections
// Returns
// Returns the parameter of success or error based on the page ID being valid.

// UpdatePageCollection ...
// Definition
// PUT https://api.domo.com/v1/pages/{PAGE_ID}/collections/{PAGE_COLLECTION_ID}
// Returns
// Returns the parameter of success or error based on the page collection ID being valid.

// DeletePageCollection Permanently deletes a page collection from your Domo instance.
// WARNING: This is destructive and cannot be reversed.
// Definition
// DELETE https://api.domo.com/v1/pages/{PAGE_ID}/collections/{COLLECTION_ID}
// Returns
// Returns the parameter of success or error based on the page collection ID being valid.

func bytesToPage(bodyBytes []byte) (page Page, err error) {
	return page, json.Unmarshal(bodyBytes, &page)
}

func bytesToPages(bodyBytes []byte) (pages Pages, err error) {
	return pages, json.Unmarshal(bodyBytes, &pages)
}
