package domo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GroupService Group API service
type GroupService service

//GroupUsers list of userID
type GroupUsers []int

// Retrieve Retrieves the details of an existing group.
//
// Definition GET https://api.domo.com/v1/groups/{GROUP_ID}
//
// Returns a group object if valid group ID was provided. When requesting,
// if the group ID is related to a customer that has been deleted,
// a subset of the group's information will be returned, including a deleted property, which will be true.
func (g *GroupService) Retrieve(groupID int) (group Group, err error) {

	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups/%d", baseURL, groupID)
	bodyBytes, _, err := g.client.genericGET(url, nil)

	err = json.Unmarshal(bodyBytes, &group)
	if err != nil {
		return group, fmt.Errorf(
			"Unable to Retrieve response from Domo API %s",
			err,
		)
	}

	logger(fmt.Sprintf("[GroupService] Retrieve : groupID %d name '%s' found !", groupID, group.Name))

	return
}

// Create Creates a new group in your Domo instance.
// Definition
// POST https://api.domo.com/v1/groups
// Returns
// Returns a group object when successful.
// The returned group will have user attributes based on the information that was provided when group was created.
func (g *GroupService) Create(name string, isDefault bool) (group *Group, err error) {
	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups", baseURL)
	payload := fmt.Sprintf("{\"name\": \"%s\",\"default\": %t}", name, isDefault)
	body := strings.NewReader(payload)

	header := map[string]string{
		"Content-Type": "application/json",
	}

	bodyBytes, statusCode, err := g.client.genericPOST(url, body, header)

	if err != nil {
		return group, fmt.Errorf(
			"Unable to Create response from Domo API %s",
			err,
		)
	}

	logger(fmt.Sprintf("createStream Status Code : %d", statusCode))
	err = json.Unmarshal(bodyBytes, &group)
	bodyString := string(bodyBytes)

	logger("createStream" + bodyString)

	return

}

// Update Updates the specified group by providing values to parameters passed.
// Any parameter left out of the request will cause the specific group’s attribute to remain unchanged.
// Definition
// PUT https://api.domo.com/v1/groups/{GROUP_ID}
// Returns
// Returns the parameter of success or error based on the group ID being valid.
func (g *GroupService) Update(groupID int, name string, isActive bool, isDefault bool) (err error) {

	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups/%d", baseURL, groupID)
	payload := fmt.Sprintf("{\"name\": \"%s\",\"active\": %t ,\"default\": %t}", name, isActive, isDefault)

	body := strings.NewReader(payload)

	header := map[string]string{
		"Content-Type": "application/json",
	}

	_, statusCode, err := g.client.genericPUT(url, body, header)

	logger(fmt.Sprintf("updategroup Status Code : %d", statusCode))

	if err != nil {
		err = fmt.Errorf(
			"Unable to Create response from Domo API %s",
			err,
		)
	}

	return

}

// Delete Permanently deletes a group from your Domo instance.
// This is destructive and cannot be reversed.
// Definition
// DELETE https://api.domo.com/v1/groups/{GROUP_ID}
// Returns
// Returns the parameter of success or error based on the group ID being valid.
func (g *GroupService) Delete(groupID int) (err error) {

	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups/%d", baseURL, groupID)
	statusCode, err := g.client.genericDELETE(url, nil)

	if err != nil {
		err = fmt.Errorf(
			"Unable to delete, response from Domo API %s",
			err,
		)
	}

	if statusCode != 204 {
		// it failed need to do something
		err = fmt.Errorf("Failed to delete dataset")
	}

	logger(fmt.Sprintf("[GroupService] Delete : groupID %d status %d", groupID, statusCode))

	return
}

// List Get a list of all groups in your Domo instance.
// Definition
// GET https://api.domo.com/v1/groups
// Returns
// Returns all group objects that meet argument criteria from original request.
func (g *GroupService) List() (userGroups Groups, err error) {

	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups", baseURL)
	bodyBytes, _, err := g.client.genericGET(url, nil)

	if err != nil {
		return userGroups, fmt.Errorf(
			"Unable to list, response from Domo API %s",
			err,
		)
	}

	err = json.Unmarshal(bodyBytes, &userGroups)
	if err != nil {
		err = fmt.Errorf(
			"Unable to unmarshal list %s",
			err,
		)
	}

	logger(fmt.Sprintf("[GroupService] List : %d groups found", len(userGroups)))

	return
}

// Find locates group by name from your Domo instance.
// Returns
// Returns the ID og the group
func (g *GroupService) Find(name string) (groupID int, err error) {

	g.client.getAccessToken("user")
	data, err := g.List()

	if err != nil {
		return
	}

	for _, r := range data {
		if r.Name == name {
			groupID = r.ID
		}
	}

	logger(fmt.Sprintf("[GroupService] FindGroup : '%s' found ID %d", name, groupID))

	return
}

// AddUser Add to a group Add user to a group in your Domo instance.
// Definition
// PUT https://api.domo.com/v1/groups/{GROUP_ID}/users/{USER_ID}
// Returns
// Returns the parameter of success or error based on the group ID being valid.
func (g *GroupService) AddUser(groupID int, userID int) (err error) {
	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups/%d/users/%d", baseURL, groupID, userID)

	header := map[string]string{
		"Content-Type": "application/json",
	}

	_, statusCode, err := g.client.genericPUT(url, nil, header)

	if err != nil {
		err = fmt.Errorf(
			"Unable to Add User from Domo API %s",
			err,
		)
	}

	logger(fmt.Sprintf("[GroupService] AddUser : add userID %d from groupID %d status %d", userID, groupID, statusCode))

	return
}

// ListUsers in a group List the users in a group in your Domo instance.
// Definition
// GET https://api.domo.com/v1/groups/{GROUP_ID}/users
// Returns
// Returns IDs of users that are a part of the requested group.
func (g *GroupService) ListUsers(groupID int) (groupUsers GroupUsers, err error) {

	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups/%d/users", baseURL, groupID)
	bodyBytes, _, err := g.client.genericGET(url, nil)

	if err != nil {
		return groupUsers, fmt.Errorf(
			"Unable to ListUsers from Domo API %s",
			err,
		)
	}

	err = json.Unmarshal(bodyBytes, &groupUsers)

	if err != nil {
		err = fmt.Errorf(
			"Unable to unmarshal GroupUsers %s",
			err,
		)
	}

	logger(fmt.Sprintf("[GroupService] ListUsers : %d users found", len(groupUsers)))

	return
}

// RemoveUser from a group Remove a user from a group in your Domo instance.
// Definition
// DELETE https://api.domo.com/v1/groups/{GROUP_ID}/users/{USER_ID}
// Returns
// Returns the parameter of success or error based on the group ID being valid.
func (g *GroupService) RemoveUser(groupID int, userID int) (err error) {

	g.client.getAccessToken("user")

	url := fmt.Sprintf("%s/v1/groups/%d/users/%d", baseURL, groupID, userID)

	statusCode, err := g.client.genericDELETE(url, nil)

	if err != nil {
		err = fmt.Errorf(
			"Unable to RemoveUser from Domo API %s",
			err,
		)
	}

	if statusCode != 204 {
		// it failed need to do something
		err = fmt.Errorf("Failed to delete dataset")
	}

	logger(fmt.Sprintf("[GroupService] RemoveUser : remove userID %d from groupID %d err %s", userID, groupID, err))

	return err
}
