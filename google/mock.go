package google

import (
    "github.com/pomack/jsonhelper.go/jsonhelper"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "bytes"
    "http"
    "io/ioutil"
    "json"
    "os"
    "strconv"
    "strings"
    "time"
)

var (
    contactsById map[string]*Contact
    groupsById map[string]*ContactGroup
)

func NewMockGoogleClient() oauth2_client.MockClient {
    client := oauth2_client.NewMockOAuthClient()
    client.SetRequestHandler(handleGoogleClientRequests)
    if contactsById == nil {
        contactsById = make(map[string]*Contact)
    }
    if groupsById == nil {
        groupsById = make(map[string]*ContactGroup)
    }
    return client
}

func AddMockContact(contact *Contact) {
    if contact != nil && contact.Id.Value != "" {
        contactsById[contact.Id.Value] = contact
    }
}

func AddMockGroup(group *ContactGroup) {
    if group != nil && group.Id.Value != "" {
        groupsById[group.Id.Value] = group
    }
}

func AddMockContactJSON(s string) {
    c := new(Contact)
    if err := json.Unmarshal([]byte(s), c); err == nil {
        AddMockContact(c)
    }
}

func AddMockContactJSONEntry(s string) {
    c := new(ContactEntryResponse)
    if err := json.Unmarshal([]byte(s), c); err == nil {
        AddMockContact(c.Entry)
    }
}

func AddMockContactJSONFeed(s string) {
    f := new(ContactFeedResponse)
    if err := json.Unmarshal([]byte(s), f); err == nil {
        if f.Feed != nil && f.Feed.Entries != nil {
            for _, c := range f.Feed.Entries {
                AddMockContact(&c)
            }
        }
    }
}

func AddMockGroupJSONEntry(s string) {
    c := new(GroupResponse)
    if err := json.Unmarshal([]byte(s), c); err == nil {
        AddMockGroup(c.Entry)
    }
}

func AddMockGroupJSONFeed(s string) {
    f := new(GroupsFeedResponse)
    if err := json.Unmarshal([]byte(s), f); err == nil {
        if f.Feed != nil && f.Feed.Entries != nil {
            for _, g := range f.Feed.Entries {
                AddMockGroup(&g)
            }
        }
    }
}

func handleGoogleClientRequests(client oauth2_client.MockClient, req *http.Request) (resp *http.Response, err os.Error) {
    if req == nil {
        return nil, nil
    }
    if req.URL.Host != "www.google.com" || !strings.HasPrefix(req.URL.Path, "/m8/feeds/") {
        return nil, nil
    }
    pathParts := strings.Split(req.URL.Path, "/")
    pathPartsLen := len(pathParts)
    if strings.HasPrefix(req.URL.Path, "/m8/feeds/contacts/") {
        if req.Method == oauth2_client.GET {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/contacts/default/full or
                case 5: return handleGetContactList(client, req)
                // handle /m8/feeds/contacts/default/full/34535483485345384
                case 6: return handleGetContact(client, req, pathParts[5])
                }
            }
        }
        return nil, nil
    }
    if strings.HasPrefix(req.URL.Path, "/m8/feeds/groups/") {
        if req.Method == oauth2_client.GET {
            // handle /m8/feeds/contacts/default/full or
            // handle /m8/feeds/contacts/default/full/ or
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/groups/default/full or
                case 5: return handleGetGroupList(client, req)
                // handle /m8/feeds/groups/default/full/34535483485345384
                case 6: return handleGetGroup(client, req, pathParts[5])
                }
            }
        }
        return nil, nil
    }
    if strings.HasPrefix(req.URL.Path, "/m8/feeds/photos/") {
        
    }
    return nil, nil
}

