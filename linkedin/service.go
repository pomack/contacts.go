package linkedin

import (
	"github.com/pomack/oauth2_client"
	"http"
	"io/ioutil"
	"json"
	"os"
	"strings"
	"url"
)

func retrieveInfo(client oauth2_client.OAuth2Client, scope, userId, resourceName, resourceId, subcategory, subcategoryId, projection string, m url.Values, value interface{}) (err os.Error) {
	var useUserId string
	if len(userId) <= 0 {
		useUserId = LINKEDIN_DEFAULT_USER_ID
	} else {
		useUserId = url.QueryEscape(userId)
	}
	uri := LINKEDIN_API_ENDPOINT
	for _, s := range []string{scope, useUserId, resourceName, resourceId, subcategory, subcategoryId} {
		if len(s) > 0 {
			if uri[len(uri)-1] != '/' {
				uri += "/"
			}
			uri += s
		}
	}
	if len(projection) > 0 {
		uri += ":(" + projection + ")"
	}
	headers := make(http.Header)
	headers.Set("Accept", "application/json")
	headers.Set("X-Li-Format", "json")
	if m == nil {
		m = make(url.Values)
	}
	resp, _, err := oauth2_client.AuthorizedGetRequest(client, headers, uri, m)
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode >= 400 {
			e := new(ErrorResponse)
			b, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(b, e)
			if len(e.Message) > 0 {
				err = e
			} else {
				err = os.NewError(string(b))
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(value)
		}
	}
	return err
}

func RetrieveSelfProfile(client oauth2_client.OAuth2Client, fields []string, m url.Values) (*Contact, os.Error) {
	return RetrieveProfile(client, "", fields, m)
}

func RetrieveProfile(client oauth2_client.OAuth2Client, userId string, fields []string, m url.Values) (*Contact, os.Error) {
	resp := new(Contact)
	var projection string
	if fields == nil || len(fields) == 0 {
		projection = LINKEDIN_DEFAULT_PROFILE_FIELDS
	} else {
		projection = strings.Join(fields, ",")
	}
	err := retrieveInfo(client, "people", userId, "", "", "", "", projection, m, resp)
	return resp, err
}

func RetrieveConnections(client oauth2_client.OAuth2Client, fields []string, m url.Values) (*ContactList, os.Error) {
	return RetrieveConnectionsForUser(client, "", fields, m)
}

func RetrieveConnectionsForUser(client oauth2_client.OAuth2Client, userId string, fields []string, m url.Values) (*ContactList, os.Error) {
	resp := new(ContactList)
	var projection string
	if fields == nil || len(fields) == 0 {
		projection = LINKEDIN_DEFAULT_PROFILE_FIELDS
	} else {
		projection = strings.Join(fields, ",")
	}
	err := retrieveInfo(client, "people", userId, "connections", "", "", "", projection, m, resp)
	return resp, err
}
