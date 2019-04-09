package domo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// UserService User API service
type UserService service

// Retrieve Retrieves the details of an existing user.
// Definition
// GET https://api.domo.com/v1/users/{id}
// Returns
// Returns a user object if valid user ID was provided.
// When requesting, if the user ID is related to a user that has been deleted,
// a subset of the user information will be returned, including a deleted property, which will be true.
func (u *UserService) Retrieve(userid int) (user User, err error) {
	u.client.getAccessToken("user")
	url := fmt.Sprintf("%s/v1/users/%d", baseURL, userid)
	bodyBytes, _, err := u.client.genericGET(url, nil)

	if err != nil {
		return user, fmt.Errorf(
			"Unable to retreive user from Domo API %s",
			err,
		)
	}

	return bytesToUser(bodyBytes)
}

// Create Creates a new user in your Domo instance.
// Definition
// POST https://api.domo.com/v1/users
// Returns
// Returns a user object when successful.
// The returned object will have user attributes based on the information that was provided when user was created.
// The two exceptions of attributes not returned are the user's timezone and locale.
func (u *UserService) Create(name string, email string, role string, sendInvite bool) (user User, err error) {
	u.client.getAccessToken("user")
	if CheckRole(role) {
		err = fmt.Errorf("%s is not a valid role , (available roles are: 'Admin', 'Privileged', 'Participant')", role)
		return
	}

	url := fmt.Sprintf("%s/v1/users", baseURL)
	payload := fmt.Sprintf("{\"name\": \"%s\",\"email\": \"%s\",\"role\": \"%s\",\"sendInvite\": %t}", name, email, role, sendInvite)

	body := strings.NewReader(payload)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	bodyBytes, statusCode, err := u.client.genericPOST(url, body, header)

	if err != nil {
		err = fmt.Errorf("Unable to create user from domo API %s", err)
		return
	}

	logger(fmt.Sprintf("createUser Status Code : %d", statusCode))
	user, err = bytesToUser(bodyBytes)

	if err != nil {
		err = fmt.Errorf("Unable to convert to user %s", err)
	}
	return

}

// Update Updates the specified user by providing values to parameters passed.
// Any parameter left out of the request will cause the specific user’s attribute to remain unchanged.
// KNOWN LIMITATION
// Currently all user fields are required
// Definition
// PUT https://api.domo.com/v1/users/{id}
// Returns
// Returns a 200 response code when successful.
func (u *UserService) Update(userid int, name string, email string, role string) (err error) {
	u.client.getAccessToken("user")
	if CheckRole(role) {
		return fmt.Errorf(
			"%s is not a valid role , (available roles are: 'Admin', 'Privileged', 'Participant')",
			role,
		)
	}

	url := fmt.Sprintf("%s/v1/groups/%d", baseURL, userid)
	payload := fmt.Sprintf(
		"{\"name\": \"%s\",\"email\": \"%s\",\"role\": \"%s\"}",
		name,
		email,
		role,
	)
	body := strings.NewReader(payload)

	header := map[string]string{"Content-Type": "application/json"}
	_, statusCode, err := u.client.genericPUT(url, body, header)

	if err != nil {
		err = fmt.Errorf("Unable to update user from dom API %s", err)
	}

	logger(fmt.Sprintf("updategroup Status Code : %d", statusCode))

	return

}

// Delete Permanently deletes a user from your Domo instance.
// WARNING
// This is destructive and cannot be reversed.
// Definition
// DELETE https://api.domo.com/v1/users/{id}
// Returns
// Returns a 204 response code when successful or error based on whether the user ID being valid.
func (u *UserService) Delete(userid int) (err error) {
	u.client.getAccessToken("user")
	url := fmt.Sprintf("%s/v1/users/%d", baseURL, userid)
	statusCode, err := u.client.genericDELETE(url, nil)

	if err != nil {
		err = fmt.Errorf("Unable delete user from dom API %s", err)
	}

	if statusCode != 204 {
		err = fmt.Errorf("Failed to delete userID %d : %s", userid, err)
	}
	return
}

// List Get a list of all users in your Domo instance.
// Definition
// GET https://api.domo.com/v1/users
// Returns
// Returns all user objects that meet argument criteria from original request.
func (u *UserService) List() (users Users, err error) {
	u.client.getAccessToken("user")
	url := fmt.Sprintf("%s/v1/users", baseURL)
	bodyBytes, _, err := u.client.genericGET(url, nil)

	if err != nil {
		return users, fmt.Errorf("Unable to get users from dom API %s", err)
	}

	users, err = bytesToUsers(bodyBytes)

	if err != nil {
		err = fmt.Errorf("Unable convert to user struct %s", err)
	}

	return
}

// Find locates user by name from your Domo instance.
// Returns
// Returns the ID of the user
func (u *UserService) Find(name string) (userID int, err error) {

	u.client.getAccessToken("user")
	data, err := u.List()

	if err != nil {
		return
	}

	for _, r := range data {
		if r.Name == name {
			return r.ID, err
		}
	}
	return
}

// CheckRole The role of the user created (available roles are: 'Admin', 'Privileged', 'Participant')
// Returns bool
func CheckRole(role string) bool {
	switch role {
	case "Admin":
		return true
	case "Privileged":
		return true
	case "Participant":
		return true
	default:
		return false
	}
}

func bytesToUsers(bodyBytes []byte) (users Users, err error) {
	err = json.Unmarshal(bodyBytes, &users)
	if err != nil {
		str := fmt.Sprintf("%s", bodyBytes)
		err = fmt.Errorf("Failed Unmarshal Users %s JSON string : %s", err, str)
	}
	return
}

func bytesToUser(bodyBytes []byte) (user User, err error) {
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		str := fmt.Sprintf("%s", bodyBytes)
		err = fmt.Errorf("Failed Unmarshal User %s JSON string : %s", err, str)
	}
	return
}
