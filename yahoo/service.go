package yahoo

import (
    "github.com/pomack/oauth2_client"
    "http"
    "io/ioutil"
    "json"
    "os"
    "strconv"
    "url"
)

type Guider interface {
    Guid() string
}

func getGuid(client oauth2_client.OAuth2Client) (string, os.Error) {
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
            return "", os.NewError("Cannot determine GUID for the current user")
        }
    }
    return guid, nil
}

func retrieveInfo(client oauth2_client.OAuth2Client, scope, userId, resourceName, resourceId, subcategory, subcategoryId string, m url.Values, value interface{}) (err os.Error) {
    var useUserId string
    if len(userId) <= 0 {
        useUserId = YAHOO_DEFAULT_USER_ID
    } else {
        useUserId = url.QueryEscape(userId)
    }
    uri := YAHOO_SOCIAL_API_ENDPOINT
    for _, s := range []string{scope, useUserId, resourceName, resourceId, subcategory, subcategoryId} {
        if len(s) > 0 {
            if uri[len(uri) - 1] != '/' {
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
            var e ErrorResponse
            b, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(b, &e)
            if len(e.Error.Description) > 0 {
                err = os.NewError(e.Error.Description)
            } else {
                err = os.NewError(string(b))
            }
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}


func RetrieveMeGuid(client oauth2_client.OAuth2Client, m url.Values) (*MeGuidResponse, os.Error) {
    resp := new(MeGuidResponse)
    err := retrieveInfo(client, "", "me", "guid", "", "", "", m, resp)
    return resp, err
}

func RetrieveContacts(client oauth2_client.OAuth2Client, m url.Values) (*ContactsResponse, os.Error) {
    return RetrieveContactsForUser(client, "", m)
}

func RetrieveContactsForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ContactsResponse, os.Error) {
    resp := new(ContactsResponse)
    err := retrieveInfo(client, "user", userId, "contacts", "", "", "", m, resp)
    return resp, err
}

func RetrieveContact(client oauth2_client.OAuth2Client, id string, m url.Values) (*ContactResponse, os.Error) {
    return RetrieveContactForUser(client, "", id, m)
}

func RetrieveContactForUser(client oauth2_client.OAuth2Client, userId, id string, m url.Values) (*ContactResponse, os.Error) {
    resp := new(ContactResponse)
    err := retrieveInfo(client, "user", userId, "contact", id, "", "", m, resp)
    return resp, err
}


func RetrieveCategories(client oauth2_client.OAuth2Client, m url.Values) (*CategoriesResponse, os.Error) {
    return RetrieveCategoriesForUser(client, "", m)
}

func RetrieveCategoriesForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*CategoriesResponse, os.Error) {
    resp := new(CategoriesResponse)
    err := retrieveInfo(client, "user", userId, "categories", "", "", "", m, resp)
    return resp, err
}


func RetrieveCategory(client oauth2_client.OAuth2Client, categoryId string, m url.Values) (*CategoryResponse, os.Error) {
    return RetrieveCategoryForUser(client, "", categoryId, m)
}

func RetrieveCategoryForUser(client oauth2_client.OAuth2Client, userId, categoryId string, m url.Values) (*CategoryResponse, os.Error) {
    resp := new(CategoryResponse)
    err := retrieveInfo(client, "user", userId, "category", categoryId, "", "", m, resp)
    return resp, err
}


func RetrieveContactSync(client oauth2_client.OAuth2Client, revno int64, m url.Values) (*ContactSyncResponse, os.Error) {
    return RetrieveContactSyncForUser(client, "", revno, m)
}

func RetrieveContactSyncForUser(client oauth2_client.OAuth2Client, userId string, revno int64, m url.Values) (*ContactSyncResponse, os.Error) {
    resp := new(ContactSyncResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("view", "sync")
    m.Set("rev", strconv.Itoa64(revno))
    err := retrieveInfo(client, "user", userId, "contacts", "", "", "", m, resp)
    return resp, err
}




func RetrieveSelfProfile(client oauth2_client.OAuth2Client, m url.Values) (*ProfileResponse, os.Error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveProfile(client, guid, m)
}

func RetrieveProfile(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ProfileResponse, os.Error) {
    resp := new(ProfileResponse)
    err := retrieveInfo(client, "user", userId, "profile", "", "", "", m, resp)
    return resp, err
}


func RetrieveNotifications(client oauth2_client.OAuth2Client, m url.Values) (*NotificationsResponse, os.Error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveNotificationsForUser(client, guid, m)
}

func RetrieveNotificationsForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*NotificationsResponse, os.Error) {
    resp := new(NotificationsResponse)
    err := retrieveInfo(client, "user", userId, "received_notifications", "", "", "", m, resp)
    return resp, err
}

func RetrieveNotification(client oauth2_client.OAuth2Client, id string, m url.Values) (*NotificationResponse, os.Error) {
    resp := new(NotificationResponse)
    err := retrieveInfo(client, "notification", id, "", "", "", "", m, resp)
    return resp, err
}

func RetrieveConnections(client oauth2_client.OAuth2Client, m url.Values) (*ConnectionsResponse, os.Error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveConnectionsForUser(client, guid, m)
}

func RetrieveConnectionsForUser(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ConnectionsResponse, os.Error) {
    resp := new(ConnectionsResponse)
    err := retrieveInfo(client, "user", userId, "connections", "", "", "", m, resp)
    return resp, err
}

func RetrieveConnection(client oauth2_client.OAuth2Client, connectionId string, m url.Values) (*ConnectionResponse, os.Error) {
    guid, err := getGuid(client)
    if err != nil {
        return nil, err
    }
    return RetrieveConnectionForUser(client, guid, connectionId, m)
}

func RetrieveConnectionForUser(client oauth2_client.OAuth2Client, userId, connectionId string, m url.Values) (*ConnectionResponse, os.Error) {
    resp := new(ConnectionResponse)
    err := retrieveInfo(client, "user", userId, "connection", connectionId, "", "", m, resp)
    return resp, err
}

