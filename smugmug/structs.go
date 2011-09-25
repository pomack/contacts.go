package smugmug


type GetFamilyResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Family                  []PersonReference       `json:"Family"`
}

type GetUserInfoResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    User                    PersonReference         `json:"User"`
}

type PersonReference struct {
    Name                    string                  `json:"Name"`
    NickName                string                  `json:"NickName"`
    Url                     string                  `json:"URL"`
}



type GetFansResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Fans                    []PersonReference       `json:"Fans"`
}

type GetFriendsResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Friends                 []PersonReference       `json:"Friends"`
}

type GetFeaturedAlbumsResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Featured                FeaturedAlbums          `json:"Featured"`
}

type FeaturedAlbums struct {
    Albums                  []AlbumReference        `json:"Albums"`
}

type AlbumReference struct {
    Id                      int64                   `json:"id"`
    Key                     string                  `json:"Key"`
}


type ImagesResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Album                   Album                   `json:"Album"`
}

type Album struct {
    Id                      int64                   `json:"id"`
    Key                     string                  `json:"Key"`
    Backprinting            *string                  `json:"Backprinting"`
    BoutiquePackaging       *int64                   `json:"BoutiquePackaging"`
    CanRank                 *bool                    `json:"CanRank"`
    Category                *Category                `json:"Category"`
    Clean                   *bool                    `json:"Clean"`
    ColorCorrection         *int64                   `json:"ColorCorrection"`
    Comments                *bool                    `json:"Comments"`
    Community               *Community               `json:"Community"`
    Description             *string                  `json:"Description"`
    Exif                    *bool                    `json:"EXIF"`
    External                *bool                    `json:"External"`
    FamilyEdit              *bool                    `json:"FamilyEdit"`
    Filenames               *bool                    `json:"Filenames"`
    FriendEdit              *bool                    `json:"FriendEdit"`
    Geography               *bool                    `json:"Geography"`
    Header                  *bool                    `json:"Header"`
    HideOwner               *bool                    `json:"HideOwner"`
    Highlight               *Highlight               `json:"Highlight"`
    ImageCount              *int64                   `json:"ImageCount"`
    InterceptShipping       *int64                   `json:"InterceptShipping"`
    Keywords                *string                  `json:"Keywords"`
    Larges                  *bool                    `json:"Larges"`
    LastUpdated             *string                  `json:"LastUpdated"`
    NiceName                *string                  `json:"NiceName"`
    Originals               *bool                    `json:"Originals"`
    PackagingBranding       *bool                    `json:"PackagingBranding"`
    Password                *string                  `json:"Password"`
    PasswordHint            *string                  `json:"PasswordHint"`
    Passworded              *bool                    `json:"Passworded"`
    Position                *int                     `json:"Position"`
    Printable               *bool                    `json:"Printable"`
    Printmark               *Printmark               `json:"Printmark"`
    ProofDays               *int64                   `json:"ProofDays"`
    Protected               *bool                    `json:"Protected"`
    Public                  *bool                    `json:"Public"`
    Share                   *bool                    `json:"Share"`
    SmugSearchable          *bool                    `json:"SmugSearchable"`
    SortDirection           *bool                    `json:"SortDirection"`
    SortMethod              *string                  `json:"SortMethod"`
    SquareThumbs            *bool                    `json:"SquareThumbs"`
    SubCategory             *SubCategory             `json:"SubCategory"`
    Template                *Template                `json:"Template"`
    Theme                   *Theme                   `json:"Theme"`
    Title                   *string                  `json:"Title"`
    Url                     *string                  `json:"URL"`
    UnsharpAmount           *float64                 `json:"UnsharpAmount"`
    UnsharpRadius           *float64                 `json:"UnsharpRadius"`
    UnsharpSigma            *float64                 `json:"UnsharpSigma"`
    UnsharpThreshold        *float64                 `json:"UnsharpThreshold"`
    Watermark               *Watermark               `json:"Watermark"`
    WorldSearchable         *bool                    `json:"WorldSearchable"`
    X2Larges                *bool                    `json:"X2Larges"`
    X3Larges                *bool                    `json:"X3Larges"`
    XLarges                 *bool                    `json:"XLarges"`
    Images                  *[]ImageReference        `json:"Images"`
}

