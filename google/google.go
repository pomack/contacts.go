package google

import (
    //"github.com/pomack/jsonhelper"
    //"github.com/pomack/oauth2_client"
)

type AtomEntry interface{}

type AtomText struct {
    Value           string  `json:"$t"`
}

type AtomId AtomText
type AtomUpdated AtomText

type AtomCategory struct {
    Scheme          string  `json:"scheme"`
    Term            string  `json:"term"`
}

type AtomLink struct {
    Rel             string  `json:"rel"`
    Type            string  `json:"type"`
    Href            string  `json:"href"`
    Etag            string  `json:"gd$etag"`
}

type AtomTitle struct {
    Value           string  `json:"$t"`
    Type            string  `json:"type"`
}

type AtomContent struct {
    Value           string  `json:"$t"`
    Type            string  `json:"type"`
}

type AtomAuthor struct {
    Name            AtomText    `json:"name"`
    Email           AtomText    `json:""`
}

type AtomGenerator struct {
    Value           string              `json:"$t"`
    Version         string              `json:"version"`
    Uri             string              `json:"uri"`
}

type ContactFeedResponse struct {
    Version         string              `json:"version"`
    Encoding        string              `json:"encoding"`
    Feed            *ContactFeed        `json:"feed"`
}

type ContactEntryResponse struct {
    Version         string              `json:"version"`
    Encoding        string              `json:"encoding"`
    Entry           *Contact            `json:"entry"`
}

type ContactFeed struct {
    Id              AtomId             `json:"id"`
    Updated         AtomUpdated        `json:"updated"`
    Categories      []AtomCategory     `json:"category"`
    Title           AtomTitle          `json:"title"`
    Links           []AtomLink         `json:"link"`
    Generator       AtomGenerator      `json:"generator"`
    TotalResults    AtomText           `json:"openSearch$totalResults"`
    StartIndex      AtomText           `json:"openSearch$startIndex"`
    ItemsPerPage    AtomText           `json:"openSearch$itemsPerPage"`
    Entries         []Contact          `json:"entry"`
}

type Contact struct {
    Id              AtomId             `json:"id"`
    Updated         AtomUpdated        `json:"updated"`
    Categories      []AtomCategory     `json:"category"`
    Title           AtomTitle          `json:"title"`
    Content         AtomContent        `json:"content"`
    Links           []AtomLink         `json:"link"`
    Deleted         *DeletedMarker     `json:"gd$deleted"`
    Name            Name               `json:"gd$name"`
    Organizations   []Organization     `json:"gd$organization"`
    Emails          []Email            `json:"gd$email"`
    Ims             []Im               `json:"gd$im"`
    PhoneNumbers    []PhoneNumber      `json:"gd$phoneNumber"`
    PostalAddresses []PostalAddress    `json:"gd$postalAddress"`
    StructuredPostalAddresses []StructuredPostalAddress    `json:"gd$structuredPostalAddress"`
    Where           Where              `json:"gd$where"`
    BillingInformation  BillingInformation `json:"gContact$billingInformation"`
    Birthday        Birthday           `json:"gContact$birthday"`
    CalendarLinks   []CalendarLink     `json:"gContact$calendarLink"`
    DirectoryServer DirectoryServer    `json:"gContact$directoryServer"`
    Events          []Event            `json:"gContact$event"`
    ExternalId      []ExternalId       `json:"gContact$externalId"`
    Gender          Gender             `json:"gContact$gender"`
    Groups          []GroupMembershipInfo  `json:"gContact$groupMembershipInfo"`
    Hobbies         []Hobby            `json:"gContact$hobby"`
    Initials        Initials           `json:"gContact$initials"`
    Jots            []Jot              `json:"gContact$jot"`
    Languages       []Language         `json:"gContact$language"`
    MaidenName      MaidenName         `json:"gContact$maidenName"`
    Mileage         Mileage            `json:"gContact$mileage"`
    Nicnkame        Nickname           `json:"gContact$nickname"`
    Occupation      Occupation         `json:"gContact$occupation"`
    Priority        Priority           `json:"gContact$priority"`
    Relationships   []Relation         `json:"gContact$relation"`
    Sensitivity     Sensitivity        `json:"gContact$sensitivity"`
    ShortName       ShortName          `json:"gContact$shortName"`
    Subject         Subject            `json:"gContact$subject"`
    UserDefinedFields   []UserDefinedField `json:"gContact:$userDefinedField"`
    Websites        []Website          `json:"gContact$website"`
}

type DeletedMarker struct {}

type Name struct {
    GivenName       AtomText            `json:"gd$givenName"`
    AdditionalName  AtomText            `json:"gd$additionalName"`
    FamilyName      AtomText            `json:"gd$familyName"`
    NamePrefix      AtomText            `json:"gd$namePrefix"`
    NameSuffix      AtomText            `json:"gd$nameSuffix"`
    FullName        AtomText            `json:"gd$fullName"`
}

