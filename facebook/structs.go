package facebook

type Contact struct {
    Id                  string             `json:"id,omitempty"`
    Name                string             `json:"name,omitempty"`
    FirstName           string             `json:"first_name,omitempty"`
    MiddleName          string             `json:"middle_name,omitempty"`
    LastName            string             `json:"last_name,omitempty"`
    Link                string             `json:"link,omitempty"`
    Username            string             `json:"username,omitempty"`
    Hometown            Location           `json:"hometown,omitempty"`
    Location            Location           `json:"location,omitempty"`
    Bio                 string             `json:"bio,omitempty"`
    Quotes              string             `json:"quotes,omitempty"`
    WorkPlaces          []WorkPlace        `json:"work,omitempty"`
    InspirationalPeople []ContactReference `json:"inspirational_people,omitempty"`
    Educations          []Education        `json:"education,omitempty"`
    Gender              string             `json:"gender,omitempty"`
    Website             string             `json:"website,omitempty"`
    InterestedIn        []string           `json:"interested_in,omitempty"`
    RelationshipStatus  string             `json:"relationship_status,omitempty"`
    Political           string             `json:"political,omitempty"`
    Timezone            int                `json:"timezone,omitempty"`
    Locale              string             `json:"locale,omitempty"`
    Languages           []Language         `json:"languages,omitempty"`
    Verified            bool               `json:"verified,omitempty"`
    UpdatedTime         string             `json:"updated_time,omitempty"`
    Type                string             `json:"type,omitempty"`
    ThirdPartyId        string             `json:"third_party_id,omitempty"`
    Birthday            string             `json:"birthday,omitempty"`
    Email               string             `json:"email,omitempty"`
    FavoriteAthletes    []ContactReference `json:"favorite_athletes,omitempty"`
    FavoriteTeams       []ContactReference `json:"favorite_teams,omitempty"`
    Religion            string             `json:"religion,omitempty"`
    SignificantOther    ContactReference   `json:"significant_other,omitempty"`
    VideoUploadLimits   SizeLimit          `json:"video_upload_limits,omitempty"`
}

