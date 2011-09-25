package yahoo

import (
	"github.com/pomack/jsonhelper"
	"json"
	"os"
)

type ContactsResponse struct {
	Contacts ContactsList `json:"contacts"`
}

type ContactResponse struct {
	Contact Contact `json:"contact"`
}

type CategoriesResponse struct {
	Categories CategoriesList `json:"categories"`
}

type CategoryResponse struct {
	Category Category `json:"category"`
}

type ContactsList struct {
	Start    int64     `json:"start"`
	Count    int64     `json:"count"`
	Total    int64     `json:"total"`
	Uri      string    `json:"uri"`
	Contacts []Contact `json:"contact"`
}

type CategoriesList struct {
	Start      int64      `json:"start"`
	Count      int64      `json:"count"`
	Total      int64      `json:"total"`
	Uri        string     `json:"uri"`
	Categories []Category `json:"category"`
}

type Contact struct {
	Uri          string         `json:"uri"`
	Created      string         `json:"created"`
	Updated      string         `json:"updated"`
	Op           string         `json:"op"`
	IsConnection bool           `json:"isConnection"`
	Id           int64          `json:"id"`
	Fields       []ContactField `json:"fields"`
	Categories   []Category     `json:"categories"`
}

type Address struct {
	Street          string `json:"street"`
	City            string `json:"city"`
	StateOrProvince string `json:"stateOrProvince"`
	PostalCode      string `json:"postalCode"`
	Country         string `json:"country"`
	CountryCode     string `json:"countryCode"`
}

type Anniversary struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Birthday Anniversary

type Name struct {
	GivenName       string `json:"givenName"`
	MiddleName      string `json:"middleName"`
	FamilyName      string `json:"familyName"`
	Prefix          string `json:"prefix"`
	Suffix          string `json:"suffix"`
	GivenNameSound  string `json:"givenNameSound"`
	FamilyNameSound string `json:"familyNameSound"`
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
	Uri        string      `json:"uri"`
	Created    string      `json:"created"`
	Updated    string      `json:"updated"`
	Id         int64       `json:"id"`
	Type       string      `json:"type"`
	EditedBy   string      `json:"editedBy"`
	Flags      []string    `json:"flags"`
	Categories []string    `json:"categories"`
	Value      interface{} `json:"value"`
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
	Uri     string `json:"uri"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Id      int64  `json:"id"`
	Name    string `json:"name"`
}

type ContactSyncResponse struct {
	ContactSync ContactSync `json:"contactsync"`
}

type ContactSync struct {
	Rev              int64      `json:"rev"`
	ClientRev        int64      `json:"clientrev"`
	AddressBookReset int64      `json:"addressbookreset"`
	LastModifiedTime string     `json:"lmt"`
	Categories       []Category `json:"categories"`
	Contacts         []Contact  `json:"contacts"`
}

type Flag string

type MeGuidResponse struct {
	Guid MeGuid `json:"guid"`
}

type MeGuid struct {
	Uri   string `json:"uri"`
	Value string `json:"value"`
}

type ProfileResponse struct {
	Profile *Profile `json:"profile"`
}

type Profile struct {
	Uri                string          `json:"uri"`
	Guid               string          `json:"guid"`
	Nickname           string          `json:"nickname"`
	Image              ImageInfo       `json:"image"`
	MemberSince        string          `json:"memberSince"`
	ProfileUrl         string          `json:"profileUrl"`
	DisplayAge         int             `json:"displayAge"`
	BirthYear          int             `json:"birthYear"`
	Birthdate          string          `json:"birthdate"`
	Gender             string          `json:"gender"`
	Location           string          `json:"location"`
	Status             Status          `json:"status"`
	IsConnected        bool            `json:"isConnected"`
	GivenName          string          `json:"givenName"`
	FamilyName         string          `json:"familyName"`
	Addresses          []SocialAddress `json:"addresses"`
	Ims                []IM            `json:"ims"`
	Emails             []Email         `json:"emails"`
	WorkPlaces         []WorkObject    `json:"works"`
	Schools            []SchoolObject  `json:"schools"`
	Websites           []Website       `json:"websites"`
	Disclosures        []Disclosure    `json:"disclosures"`
	TimeZone           string          `json:"timeZone"`
	Lang               string          `json:"lang"`
	RelationshipStatus string          `json:"relationshipStatus"`
	LookingFor         string          `json:"lookingFor"`
	Interests          []Interest      `json:"interests"`
	Created            string          `json:"created"`
	Updated            string          `json:"updated"`
	Searchable         bool            `json:"searchable"`
}

type ImageInfo struct {
	Size     string `json:"size"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	ImageUrl string `json:"imageUrl"`
}

type Status struct {
	Uri                string `json:"uri"`
	Message            string `json:"message"`
	LastStatusModified string `json:"lastStatusModified"`
}

