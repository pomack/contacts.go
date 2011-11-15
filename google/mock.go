package google

import (
    "github.com/pomack/jsonhelper.go/jsonhelper"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "bytes"
    "fmt"
    "http"
    "io/ioutil"
    "json"
    "os"
    "rand"
    "strconv"
    "strings"
    "time"
)

type MockGoogleClient struct {
    oauth2_client.MockClient
    contactsById map[string]*Contact
    groupsById map[string]*ContactGroup
}

func NewMockGoogleClient() *MockGoogleClient {
    client := oauth2_client.NewMockOAuthClient()
    m := &MockGoogleClient{
        MockClient: client,
        contactsById: make(map[string]*Contact),
        groupsById: make(map[string]*ContactGroup),
    }
    gclient := oauth2_client.NewGoogleClient()
    client.SetServiceId(gclient.ServiceId())
    client.SetRequestHandler(makeHandlerForGoogleClientRequests(m))
    return m
}

func makeHandlerForGoogleClientRequests(gclient *MockGoogleClient) oauth2_client.MockRequestHandler {
    return func(client oauth2_client.MockClient, req *http.Request) (*http.Response, os.Error) {
        return gclient.handleGoogleClientRequests(req)
    }
}

func (p* MockGoogleClient) AddMockContact(contact *Contact) {
    if contact != nil && contact.Id.Value != "" {
        c := new(Contact)
        *c = *contact
        p.contactsById[contact.Id.Value] = c
    }
}

func (p* MockGoogleClient) AddMockGroup(group *ContactGroup) {
    if group != nil && group.Id.Value != "" {
        g := new(ContactGroup)
        *g = *group
        p.groupsById[group.Id.Value] = g
    }
}

func (p* MockGoogleClient) AddMockContactJSON(s string) {
    c := new(Contact)
    var err os.Error
    if err = json.Unmarshal([]byte(s), c); err == nil {
        p.AddMockContact(c)
    } else {
        fmt.Printf("error adding mock contact JSON: %s\n\n", err.String())
    }
}

func (p* MockGoogleClient) AddMockContactJSONEntry(s string) {
    c := new(ContactEntryResponse)
    var err os.Error
    if err = json.Unmarshal([]byte(s), c); err == nil {
        p.AddMockContact(c.Entry)
    } else {
        fmt.Printf("error adding mock contact JSON Entry: %s\n\n", err.String())
    }
}

func (p* MockGoogleClient) AddMockContactJSONFeed(s string) {
    f := new(ContactFeedResponse)
    var err os.Error
    if err = json.Unmarshal([]byte(s), f); err == nil {
        if f.Feed != nil && f.Feed.Entries != nil {
            for _, c := range f.Feed.Entries {
                p.AddMockContact(&c)
            }
        }
    } else {
        fmt.Printf("error adding mock contact json feed: %s\n\n", err.String())
    }
}

func (p* MockGoogleClient) AddMockGroupJSONEntry(s string) {
    c := new(GroupResponse)
    var err os.Error
    if err = json.Unmarshal([]byte(s), c); err == nil {
        p.AddMockGroup(c.Entry)
    } else {
        fmt.Printf("error adding mock group json entry: %s\n\n", err.String())
    }
}

func (p* MockGoogleClient) AddMockGroupJSONFeed(s string) {
    f := new(GroupsFeedResponse)
    var err os.Error
    if err = json.Unmarshal([]byte(s), f); err == nil {
        if f.Feed != nil && f.Feed.Entries != nil {
            for _, g := range f.Feed.Entries {
                p.AddMockGroup(&g)
            }
        }
    } else {
        fmt.Printf("error adding mock group json feed: %s\n\n", err.String())
    }
}

