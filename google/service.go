package google

import (
    "github.com/pomack/oauth2_client"
    "http"
    "image"
    "io/ioutil"
    "json"
    "os"
    "strings"
    "url"
)

func retrieveInfo(client oauth2_client.OAuth2Client, scope, userId, projection, id string, m url.Values, value interface{}) (err os.Error) {
    var useUserId string
    if len(userId) <= 0 {
        useUserId = GOOGLE_DEFAULT_USER_ID
    } else {
        useUserId = url.QueryEscape(userId)
    }
    if len(projection) <= 0 {
        projection = GOOGLE_DEFAULT_PROJECTION
    }
    headers := make(http.Header)
    headers.Set("GData-Version", "3.0")
    if m == nil {
        m = make(url.Values)
    }
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_FEEDS_API_ENDPOINT
    for _, s := range []string{scope, useUserId, projection, id} {
        if len(s) > 0 {
            if uri[len(uri) - 1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, headers, uri, m)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            var e ErrorResponse
            b, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(b, &e)
            if len(e.Error.Description) > 0 {
                err = os.NewError(e.Error.Message)
            } else {
                err = os.NewError(string(b))
            }
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveContacts(client oauth2_client.OAuth2Client, m url.Values) (*ContactFeed, os.Error) {
    return RetrieveContactsWithProjection(client, "", m)
}

func RetrieveContactsWithProjection(client oauth2_client.OAuth2Client, projection string, m url.Values) (*ContactFeed, os.Error) {
    return RetrieveContactsWithUserIdAndProjection(client, "", projection, m)
}

func RetrieveContactsWithUserIdAndProjection(client oauth2_client.OAuth2Client, userId, projection string, m url.Values) (*ContactFeed, os.Error) {
    resp := new(ContactFeedResponse)
    err := retrieveInfo(client, GOOGLE_CONTACTS_SCOPE, userId, projection, "", m, resp)
    return resp.Feed, err
}

func RetrieveContact(client oauth2_client.OAuth2Client, id string, m url.Values) (*Contact, os.Error) {
    return RetrieveContactWithProjection(client, "", id, m)
}

func RetrieveContactWithProjection(client oauth2_client.OAuth2Client, projection, id string, m url.Values) (*Contact, os.Error) {
    return RetrieveContactWithUserIdAndProjection(client, "", projection, id, m)
}

func RetrieveContactWithUserIdAndProjection(client oauth2_client.OAuth2Client, userId, projection, id string, m url.Values) (*Contact, os.Error) {
    resp := new(ContactEntryResponse)
    err := retrieveInfo(client, GOOGLE_CONTACTS_SCOPE, userId, projection, id, m, resp)
    return resp.Entry, err
}

func RetrieveContactPhoto(client oauth2_client.OAuth2Client, id string) (theimage image.Image, format string, err os.Error) {
    uri := GOOGLE_FEEDS_API_ENDPOINT + strings.Replace(strings.Join([]string{GOOGLE_PHOTOS_SCOPE, GOOGLE_DEFAULT_USER_ID, id}, "/"), "//", "/", -1)
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, nil, uri, nil)
    if err != nil {
        return theimage, format, err
    }
    if resp != nil {
        theimage, format, err = image.Decode(resp.Body)
    }
    return theimage, format, err
}

func RetrieveGroups(client oauth2_client.OAuth2Client, m url.Values) (*GroupsFeedResponse, os.Error) {
    return RetrieveGroupsWithUserId(client, "", m)
}

func RetrieveGroupsWithUserId(client oauth2_client.OAuth2Client, userId string, m url.Values) (*GroupsFeedResponse, os.Error) {
    resp := new(GroupsFeedResponse)
    err := retrieveInfo(client, GOOGLE_GROUPS_SCOPE, userId, "", "", m, resp)
    return resp, err
}

func RetrieveGroup(client oauth2_client.OAuth2Client, id string, m url.Values) (*GroupResponse, os.Error) {
    return RetrieveGroupWithUserId(client, "", id, m)
}

func RetrieveGroupWithUserId(client oauth2_client.OAuth2Client, userId, id string, m url.Values) (*GroupResponse, os.Error) {
    resp := new(GroupResponse)
    err := retrieveInfo(client, GOOGLE_GROUPS_SCOPE, userId, "", id, m, resp)
    return resp, err
}