type IdAndName struct {
    Id   string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

type StandardMedia struct {
    Name        string `json:"name,omitempty"`
    Category    string `json:"category,omitempty"`
    Id          string `json:"id,omitempty"`
    CreatedTime string `json:"created_time,omitempty"`
}

type Location IdAndName

type WorkPlace struct {
    Employer    Company            `json:"employer,omitempty"`
    Location    Location           `json:"location,omitempty"`
    Position    Position           `json:"position,omitempty"`
    With        []ContactReference `json:"with,omitempty"`
    Description string             `json:"description,omitempty"`
    StartDate   string             `json:"start_date,omitempty"`
    EndDate     string             `json:"end_date,omitempty"`
}

type Company IdAndName
type Position IdAndName
type ContactReference IdAndName

type Education struct {
    School         School             `json:"school,omitempty"`
    Year           GraduationClass    `json:"year,omitempty"`
    Type           string             `json:"type,omitempty"`
    Degree         Degree             `json:"degree,omitempty"`
    Concentrations []Concentration    `json:"concentration,omitempty"`
    With           []ContactReference `json:"with,omitempty"`
}

type School IdAndName
type GraduationClass IdAndName
type Degree IdAndName
type Concentration IdAndName

type Language IdAndName

type SizeLimit struct {
    Length string `json:"length,omitempty"`
    Size   string `json:"size,omitempty"`
}

type ActivitiesResponse struct {
    Data []Activity `json:"data,omitempty"`
}

type Activity StandardMedia

type AlbumsResponse struct {
    Data   []Album `json:"data,omitempty"`
    Paging Paging  `json:"paging,omitempty"`
}

type Album struct {
    Id          string           `json:"id,omitempty"`
    From        ContactReference `json:"from,omitempty"`
    Name        string           `json:"name,omitempty"`
    Link        string           `json:"link,omitempty"`
    CoverPhoto  string           `json:"cover_photo,omitempty"`
    Privacy     string           `json:"privacy,omitempty"`
    Count       int              `json:"count,omitempty"`
    Type        string           `json:"type,omitempty"`
    CreatedTime string           `json:"created_time,omitempty"`
    UpdatedTime string           `json:"updated_time,omitempty"`
}

type Paging struct {
    Previous string `json:"previous,omitempty"`
    Next     string `json:"next,omitempty"`
}

type BooksResponse struct {
    Data []Book `json:"data,omitempty"`
}

type Book StandardMedia

type CheckinsResponse struct {
    Data   []Checkin `json:"data,omitempty"`
    Paging Paging    `json:"paging,omitempty"`
}

type Checkin struct {
    Id          string                `json:"id,omitempty"`
    From        ContactReference      `json:"from,omitempty"`
    Tags        ContactReferenceArray `json:"tags,omitempty"`
    Message     string                `json:"message,omitempty"`
    Place       CheckinLocation       `json:"place,omitempty"`
    Application Application           `json:"application,omitempty"`
    CreatedTime string                `json:"created_time,omitempty"`
    Likes       LikeArray             `json:"likes,omitempty"`
    Comments    CommentArray          `json:"comments,omitempty"`
}

type ContactReferenceArray struct {
    Data []ContactReference `json:"data,omitempty"`
}

type CheckinLocation struct {
    Id       string                  `json:"id,omitempty"`
    Name     string                  `json:"name,omitempty"`
    Location CheckinPhysicalLocation `json:"location,omitempty"`
}

type CheckinPhysicalLocation struct {
    Street    string  `json:"street,omitempty"`
    City      string  `json:"city,omitempty"`
    State     string  `json:"state,omitempty"`
    Country   string  `json:"country,omitempty"`
    Zip       string  `json:"zip,omitempty"`
    Latitude  float64 `json:"latitude,omitempty"`
    Longitude float64 `json:"longitude,omitempty"`
}

type Application IdAndName

type LikeArray struct {
    Data []Like `json:"data,omitempty"`
}

type Like IdAndName

type CommentArray struct {
    Data []Comment `json:"data,omitempty"`
}

type Comment struct {
    Id          string           `json:"id,omitempty"`
    From        ContactReference `json:"from,omitempty"`
    Message     string           `json:"message,omitempty"`
    CreatedTime string           `json:"created_time,omitempty"`
    Likes       int              `json:"likes,omitempty"`
}

type EventsResponse struct {
    Data   []Event `json:"data,omitempty"`
    Paging Paging  `json:"paging,omitempty"`
}

type Event struct {
    Name       string `json:"name,omitempty"`
    StartTime  string `json:"start_time,omitempty"`
    EndTime    string `json:"end_time,omitempty"`
    Location   string `json:"location,omitempty"`
    Id         string `json:"id,omitempty"`
    RsvpStatus string `json:"rsvp_status,omitempty"`
}

type FamilyResponse struct {
    Data []Relationship `json:"data,omitempty"`
}

type Relationship struct {
    Name         string `json:"name,omitempty"`
    Id           string `json:"id,omitempty"`
    Relationship string `json:"relationship,omitempty"`
}

type FriendsResponse ContactReferenceArray

type GamesResponse struct {
    Data []Game `json:"data,omitempty"`
}

type Game StandardMedia

type GroupsResponse struct {
    Data []Group `json:"data,omitempty"`
}

type Group struct {
    Version       int    `json:"version,omitempty"`
    Name          string `json:"name,omitempty"`
    Id            string `json:"id,omitempty"`
    BookmarkOrder int64  `json:"bookmark_order,omitempty"`
}

type InterestsResponse struct {
    Data []Interest `json:"data,omitempty"`
}

type Interest StandardMedia

type LikesResponse struct {
    Data []LikeInfo `json:"data,omitempty"`
}

type LikeInfo StandardMedia

type MoviesResponse struct {
    Data []Movie `json:"data,omitempty"`
}

type Movie StandardMedia

type MusicResponse struct {
    Data []Music `json:"data,omitempty"`
}

type Music StandardMedia

type NotesResponse struct {
    Data   []Note `json:"data,omitempty"`
    Paging Paging `json:"paging,omitempty"`
}

type Note struct {
    Id          string           `json:"id,omitempty"`
    From        ContactReference `json:"from,omitempty"`
    Subject     string           `json:"subject,omitempty"`
    Message     string           `json:"message,omitempty"`
    Icon        string           `json:"icon,omitempty"`
    CreatedTime string           `json:"created_time,omitempty"`
    UpdatedTime string           `json:"updated_time,omitempty"`
    Comments    CommentArray     `json:"comments,omitempty"`
}

type PhotosResponse struct {
    Data   []Photo `json:"data,omitempty"`
    Paging Paging  `json:"paging,omitempty"`
}

type Photo struct {
    Id          string            `json:"id,omitempty"`
    From        ContactReference  `json:"from,omitempty"`
    Tags        UserPhotoTagArray `json:"tags,omitempty"`
    Picture     string            `json:"picture,omitempty"`
    Source      string            `json:"source,omitempty"`
    Height      int               `json:"height,omitempty"`
    Width       int               `json:"width,omitempty"`
    Images      []Image           `json:"images,omitempty"`
    Link        string            `json:"link,omitempty"`
    Icon        string            `json:"icon,omitempty"`
    CreatedTime string            `json:"created_time,omitempty"`
    Position    int               `json:"position,omitempty"`
    UpdatedTime string            `json:"updated_time,omitempty"`
    Comments    CommentArray      `json:"comments,omitempty"`
}

type UserPhotoTagArray struct {
    Data []UserPhotoTag `json:"data,omitempty"`
}

type UserPhotoTag struct {
    Id          string  `json:"id,omitempty"`
    Name        string  `json:"name,omitempty"`
    X           float64 `json:"x,omitempty"`
    Y           float64 `json:"y,omitempty"`
    CreatedTime string  `json:"created_time,omitempty"`
}

type Image struct {
    Height int    `json:"height,omitempty"`
    Width  int    `json:"width,omitempty"`
    Source string `json:"source,omitempty"`
}

type StatusesResponse struct {
    Data []Status `json:"data,omitempty"`
}

type Status struct {
    Id          string           `json:"id,omitempty"`
    From        ContactReference `json:"from,omitempty"`
    Message     string           `json:"message,omitempty"`
    UpdatedTime string           `json:"updated_time,omitempty"`
    Likes       LikeArray        `json:"likes,omitempty"`
    Comments    CommentArray     `json:"comments,omitempty"`
}

type TelevisionResponse struct {
    Data []TVShow `json:"data,omitempty"`
}

type TVShow StandardMedia

type ErrorResponse struct {
    Error Error `json:"error,omitempty"`
}

func (p *ErrorResponse) String() string { return p.Error.Message }

type Error struct {
    Message string `json:"message,omitempty"`
    Type    string `json:"type,omitempty"`
}

func (p *Error) String() string { return p.Message }
