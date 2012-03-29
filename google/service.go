package google

import (
    "encoding/json"
    "errors"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "image"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
)

func retrieveInfo(client oauth2_client.OAuth2Client, scope, userId, projection, id string, m url.Values, value interface{}) (err error) {
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
            if uri[len(uri)-1] != '/' {
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
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveContacts(client oauth2_client.OAuth2Client, m url.Values) (*ContactFeed, error) {
    return RetrieveContactsWithProjection(client, "", m)
}

func RetrieveContactsWithProjection(client oauth2_client.OAuth2Client, projection string, m url.Values) (*ContactFeed, error) {
    return RetrieveContactsWithUserIdAndProjection(client, "", projection, m)
}

func RetrieveContactsWithUserIdAndProjection(client oauth2_client.OAuth2Client, userId, projection string, m url.Values) (*ContactFeed, error) {
    resp := new(ContactFeedResponse)
    err := retrieveInfo(client, GOOGLE_CONTACTS_SCOPE, userId, projection, "", m, resp)
    return resp.Feed, err
}

func RetrieveContact(client oauth2_client.OAuth2Client, id string, m url.Values) (*Contact, error) {
    return RetrieveContactWithProjection(client, "", id, m)
}

func RetrieveContactWithProjection(client oauth2_client.OAuth2Client, projection, id string, m url.Values) (*Contact, error) {
    return RetrieveContactWithUserIdAndProjection(client, "", projection, id, m)
}

func RetrieveContactWithUserIdAndProjection(client oauth2_client.OAuth2Client, userId, projection, id string, m url.Values) (*Contact, error) {
    resp := new(ContactEntryResponse)
    err := retrieveInfo(client, GOOGLE_CONTACTS_SCOPE, userId, projection, id, m, resp)
    return resp.Entry, err
}

func RetrieveContactPhoto(client oauth2_client.OAuth2Client, id string) (theimage image.Image, format string, err error) {
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

func RetrieveAllGroups(client oauth2_client.OAuth2Client) (*GroupsFeedResponse, error) {
    return RetrieveGroupsWithUserId(client, "", nil)
}

func RetrieveGroups(client oauth2_client.OAuth2Client, m url.Values) (*GroupsFeedResponse, error) {
    return RetrieveGroupsWithUserId(client, "", m)
}

func RetrieveGroupsWithUserId(client oauth2_client.OAuth2Client, userId string, m url.Values) (*GroupsFeedResponse, error) {
    resp := new(GroupsFeedResponse)
    err := retrieveInfo(client, GOOGLE_GROUPS_SCOPE, userId, "", "", m, resp)
    return resp, err
}

func RetrieveGroup(client oauth2_client.OAuth2Client, id string, m url.Values) (*GroupResponse, error) {
    return RetrieveGroupWithUserId(client, "", id, m)
}

func RetrieveGroupWithUserId(client oauth2_client.OAuth2Client, userId, id string, m url.Values) (*GroupResponse, error) {
    resp := new(GroupResponse)
    err := retrieveInfo(client, GOOGLE_GROUPS_SCOPE, userId, "", id, m, resp)
    return resp, err
}

func CreateContact(client oauth2_client.OAuth2Client, userId, projection string, value *Contact) (response *ContactEntryResponse, err error) {
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
    m := make(url.Values)
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_CONTACTS_API_ENDPOINT
    for _, s := range []string{useUserId, projection} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    entry := &ContactEntryInsertRequest{Version: "1.0", Encoding: "UTF-8", Entry: value}
    value.Xmlns = XMLNS_ATOM
    value.XmlnsGcontact = XMLNS_GCONTACT
    value.XmlnsBatch = XMLNS_GDATA_BATCH
    value.XmlnsGd = XMLNS_GD
    buf, err := json.Marshal(entry)
    if err != nil {
        return
    }
    resp, _, err := oauth2_client.AuthorizedPostRequestBytes(client, headers, uri, m, buf)
    if err != nil {
        return
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        } else {
            response = new(ContactEntryResponse)
            err = json.NewDecoder(resp.Body).Decode(response)
        }
    }
    return
}

func UpdateContact(client oauth2_client.OAuth2Client, userId, projection string, original, value *Contact) (response *ContactEntryResponse, err error) {
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
    headers.Set("If-Match", original.Etag)
    m := make(url.Values)
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_CONTACTS_API_ENDPOINT
    for _, s := range []string{useUserId, projection, original.ContactId()} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    entry := &ContactEntryUpdateRequest{Entry: value}
    value.Xmlns = XMLNS_ATOM
    value.XmlnsGcontact = XMLNS_GCONTACT
    value.XmlnsBatch = XMLNS_GDATA_BATCH
    value.XmlnsGd = XMLNS_GD
    value.Etag = original.Etag
    value.Id = original.Id
    value.Updated.Value = time.Now().UTC().Format(GOOGLE_DATETIME_FORMAT)
    value.Categories = make([]AtomCategory, 1)
    value.Categories[0].Scheme = ATOM_CATEGORY_SCHEME_KIND
    value.Categories[0].Term = ATOM_CATEGORY_TERM_CONTACT
    buf, err := json.Marshal(entry)
    if err != nil {
        return
    }
    resp, _, err := oauth2_client.AuthorizedPutRequestBytes(client, headers, uri, m, buf)
    if err != nil {
        return
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        } else {
            response = new(ContactEntryResponse)
            err = json.NewDecoder(resp.Body).Decode(response)
        }
    }
    return
}

func DeleteContact(client oauth2_client.OAuth2Client, userId, projection string, original *Contact) (err error) {
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
    headers.Set("If-Match", original.Etag)
    m := make(url.Values)
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_CONTACTS_API_ENDPOINT
    for _, s := range []string{useUserId, projection, original.ContactId()} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    resp, _, err := oauth2_client.AuthorizedDeleteRequest(client, headers, uri, nil)
    if err != nil {
        return
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        }
    }
    return
}

