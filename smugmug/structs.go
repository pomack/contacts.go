package smugmug

type GetFamilyResponse struct {
    Stat   string            `json:"stat,omitempty"`
    Method string            `json:"method,omitempty"`
    Family []PersonReference `json:"Family,omitempty"`
}

type GetUserInfoResponse struct {
    Stat   string          `json:"stat,omitempty"`
    Method string          `json:"method,omitempty"`
    User   PersonReference `json:"User,omitempty"`
}

type PersonReference struct {
    Name     string `json:"Name,omitempty"`
    NickName string `json:"NickName,omitempty"`
    Url      string `json:"URL,omitempty"`
}

type GetFansResponse struct {
    Stat   string            `json:"stat,omitempty"`
    Method string            `json:"method,omitempty"`
    Fans   []PersonReference `json:"Fans,omitempty"`
}

type GetFriendsResponse struct {
    Stat    string            `json:"stat,omitempty"`
    Method  string            `json:"method,omitempty"`
    Friends []PersonReference `json:"Friends,omitempty"`
}

type GetFeaturedAlbumsResponse struct {
    Stat     string         `json:"stat,omitempty"`
    Method   string         `json:"method,omitempty"`
    Featured FeaturedAlbums `json:"Featured,omitempty"`
}

type FeaturedAlbums struct {
    Albums []AlbumReference `json:"Albums,omitempty"`
}

type AlbumReference struct {
    Id  int64  `json:"id,omitempty"`
    Key string `json:"Key,omitempty"`
}

type ImagesResponse struct {
    Stat   string `json:"stat,omitempty"`
    Method string `json:"method,omitempty"`
    Album  Album  `json:"Album,omitempty"`
}

type Album struct {
    Id                int64             `json:"id,omitempty"`
    Key               string            `json:"Key,omitempty"`
    Backprinting      *string           `json:"Backprinting,omitempty"`
    BoutiquePackaging *int64            `json:"BoutiquePackaging,omitempty"`
    CanRank           *bool             `json:"CanRank,omitempty"`
    Category          *Category         `json:"Category,omitempty"`
    Clean             *bool             `json:"Clean,omitempty"`
    ColorCorrection   *int64            `json:"ColorCorrection,omitempty"`
    Comments          *bool             `json:"Comments,omitempty"`
    Community         *Community        `json:"Community,omitempty"`
    Description       *string           `json:"Description,omitempty"`
    Exif              *bool             `json:"EXIF,omitempty"`
    External          *bool             `json:"External,omitempty"`
    FamilyEdit        *bool             `json:"FamilyEdit,omitempty"`
    Filenames         *bool             `json:"Filenames,omitempty"`
    FriendEdit        *bool             `json:"FriendEdit,omitempty"`
    Geography         *bool             `json:"Geography,omitempty"`
    Header            *bool             `json:"Header,omitempty"`
    HideOwner         *bool             `json:"HideOwner,omitempty"`
    Highlight         *Highlight        `json:"Highlight,omitempty"`
    ImageCount        *int64            `json:"ImageCount,omitempty"`
    InterceptShipping *int64            `json:"InterceptShipping,omitempty"`
    Keywords          *string           `json:"Keywords,omitempty"`
    Larges            *bool             `json:"Larges,omitempty"`
    LastUpdated       *string           `json:"LastUpdated,omitempty"`
    NiceName          *string           `json:"NiceName,omitempty"`
    Originals         *bool             `json:"Originals,omitempty"`
    PackagingBranding *bool             `json:"PackagingBranding,omitempty"`
    Password          *string           `json:"Password,omitempty"`
    PasswordHint      *string           `json:"PasswordHint,omitempty"`
    Passworded        *bool             `json:"Passworded,omitempty"`
    Position          *int              `json:"Position,omitempty"`
    Printable         *bool             `json:"Printable,omitempty"`
    Printmark         *Printmark        `json:"Printmark,omitempty"`
    ProofDays         *int64            `json:"ProofDays,omitempty"`
    Protected         *bool             `json:"Protected,omitempty"`
    Public            *bool             `json:"Public,omitempty"`
    Share             *bool             `json:"Share,omitempty"`
    SmugSearchable    *bool             `json:"SmugSearchable,omitempty"`
    SortDirection     *bool             `json:"SortDirection,omitempty"`
    SortMethod        *string           `json:"SortMethod,omitempty"`
    SquareThumbs      *bool             `json:"SquareThumbs,omitempty"`
    SubCategory       *SubCategory      `json:"SubCategory,omitempty"`
    Template          *Template         `json:"Template,omitempty"`
    Theme             *Theme            `json:"Theme,omitempty"`
    Title             *string           `json:"Title,omitempty"`
    Url               *string           `json:"URL,omitempty"`
    UnsharpAmount     *float64          `json:"UnsharpAmount,omitempty"`
    UnsharpRadius     *float64          `json:"UnsharpRadius,omitempty"`
    UnsharpSigma      *float64          `json:"UnsharpSigma,omitempty"`
    UnsharpThreshold  *float64          `json:"UnsharpThreshold,omitempty"`
    Watermark         *Watermark        `json:"Watermark,omitempty"`
    WorldSearchable   *bool             `json:"WorldSearchable,omitempty"`
    X2Larges          *bool             `json:"X2Larges,omitempty"`
    X3Larges          *bool             `json:"X3Larges,omitempty"`
    XLarges           *bool             `json:"XLarges,omitempty"`
    Images            *[]ImageReference `json:"Images,omitempty"`
}

