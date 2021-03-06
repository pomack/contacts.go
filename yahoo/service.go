package yahoo

import (
    "encoding/json"
    "errors"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
)

type Guider interface {
    Guid() string
}

func getGuid(client oauth2_client.OAuth2Client) (string, error) {
    var guid string
    if guider, ok := client.(Guider); ok {
        guid = guider.Guid()
    }
    if len(guid) <= 0 {
        resp, err := RetrieveMeGuid(client, nil)
        if err != nil {
            return "", err
        }
        if resp != nil {
            guid = resp.Guid.Value
        }
        if len(guid) <= 0 {
            return "", errors.New("Cannot determine GUID for the current user")
        }
    }
    return guid, nil
}

func retrieveInfo(client oauth2_client.OAuth2Client, scope, userId, resourceName, resourceId, subcategory, subcategoryId string, m url.Values, value interface{}) (err error) {
    var useUserId string
    if len(userId) <= 0 {
        useUserId = YAHOO_DEFAULT_USER_ID
    } else {
        useUserId = url.QueryEscape(userId)
    }
    uri := YAHOO_SOCIAL_API_ENDPOINT
    for _, s := range []string{scope, useUserId, resourceName, resourceId, subcategory, subcategoryId} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    headers := make(http.Header)
    headers.Set("Accept", "application/json")
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
            if len(e.ErrorField.Description) > 0 {
                err = e
            } else {
                err = errors.New(string(b))
            }
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveMeGuid(client oauth2_client.OAuth2Client, m url.Values) (*MeGuidResponse, error) {
    resp := new(MeGuidResponse)
    err := retrieveInfo(client, "", "me", "guid", "", "", "", m, resp)
    return resp, err
}

func RetrieveContacts(client oauth2_client.OAuth2Client, m url.Values) (*ContactsResponse, error) {
    return RetrieveContactsForUser(client, "", m)
}

func RetrieveContactsForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ContactsResponse, error) {
    resp := new(ContactsResponse)
    err := retrieveInfo(client, "user", userId, "contacts", "", "", "", m, resp)
    return resp, err
}

func RetrieveContact(client oauth2_client.OAuth2Client, id string, m url.Values) (*ContactResponse, error) {
    return RetrieveContactForUser(client, "", id, m)
}

func RetrieveContactForUser(client oauth2_client.OAuth2Client, userId, id string, m url.Values) (*ContactResponse, error) {
    resp := new(ContactResponse)
    err := retrieveInfo(client, "user", userId, "contact", id, "", "", m, resp)
    return resp, err
}

func RetrieveCategories(client oauth2_client.OAuth2Client, m url.Values) (*CategoriesResponse, error) {
    return RetrieveCategoriesForUser(client, "", m)
}

func RetrieveCategoriesForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*CategoriesResponse, error) {
    resp := new(CategoriesResponse)
    err := retrieveInfo(client, "user", userId, "categories", "", "", "", m, resp)
    return resp, err
}

func RetrieveCategory(client oauth2_client.OAuth2Client, categoryId string, m url.Values) (*CategoryResponse, error) {
    return RetrieveCategoryForUser(client, "", categoryId, m)
}

func RetrieveCategoryForUser(client oauth2_client.OAuth2Client, userId, categoryId string, m url.Values) (*CategoryResponse, error) {
    resp := new(CategoryResponse)
    err := retrieveInfo(client, "user", userId, "category", categoryId, "", "", m, resp)
    return resp, err
}

func RetrieveContactSync(client oauth2_client.OAuth2Client, revno int64, m url.Values) (*ContactSyncResponse, error) {
    return RetrieveContactSyncForUser(client, "", revno, m)
}

func RetrieveContactSyncForUser(client oauth2_client.OAuth2Client, userId string, revno int64, m url.Values) (*ContactSyncResponse, error) {
    resp := new(ContactSyncResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("view", "sync")
    m.Set("rev", strconv.FormatInt(revno, 10))
    err := retrieveInfo(client, "user", userId, "contacts", "", "", "", m, resp)
    return resp, err
}

func RetrieveSelfProfile(client oauth2_client.OAuth2Client, m url.Values) (*ProfileResponse, error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveProfile(client, guid, m)
}

func RetrieveProfile(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ProfileResponse, error) {
    resp := new(ProfileResponse)
    err := retrieveInfo(client, "user", userId, "profile", "", "", "", m, resp)
    return resp, err
}

func RetrieveNotifications(client oauth2_client.OAuth2Client, m url.Values) (*NotificationsResponse, error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveNotificationsForUser(client, guid, m)
}

func RetrieveNotificationsForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*NotificationsResponse, error) {
    resp := new(NotificationsResponse)
    err := retrieveInfo(client, "user", userId, "received_notifications", "", "", "", m, resp)
    return resp, err
}

