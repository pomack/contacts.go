package linkedin

type Contact struct {
    Associations          string             `json:"associations,omitempty"`
    Certifications        CertificationList  `json:"certifications,omitempty"`
    DateOfBirth           Date               `json:"dateOfBirth,omitempty"`
    Educations            EducationList      `json:"educations,omitempty"`
    FirstName             string             `json:"firstName,omitempty"`
    Headline              string             `json:"headline,omitempty"`
    Honors                string             `json:"honors,omitempty"`
    Id                    string             `json:"id,omitempty"`
    ImAccounts            ImAccountList      `json:"imAccounts,omitempty"`
    Industry              string             `json:"industry,omitempty"`
    Interests             string             `json:"interests,omitempty"`
    Languages             LanguageList       `json:"languages,omitempty"`
    LastName              string             `json:"lastName,omitempty"`
    Location              Location           `json:"location,omitempty"`
    MainAddress           string             `json:"mainAddress,omitempty"`
    Urls                  UrlList            `json:"memberUrlResources,omitempty"`
    NumConnections        int                `json:"numConnections,omitempty"`
    NumConnnectionsCapped bool               `json:"numConnectionsCapped,omitempty"`
    PhoneNumbers          PhoneNumberList    `json:"phoneNumbers,omitempty"`
    PictureUrl            string             `json:"pictureUrl,omitempty"`
    Positions             PositionList       `json:"positions,omitempty"`
    PublicProfileUrl      string             `json:"publicProfileUrl,omitempty"`
    Skills                SkillList          `json:"skills,omitempty"`
    Specialties           string             `json:"specialties,omitempty"`
    Summary               string             `json:"summary,omitempty"`
    TwitterAccounts       TwitterAccountList `json:"twitterAccounts,omitempty"`
}

type Date struct {
    Year  int `json:"year,omitempty"`
    Month int `json:"month,omitempty"`
    Day   int `json:"day,omitempty"`
}

type CertificationList struct {
    Total  int             `json:"_total,omitempty"`
    Values []Certification `json:"values,omitempty"`
}

type Certification struct {
    Id        int       `json:"id,omitempty"`
    Name      string    `json:"name,omitempty"`
    Authority Authority `json:"authorityName,omitempty"`
    Number    string    `json:"number,omitempty"`
    StartDate Date      `json:"startDate,omitempty"`
    EndDate   Date      `json:"endDate,omitempty"`
}

type Authority struct {
    Name string `json:"name,omitempty"`
}

type EducationList struct {
    Total  int         `json:"_total,omitempty"`
    Values []Education `json:"values,omitempty"`
}

type Education struct {
    Activities   string `json:"activities,omitempty"`
    Degree       string `json:"degree,omitempty"`
    EndDate      Date   `json:"endDate,omitempty"`
    FieldOfStudy string `json:"fieldOfStudy,omitempty"`
    Id           int    `json:"id,omitempty"`
    Notes        string `json:"notes,omitempty"`
    SchoolName   string `json:"schoolName,omitempty"`
    StartDate    Date   `json:"startDate,omitempty"`
}

type ImAccountList struct {
    Total  int         `json:"_total,omitempty"`
    Values []ImAccount `json:"values,omitempty"`
}

type ImAccount struct {
    Name string `json:"imAccountName,omitempty"`
    Type string `json:"imAccountType,omitempty"`
}

type LanguageList struct {
    Total  int               `json:"_total,omitempty"`
    Values []LanguageWrapper `json:"values,omitempty"`
}

type LanguageWrapper struct {
    Id          int         `json:"id,omitempty"`
    Language    Language    `json:"language,omitempty"`
    Proficiency Proficiency `json:"proficiency,omitempty"`
}

type Language struct {
    Name string `json:"name,omitempty"`
}

type Location struct {
    Country Country `json:"country,omitempty"`
    Name    string  `json:"name,omitempty"`
}

type Country struct {
    Code string `json:"code,omitempty"`
}

type UrlList struct {
    Total  int   `json:"_total,omitempty"`
    Values []Url `json:"values,omitempty"`
}

type Url struct {
    Name string `json:"name,omitempty"`
    Url  string `json:"url,omitempty"`
}

type PhoneNumberList struct {
    Total  int           `json:"_total,omitempty"`
    Values []PhoneNumber `json:"values,omitempty"`
}

type PhoneNumber struct {
    Number string `json:"phoneNumber,omitempty"`
    Type   string `json:"phoneType,omitempty"`
}

type PositionList struct {
    Total  int        `json:"_total,omitempty"`
    Values []Position `json:"values,omitempty"`
}

type Position struct {
    Company   Company `json:"company,omitempty"`
    EndDate   Date    `json:"endDate,omitempty"`
    Id        int64   `json:"id,omitempty"`
    IsCurrent bool    `json:"isCurrent,omitempty"`
    StartDate Date    `json:"startDate,omitempty"`
    Summary   string  `json:"summary,omitempty"`
    Title     string  `json:"title,omitempty"`
}

type Company struct {
    Id       int64  `json:"id,omitempty"`
    Industry string `json:"industry,omitempty"`
    Name     string `json:"name,omitempty"`
    Size     string `json:"size,omitempty"`
    Ticker   string `json:"ticker,omitempty"`
    Type     string `json:"type,omitempty"`
}

type SkillList struct {
    Total  int            `json:"_total,omitempty"`
    Values []SkillWrapper `json:"values,omitempty"`
}

type SkillWrapper struct {
    Id          int             `json:"id,omitempty"`
    Skill       Skill           `json:"skill,omitempty"`
    Proficiency Proficiency     `json:"proficiency,omitempty"`
    Years       ExperienceYears `json:"years,omitempty"`
}

type Skill struct {
    Name string `json:"name,omitempty"`
}

type Proficiency struct {
    Level string `json:"level,omitempty"`
    Name  string `json:"name,omitempty"`
}

type ExperienceYears struct {
    Id   int    `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

type TwitterAccountList struct {
    Total  int              `json:"_total,omitempty"`
    Values []TwitterAccount `json:"values,omitempty"`
}

type TwitterAccount struct {
    ProviderAccountId   string `json:"providerAccountId,omitempty"`
    ProviderAccountName string `json:"providerAccountName,omitempty"`
}

type ContactList struct {
    Total  int       `json:"_total,omitempty"`
    Values []Contact `json:"contact,omitempty"`
}

type ErrorResponse struct {
    ErrorCode int    `json:"errorCode,omitempty"`
    Message   string `json:"message,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    Status    int    `json:"status,omitempty"`
    Timestamp int64  `json:"timestamp,omitempty"`
}

func (p *ErrorResponse) String() string {
    return p.Message
}