type IdAndName struct {
    Id   int64  `json:"id,omitempty"`
    Name string `json:"Name,omitempty"`
}

type Category IdAndName
type SubCategory IdAndName
type Community IdAndName
type Printmark IdAndName
type Watermark IdAndName

type Template struct {
    Id int64 `json:"id,omitempty"`
}

type Theme struct {
    Id   int64  `json:"id,omitempty"`
    Name string `json:"Name,omitempty"`
    Type string `json:"Type,omitempty"`
}

type Highlight struct {
    Id   int64  `json:"id,omitempty"`
    Key  string `json:"Key,omitempty"`
    Type string `json:"Type,omitempty"`
}

type ImageReference struct {
    Id  int64  `json:"id,omitempty"`
    Key string `json:"Key,omitempty"`
}

type AuthTokenResponse struct {
    Stat   string     `json:"stat,omitempty"`
    Method string     `json:"method,omitempty"`
    Auth   AuthObject `json:"Auth,omitempty"`
}

type AuthObject struct {
    Token TokenObject `json:"Token,omitempty"`
    User  UserObject  `json:"User,omitempty"`
}

type TokenObject struct {
    Id          string `json:"id,omitempty"`
    Secret      string `json:"Secret,omitempty"`
    Access      string `json:"Access,omitempty"`
    Permissions string `json:"Permissions,omitempty"`
}

type UserObject struct {
    Id            int64  `json:"id,omitempty"`
    AccountStatus string `json:"AccountStatus,omitempty"`
    AccountType   string `json:"AccountType,omitempty"`
    FileSizeLimit int64  `json:"FileSizeLimit,omitempty"`
    Name          string `json:"Name,omitempty"`
    NickName      string `json:"NickName,omitempty"`
    SmugVault     bool   `json:"SmugVault,omitempty"`
    Url           string `json:"URL,omitempty"`
}

type GetAlbumInfoResponse struct {
    Stat   string `json:"stat,omitempty"`
    Method string `json:"method,omitempty"`
    Album  Album  `json:"Album,omitempty"`
}

type GetAlbumsResponse struct {
    Stat   string  `json:"stat,omitempty"`
    Method string  `json:"method,omitempty"`
    Albums []Album `json:"Albums,omitempty"`
}

type AlbumStatsResponse struct {
    Stat   string     `json:"stat,omitempty"`
    Method string     `json:"method,omitempty"`
    Album  AlbumStats `json:"Album,omitempty"`
}

type AlbumStats struct {
    Id        int64         `json:"id,omitempty"`
    Key       string        `json:"Key,omitempty"`
    Bytes     int64         `json:"Bytes,omitempty"`
    Hits      int64         `json:"Hits,omitempty"`
    Large     int64         `json:"Large,omitempty"`
    Medium    int64         `json:"Medium,omitempty"`
    Original  int64         `json:"Original,omitempty"`
    Small     int64         `json:"Small,omitempty"`
    Video1280 *int64        `json:"Video1280,omitempty"`
    Video1920 *int64        `json:"Video1920,omitempty"`
    Video320  *int64        `json:"Video320,omitempty"`
    Video640  *int64        `json:"Video640,omitempty"`
    Video960  *int64        `json:"Video960,omitempty"`
    X2Large   int64         `json:"X2Large,omitempty"`
    X3Large   int64         `json:"X3Large,omitempty"`
    XLarge    int64         `json:"XLarge,omitempty"`
    Images    *[]ImageStats `json:"Images,omitempty"`
}

type ImageStats struct {
    Id        int64  `json:"id,omitempty"`
    Key       string `json:"Key,omitempty"`
    Bytes     int64  `json:"Bytes,omitempty"`
    Hits      int64  `json:"Hits,omitempty"`
    Large     int64  `json:"Large,omitempty"`
    Medium    int64  `json:"Medium,omitempty"`
    Original  int64  `json:"Original,omitempty"`
    Small     int64  `json:"Small,omitempty"`
    Video1280 int64  `json:"Video1280,omitempty"`
    Video1920 int64  `json:"Video1920,omitempty"`
    Video320  int64  `json:"Video320,omitempty"`
    Video640  int64  `json:"Video640,omitempty"`
    Video960  int64  `json:"Video960,omitempty"`
    X2Large   int64  `json:"X2Large,omitempty"`
    X3Large   int64  `json:"X3Large,omitempty"`
    XLarge    int64  `json:"XLarge,omitempty"`
}

type ErrorResponse struct {
    Stat    string `json:"stat,omitempty"`
    Method  string `json:"method,omitempty"`
    Code    int    `json:"code,omitempty"`
    Message string `json:"message,omitempty"`
}

func (p *ErrorResponse) Error() string  { return p.Message }
func (p *ErrorResponse) String() string { return p.Message }

type GetCategoriesResponse struct {
    Stat       string         `json:"stat,omitempty"`
    Method     string         `json:"method,omitempty"`
    Categories []CategoryInfo `json:"Categories,omitempty"`
}

type CategoryInfo struct {
    Id       int64  `json:"id,omitempty"`
    Name     string `json:"Name,omitempty"`
    NiceName string `json:"NiceName,omitempty"`
    Type     string `json:"Type,omitempty"`
}