func (p* MockGoogleClient) handleGoogleClientRequests(req *http.Request) (resp *http.Response, err os.Error) {
    if req == nil {
        return nil, nil
    }
    if req.URL.Host != "www.google.com" || !strings.HasPrefix(req.URL.Path, "/m8/feeds/") {
        return nil, nil
    }
    pathParts := strings.Split(req.URL.Path, "/")
    pathPartsLen := len(pathParts)
    //fmt.Printf("Mock handling %s with parts: %v\n", req.URL.String(), pathParts)
    if strings.HasPrefix(req.URL.Path, "/m8/feeds/contacts/") {
        if req.Method == oauth2_client.GET {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/contacts/default/full or
                case 6:
                    resp, err = p.handleGetContactList(req)
                // handle /m8/feeds/contacts/default/full/34535483485345384
                case 7:
                    resp, err = p.handleGetContact(req, pathParts[5])
                }
            }
        } else if req.Method == oauth2_client.POST {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/contacts/default/full or
                case 6:
                    resp, err = p.handleCreateContact(req)
                }
            }
        } else if req.Method == oauth2_client.PUT {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/contacts/default/full/34535483485345384
                case 7:
                    resp, err = p.handleUpdateContact(req, pathParts[5])
                }
            }
        } else if req.Method == oauth2_client.DELETE {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/contacts/default/full/34535483485345384
                case 7:
                    resp, err = p.handleDeleteContact(req, pathParts[5])
                }
            }
        }
    } else if strings.HasPrefix(req.URL.Path, "/m8/feeds/groups/") {
        if req.Method == oauth2_client.GET {
            // handle /m8/feeds/groups/default/full or
            // handle /m8/feeds/groups/default/full/ or
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/groups/default/full or
                case 6: 
                    resp, err = p.handleGetGroupList(req)
                // handle /m8/feeds/groups/default/full/34535483485345384
                case 7: 
                    resp, err = p.handleGetGroup(req, pathParts[5])
                }
            }
        } else if req.Method == oauth2_client.POST {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/groups/default/full or
                case 6:
                    resp, err = p.handleCreateGroup(req)
                }
            }
        } else if req.Method == oauth2_client.PUT {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/groups/default/full/34535483485345384
                case 7:
                    resp, err = p.handleUpdateGroup(req, pathParts[5])
                }
            }
        } else if req.Method == oauth2_client.DELETE {
            if pathParts[0] == "" {
                switch pathPartsLen {
                // handle /m8/feeds/groups/default/full/34535483485345384
                case 7:
                    resp, err = p.handleDeleteGroup(req, pathParts[5])
                }
            }
        }
    } else if strings.HasPrefix(req.URL.Path, "/m8/feeds/photos/") {
        
    }
    fmt.Printf("[MOCK-GOOGLE]: Done mock handle %s %s\n", req.Method, req.URL.String())
    return
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

func (p* MockGoogleClient) handleGetContactList(req *http.Request) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleGetContactList\n")
    numContacts := len(p.contactsById)
    entries := make([]Contact, numContacts)
    r := NewContactFeedResponse()
    f := r.Feed
    f.Entries = entries
    i := 0
    for _, v := range p.contactsById {
        entries[i] = *v
        i++
    }
    userInfo, _ := p.RetrieveUserInfo()
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
    //fmt.Printf("Contact List Entry Count: %d\n", numContacts)
    return createJSONHttpResponseFromObj(req, r)
}

func (p* MockGoogleClient) handleGetContact(req *http.Request, contactId string) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleGetContact\n")
    c, ok := p.contactsById[contactId]
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
    //fmt.Printf("Contact Entry: %v\n", c.Name)
    return createJSONHttpResponseFromObj(req, r)
}

func (p *MockGoogleClient) generateId() string {
    arr := make([]string, 8)
    for i := 0; i < 8; i++ {
        arr[i] = fmt.Sprintf("%02x", rand.Intn(256))
    }
    return strings.Join(arr, "")
}

func (p *MockGoogleClient) generateEtag() string {
    arr := make([]string, 14)
    for i := 0; i < 14; i++ {
        arr[i] = fmt.Sprintf("%02x", rand.Intn(256))
    }
    return "\"" + strings.Join(arr, "") + "\""
}

