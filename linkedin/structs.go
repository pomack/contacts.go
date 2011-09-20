package linkedin

type Contact struct {
    Associations    string              `json:"associations"`
    Certifications  CertificationList   `json:"certifications"`
    DateOfBirth     Date                `json:"dateOfBirth"`
    Educations      EducationList       `json:"educations"`
    FirstName       string              `json:"firstName"`
    Headline        string              `json:"headline"`
    Honors          string              `json:"honors"`
    Id              string              `json:"id"`
    ImAccounts      ImAccountList       `json:"imAccounts"`
    Industry        string              `json:"industry"`
    Interests       string              `json:"interests"`
    Languages       LanguageList        `json:"languages"`
    LastName        string              `json:"lastName"`
    Location        Location            `json:"location"`
    MainAddress     string              `json:"mainAddress"`
    Urls            UrlList             `json:"memberUrlResources"`
    NumConnections  int                 `json:"numConnections"`
    NumConnnectionsCapped   bool        `json:"numConnectionsCapped"`
    PhoneNumbers    PhoneNumberList     `json:"phoneNumbers"`
    PictureUrl      string              `json:"pictureUrl"`
    Positions       PositionList        `json:"positions"`
    PublicProfileUrl    string          `json:"publicProfileUrl"`
    Skills          SkillList           `json:"skills"`
    Specialties     string              `json:"specialties"`
    Summary         string              `json:"summary"`
    TwitterAccounts TwitterAccountList  `json:"twitterAccounts"`
}


type Date struct {
    Year            int             `json:"year"`
    Month           int             `json:"month"`
    Day             int             `json:"day"`
}

type CertificationList struct {
    Total           int             `json:"_total"`
    Values          []Certification `json:"values"`
}

type Certification struct {
    Id              int             `json:"id"`
    Name            string          `json:"name"`
}

type EducationList struct {
    Total           int             `json:"_total"`
    Values          []Education     `json:"values"`
}

type Education struct {
    Activities      string          `json:"activities"`
    Degree          string          `json:"degree"`
    EndDate         Date            `json:"endDate"`
    FieldOfStudy    string          `json:"fieldOfStudy"`
    Id              int             `json:"id"`
    Notes           string          `json:"notes"`
    SchoolName      string          `json:"schoolName"`
    StartDate       Date            `json:"startDate"`
}

type ImAccountList struct {
    Total           int             `json:"_total"`
    Values          []ImAccount     `json:"values"`
}

type ImAccount struct {
    Name            string          `json:"imAccountName"`
    Type            string          `json:"imAccountType"`
}

type LanguageList struct {
    Total           int                 `json:"_total"`
    Values          []LanguageWrapper   `json:"values"`
}

type LanguageWrapper struct {
    Id              int                 `json:"id"`
    Language        Language            `json:"language"`
}

type Language struct {
    Name            string              `json:"name"`
}

type Location struct {
    Country         Country         `json:"country"`
    Name            string          `json:"name"`
}

type Country struct {
    Code            string          `json:"code"`
}

type UrlList struct {
    Total           int             `json:"_total"`
    Values          []Url           `json:"values"`
}

type Url struct {
    Name            string          `json:"name"`
    Url             string          `json:"url"`
}

type PhoneNumberList struct {
    Total           int             `json:"_total"`
    Values          []PhoneNumber   `json:"values"`
}

type PhoneNumber struct {
    Number          string          `json:"phoneNumber"`
    Type            string          `json:"phoneType"`
}

type PositionList struct {
    Total           int             `json:"_total"`
    Values          []Position      `json:"values"`
}

type Position struct {
    Company         Company         `json:"company"`
    EndDate         Date            `json:"endDate"`
    Id              int64           `json:"id"`
    IsCurrent       bool            `json:"isCurrent"`
    StartDate       Date            `json:"startDate"`
    Summary         string          `json:"summary"`
    Title           string          `json:"title"`
}

type Company struct {
    Id              int64           `json:"id"`
    Industry        string          `json:"industry"`
    Name            string          `json:"name"`
    Size            string          `json:"size"`
    Ticker          string          `json:"ticker"`
    Type            string          `json:"type"`
}

type SkillList struct {
    Total           int             `json:"_total"`
    Values          []SkillWrapper  `json:"values"`
}

type SkillWrapper struct {
    Id              int             `json:"id"`
    Skill           Skill           `json:"skill"`
}

type Skill struct {
    Name            string          `json:"name"`
}

type TwitterAccountList struct {
    Total           int                 `json:"_total"`
    Values          []TwitterAccount    `json:"values"`
}

type TwitterAccount struct {
    ProviderAccountId   string          `json:"providerAccountId"`
    ProviderAccountName string          `json:"providerAccountName"`
}