type WorkObject struct {
	WorkName   string `json:"workName"`
	Title      string `json:"title"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	Id         int64  `json:"id"`
	Current    bool   `json:"current"`
}

type SchoolObject struct {
	SchoolName  string `json:"schoolName"`
	SchoolType  string `json:"schoolType"`
	SchoolMajor string `json:"schoolMajor"`
	SchoolYear  string `json:"schoolYear"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Id          int64  `json:"id"`
	Current     bool   `json:"current"`
}

type Interest struct {
	InterestCategory  string   `json:"interestCategory"`
	DeclaredInterests []string `json:"declaredInterests"`
}

type Phone struct {
	Type    string `json:"type"`
	Number  string `json:"number"`
	Primary bool   `json:"primary"`
	Id      int64  `json:"id"`
}

type UpdateProfileRequest struct {
	UpdateProfile UpdateProfile `json:"updateprofile"`
}

type UpdateProfile struct {
	Nickname           string          `json:"nickname"`
	DisplayAge         int             `json:"displayAge"`
	Gender             string          `json:"gender"`
	Location           string          `json:"location"`
	GivenName          string          `json:"givenName"`
	FamilyName         string          `json:"familyName"`
	Works              []WorkObject    `json:"works"`
	Schools            []SchoolObject  `json:"schools"`
	Websites           []string        `json:"websites"`
	Addresses          []SocialAddress `json:"addresses"`
	Phones             []Phone         `json:"phones"`
	Ims                []IM            `json:"ims"`
	Emails             []Email         `json:"emails"`
	RelationshipStatus []string        `json:"relationshipStatus"`
	LookingFor         []string        `json:"lookingFor"`
	LanguagesSpoken    []string        `json:"languagesSpoken"`
	Interests          []Interest      `json:"interests"`
	ProfileMode        string          `json:"profileMode"`
	ProfileHidden      bool            `json:"profileHidden"`
	ProfileStatus      string          `json:"profileStatus"`
	Birthdate          string          `json:"birthdate"`
	AboutMe            string          `json:"aboutMe"`
}

type SocialAddress struct {
	Type       string `json:"type"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
	Id         int64  `json:"id"`
	Current    bool   `json:"current"`
}

type IM struct {
	Type   string `json:"type"`
	Handle string `json:"handle"`
	Id     int64  `json:"id"`
}

type Email struct {
	Type    string `json:"type"`
	Handle  string `json:"handle"`
	Primary bool   `json:"primary"`
	Id      int64  `json:"id"`
}

type Website struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

type Disclosure struct {
	Name    string `json:"name"`
	Seen    string `json:"seen"`
	Version string `json:"version"`
}

type NotificationsResponse struct {
	Notifications NotificationList `json:"notifications"`
}

type NotificationList struct {
	Start        int64          `json:"start"`
	Count        int64          `json:"count"`
	Total        int64          `json:"total"`
	Notification []Notification `json:"notification"`
}

type NotificationResponse struct {
	Notification Notification `json:"notification"`
}

type Notification struct {
	Id          string    `json:"id"`
	AppId       string    `json:"appid"`
	State       string    `json:"string"`
	Created     string    `json:"created"`
	Updated     string    `json:"updated"`
	Recipient   string    `json:"recipient"`
	Sender      string    `json:"sender"`
	AppName     string    `json:"appname"`
	Favicon     string    `json:"favicon"`
	AppUrl      string    `json:"appurl"`
	New         bool      `json:"new"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Nfsimple    string    `json:"nfsimple"`
	AppMsg      string    `json:"appmsg"`
	Category    string    `json:"category"`
	UserMessage string    `json:"usrmsg"`
	Image       ImageInfo `json:"image"`
	SendMsg     bool      `json:"sendmsg"`
	Expiry      string    `json:"expiry"`
	Choice      string    `json:"choice"`
	Choices     []Choice  `json:"choices"`
	Uri         string    `json:"uri"`
}

type Choice struct {
	Label       string `json:"label"`
	UrlTemplate string `json:"urltemplate"`
}

type ConnectionsResponse struct {
	Connections ConnectionList `json:"connections"`
}

type ConnectionList struct {
	Start      int64        `json:"start"`
	Count      int64        `json:"count"`
	Total      int64        `json:"total"`
	Connection []Connection `json:"connection"`
}

type ConnectionResponse struct {
	Connection Connection `json:"connection"`
}

type Connection struct {
	ContactId int64  `json:"contactid"`
	Guid      string `json:"guid"`
	Uri       string `json:"uri"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func (p *ErrorResponse) String() string { return p.Error.Description }

type Error struct {
	Uri         string `json:"uri"`
	Lang        string `json:"lang"`
	Description string `json:"description"`
}

func (p *Error) String() string { return p.Description }
