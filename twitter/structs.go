package twitter

import (
    "github.com/pomack/jsonhelper"
)

type User struct {
    Id                             int64   `json:"id,omitempty"`
    IdStr                          string  `json:"id_str,omitempty"`
    Name                           string  `json:"name,omitempty"`
    Location                       string  `json:"location,omitempty"`
    ContributorsEnabled            bool    `json:"contributors_enabled,omitempty"`
    Description                    string  `json:"description,omitempty"`
    FavoritesCount                 int64   `json:"favourites_count,omitempty"`
    Following                      bool    `json:"following,omitempty"`
    FollowRequestSent              bool    `json:"follow_request_sent,omitempty"`
    FollowersCount                 int64   `json:"followers_count,omitempty"`
    FriendsCount                   int64   `json:"friends_count,omitempty"`
    GeoEnabled                     bool    `json:"geo_enabled,omitempty"`
    IsTranslator                   bool    `json:"is_translator,omitempty"`
    Lang                           string  `json:"lang,omitempty"`
    ListedCount                    int64   `json:"listed_count,omitempty"`
    Notifications                  bool    `json:"notifications,omitempty"`
    Protected                      bool    `json:"protected,omitempty"`
    ScreenName                     string  `json:"screen_name,omitempty"`
    ShowAllInlineMedia             bool    `json:"show_all_inline_media,omitempty"`
    Status                         Status  `json:"status,omitempty"`
    StatusesCount                  int64   `json:"statuses_count,omitempty"`
    TimeZone                       string  `json:"time_zone,omitempty"`
    Url                            *string `json:"url,omitempty"`
    UtcOffset                      int     `json:"utc_offset,omitempty"`
    Verified                       bool    `json:"verified,omitempty"`
    ProfileBackgroundColor         string  `json:"profile_background_color,omitempty"`
    ProfileBackgroundImageUrl      string  `json:"profile_background_image_url,omitempty"`
    ProfileBackgroundImageUrlHttps string  `json:"profile_background_image_url_https,omitempty"`
    ProfileImageUrl                string  `json:"profile_image_url,omitempty"`
    ProfileImageUrlHttps           string  `json:"profile_image_url_https,omitempty"`
    ProfileLinkColor               string  `json:"profile_link_color,omitempty"`
    ProfileSidebarFillColor        string  `json:"profile_sidebar_fill_color,omitempty"`
    ProfileSidebarBorderColor      string  `json:"profile_sidebar_border_color,omitempty"`
    ProfileTextColor               string  `json:"profile_text_color,omitempty"`
    ProfileBackgroundTile          bool    `json:"profile_background_tile,omitempty"`
    ProfileUseBackgroundImage      bool    `json:"profile_use_background_image,omitempty"`
}

type Status struct {
    Id                   int64        `json:"id,omitempty"`
    IdStr                string       `json:"id_str,omitempty"`
    Text                 string       `json:"text,omitempty"`
    Annotations          interface{}  `json:"annotations,omitempty"`
    Coordinates          *GeoLocation `json:"coordinates,omitempty"`
    CreatedAt            string       `json:"created_at,omitempty"`
    Contributors         *[]int64     `json:"contributors,omitempty"`
    Geo                  *GeoLocation `json:"geo,omitempty"`
    InReplyToScreenName  *string      `json:"in_reply_to_screen_name,omitempty"`
    InReplyToStatusId    *int64       `json:"in_reply_to_status_id,omitempty"`
    InReplyToStatusIdStr *string      `json:"in_reply_to_status_id_str,omitempty"`
    InReplyToUserId      *int64       `json:"in_reply_to_user_id,omitempty"`
    InReplyToUserIdStr   *string      `json:"in_reply_to_user_id_str,omitempty"`
    Source               string       `json:"source,omitempty"`
    Place                interface{}  `json:"place,omitempty"`
    RetweetCount         int64        `json:"retweet_count,omitempty"`
    Favorited            bool         `json:"favorited,omitempty"`
    Truncated            bool         `json:"truncated,omitempty"`
    Retweeted            bool         `json:"retweeted,omitempty"`
    Entities             *Entity      `json:"entities,omitempty"`
    User                 *User        `json:"user,omitempty"`
}

type GeoLocation struct {
    Type        string `json:"type,omitempty"`
    Coordinates *[]int `json:"coordinates,omitempty"`
}

type Entity struct {
    Media        []MediaEntity `json:"media,omitempty"`
    Urls         []UrlInfo     `json:"urls,omitempty"`
    Hashtags     []Hashtag     `json:"hashtags,omitempty"`
    UserMentions []UserMention `json:"user_mentions,omitempty"`
}

type MediaEntity struct {
    Id               int64      `json:"id,omitempty"`
    IdStr            string     `json:"id_str,omitempty"`
    MediaUrl         *string    `json:"media_url,omitempty"`
    MediaUrlHttps    *string    `json:"media_url_https,omitempty"`
    Url              string     `json:"url,omitempty"`
    DisplayUrl       *string    `json:"display_url,omitempty"`
    DisplayUrlHttps  *string    `json:"display_url_https,omitempty"`
    ExpandedUrl      *string    `json:"expanded_url,omitempty"`
    ExpandedUrlHttps *string    `json:"expanded_url_https,omitempty"`
    Type             string     `json:"type,omitempty"`
    Indices          []int      `json:"indices,omitempty"`
    Sizes            ImageSizes `json:"sizes,omitempty"`
}

type UrlInfo struct {
    Url         string  `json:"url,omitempty"`
    DisplayUrl  *string `json:"display_url,omitempty"`
    ExpandedUrl *string `json:"expanded_url,omitempty"`
    Indices     []int   `json:"indices,omitempty"`
}

type Hashtag struct {
    Text    string `json:"text,omitempty"`
    Indices []int  `json:"indices,omitempty"`
}

type UserMention struct {
    Id         int64  `json:"id,omitempty"`
    IdStr      string `json:"id_str,omitempty"`
    ScreenName string `json:"screen_name,omitempty"`
    Name       string `json:"name,omitempty"`
    Indices    []int  `json:"indices,omitempty"`
}

type ImageSizes struct {
    Large  ImageSize `json:"large,omitempty"`
    Medium ImageSize `json:"medium,omitempty"`
    Small  ImageSize `json:"small,omitempty"`
    Thumb  ImageSize `json:"thumb,omitempty"`
}

type ImageSize struct {
    W      int    `json:"w,omitempty"`
    H      int    `json:"h,omitempty"`
    Resize string `json:"resize,omitempty"`
}

type GeoPlace struct {
    Id              string                `json:"id"`
    Name            string                `json:"name,omitempty"`
    FullName        string                `json:"full_name,omitempty"`
    Url             string                `json:"url,omitempty"`
    PlaceType       string                `json:"place_type,omitempty"`
    Polylines       []string              `json:"polylines,omitempty"`
    Country         string                `json:"country,omitempty"`
    CountryCode     string                `json:"country_code,omitempty"`
    Attributes      jsonhelper.JSONObject `json:"attributes,omitempty"`
    BoundingBox     BoundingBox           `json:"bounding_box,omitempty"`
    Geometry        BoundingBox           `json:"geometry,omitempty"`
    ContainedWithin []GeoPlace            `json:"contained_within,omitempty"`
}

type BoundingBox struct {
    Coordinates [][]int `json:"coordinates,omitempty"`
    Type        string  `json:"type,omitempty"`
}

type ErrorResponse struct {
    Error   string `json:"error,omitempty"`
    Request string `json:"request,omitempty"`
}

func (p *ErrorResponse) String() string { return p.Error }
