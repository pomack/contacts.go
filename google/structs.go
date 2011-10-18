package google

import (
    "strings"
    "url"
)

type AtomEntry interface{}

type AtomText struct {
    Value string `json:"$t,omitempty"`
}

type AtomName struct {
    Value string `json:"$t,omitempty"`
    Yomi  string `json:"yomi,omitempty"`
}

type AtomId AtomText
type AtomUpdated AtomText

type AtomCategory struct {
    Scheme string `json:"scheme,omitempty"`
    Term   string `json:"term,omitempty"`
}

type AtomLink struct {
    Rel  string `json:"rel,omitempty"`
    Type string `json:"type,omitempty"`
    Href string `json:"href,omitempty"`
    Etag string `json:"gd$etag,omitempty"`
}

type AtomTitle struct {
    Value string `json:"$t,omitempty"`
    Type  string `json:"type,omitempty"`
}

type AtomContent struct {
    Value string `json:"$t,omitempty"`
    Type  string `json:"type,omitempty"`
}

type AtomAuthor struct {
    Name  AtomText `json:"name,omitempty"`
    Email AtomText `json:",omitempty"`
}

type AtomFeed struct {
    Id      AtomText    `json:"id,omitempty"`
    Title   AtomText    `json:"title,omitempty"`
    Content AtomContent `json:"content,omitempty"`
    Entries []AtomEntry `json:"entry,omitempty"`
}

type AtomGenerator struct {
    Value   string `json:"$t,omitempty"`
    Version string `json:"version,omitempty"`
    Uri     string `json:"uri,omitempty"`
}

type ContactFeedResponse struct {
    Version  string       `json:"version,omitempty"`
    Encoding string       `json:"encoding,omitempty"`
    Feed     *ContactFeed `json:"feed,omitempty"`
}

type ContactEntryResponse struct {
    Version  string   `json:"version,omitempty"`
    Encoding string   `json:"encoding,omitempty"`
    Entry    *Contact `json:"entry,omitempty"`
}

type ContactFeed struct {
    Id           AtomId         `json:"id,omitempty"`
    Updated      AtomUpdated    `json:"updated,omitempty"`
    Categories   []AtomCategory `json:"category,omitempty"`
    Title        AtomTitle      `json:"title,omitempty"`
    Links        []AtomLink     `json:"link,omitempty"`
    Generator    AtomGenerator  `json:"generator,omitempty"`
    TotalResults AtomText       `json:"openSearch$totalResults,omitempty"`
    StartIndex   AtomText       `json:"openSearch$startIndex,omitempty"`
    ItemsPerPage AtomText       `json:"openSearch$itemsPerPage,omitempty"`
    Entries      []Contact      `json:"entry,omitempty"`
}

