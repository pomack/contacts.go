package facebook



type Contact struct {
    Id                  string              `json:"id"`
    Name                string              `json:"name"`
    FirstName           string              `json:"first_name"`
    MiddleName          string              `json:"middle_name"`
    LastName            string              `json:"last_name"`
    Link                string              `json:"link"`
    Username            string              `json:"pomack"`
    Hometown            Location            `json:"hometown"`
    Location            Location            `json:"location"`
    Bio                 string              `json:"bio"`
    Quotes              string              `json:"quotes"`
    WorkPlaces          []WorkPlace         `json:"work"`
    InspirationalPeople []ContactReference  `json:"inspirational_people"`
    Educations          []Education         `json:"education"`
    Gender              string              `json:"gender"`
    Website             string              `json:"website"`
    InterestedIn        []string            `json:"interested_in"`
    RelationshipStatus  string              `json:"relationship_status"`
    Political           string              `json:"political"`
    Timezone            int                 `json:"timezone"`
    Locale              string              `json:"locale"`
    Languages           []Language          `json:"languages"`
    Verified            bool                `json:"verified"`
    UpdatedTime         string              `json:"updated_time"`
    Type                string              `json:"type"`
    ThirdPartyId        string              `json:"third_party_id"`
    Birthday            string              `json:"birthday"`
    Email               string              `json:"email"`
    FavoriteAthletes    []ContactReference  `json:"favorite_athletes"`
    FavoriteTeams       []ContactReference  `json:"favorite_teams"`
    Religion            string              `json:"religion"`
    SignificantOther    ContactReference    `json:"significant_other"`
    VideoUploadLimits   SizeLimit           `json:"video_upload_limits"`
}

type IdAndName struct {
    Id                  string              `json:"id"`
    Name                string              `json:"name"`
}

type StandardMedia struct {
    Name                string              `json:"name"`
    Category            string              `json:"category"`
    Id                  string              `json:"id"`
    CreatedTime         string              `json:"created_time"`
}

type Location IdAndName

type WorkPlace struct {
    Employer            Company             `json:"employer"`
    Location            Location            `json:"location"`
    Position            Position            `json:"position"`
    With                []ContactReference  `json:"with"`
    Description         string              `json:"description"`
    StartDate           string              `json:"start_date"`
    EndDate             string              `json:"end_date"`
}

type Company  IdAndName
type Position IdAndName
type ContactReference IdAndName

type Education struct {
    School              School              `json:"school"`   
    Year                GraduationClass     `json:"year"`
    Type                string              `json:"type"`
    Degree              Degree              `json:"degree"`
    Concentrations      []Concentration     `json:"concentration"`
    With                []ContactReference  `json:"with"`
}

type School IdAndName
type GraduationClass IdAndName
type Degree IdAndName
type Concentration IdAndName

type Language IdAndName

type SizeLimit struct {
    Length              string                  `json:"length"`
    Size                string                  `json:"size"`
}

type ActivitiesResponse struct {
    Data                []Activity              `json:"data"`
}

type Activity StandardMedia

type AlbumsResponse struct {
    Data                []Album                 `json:"data"`
    Paging              Paging                  `json:"paging"`
}

type Album struct {
    Id                  string              `json:"id"`
    From                ContactReference    `json:"from"`
    Name                string              `json:"name"`
    Link                string              `json:"link"`
    CoverPhoto          string              `json:"cover_photo"`
    Privacy             string              `json:"privacy"`
    Count               int                 `json:"count"`
    Type                string              `json:"type"`
    CreatedTime         string              `json:"created_time"`
    UpdatedTime         string              `json:"updated_time"`
}

type Paging struct {
    Previous            string              `json:"previous"`
    Next                string              `json:"next"`
}

type BooksResponse struct {
    Data                []Book                  `json:"data"`
}

type Book StandardMedia

type CheckinsResponse struct {
    Data                []Checkin               `json:"data"`
    Paging              Paging                  `json:"paging"`
}

type Checkin struct {
    Id                  string                  `json:"id"`
    From                ContactReference        `json:"from"`
    Tags                ContactReferenceArray   `json:"tags"`
    Message             string                  `json:"message"`
    Place               CheckinLocation         `json:"place"`
    Application         Application             `json:"application"`
    CreatedTime         string                  `json:"created_time"`
    Likes               LikeArray               `json:"likes"`
    Comments            CommentArray            `json:"comments"`
}

type ContactReferenceArray struct {
    Data                []ContactReference      `json:"data"`
}

type CheckinLocation struct {
    Id                  string                  `json:"id"`
    Name                string                  `json:"name"`
    Location            CheckinPhysicalLocation `json:"location"`
}