func (p* MockGoogleClient) handleCreateContact(req *http.Request) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleCreateContact\n")
    b := NewContactEntryResponse()
    err = json.NewDecoder(req.Body).Decode(b)
    if err != nil {
        panic(fmt.Sprintf("Error creating contact: %s\n", err.String()))
        resp, _ = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "500 INTERNAL SERVER ERROR"
        resp.StatusCode = http.StatusInternalServerError
        return
    }
    if b.Entry == nil {
        resp, err = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "400 BAD REQUEST"
        resp.StatusCode = http.StatusBadRequest
        return
    }
    e := b.Entry
    contactId := p.generateId()
    e.Id.Value = contactId
    e.Etag = p.generateEtag()
    e.Updated.Value = time.UTC().Format(GOOGLE_DATETIME_FORMAT)
    p.contactsById[contactId] = e
    r := NewContactEntryResponse()
    r.Entry = e
    r.Entry.Xmlns = XMLNS_ATOM
    r.Entry.XmlnsGcontact = XMLNS_GCONTACT
    r.Entry.XmlnsBatch = XMLNS_GDATA_BATCH
    r.Entry.XmlnsGd = XMLNS_GD
    //fmt.Printf("Contact Entry: %v\n", c.Name)
    return createJSONHttpResponseFromObj(req, r)
}

func (p* MockGoogleClient) handleUpdateContact(req *http.Request, contactId string) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleUpdateContact\n")
    c, ok := p.contactsById[contactId]
    if !ok {
        resp, err = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "404 NOT FOUND"
        resp.StatusCode = http.StatusNotFound
        return
    }
    ifmatch := req.Header.Get("If-Match")
    if ifmatch != "*" && ifmatch != c.Etag {
        resp, err = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "412 PRECONDITION FAILED"
        resp.StatusCode = http.StatusPreconditionFailed
        return
    }
    b := NewContactEntryResponse()
    err = json.NewDecoder(req.Body).Decode(b)
    if err != nil {
        panic(fmt.Sprintf("Error updating contact: %s\n", err.String()))
        resp, _ = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "500 INTERNAL SERVER ERROR"
        resp.StatusCode = http.StatusInternalServerError
        return
    }
    if b.Entry == nil {
        resp, err = createJSONHttpResponseFromObj(req, NewContactEntryResponse())
        resp.Status = "400 BAD REQUEST"
        resp.StatusCode = http.StatusBadRequest
        return
    }
    e := b.Entry
    e.Id.Value = contactId
    e.Etag = p.generateEtag()
    e.Updated.Value = time.UTC().Format(GOOGLE_DATETIME_FORMAT)
    p.contactsById[contactId] = e
    r := NewContactEntryResponse()
    r.Entry = e
    r.Entry.Xmlns = XMLNS_ATOM
    r.Entry.XmlnsGcontact = XMLNS_GCONTACT
    r.Entry.XmlnsBatch = XMLNS_GDATA_BATCH
    r.Entry.XmlnsGd = XMLNS_GD
    //fmt.Printf("Contact Entry: %v\n", c.Name)
    return createJSONHttpResponseFromObj(req, r)
}


func (p* MockGoogleClient) handleDeleteContact(req *http.Request, contactId string) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleDeleteContact\n")
    c, ok := p.contactsById[contactId]
    if !ok {
        resp, err = createJSONHttpResponseFromObj(req, nil)
        resp.Status = "404 NOT FOUND"
        resp.StatusCode = http.StatusNotFound
        return
    }
    ifmatch := req.Header.Get("If-Match")
    if ifmatch != "*" && ifmatch != c.Etag {
        resp, err = createJSONHttpResponseFromObj(req, nil)
        resp.Status = "412 PRECONDITION FAILED"
        resp.StatusCode = http.StatusPreconditionFailed
        return
    }
    p.contactsById[contactId] = c, false
    resp, err = createJSONHttpResponseFromObj(req, nil)
    return
}

func (p* MockGoogleClient) handleGetGroupList(req *http.Request) (resp *http.Response, err os.Error) {
    numGroups := len(p.groupsById)
    entries := make([]ContactGroup, numGroups)
    r := NewGroupsFeedResponse()
    f := r.Feed
    f.Entries = entries
    i := 0
    for _, v := range p.groupsById {
        entries[i] = *v
        i++
    }
    userInfo, _ := p.RetrieveUserInfo()
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
    //fmt.Printf("Group List Entry Count: %d\n", numGroups)
    return createJSONHttpResponseFromObj(req, r)
}