type Contact struct {
    Id                        AtomId                    `json:"id,omitempty"`
    Xmlns                     string                    `json:"xmlns,omitempty"`
    XmlnsGcontact             string                    `json:"xmlns$gContact,omitempty"`
    XmlnsBatch                string                    `json:"xmlns$batch,omitempty"`
    XmlnsGd                   string                    `json:"xmlns$gd,omitempty"`
    Etag                      string                    `json:"gd$etag,omitempty"`
    Updated                   AtomUpdated               `json:"updated,omitempty"`
    Categories                []AtomCategory            `json:"category,omitempty"`
    Title                     AtomTitle                 `json:"title,omitempty"`
    Content                   AtomContent               `json:"content,omitempty"`
    Links                     []AtomLink                `json:"link,omitempty"`
    Deleted                   *DeletedMarker            `json:"gd$deleted,omitempty"`
    Name                      Name                      `json:"gd$name,omitempty"`
    Organizations             []Organization            `json:"gd$organization,omitempty"`
    Emails                    []Email                   `json:"gd$email,omitempty"`
    Ims                       []Im                      `json:"gd$im,omitempty"`
    PhoneNumbers              []PhoneNumber             `json:"gd$phoneNumber,omitempty"`
    PostalAddresses           []PostalAddress           `json:"gd$postalAddress,omitempty"`
    StructuredPostalAddresses []StructuredPostalAddress `json:"gd$structuredPostalAddress,omitempty"`
    Where                     Where                     `json:"gd$where,omitempty"`
    BillingInformation        BillingInformation        `json:"gContact$billingInformation,omitempty"`
    Birthday                  Birthday                  `json:"gContact$birthday,omitempty"`
    CalendarLinks             []CalendarLink            `json:"gContact$calendarLink,omitempty"`
    DirectoryServer           DirectoryServer           `json:"gContact$directoryServer,omitempty"`
    Events                    []Event                   `json:"gContact$event,omitempty"`
    ExternalId                []ExternalId              `json:"gContact$externalId,omitempty"`
    Gender                    Gender                    `json:"gContact$gender,omitempty"`
    Groups                    []GroupMembershipInfo     `json:"gContact$groupMembershipInfo,omitempty"`
    Hobbies                   []Hobby                   `json:"gContact$hobby,omitempty"`
    Initials                  Initials                  `json:"gContact$initials,omitempty"`
    Jots                      []Jot                     `json:"gContact$jot,omitempty"`
    Languages                 []Language                `json:"gContact$language,omitempty"`
    MaidenName                MaidenName                `json:"gContact$maidenName,omitempty"`
    Mileage                   Mileage                   `json:"gContact$mileage,omitempty"`
    Nickname                  Nickname                  `json:"gContact$nickname,omitempty"`
    Occupation                Occupation                `json:"gContact$occupation,omitempty"`
    Priority                  Priority                  `json:"gContact$priority,omitempty"`
    Relationships             []Relation                `json:"gContact$relation,omitempty"`
    Sensitivity               Sensitivity               `json:"gContact$sensitivity,omitempty"`
    ShortName                 ShortName                 `json:"gContact$shortName,omitempty"`
    Subject                   Subject                   `json:"gContact$subject,omitempty"`
    UserDefinedFields         []UserDefinedField        `json:"gContact$userDefinedField,omitempty"`
    Websites                  []Website                 `json:"gContact$website,omitempty"`
}

func (p *Contact) ContactId() string {
    arr := strings.Split(p.Id.Value, "/")
    if len(arr) > 0 {
        return arr[len(arr)-1]
    }
    return ""
}
func (p *Contact) SetContactId(contactId string) {
    p.Id.Value = strings.Join([]string{GOOGLE_CONTACTS_API_ENDPOINT, "default/full/", url.QueryEscape(contactId)}, "")
}

type AdditionalName AtomName

type Comments struct {
    Rel      string   `json:"rel,omitempty"`
    FeedLink FeedLink `json:"gd$feedLink,omitempty"`
}

type Country struct {
    Value string `json:"$t,omitempty"`
    Code  string `json:"code,omitempty"`
}

type DeletedMarker struct{}

type Email struct {
    Address     string `json:"address,omitempty"`
    DisplayName string `json:"displayName,omitempty"`
    Label       string `json:"label,omitempty"`
    Rel         string `json:"rel,omitempty"`
    Primary     string `json:"primary,omitempty"`
}

type EntryLink struct {
    Href     string    `json:"href,omitempty"`
    ReadOnly string    `json:"readOnly,omitempty"`
    Rel      string    `json:"rel,omitempty"`
    Entry    AtomEntry `json:"entry,omitempty"`
}

type ExtendedProperty struct {
    Name        string   `json:"name,omitempty"`
    Description AtomText `json:"description,omitempty"`
}

type FamilyName AtomName

type FeedLink struct {
    CountHint string   `json:"countHint,omitempty"`
    Href      string   `json:"href,omitempty"`
    ReadOnly  string   `json:"readOnly,omitempty"`
    Rel       string   `json:"rel,omitempty"`
    AtomFeed  AtomFeed `json:"feed,omitempty"`
}

type GivenName AtomName

