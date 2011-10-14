package yahoo

import (
    "github.com/pomack/jsonhelper.go/jsonhelper"
    "json"
    "os"
)

type ContactsResponse struct {
    Contacts ContactsList `json:"contacts,omitempty"`
}

type ContactResponse struct {
    Contact Contact `json:"contact,omitempty"`
}

type CategoriesResponse struct {
    Categories CategoriesList `json:"categories,omitempty"`
}

type CategoryResponse struct {
    Category Category `json:"category,omitempty"`
}

type ContactsList struct {
    Start    int64     `json:"start,omitempty"`
    Count    int64     `json:"count,omitempty"`
    Total    int64     `json:"total,omitempty"`
    Uri      string    `json:"uri,omitempty"`
    Contacts []Contact `json:"contact,omitempty"`
}

type CategoriesList struct {
    Start      int64      `json:"start,omitempty"`
    Count      int64      `json:"count,omitempty"`
    Total      int64      `json:"total,omitempty"`
    Uri        string     `json:"uri,omitempty"`
    Categories []Category `json:"category,omitempty"`
}

type Contact struct {
    Uri          string         `json:"uri,omitempty"`
    Created      string         `json:"created,omitempty"`
    Updated      string         `json:"updated,omitempty"`
    Op           string         `json:"op,omitempty"`
    IsConnection bool           `json:"isConnection,omitempty"`
    Id           int64          `json:"id,omitempty"`
    Fields       []ContactField `json:"fields,omitempty"`
    Categories   []Category     `json:"categories,omitempty"`
}

type Anniversary struct {
    Day   int `json:"day,omitempty"`
    Month int `json:"month,omitempty"`
    Year  int `json:"year,omitempty"`
}

type Birthday Anniversary

