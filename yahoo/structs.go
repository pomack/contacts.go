package yahoo

import (
    "github.com/pomack/jsonhelper"
    "json"
    "os"
)


type ContactsResponse struct {
    Contacts        ContactsList        `json:"contacts"`
}

type ContactsList struct {
    Start           int                 `json:"start"`
    Count           int                 `json:"count"`
    Total           int                 `json:"total"`
    Uri             string              `json:"uri"`
    Contacts        []Contact           `json:"contact"`
}

type Contact struct {
    Uri             string              `json:"uri"`
    Created         string              `json:"created"`
    Updated         string              `json:"updated"`
    Op              string              `json:"op"`
    IsConnection    bool                `json:"isConnection"`
    Id              int64               `json:"id"`
    Fields          []ContactField      `json:"fields"`
    Categories      []Category          `json:"categories"`
}


type Address struct {
    Street          string              `json:"street"`
    City            string              `json:"city"`
    StateOrProvince string              `json:"stateOrProvince"`
    PostalCode      string              `json:"postalCode"`
    Country         string              `json:"country"`
    CountryCode     string              `json:"countryCode"`
}

type Anniversary struct {
    Day             int                 `json:"day"`
    Month           int                 `json:"month"`
    Year            int                 `json:"year"`
}

type Birthday Anniversary

type Name struct {
    GivenName       string              `json:"givenName"`
    MiddleName      string              `json:"middleName"`
    FamilyName      string              `json:"familyName"`
    Prefix          string              `json:"prefix"`
    Suffix          string              `json:"suffix"`
    GivenNameSound  string              `json:"givenNameSound"`
    FamilyNameSound string              `json:"familyNameSound"`
}
func (p *Name) FromJSON(o jsonhelper.JSONObject) {
    p.GivenName = o.GetAsString("givenName")
    p.MiddleName = o.GetAsString("middleName")
    p.FamilyName = o.GetAsString("familyName")
    p.Prefix = o.GetAsString("prefix")
    p.Suffix = o.GetAsString("suffix")
    p.GivenNameSound = o.GetAsString("givenNameSound")
    p.FamilyNameSound = o.GetAsString("familyNameSound")
}

type ContactField struct {
    Uri             string              `json:"uri"`
    Created         string              `json:"created"`
    Updated         string              `json:"updated"`
    Id              int64               `json:"id"`
    Type            string              `json:"type"`
    EditedBy        string              `json:"editedBy"`
    Flags           []string            `json:"flags"`
    Categories      []string            `json:"categories"`
    Value           interface{}         `json:"value"`
}
func (p *ContactField) UnmarshalJSON(data []byte) os.Error {
    o := jsonhelper.NewJSONObject()
    err := json.Unmarshal(data, &o)
    p.Uri = o.GetAsString("uri")
    p.Created = o.GetAsString("created")
    p.Updated = o.GetAsString("updated")
    p.Id = o.GetAsInt64("id")
    p.Type = o.GetAsString("type")
    p.EditedBy = o.GetAsString("editedBy")
    flags := o.GetAsArray("flags")
    lf := flags.Len()
    p.Flags = make([]string, lf)
    for i := 0; i < lf; i++ {
        p.Flags[i] = flags.GetAsString(i)
    }
    categories := o.GetAsArray("categories")
    lc := categories.Len()
    p.Categories = make([]string, lc)
    for i := 0; i < lc; i++ {
        p.Categories[i] = categories.GetAsString(i)
    }
    if p.Type == "name" {
        var name Name
        name.FromJSON(o.GetAsObject("value"))
        p.Value = name
    } else {
        p.Value = o.GetAsString("value")
    }
    return err
}


type Category struct {
    Uri             string              `json:"uri"`
    Created         string              `json:"created"`
    Updated         string              `json:"updated"`
    Id              int64               `json:"id"`
    Name            string              `json:"name"`
}

type ContactSyncResponse struct {
    ContactSync         ContactSync         `json:"contactsync"`
}

type ContactSync struct {
    Rev                 int64               `json:"rev"`
    ClientRev           int64               `json:"clientrev"`
    AddressBookReset    int64               `json:"addressbookreset"`
    LastModifiedTime    string              `json:"lmt"`
    Categories          []Category          `json:"categories"`
    Contacts            []Contact           `json:"contacts"`
}

type Flag string