type Im struct {
    Address  string `json:"address,omitempty"`
    Label    string `json:"label,omitempty"`
    Rel      string `json:"rel,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    Primary  string `json:"primary,omitempty"`
}

type Money struct {
    Amount       float64 `json:"amount,omitempty"`
    CurrencyCode string  `json:"currencyCode,omitempty"`
}

type Name struct {
    GivenName      GivenName      `json:"gd$givenName,omitempty"`
    AdditionalName AdditionalName `json:"gd$additionalName,omitempty"`
    FamilyName     FamilyName     `json:"gd$familyName,omitempty"`
    NamePrefix     AtomText       `json:"gd$namePrefix,omitempty"`
    NameSuffix     AtomText       `json:"gd$nameSuffix,omitempty"`
    FullName       AtomText       `json:"gd$fullName,omitempty"`
}

type Organization struct {
    Label             string            `json:"label,omitempty"`
    OrgDepartment     OrgDepartment     `json:"gd$orgDepartment,omitempty"`
    OrgJobDescription OrgJobDescription `json:"gd$orgJobDescription,omitempty"`
    OrgName           OrgName           `json:"gd$orgName,omitempty"`
    OrgSymbol         OrgSymbol         `json:"gd$orgSymbol,omitempty"`
    OrgTitle          OrgTitle          `json:"gd$orgTitle,omitempty"`
    Primary           string            `json:"primary,omitempty"`
    Rel               string            `json:"rel,omitempty"`
    Where             Where             `json:"gd$where,omitempty"`
}

type OrgDepartment AtomText
type OrgJobDescription AtomText
type OrgName AtomText
type OrgSymbol AtomText
type OrgTitle AtomText

type OriginalEvent struct {
    Id   string `json:"id,omitempty"`
    Href string `json:"href,omitempty"`
    When When   `json:"gd$when,omitempty"`
}

type PhoneNumber struct {
    Value   string `json:"$t,omitempty"`
    Label   string `json:"label,omitempty"`
    Rel     string `json:"rel,omitempty"`
    Uri     string `json:"uri,omitempty"`
    Primary string `json:"primary,omitempty"`
}

type PostalAddress struct {
    Value   string `json:"$t,omitempty"`
    Label   string `json:"label,omitempty"`
    Rel     string `json:"rel,omitempty"`
    Primary string `json:"primary,omitempty"`
}

type Rating struct {
    Average   float64 `json:"average,omitempty"`
    Max       int     `json:"max,omitempty"`
    Min       int     `json:"min,omitempty"`
    NumRaters int     `json:"numRaters,omitempty"`
    Rel       string  `json:"rel,omitempty"`
    Value     int     `json:"value,omitempty"`
}

type Recurrence AtomText

type RecurrenceException struct {
    Specialized   string        `json:"specialized,omitempty"`
    EntryLink     EntryLink     `json:"gd$entryLink,omitempty"`
    OriginalEvent OriginalEvent `json:"gd$originalEvent,omitempty"`
}

type Reminder struct {
    AbsoluteTime string `json:"absoluteTime,omitempty"`
    Method       string `json:"method,omitempty"`
    Days         string `json:"days,omitempty"`
    Hours        string `json:"hours,omitempty"`
    Minutes      string `json:"minutes,omitempty"`
}

type ResourceId AtomText

type StructuredPostalAddress struct {
    Rel              string           `json:"rel,omitempty"`
    MailClass        string           `json:"mailClass,omitempty"`
    Usage            string           `json:"usage,omitempty"`
    Label            string           `json:"label,omitempty"`
    Primary          string           `json:"primary,omitempty"`
    Agent            Agent            `json:"gd$agent,omitempty"`
    HouseName        HouseName        `json:"gd$housename,omitempty"`
    Street           Street           `json:"gd$street,omitempty"`
    POBox            POBox            `json:"gd$pobox,omitempty"`
    Neighborhood     Neighborhood     `json:"gd$neighborhood,omitempty"`
    City             City             `json:"gd$city,omitempty"`
    Subregion        Subregion        `json:"gd$subregion,omitempty"`
    Region           Region           `json:"gd$region,omitempty"`
    Postcode         Postcode         `json:"gd$postcode,omitempty"`
    Country          Country          `json:"gd$country,omitempty"`
    FormattedAddress FormattedAddress `json:"gd$formattedAddress,omitempty"`
}

type When struct {
    EndTime     string      `json:"endTime,omitempty"`
    StartTime   string      `json:"startTime,omitempty"`
    ValueString string      `json:"valueString,omitempty"`
    Reminders   []*Reminder `json:"gd$reminder,omitempty"`
}

type Where struct {
    Label       string    `json:"label,omitempty"`
    Rel         string    `json:"rel,omitempty"`
    ValueString string    `json:"valueString,omitempty"`
    EntryLink   EntryLink `json:"gd$entryLink,omitempty"`
}

type Who struct {
    Email          string    `json:"email,omitempty"`
    Rel            string    `json:"rel,omitempty"`
    ValueString    string    `json:"valueString,omitempty"`
    AttendeeStatus AtomText  `json:"gd$attendeeStatus,omitempty"`
    AttendeeType   AtomText  `json:"gd$attendeeType,omitempty"`
    EntryLink      EntryLink `json:"gd$entryLink,omitempty"`
}
type Agent AtomText
type HouseName AtomText
type Street AtomText
type POBox AtomText
type Neighborhood AtomText
type City AtomText
type Subregion AtomText
type Region AtomText
type Postcode AtomText
type FormattedAddress AtomText

type ContactGroup struct {
    Etag               string             `json:"gd$etag,omitempty"`
    Id                 AtomId             `json:"id,omitempty"`
    Xmlns              string             `json:"xmlns,omitempty"`
    XmlnsGcontact      string             `json:"xmlns$gContact,omitempty"`
    XmlnsBatch         string             `json:"xmlns$batch,omitempty"`
    XmlnsGd            string             `json:"xmlns$gd,omitempty"`
    Updated            AtomUpdated        `json:"updated,omitempty"`
    Categories         []AtomCategory     `json:"category,omitempty"`
    Title              AtomText           `json:"title,omitempty"`
    Content            AtomContent        `json:"content,omitempty"`
    Links              []AtomLink         `json:"link,omitempty"`
    Deleted            *DeletedMarker     `json:"gd$deleted,omitempty"`
    ExtendedProperties []ExtendedProperty `json:"gd$extendedProperty,omitempty"`
    SystemGroup        SystemGroup        `json:"gContact$systemGroup,omitempty"`
}

func (p *ContactGroup) GroupId() string {
    arr := strings.Split(p.Id.Value, "/")
    if len(arr) > 0 {
        return arr[len(arr)-1]
    }
    return ""
}
func (p *ContactGroup) SetGroupId(groupId string) {
    p.Id.Value = strings.Join([]string{GOOGLE_GROUPS_API_ENDPOINT, "default/full/", url.QueryEscape(groupId)}, "")
}

type BillingInformation AtomText

type Birthday struct {
    When string `json:"when,omitempty"`
}

type CalendarLink struct {
    Rel     string `json:"rel,omitempty"`
    Label   string `json:"label,omitempty"`
    Primary string `json:"primary,omitempty"`
    Href    string `json:"href,omitempty"`
}

type DirectoryServer AtomText

type Event struct {
    Label string `json:"label,omitempty"`
    Rel   string `json:"rel,omitempty"`
    When  When   `json:"gd$when,omitempty"`
}

type ExternalId struct {
    Label string `json:"label,omitempty"`
    Rel   string `json:"rel,omitempty"`
    Value string `json:"value,omitempty"`
}

type Gender struct {
    Value string `json:"value,omitempty"`
}

type GroupMembershipInfo struct {
    Href    string `json:"href,omitempty"`
    Deleted string `json:"deleted,omitempty"`
}

type Hobby AtomText
type Initials AtomText

type Jot struct {
    Rel string `json:"rel,omitempty"`
}

type Language struct {
    Code  string `json:"code,omitempty"`
    Label string `json:"label,omitempty"`
}

type MaidenName AtomText
type Mileage AtomText
type Nickname AtomText
type Occupation AtomText

type Priority struct {
    Rel string `json:"rel,omitempty"`
}

type Relation struct {
    Value string `json:"$t,omitempty"`
    Label string `json:"label,omitempty"`
    Rel   string `json:"rel,omitempty"`
}

type Sensitivity struct {
    Rel string `json:"rel,omitempty"`
}

type ShortName AtomText
type Subject AtomText

type SystemGroup struct {
    Id string `json:"id,omitempty"`
}

type UserDefinedField struct {
    Key   string `json:"key,omitempty"`
    Value string `json:"value,omitempty"`
}

type Website struct {
    Href    string `json:"href,omitempty"`
    Label   string `json:"label,omitempty"`
    Primary string `json:"primary,omitempty"`
    Rel     string `json:"rel,omitempty"`
}

type GroupsFeedResponse struct {
    Version  string      `json:"version,omitempty"`
    Encoding string      `json:"encoding,omitempty"`
    Feed     *GroupsFeed `json:"feed,omitempty"`
}

type GroupsFeed struct {
    Id              AtomId         `json:"id,omitempty"`
    Xmlns           string         `json:"xmlns,omitempty"`
    XmlnsOpenSearch string         `json:"xmlns$openSearch,omitempty"`
    XmlnsGcontact   string         `json:"xmlns$gContact,omitempty"`
    XmlnsBatch      string         `json:"xmlns$batch,omitempty"`
    XmlnsGd         string         `json:"xmlns$gd,omitempty"`
    Updated         AtomUpdated    `json:"updated,omitempty"`
    Category        []AtomCategory `json:"category,omitempty"`
    Title           AtomTitle      `json:"title,omitempty"`
    Links           []AtomLink     `json:"link,omitempty"`
    Author          []AtomAuthor   `json:"author,omitempty"`
    Generator       AtomGenerator  `json:"generator,omitempty"`
    TotalResults    AtomText       `json:"openSearch$totalResults,omitempty"`
    StartIndex      AtomText       `json:"openSearch$startIndex,omitempty"`
    ItemsPerPage    AtomText       `json:"openSearch$itemsPerPage,omitempty"`
    Entries         []ContactGroup `json:"entry,omitempty"`
}

func NewGroupsFeed() *GroupsFeed {
    return &GroupsFeed{
        Xmlns:           XMLNS_ATOM,
        XmlnsOpenSearch: XMLNS_OPENSEARCH,
        XmlnsGcontact:   XMLNS_GCONTACT,
        XmlnsBatch:      XMLNS_GDATA_BATCH,
        XmlnsGd:         XMLNS_GD,
    }
}

type GroupResponse struct {
    Version  string        `json:"version,omitempty"`
    Encoding string        `json:"encoding,omitempty"`
    Entry    *ContactGroup `json:"entry,omitempty"`
}

type ContactEntryInsertRequest ContactEntryResponse
type ContactEntryUpdateRequest ContactEntryResponse
type GroupEntryInsertRequest GroupResponse
type GroupEntryUpdateRequest GroupResponse

type ContactQuery struct {
    Alt               string `json:"alt,omitempty"`
    Q                 string `json:"q,omitempty"`
    MaxResults        int64  `json:"max-results,omitempty"`
    StartIndex        int64  `json:"start-index,omitempty"`
    UpdatedMin        string `json:"updated-min,omitempty"`
    OrderBy           string `json:"orderby,omitempty"`
    ShowDeleted       bool   `json:"showdeleted,omitempty"`
    RequireAllDeleted bool   `json:"requiredalldeleted,omitempty"`
    SortOrder         string `json:"sortorder,omitempty"`
    Group             string `json:"group,omitempty"`
}

type GroupQuery struct {
    Alt               string `json:"alt,omitempty"`
    Q                 string `json:"q,omitempty"`
    MaxResults        int64  `json:"max-results,omitempty"`
    StartIndex        int64  `json:"start-index,omitempty"`
    UpdatedMin        string `json:"updated-min,omitempty"`
    OrderBy           string `json:"orderby,omitempty"`
    ShowDeleted       bool   `json:"showdeleted,omitempty"`
    RequireAllDeleted bool   `json:"requiredalldeleted,omitempty"`
    SortOrder         string `json:"sortorder,omitempty"`
}