func createJSONHttpResponseFromObj(req *http.Request, obj interface{}) (resp *http.Response, err os.Error) {
    var b []byte
    if obj != nil {
        obj1, _ := jsonhelper.MarshalWithOptions(obj, GOOGLE_DATETIME_FORMAT)
        jsonobj := obj1.(jsonhelper.JSONObject)
        b, err = json.Marshal(jsonobj)
    }
    if b == nil {
        b = make([]byte, 0)
    }
    buf := bytes.NewBuffer(b)
    headers := make(http.Header)
    headers.Add("Content-Type", "application/json; charset=utf-8")
    resp = new(http.Response)
    resp.Status = "200 OK"
    resp.StatusCode = http.StatusOK
    resp.Proto = "HTTP/1.1"
    resp.ProtoMajor = 1
    resp.ProtoMinor = 1
    resp.Header = headers
    resp.Body = ioutil.NopCloser(buf)
    resp.ContentLength = int64(len(b))
    resp.TransferEncoding = []string{"identity"}
    resp.Close = true
    resp.Request = req
    return
}

func handleGetContactList(client oauth2_client.MockClient, req *http.Request) (resp *http.Response, err os.Error) {
    numContacts := len(contactsById)
    entries := make([]Contact, numContacts)
    r := NewContactFeedResponse()
    f := r.Feed
    f.Entries = entries
    i := 0
    for _, v := range contactsById {
        entries[i] = *v
        i++
    }
    userInfo, _ := client.RetrieveUserInfo()
    if userInfo != nil {
        f.Id.Value = userInfo.Guid()
        f.Author = make([]AtomAuthor, 1)
        f.Author[0].Email.Value = userInfo.Username()
        f.Author[0].Name.Value = userInfo.GivenName()
        f.Title.Value = userInfo.GivenName() + "'s Contacts"
    }
    f.Updated.Value = time.UTC().Format(GOOGLE_DATETIME_FORMAT)
    f.TotalResults.Value = strconv.Itoa(numContacts)
    f.StartIndex.Value = "1"
    f.ItemsPerPage.Value = "25"
    return createJSONHttpResponseFromObj(req, r)
}

func handleGetContact(client oauth2_client.MockClient, req *http.Request, contactId string) (resp *http.Response, err os.Error) {
    c, ok := contactsById[contactId]
    if !ok {
        resp, err = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "404 NOT FOUND"
        resp.StatusCode = http.StatusNotFound
        return
    }
    r := NewContactEntryResponse()
    r.Entry = c
    r.Entry.Xmlns = XMLNS_ATOM
    r.Entry.XmlnsGcontact = XMLNS_GCONTACT
    r.Entry.XmlnsBatch = XMLNS_GDATA_BATCH
    r.Entry.XmlnsGd = XMLNS_GD
    return createJSONHttpResponseFromObj(req, r)
}

func handleGetGroupList(client oauth2_client.MockClient, req *http.Request) (resp *http.Response, err os.Error) {
    numGroups := len(groupsById)
    entries := make([]ContactGroup, numGroups)
    r := NewGroupsFeedResponse()
    f := r.Feed
    f.Entries = entries
    i := 0
    for _, v := range groupsById {
        entries[i] = *v
        i++
    }
    userInfo, _ := client.RetrieveUserInfo()
    if userInfo != nil {
        f.Id.Value = userInfo.Guid()
        f.Author = make([]AtomAuthor, 1)
        f.Author[0].Email.Value = userInfo.Username()
        f.Author[0].Name.Value = userInfo.GivenName()
        f.Title.Value = userInfo.GivenName() + "'s Contact Groups"
    }
    f.Updated.Value = time.UTC().Format(GOOGLE_DATETIME_FORMAT)
    f.TotalResults.Value = strconv.Itoa(numGroups)
    f.StartIndex.Value = "1"
    f.ItemsPerPage.Value = "25"
    return createJSONHttpResponseFromObj(req, r)
}

func handleGetGroup(client oauth2_client.MockClient, req *http.Request, groupId string) (resp *http.Response, err os.Error) {
    g, ok := groupsById[groupId]
    if !ok {
        resp, err = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "404 NOT FOUND"
        resp.StatusCode = http.StatusNotFound
        return
    }
    r := NewGroupResponse()
    r.Entry = g
    r.Entry.Xmlns = XMLNS_ATOM
    r.Entry.XmlnsGcontact = XMLNS_GCONTACT
    r.Entry.XmlnsBatch = XMLNS_GDATA_BATCH
    r.Entry.XmlnsGd = XMLNS_GD
    return createJSONHttpResponseFromObj(req, r)
}