func CreateGroup(client oauth2_client.OAuth2Client, userId, projection string, value *ContactGroup) (response *GroupResponse, err error) {
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
    m := make(url.Values)
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_GROUPS_API_ENDPOINT
    for _, s := range []string{useUserId, projection} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    entry := &GroupEntryInsertRequest{Version: "1.0", Encoding: "UTF-8", Entry: value}
    value.Xmlns = XMLNS_ATOM
    value.XmlnsGcontact = XMLNS_GCONTACT
    value.XmlnsBatch = XMLNS_GDATA_BATCH
    value.XmlnsGd = XMLNS_GD
    buf, err := json.Marshal(entry)
    if err != nil {
        return
    }
    resp, _, err := oauth2_client.AuthorizedPostRequestBytes(client, headers, uri, m, buf)
    if err != nil {
        return
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        } else {
            response = new(GroupResponse)
            err = json.NewDecoder(resp.Body).Decode(response)
        }
    }
    return
}

func UpdateGroup(client oauth2_client.OAuth2Client, userId, projection string, original, value *ContactGroup) (response *GroupResponse, err error) {
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
    headers.Set("If-Match", original.Etag)
    m := make(url.Values)
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_GROUPS_API_ENDPOINT
    for _, s := range []string{useUserId, projection, original.GroupId()} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    entry := &GroupEntryUpdateRequest{Entry: value}
    value.Xmlns = XMLNS_ATOM
    value.XmlnsGcontact = XMLNS_GCONTACT
    value.XmlnsBatch = XMLNS_GDATA_BATCH
    value.XmlnsGd = XMLNS_GD
    value.Etag = original.Etag
    value.Id = original.Id
    value.Updated.Value = time.Now().UTC().Format(GOOGLE_DATETIME_FORMAT)
    value.Categories = make([]AtomCategory, 1)
    value.Categories[0].Scheme = ATOM_CATEGORY_SCHEME_KIND
    value.Categories[0].Term = ATOM_CATEGORY_TERM_GROUP
    buf, err := json.Marshal(entry)
    if err != nil {
        return
    }
    resp, _, err := oauth2_client.AuthorizedPutRequestBytes(client, headers, uri, m, buf)
    if err != nil {
        return
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        } else {
            response = new(GroupResponse)
            err = json.NewDecoder(resp.Body).Decode(response)
        }
    }
    return
}

func DeleteGroup(client oauth2_client.OAuth2Client, userId, projection string, original *ContactGroup) (err error) {
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
    headers.Set("If-Match", original.Etag)
    m := make(url.Values)
    if len(m.Get(CONTACTS_ALT_PARAM)) <= 0 {
        m.Set(CONTACTS_ALT_PARAM, "json")
    }
    uri := GOOGLE_GROUPS_API_ENDPOINT
    for _, s := range []string{useUserId, projection, original.GroupId()} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    resp, _, err := oauth2_client.AuthorizedDeleteRequest(client, headers, uri, nil)
    if err != nil {
        return
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        }
    }
    return
}