func (p* MockGoogleClient) handleGetGroup(req *http.Request, groupId string) (resp *http.Response, err os.Error) {
    g, ok := p.groupsById[groupId]
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
    //fmt.Printf("Group Entry: %s\n", g.Title.Value)
    return createJSONHttpResponseFromObj(req, r)
}

func (p* MockGoogleClient) handleCreateGroup(req *http.Request) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleCreateGroup\n")
    b := NewGroupResponse()
    err = json.NewDecoder(req.Body).Decode(b)
    if err != nil {
        panic(fmt.Sprintf("Error creating group: %s\n", err.String()))
        resp, _ = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "500 INTERNAL SERVER ERROR"
        resp.StatusCode = http.StatusInternalServerError
        return
    }
    if b.Entry == nil {
        resp, err = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "400 BAD REQUEST"
        resp.StatusCode = http.StatusBadRequest
        return
    }
    e := b.Entry
    groupId := p.generateId()
    e.Id.Value = groupId
    e.Etag = p.generateEtag()
    e.Updated.Value = time.UTC().Format(GOOGLE_DATETIME_FORMAT)
    p.groupsById[groupId] = e
    r := NewGroupResponse()
    r.Entry = e
    r.Entry.Xmlns = XMLNS_ATOM
    r.Entry.XmlnsGcontact = XMLNS_GCONTACT
    r.Entry.XmlnsBatch = XMLNS_GDATA_BATCH
    r.Entry.XmlnsGd = XMLNS_GD
    //fmt.Printf("Contact Entry: %v\n", c.Name)
    return createJSONHttpResponseFromObj(req, r)
}

func (p* MockGoogleClient) handleUpdateGroup(req *http.Request, groupId string) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleUpdateGroup\n")
    g, ok := p.groupsById[groupId]
    if !ok {
        resp, err = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "404 NOT FOUND"
        resp.StatusCode = http.StatusNotFound
        return
    }
    ifmatch := req.Header.Get("If-Match")
    if ifmatch != "*" && ifmatch != g.Etag {
        resp, err = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "412 PRECONDITION FAILED"
        resp.StatusCode = http.StatusPreconditionFailed
        return
    }
    b := NewGroupResponse()
    err = json.NewDecoder(req.Body).Decode(b)
    if err != nil {
        panic(fmt.Sprintf("Error updating group: %s\n", err.String()))
        resp, _ = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "500 INTERNAL SERVER ERROR"
        resp.StatusCode = http.StatusInternalServerError
        return
    }
    if b.Entry == nil {
        resp, err = createJSONHttpResponseFromObj(req, NewGroupResponse())
        resp.Status = "400 BAD REQUEST"
        resp.StatusCode = http.StatusBadRequest
        return
    }
    e := b.Entry
    e.Id.Value = groupId
    e.Etag = p.generateEtag()
    e.Updated.Value = time.UTC().Format(GOOGLE_DATETIME_FORMAT)
    p.groupsById[groupId] = e
    r := NewGroupResponse()
    r.Entry = e
    r.Entry.Xmlns = XMLNS_ATOM
    r.Entry.XmlnsGcontact = XMLNS_GCONTACT
    r.Entry.XmlnsBatch = XMLNS_GDATA_BATCH
    r.Entry.XmlnsGd = XMLNS_GD
    //fmt.Printf("Contact Entry: %v\n", c.Name)
    return createJSONHttpResponseFromObj(req, r)
}


func (p* MockGoogleClient) handleDeleteGroup(req *http.Request, groupId string) (resp *http.Response, err os.Error) {
    //fmt.Printf("calling handleDeleteGroup\n")
    g, ok := p.groupsById[groupId]
    if !ok {
        resp, err = createJSONHttpResponseFromObj(req, nil)
        resp.Status = "404 NOT FOUND"
        resp.StatusCode = http.StatusNotFound
        return
    }
    ifmatch := req.Header.Get("If-Match")
    if ifmatch != "*" && ifmatch != g.Etag {
        resp, err = createJSONHttpResponseFromObj(req, nil)
        resp.Status = "412 PRECONDITION FAILED"
        resp.StatusCode = http.StatusPreconditionFailed
        return
    }
    p.groupsById[groupId] = g, false
    resp, err = createJSONHttpResponseFromObj(req, nil)
    return
}