type IdAndName struct {
    Id                      int64                   `json:"id"`
    Name                    string                  `json:"Name"`
}

type Category IdAndName
type SubCategory IdAndName
type Community IdAndName
type Printmark IdAndName
type Watermark IdAndName

type Template struct {
    Id                      int64                   `json:"id"`
}

type Theme struct {
    Id                      int64                   `json:"id"`
    Name                    string                  `json:"Name"`
    Type                    string                  `json:"Type"`
}

type Highlight struct {
    Id                      int64                   `json:"id"`
    Key                     string                  `json:"Key"`
    Type                    string                  `json:"Type"`
}

type ImageReference struct {
    Id                      int64                   `json:"id"`
    Key                     string                  `json:"Key"`
}

type AuthTokenResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Auth                    AuthObject              `json:"Auth"`
}

type AuthObject struct {
    Token                   TokenObject             `json:"Token"`
    User                    UserObject              `json:"User"`
}

type TokenObject struct {
    Id                      string                  `json:"id"`
    Secret                  string                  `json:"Secret"`
    Access                  string                  `json:"Access"`
    Permissions             string                  `json:"Permissions"`
}

type UserObject struct {
    Id                      int64                   `json:"id"`
    AccountStatus           string                  `json:"AccountStatus"`
    AccountType             string                  `json:"AccountType"`
    FileSizeLimit           int64                   `json:"FileSizeLimit"`
    Name                    string                  `json:"Name"`
    NickName                string                  `json:"NickName"`
    SmugVault               bool                    `json:"SmugVault"`
    Url                     string                  `json:"URL"`
}

type GetAlbumInfoResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Album                   Album                   `json:"Album"`
}

type GetAlbumsResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Albums                  []Album                 `json:"Albums"`
}

type AlbumStatsResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Album                   AlbumStats              `json:"Album"`
}

type AlbumStats struct {
    Id                      int64                   `json:"id"`
    Key                     string                  `json:"Key"`
    Bytes                   int64                   `json:"Bytes"`
    Hits                    int64                   `json:"Hits"`
    Large                   int64                   `json:"Large"`
    Medium                  int64                   `json:"Medium"`
    Original                int64                   `json:"Original"`
    Small                   int64                   `json:"Small"`
    Video1280               *int64                  `json:"Video1280"`
    Video1920               *int64                  `json:"Video1920"`
    Video320                *int64                  `json:"Video320"`
    Video640                *int64                  `json:"Video640"`
    Video960                *int64                  `json:"Video960"`
    X2Large                 int64                   `json:"X2Large"`
    X3Large                 int64                   `json:"X3Large"`
    XLarge                  int64                   `json:"XLarge"`
    Images                  *[]ImageStats           `json:"Images"`
}

type ImageStats struct {
    Id                      int64                   `json:"id"`
    Key                     string                  `json:"Key"`
    Bytes                   int64                   `json:"Bytes"`
    Hits                    int64                   `json:"Hits"`
    Large                   int64                   `json:"Large"`
    Medium                  int64                   `json:"Medium"`
    Original                int64                   `json:"Original"`
    Small                   int64                   `json:"Small"`
    Video1280               int64                   `json:"Video1280"`
    Video1920               int64                   `json:"Video1920"`
    Video320                int64                   `json:"Video320"`
    Video640                int64                   `json:"Video640"`
    Video960                int64                   `json:"Video960"`
    X2Large                 int64                   `json:"X2Large"`
    X3Large                 int64                   `json:"X3Large"`
    XLarge                  int64                   `json:"XLarge"`
}

type ErrorResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Code                    int                     `json:"code"`
    Message                 string                  `json:"message"`
}
func (p *ErrorResponse) String() string { return p.Message }

type GetCategoriesResponse struct {
    Stat                    string                  `json:"stat"`
    Method                  string                  `json:"method"`
    Categories              []CategoryInfo          `json:"Categories"`
}

type CategoryInfo struct {
    Id                      int64                   `json:"id"`
    Name                    string                  `json:"Name"`
    NiceName                string                  `json:"NiceName"`
    Type                    string                  `json:"Type"`
}


