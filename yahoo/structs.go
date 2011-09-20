package yahoo

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
    IsConnection    bool                `json:"isConnection"`
    Id              int                 `json:"id"`
    Fields          []ContactField      `json:"fields"`
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

type ContactField struct {
    Uri             string              `json:"uri"`
    Created         string              `json:"created"`
    Updated         string              `json:"updated"`
    Id              int                 `json:"id"`
    Type            string              `json:"type"`
    EditedBy        string              `json:"editedBy"`
    Flags           []string            `json:"flags"`
    Categories      []string            `json:"categories"`
    Value           interface{}         `json:"value"`
}