type CheckinPhysicalLocation struct {
    Street              string                  `json:"street"`
    City                string                  `json:"city"`
    State               string                  `json:"state"`
    Country             string                  `json:"country"`
    Zip                 string                  `json:"zip"`
    Latitude            float64                 `json:"latitude"`
    Longitude           float64                 `json:"longitude"`
}

type Application IdAndName

type LikeArray struct {
    Data                []Like                  `json:"data"`
}

type Like IdAndName

type CommentArray struct {
    Data                []Comment               `json:"data"`
}

type Comment struct {
    Id                  string                  `json:"id"`
    From                ContactReference        `json:"from"`
    Message             string                  `json:"message"`
    CreatedTime         string                  `json:"created_time"`
    Likes               int                     `json:"likes"`
}

type EventsResponse struct {
    Data                []Event                 `json:"data"`
    Paging              Paging                  `json:"paging"`
}

type Event struct {
    Name                string                  `json:"name"`
    StartTime           string                  `json:"start_time"`
    EndTime             string                  `json:"end_time"`
    Location            string                  `json:"location"`
    Id                  string                  `json:"id"`
    RsvpStatus          string                  `json:"rsvp_status"`
}

type FamilyResponse struct {
    Data                []Relationship      `json:"data"`
}

type Relationship struct {
    Name                string              `json:"name"`
    Id                  string              `json:"id"`
    Relationship        string              `json:"relationship"`
}

type FriendsResponse ContactReferenceArray

type GamesResponse struct {
    Data                []Game              `json:"data"`
}

type Game StandardMedia

type GroupsResponse struct {
    Data                []Group                 `json:"data"`
}

type Group struct {
    Version             int                     `json:"version"`
    Name                string                  `json:"name"`
    Id                  string                  `json:"id"`
    BookmarkOrder       int64                   `json:"bookmark_order"`
}

type InterestsResponse struct {
    Data                []Interest              `json:"data"`
}

type Interest StandardMedia

type LikesResponse struct {
    Data                []LikeInfo              `json:"data"`
}

type LikeInfo StandardMedia

type MoviesResponse struct {
    Data                []Movie             `json:"data"`
}

type Movie StandardMedia

type MusicResponse struct {
    Data                []Music             `json:"data"`
}

type Music StandardMedia

type NotesResponse struct {
    Data                []Note              `json:"data"`
    Paging              Paging              `json:"paging"`
}

type Note struct {
    Id                  string              `json:"id"`
    From                ContactReference    `json:"from"`
    Subject             string              `json:"subject"`
    Message             string              `json:"message"`
    Icon                string              `json:"icon"`
    CreatedTime         string              `json:"created_time"`
    UpdatedTime         string              `json:"updated_time"`
    Comments            CommentArray        `json:"comments"`
}

type PhotosResponse struct {
    Data                []Photo             `json:"data"`
    Paging              Paging              `json:"paging"`
}

type Photo struct {
    Id                  string              `json:"id"`
    From                ContactReference    `json:"from"`
    Tags                UserPhotoTagArray   `json:"tags"`
    Picture             string              `json:"picture"`
    Source              string              `json:"source"`
    Height              int                 `json:"height"`
    Width               int                 `json:"width"`
    Images              []Image             `json:"images"`
    Link                string              `json:"link"`
    Icon                string              `json:"icon"`
    CreatedTime         string              `json:"created_time"`
    Position            int                 `json:"position"`
    UpdatedTime         string              `json:"updated_time"`
    Comments            CommentArray        `json:"comments"`
}

type UserPhotoTagArray struct {
    Data                []UserPhotoTag      `json:"data"`
}

type UserPhotoTag struct {
    Id                  string              `json:"id"`
    Name                string              `json:"name"`
    X                   float64             `json:"x"`
    Y                   float64             `json:"y"`
    CreatedTime         string              `json:"created_time"`
}

type Image struct {
    Height              int                 `json:"height"`
    Width               int                 `json:"width"`
    Source              string              `json:"source"`
}

type StatusesResponse struct {
    Data                []Status            `json:"data"`
}

type Status struct {
    Id                  string              `json:"id"`
    From                ContactReference    `json:"from"`
    Message             string              `json:"message"`
    UpdatedTime         string              `json:"updated_time"`
    Likes               LikeArray           `json:"likes"`
    Comments            CommentArray        `json:"comments"`
}


type TelevisionResponse struct {
    Data                []TVShow            `json:"data"`
}

type TVShow StandardMedia


type ErrorResponse struct {
    Error               Error               `json:"error"`
}

type Error struct {
    Message             string              `json:"message"`
    Type                string              `json:"type"`
}