func RetrieveNotification(client oauth2_client.OAuth2Client, id string, m url.Values) (*NotificationResponse, error) {
    resp := new(NotificationResponse)
    err := retrieveInfo(client, "notification", id, "", "", "", "", m, resp)
    return resp, err
}

func RetrieveConnections(client oauth2_client.OAuth2Client, m url.Values) (*ConnectionsResponse, error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveConnectionsForUser(client, guid, m)
}

func RetrieveConnectionsForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ConnectionsResponse, error) {
    resp := new(ConnectionsResponse)
    err := retrieveInfo(client, "user", userId, "connections", "", "", "", m, resp)
    return resp, err
}

func RetrieveConnection(client oauth2_client.OAuth2Client, connectionId string, m url.Values) (*ConnectionResponse, error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveConnectionForUser(client, guid, connectionId, m)
}

func RetrieveConnectionForUser(client oauth2_client.OAuth2Client, userId, connectionId string, m url.Values) (*ConnectionResponse, error) {
    resp := new(ConnectionResponse)
    err := retrieveInfo(client, "user", userId, "connection", connectionId, "", "", m, resp)
    return resp, err
}

func CreateContact(client oauth2_client.OAuth2Client, userId string, contact *Contact) (err error) {
    if userId == "" {
        userId = "me"
    }
    var useUserId string
    if len(userId) <= 0 {
        useUserId = YAHOO_DEFAULT_USER_ID
    } else {
        useUserId = url.QueryEscape(userId)
    }
    uri := YAHOO_SOCIAL_API_ENDPOINT
    scope := "user"
    resourceName := "contacts"
    for _, s := range []string{scope, useUserId, resourceName} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    headers := make(http.Header)
    headers.Set("Accept", "application/json")
    headers.Set("Content-Type", "application/json; charset=utf-8")
    m := make(url.Values)
    m.Add("alt", "json")
    b, err := json.Marshal(contact)
    if err != nil {
        return err
    }
    resp, _, err := oauth2_client.AuthorizedPostRequestBytes(client, headers, uri, m, b)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            e := new(ErrorResponse)
            b, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(b, e)
            if len(e.ErrorField.Description) > 0 {
                err = e
            } else {
                err = errors.New(string(b))
            }
        }
    }
    return err
}

func UpdateContact(client oauth2_client.OAuth2Client, userId, contactId string, contact *Contact) (err error) {
    if userId == "" {
        userId = "me"
    }
    var useUserId string
    if len(userId) <= 0 {
        useUserId = YAHOO_DEFAULT_USER_ID
    } else {
        useUserId = url.QueryEscape(userId)
    }
    uri := YAHOO_SOCIAL_API_ENDPOINT
    scope := "user"
    resourceName := "contact"
    resourceId := contactId
    for _, s := range []string{scope, useUserId, resourceName, resourceId} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    headers := make(http.Header)
    headers.Set("Accept", "application/json")
    headers.Set("Content-Type", "application/json; charset=utf-8")
    m := make(url.Values)
    m.Add("alt", "json")
    b, err := json.Marshal(contact)
    if err != nil {
        return err
    }
    resp, _, err := oauth2_client.AuthorizedPutRequestBytes(client, headers, uri, m, b)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            e := new(ErrorResponse)
            b, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(b, e)
            if len(e.ErrorField.Description) > 0 {
                err = e
            } else {
                err = errors.New(string(b))
            }
        }
    }
    return err
}

func DeleteContact(client oauth2_client.OAuth2Client, userId, contactId string) (err error) {
    if userId == "" {
        userId = "me"
    }
    var useUserId string
    if len(userId) <= 0 {
        useUserId = YAHOO_DEFAULT_USER_ID
    } else {
        useUserId = url.QueryEscape(userId)
    }
    uri := YAHOO_SOCIAL_API_ENDPOINT
    scope := "user"
    resourceName := "contact"
    resourceId := contactId
    for _, s := range []string{scope, useUserId, resourceName, resourceId} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    headers := make(http.Header)
    headers.Set("Accept", "application/json")
    resp, _, err := oauth2_client.AuthorizedDeleteRequest(client, headers, uri, nil)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            e := new(ErrorResponse)
            b, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(b, e)
            if len(e.ErrorField.Description) > 0 {
                err = e
            } else {
                err = errors.New(string(b))
            }
        }
    }
    return err
}
