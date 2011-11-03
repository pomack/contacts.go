package googleplus

type Person struct {
    AboutMe            string         `json:"aboutMe,omitempty"`
    Birthday           string         `json:"birthday,omitempty"`
    CurrentLocation    string         `json:"currentLocation,omitempty"`
    DisplayName        string         `json:"displayName,omitempty"`
    Emails             []Email        `json:"emails,omitempty"`
    Gender             string         `json:"gender,omitempty"`
    HasApp             bool           `json:"hasApp,omitempty"`
    Id                 string         `json:"id,omitempty"`
    Image              Image          `json:"image,omitempty"`
    Kind               string         `json:"kind,omitempty"`
    LanguagesSpoken    []string       `json:"languagesSpoken,omitempty"`
    Name               Name           `json:"name,omitempty"`
    Nickname           string         `json:"nickname,omitempty"`
    Organizations      []Organization `json:"organizations,omitempty"`
    PlacesLived        []PlaceLived   `json:"placesLived,omitempty"`
    RelationshipStatus string         `json:"relationshipStatus,omitempty"`
    Tagline            string         `json:"tagline,omitempty"`
    Url                string         `json:"url,omitempty"`
    Urls               []Url          `json:"urls,omitempty"`
}

func NewPerson() *Person {
    return &Person{
        Kind: PERSON_KIND,
    }
}

type Email struct {
    Type    string `json:"type,omitempty"`
    Value   string `json:"value,omitempty"`
    Primary bool   `json:"primary,omitempty"`
}

type Image struct {
    Url string `json:"url,omitempty"`
}

type Name struct {
    FamilyName      string `json:"familyName,omitempty"`
    Formatted       string `json:"formatted,omitempty"`
    GivenName       string `json:"givenName,omitempty"`
    HonorificPrefix string `json:"honorificPrefix,omitempty"`
    HonorificSuffix string `json:"honorificSuffix,omitempty"`
    MiddleName      string `json:"middleName,omitempty"`
}

type Organization struct {
    Department  string `json:"department,omitempty"`
    Description string `json:"description,omitempty"`
    EndDate     string `json:"endDate,omitempty"`
    Location    string `json:"location,omitempty"`
    Name        string `json:"name,omitempty"`
    Primary     bool   `json:"primary,omitempty"`
    StartDate   string `json:"startDate,omitempty"`
    Title       string `json:"title,omitempty"`
    Type        string `json:"type,omitempty"`
}

type PlaceLived struct {
    Value   string `json:"value,omitempty"`
    Primary bool   `json:"primary,omitempty"`
}

type Url struct {
    Value   string `json:"value,omitempty"`
    Type    string `json:"type,omitempty"`
    Primary bool   `json:"primary,omitempty"`
}

type Activity struct {
    Access          Access `json:"access,omitempty"`
    Actor           Actor  `json:"actor,omitempty"`
    Address         string `json:"address,omitempty"`
    Annotation      string `json:"annotation,omitempty"`
    CrosspostSource string `json:"crosspostSource,omitempty"`
    Geocode         string `json:"geocode,omitempty"`
    Id              string `json:"id,omitempty"`
    Kind            string `json:"kind,omitempty"`
    Object          Object `json:"object,omitempty"`
    PlaceId         string `json:"placeId,omitempty"`
    PlaceName       string `json:"placeName,omitempty"`
    Placeholder     bool   `json:"placeholder,omitempty"`
    Provider        Titler `json:"provider,omitempty"`
    Published       string `json:"published,omitempty"`
    Radius          string `json:"radius,omitempty"`
    Title           string `json:"title,omitempty"`
    Updated         string `json:"updated,omitempty"`
    Url             string `json:"url,omitempty"`
    Verb            string `json:"verb,omitempty"`
}

type Access struct {
    Description string `json:"description,omitempty"`
    Items       []Item `json:"items,omitempty"`
    Kind        string `json:"kind,omitempty"`
}

type Item struct {
    Id   string `json:"id,omitempty"`
    Type string `json:"type,omitempty"`
}

type Actor struct {
    DisplayName string `json:"displayName,omitempty"`
    Id          string `json:"id,omitempty"`
    Image       Image  `json:"image,omitempty"`
    Url         string `json:"url,omitempty"`
}

type Object struct {
    Actor           Actor        `json:"actor,omitempty"`
    Attachments     []Attachment `json:"attachments,omitempty"`
    Content         string       `json:"content,omitempty"`
    Id              string       `json:"id,omitempty"`
    ObjectType      string       `json:"objectType,omitempty"`
    OriginalContent string       `json:"originalContent,omitempty"`
    Plusoners       ItemCount    `json:"plusoners,omitempty"`
    Replies         ItemCount    `json:"replies,omitempty"`
    Resharers       ItemCount    `json:"resharers,omitempty"`
    Url             string       `json:"url,omitempty"`
}

type Attachment struct {
    Content     string    `json:"content,omitempty"`
    DisplayName string    `json:"displayName,omitempty"`
    Embed       Embed     `json:"embed,omitempty"`
    FullImage   ImageInfo `json:"fullImage,omitempty"`
    Id          string    `json:"id,omitempty"`
    Image       ImageInfo `json:"image,omitempty"`
    ObjectType  string    `json:"objectType,omitempty"`
    Url         string    `json:"url,omitempty"`
}

type Embed struct {
    Type string `json:"type,omitempty"`
    Url  string `json:"url,omitempty"`
}

type ImageInfo struct {
    Height int    `json:"height,omitempty"`
    Type   string `json:"type,omitempty"`
    Url    string `json:"url,omitempty"`
    Width  int    `json:"width,omitempty"`
}

type ItemCount struct {
    TotalItems int `json:"totalItems,omitempty"`
}

type Titler struct {
    Title string `json:"title,omitempty"`
}

type ActivityListResponse struct {
    Id            string     `json:"id,omitempty"`
    Items         []Activity `json:"items,omitempty"`
    Kind          string     `json:"kind,omitempty"`
    NextLink      string     `json:"nextLink,omitempty"`
    NextPageToken string     `json:"nextPageToken,omitempty"`
    SelfLink      string     `json:"selfLink,omitempty"`
    Title         string     `json:"title,omitempty"`
    Updated       string     `json:"updated,omitempty"`
}