type Name struct {
    GivenName       string `json:"givenName,omitempty"`
    MiddleName      string `json:"middleName,omitempty"`
    FamilyName      string `json:"familyName,omitempty"`
    Prefix          string `json:"prefix,omitempty"`
    Suffix          string `json:"suffix,omitempty"`
    GivenNameSound  string `json:"givenNameSound,omitempty"`
    FamilyNameSound string `json:"familyNameSound,omitempty"`
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

type Date struct {
    Year        int `json:"year,omitempty"`
    Month       int `json:"month,omitempty"`
    Day         int `json:"day,omitempty"`
}

func (p *Date) FromJSON(o jsonhelper.JSONObject) {
    p.Year = o.GetAsInt("year")
    p.Month = o.GetAsInt("month")
    p.Day = o.GetAsInt("day")
}

type Address struct {
    Street              string `json:"street,omitempty"`
    City                string `json:"city,omitempty"`
    StateOrProvince     string `json:"stateOrProvince,omitempty"`
    PostalCode          string `json:"postalCode,omitempty"`
    Country             string `json:"country,omitempty"`
    CountryCode         string `json:"countryCode,omitempty"`
}

func (p *Address) FromJSON(o jsonhelper.JSONObject) {
    p.Street = o.GetAsString("street")
    p.City = o.GetAsString("city")
    p.StateOrProvince = o.GetAsString("stateOrProvince")
    p.PostalCode = o.GetAsString("postalCode")
    p.Country = o.GetAsString("country")
    p.CountryCode = o.GetAsString("countryCode")
}

type ContactField struct {
    Uri        string      `json:"uri,omitempty"`
    Created    string      `json:"created,omitempty"`
    Updated    string      `json:"updated,omitempty"`
    Id         int64       `json:"id,omitempty"`
    Type       string      `json:"type,omitempty"`
    EditedBy   string      `json:"editedBy,omitempty"`
    Flags      []string    `json:"flags,omitempty"`
    Categories []string    `json:"categories,omitempty"`
    Value      interface{} `json:"value,omitempty"`
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
    switch p.Type {
    case "name":
        var name Name
        name.FromJSON(o.GetAsObject("value"))
        p.Value = name
    case "birthday", "anniversary":
        var d Date
        d.FromJSON(o.GetAsObject("value"))
        p.Value = d
    case "address":
        var addr Address
        addr.FromJSON(o.GetAsObject("value"))
        p.Value = addr
    default:
        p.Value = o.GetAsString("value")
    }
    return err
}

type Category struct {
    Uri     string `json:"uri,omitempty"`
    Created string `json:"created,omitempty"`
    Updated string `json:"updated,omitempty"`
    Id      int64  `json:"id,omitempty"`
    Name    string `json:"name,omitempty"`
}

type ContactSyncResponse struct {
    ContactSync ContactSync `json:"contactsync,omitempty"`
}

type ContactSync struct {
    Rev              int64      `json:"rev,omitempty"`
    ClientRev        int64      `json:"clientrev,omitempty"`
    AddressBookReset int64      `json:"addressbookreset,omitempty"`
    LastModifiedTime string     `json:"lmt,omitempty"`
    Categories       []Category `json:"categories,omitempty"`
    Contacts         []Contact  `json:"contacts,omitempty"`
}

type Flag string

type MeGuidResponse struct {
    Guid MeGuid `json:"guid,omitempty"`
}

type MeGuid struct {
    Uri   string `json:"uri,omitempty"`
    Value string `json:"value,omitempty"`
}

type ProfileResponse struct {
    Profile *Profile `json:"profile,omitempty"`
}

type Profile struct {
    Uri                string          `json:"uri,omitempty"`
    Guid               string          `json:"guid,omitempty"`
    Nickname           string          `json:"nickname,omitempty"`
    Image              ImageInfo       `json:"image,omitempty"`
    MemberSince        string          `json:"memberSince,omitempty"`
    ProfileUrl         string          `json:"profileUrl,omitempty"`
    DisplayAge         int             `json:"displayAge,omitempty"`
    BirthYear          int             `json:"birthYear,omitempty"`
    Birthdate          string          `json:"birthdate,omitempty"`
    Gender             string          `json:"gender,omitempty"`
    Location           string          `json:"location,omitempty"`
    Status             Status          `json:"status,omitempty"`
    IsConnected        bool            `json:"isConnected,omitempty"`
    GivenName          string          `json:"givenName,omitempty"`
    FamilyName         string          `json:"familyName,omitempty"`
    Addresses          []SocialAddress `json:"addresses,omitempty"`
    Ims                []IM            `json:"ims,omitempty"`
    Emails             []Email         `json:"emails,omitempty"`
    WorkPlaces         []WorkObject    `json:"works,omitempty"`
    Schools            []SchoolObject  `json:"schools,omitempty"`
    Websites           []Website       `json:"websites,omitempty"`
    Disclosures        []Disclosure    `json:"disclosures,omitempty"`
    TimeZone           string          `json:"timeZone,omitempty"`
    Lang               string          `json:"lang,omitempty"`
    RelationshipStatus string          `json:"relationshipStatus,omitempty"`
    LookingFor         string          `json:"lookingFor,omitempty"`
    Interests          []Interest      `json:"interests,omitempty"`
    Created            string          `json:"created,omitempty"`
    Updated            string          `json:"updated,omitempty"`
    Searchable         bool            `json:"searchable,omitempty"`
}

type ImageInfo struct {
    Size     string `json:"size,omitempty"`
    Width    int64  `json:"width,omitempty"`
    Height   int64  `json:"height,omitempty"`
    ImageUrl string `json:"imageUrl,omitempty"`
}

type Status struct {
    Uri                string `json:"uri,omitempty"`
    Message            string `json:"message,omitempty"`
    LastStatusModified string `json:"lastStatusModified,omitempty"`
}

type WorkObject struct {
    WorkName   string `json:"workName,omitempty"`
    Title      string `json:"title,omitempty"`
    Address    string `json:"address,omitempty"`
    City       string `json:"city,omitempty"`
    State      string `json:"state,omitempty"`
    PostalCode string `json:"postalCode,omitempty"`
    Country    string `json:"country,omitempty"`
    StartDate  string `json:"startDate,omitempty"`
    EndDate    string `json:"endDate,omitempty"`
    Id         int64  `json:"id,omitempty"`
    Current    bool   `json:"current,omitempty"`
}

type SchoolObject struct {
    SchoolName  string `json:"schoolName,omitempty"`
    SchoolType  string `json:"schoolType,omitempty"`
    SchoolMajor string `json:"schoolMajor,omitempty"`
    SchoolYear  string `json:"schoolYear,omitempty"`
    City        string `json:"city,omitempty"`
    State       string `json:"state,omitempty"`
    Country     string `json:"country,omitempty"`
    StartDate   string `json:"startDate,omitempty"`
    EndDate     string `json:"endDate,omitempty"`
    Id          int64  `json:"id,omitempty"`
    Current     bool   `json:"current,omitempty"`
}

type Interest struct {
    InterestCategory  string   `json:"interestCategory,omitempty"`
    DeclaredInterests []string `json:"declaredInterests,omitempty"`
}

type Phone struct {
    Type    string `json:"type,omitempty"`
    Number  string `json:"number,omitempty"`
    Primary bool   `json:"primary,omitempty"`
    Id      int64  `json:"id,omitempty"`
}

type UpdateProfileRequest struct {
    UpdateProfile UpdateProfile `json:"updateprofile,omitempty"`
}

type UpdateProfile struct {
    Nickname           string          `json:"nickname,omitempty"`
    DisplayAge         int             `json:"displayAge,omitempty"`
    Gender             string          `json:"gender,omitempty"`
    Location           string          `json:"location,omitempty"`
    GivenName          string          `json:"givenName,omitempty"`
    FamilyName         string          `json:"familyName,omitempty"`
    Works              []WorkObject    `json:"works,omitempty"`
    Schools            []SchoolObject  `json:"schools,omitempty"`
    Websites           []string        `json:"websites,omitempty"`
    Addresses          []SocialAddress `json:"addresses,omitempty"`
    Phones             []Phone         `json:"phones,omitempty"`
    Ims                []IM            `json:"ims,omitempty"`
    Emails             []Email         `json:"emails,omitempty"`
    RelationshipStatus []string        `json:"relationshipStatus,omitempty"`
    LookingFor         []string        `json:"lookingFor,omitempty"`
    LanguagesSpoken    []string        `json:"languagesSpoken,omitempty"`
    Interests          []Interest      `json:"interests,omitempty"`
    ProfileMode        string          `json:"profileMode,omitempty"`
    ProfileHidden      bool            `json:"profileHidden,omitempty"`
    ProfileStatus      string          `json:"profileStatus,omitempty"`
    Birthdate          string          `json:"birthdate,omitempty"`
    AboutMe            string          `json:"aboutMe,omitempty"`
}

type SocialAddress struct {
    Type       string `json:"type,omitempty"`
    Street     string `json:"street,omitempty"`
    City       string `json:"city,omitempty"`
    State      string `json:"state,omitempty"`
    PostalCode string `json:"postalCode,omitempty"`
    Country    string `json:"country,omitempty"`
    Id         int64  `json:"id,omitempty"`
    Current    bool   `json:"current,omitempty"`
}

type IM struct {
    Type   string `json:"type,omitempty"`
    Handle string `json:"handle,omitempty"`
    Id     int64  `json:"id,omitempty"`
}

type Email struct {
    Type    string `json:"type,omitempty"`
    Handle  string `json:"handle,omitempty"`
    Primary bool   `json:"primary,omitempty"`
    Id      int64  `json:"id,omitempty"`
}

type Website struct {
    Id  int64  `json:"id,omitempty"`
    Url string `json:"url,omitempty"`
}

type Disclosure struct {
    Name    string `json:"name,omitempty"`
    Seen    string `json:"seen,omitempty"`
    Version string `json:"version,omitempty"`
}

type NotificationsResponse struct {
    Notifications NotificationList `json:"notifications,omitempty"`
}

type NotificationList struct {
    Start        int64          `json:"start,omitempty"`
    Count        int64          `json:"count,omitempty"`
    Total        int64          `json:"total,omitempty"`
    Notification []Notification `json:"notification,omitempty"`
}

type NotificationResponse struct {
    Notification Notification `json:"notification,omitempty"`
}

type Notification struct {
    Id          string    `json:"id,omitempty"`
    AppId       string    `json:"appid,omitempty"`
    State       string    `json:"string,omitempty"`
    Created     string    `json:"created,omitempty"`
    Updated     string    `json:"updated,omitempty"`
    Recipient   string    `json:"recipient,omitempty"`
    Sender      string    `json:"sender,omitempty"`
    AppName     string    `json:"appname,omitempty"`
    Favicon     string    `json:"favicon,omitempty"`
    AppUrl      string    `json:"appurl,omitempty"`
    New         bool      `json:"new,omitempty"`
    Title       string    `json:"title,omitempty"`
    Type        string    `json:"type,omitempty"`
    Nfsimple    string    `json:"nfsimple,omitempty"`
    AppMsg      string    `json:"appmsg,omitempty"`
    Category    string    `json:"category,omitempty"`
    UserMessage string    `json:"usrmsg,omitempty"`
    Image       ImageInfo `json:"image,omitempty"`
    SendMsg     bool      `json:"sendmsg,omitempty"`
    Expiry      string    `json:"expiry,omitempty"`
    Choice      string    `json:"choice,omitempty"`
    Choices     []Choice  `json:"choices,omitempty"`
    Uri         string    `json:"uri,omitempty"`
}

type Choice struct {
    Label       string `json:"label,omitempty"`
    UrlTemplate string `json:"urltemplate,omitempty"`
}

type ConnectionsResponse struct {
    Connections ConnectionList `json:"connections,omitempty"`
}

type ConnectionList struct {
    Start      int64        `json:"start,omitempty"`
    Count      int64        `json:"count,omitempty"`
    Total      int64        `json:"total,omitempty"`
    Connection []Connection `json:"connection,omitempty"`
}

type ConnectionResponse struct {
    Connection Connection `json:"connection,omitempty"`
}

type Connection struct {
    ContactId int64  `json:"contactid,omitempty"`
    Guid      string `json:"guid,omitempty"`
    Uri       string `json:"uri,omitempty"`
    Created   string `json:"created,omitempty"`
    Updated   string `json:"updated,omitempty"`
}

type ErrorResponse struct {
    Error Error `json:"error,omitempty"`
}

func (p *ErrorResponse) String() string { return p.Error.Description }

type Error struct {
    Uri         string `json:"uri,omitempty"`
    Lang        string `json:"lang,omitempty"`
    Description string `json:"description,omitempty"`
}

func (p *Error) String() string { return p.Description }
