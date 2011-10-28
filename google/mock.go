package google

import (
    "github.com/pomack/jsonhelper.go/jsonhelper"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "bytes"
    //"fmt"
    "http"
    "io/ioutil"
    "json"
    "os"
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
    if err := json.Unmarshal([]byte(s), c); err == nil {
        p.AddMockContact(c)
    }
}

func (p* MockGoogleClient) AddMockContactJSONEntry(s string) {
    c := new(ContactEntryResponse)
    if err := json.Unmarshal([]byte(s), c); err == nil {
        p.AddMockContact(c.Entry)
    }
}

func (p* MockGoogleClient) AddMockContactJSONFeed(s string) {
    f := new(ContactFeedResponse)
    if err := json.Unmarshal([]byte(s), f); err == nil {
        if f.Feed != nil && f.Feed.Entries != nil {
            for _, c := range f.Feed.Entries {
                p.AddMockContact(&c)
            }
        }
    }
}

func (p* MockGoogleClient) AddMockGroupJSONEntry(s string) {
    c := new(GroupResponse)
    if err := json.Unmarshal([]byte(s), c); err == nil {
        p.AddMockGroup(c.Entry)
    }
}

func (p* MockGoogleClient) AddMockGroupJSONFeed(s string) {
    f := new(GroupsFeedResponse)
    if err := json.Unmarshal([]byte(s), f); err == nil {
        if f.Feed != nil && f.Feed.Entries != nil {
            for _, g := range f.Feed.Entries {
                p.AddMockGroup(&g)
            }
        }
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
        }
    } else if strings.HasPrefix(req.URL.Path, "/m8/feeds/groups/") {
        if req.Method == oauth2_client.GET {
            // handle /m8/feeds/contacts/default/full or
            // handle /m8/feeds/contacts/default/full/ or
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
        }
    } else if strings.HasPrefix(req.URL.Path, "/m8/feeds/photos/") {
        
    }
    //fmt.Printf("Done mock handle %s\n", req.URL.String())
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