type Organization struct {
    Label               string              `json:"label"`
    OrgDepartment       OrgDepartment       `json:"gd$orgDepartment"`
    OrgJobDescription   OrgJobDescription   `json:"gd$orgJobDescription"`
    OrgName             OrgName             `json:"gd$orgName"`
    OrgSymbol           OrgSymbol           `json:"gd$orgSymbol"`
    OrgTitle            OrgTitle            `json:"gd$orgTitle"`
    Primary             string              `json:"primary"`
    Rel                 string              `json:"rel"`
    Where               Where               `json:"gd$where"`
}

type OrgDepartment AtomText
type OrgJobDescription AtomText
type OrgName AtomText
type OrgSymbol AtomText
type OrgTitle AtomText

type Email struct {
    Address         string              `json:"address"`
    DisplayName     string              `json:"displayName"`
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    Primary         string              `json:"primary"`
}

type Im struct {
    Address         string              `json:"address"`
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    Protocol        string              `json:"protocol"`
    Primary         string              `json:"primary"`
}

type PhoneNumber struct {
    Value           string              `json:"$t"`
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    Uri             string              `json:"uri"`
    Primary         string              `json:"primary"`
}

type PostalAddress struct {
    Value           string              `json:"$t"`
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    Primary         string              `json:"primary"`
}

type StructuredPostalAddress struct {
    Rel             string              `json:"rel"`
    MailClass       string              `json:"mailClass"`
    Usage           string              `json:"usage"`
    Label           string              `json:"label"`
    Primary         string              `json:"primary"`
    Agent           Agent               `json:"gd$agent"`
    HouseName       HouseName           `json:"gd$housename"`
    Street          Street              `json:"gd$street"`
    POBox           POBox               `json:"gd$pobox"`
    Neighborhoood   Neighborhood        `json:"gd$neighborhood"`
    City            City                `json:"gd$city"`
    Subregion       Subregion           `json:"gd$subregion"`
    Region          Region              `json:"gd$region"`
    Postcode        Postcode            `json:"gd$postcode"`
    Country         Country             `json:"gd$country"`
    FormattedAddress    FormattedAddress    `json:"gd$formattedAddress"`
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

type Country struct {
    Value           string              `json:"$t"`
    Code            string              `json:"code"`
}

type Where struct {
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    ValueString     string              `json:"valueString"`
    EntryLink       EntryLink           `json:"gd$entryLink"`
}

type When struct {
    EndTime         string              `json:"endTime"`
    StartTime       string              `json:"startTime"`
    ValueString     string              `json:"valueString"`
}

type GroupMembershipInfo struct {
    Href            string              `json:"href"`
    Deleted         string              `json:"deleted"`
}

type Money struct {
    Amount          float64             `json:"amount"`
    CurrencyCode    string              `json:"currencyCode"`
}

type EntryLink struct {
    Href            string              `json:"href"`
    ReadOnly        string              `json:"readOnly"`
    Rel             string              `json:"rel"`
    Entry           AtomEntry           `json:"entry"`
}

type Rating struct {
    Average         float64             `json:"average"`
    Max             int                 `json:"max"`
    Min             int                 `json:"min"`
    NumRaters       int                 `json:"numRaters"`
    Rel             string              `json:"rel"`
    Value           int                 `json:"value"`
}

type ResourceId AtomText

type ExtendedProperty struct {
    Name            string              `json:"name"`
    Description     AtomText            `json:"description"`
}

type ContactGroup struct {
    Category        AtomCategory        `json:"category"`
    Updated         AtomText            `json:"updated"`
    Title           AtomText            `json:"title"`
    Content         AtomContent         `json:"content"`
    Deleted         *DeletedMarker      `json:"gd$deleted"`
    ExtendedProperties  []ExtendedProperty `json:"gd$extendedProperty"`
    SystemGroup     SystemGroup         `json:"gContact$systemGroup"`
}

type BillingInformation AtomText

type Birthday struct {
    When            string              `json:"when"`
}

type CalendarLink struct {
    Rel             string              `json:"rel"`
    Label           string              `json:"label"`
    Primary         string              `json:"primary"`
    Href            string              `json:"href"`
}

type DirectoryServer AtomText

type Event struct {
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    When            When                `json:"gd$when"`
}

type ExternalId struct {
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
    Value           string              `json:"value"`
}

type Gender struct {
    Value           string              `json:"value"`
}

type Hobby AtomText
type Initials AtomText

type Jot struct {
    Rel             string              `json:"rel"`
}

type Language struct {
    Code            string              `json:"code"`
    Label           string              `json:"label"`
}

type MaidenName AtomText
type Mileage AtomText
type Nickname AtomText
type Occupation AtomText

type Priority struct {
    Rel             string              `json:"rel"`
}

type Relation struct {
    Value           string              `json:"$t"`
    Label           string              `json:"label"`
    Rel             string              `json:"rel"`
}

type Sensitivity struct {
    Rel             string              `json:"rel"`
}

type ShortName AtomText
type Subject AtomText

type SystemGroup struct {
    Id              string              `json:"id"`
}

type UserDefinedField struct {
    Key             string              `json:"key"`
    Value           string              `json:"value"`
}

type Website struct {
    Href            string              `json:"href"`
    Label           string              `json:"label"`
    Primary         string              `json:"primary"`
    Rel             string              `json:"rel"`
}
