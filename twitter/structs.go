package twitter

import (
	"github.com/pomack/jsonhelper"
)

type User struct {
	Id                             int64   `json:"id"`
	IdStr                          string  `json:"id_str"`
	Name                           string  `json:"name"`
	Location                       string  `json:"location"`
	ContributorsEnabled            bool    `json:"contributors_enabled"`
	Description                    string  `json:"description"`
	FavoritesCount                 int64   `json:"favourites_count"`
	Following                      bool    `json:"following"`
	FollowRequestSent              bool    `json:"follow_request_sent"`
	FollowersCount                 int64   `json:"followers_count"`
	FriendsCount                   int64   `json:"friends_count"`
	GeoEnabled                     bool    `json:"geo_enabled"`
	IsTranslator                   bool    `json:"is_translator"`
	Lang                           string  `json:"lang"`
	ListedCount                    int64   `json:"listed_count"`
	Notifications                  bool    `json:"notifications"`
	Protected                      bool    `json:"protected"`
	ScreenName                     string  `json:"screen_name"`
	ShowAllInlineMedia             bool    `json:"show_all_inline_media"`
	Status                         Status  `json:"status"`
	StatusesCount                  int64   `json:"statuses_count"`
	TimeZone                       string  `json:"time_zone"`
	Url                            *string `json:"url"`
	UtcOffset                      int     `json:"utc_offset"`
	Verified                       bool    `json:"verified"`
	ProfileBackgroundColor         string  `json:"profile_background_color"`
	ProfileBackgroundImageUrl      string  `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string  `json:"profile_background_image_url_https"`
	ProfileImageUrl                string  `json:"profile_image_url"`
	ProfileImageUrlHttps           string  `json:"profile_image_url_https"`
	ProfileLinkColor               string  `json:"profile_link_color"`
	ProfileSidebarFillColor        string  `json:"profile_sidebar_fill_color"`
	ProfileSidebarBorderColor      string  `json:"profile_sidebar_border_color"`
	ProfileTextColor               string  `json:"profile_text_color"`
	ProfileBackgroundTile          bool    `json:"profile_background_tile"`
	ProfileUseBackgroundImage      bool    `json:"profile_use_background_image"`
}

type Status struct {
	Id                   int64        `json:"id"`
	IdStr                string       `json:"id_str"`
	Text                 string       `json:"text"`
	Annotations          interface{}  `json:"annotations"`
	Coordinates          *GeoLocation `json:"coordinates"`
	CreatedAt            string       `json:"created_at"`
	Contributors         *[]int64     `json:"contributors"`
	Geo                  *GeoLocation `json:"geo"`
	InReplyToScreenName  *string      `json:"in_reply_to_screen_name"`
	InReplyToStatusId    *int64       `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr *string      `json:"in_reply_to_status_id_str"`
	InReplyToUserId      *int64       `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   *string      `json:"in_reply_to_user_id_str"`
	Source               string       `json:"source"`
	Place                interface{}  `json:"place"`
	RetweetCount         int64        `json:"retweet_count"`
	Favorited            bool         `json:"favorited"`
	Truncated            bool         `json:"truncated"`
	Retweeted            bool         `json:"retweeted"`
	Entities             *Entity      `json:"entities"`
	User                 *User        `json:"user"`
}

type GeoLocation struct {
	Type        string `json:"type"`
	Coordinates *[]int `json:"coordinates"`
}

type Entity struct {
	Media        []MediaEntity `json:"media"`
	Urls         []UrlInfo     `json:"urls"`
	Hashtags     []Hashtag     `json:"hashtags"`
	UserMentions []UserMention `json:"user_mentions"`
}

type MediaEntity struct {
	Id               int64      `json:"id"`
	IdStr            string     `json:"id_str"`
	MediaUrl         *string    `json:"media_url"`
	MediaUrlHttps    *string    `json:"media_url_https"`
	Url              string     `json:"url"`
	DisplayUrl       *string    `json:"display_url"`
	DisplayUrlHttps  *string    `json:"display_url_https"`
	ExpandedUrl      *string    `json:"expanded_url"`
	ExpandedUrlHttps *string    `json:"expanded_url_https"`
	Type             string     `json:"type"`
	Indices          []int      `json:"indices"`
	Sizes            ImageSizes `json:"sizes"`
}

type UrlInfo struct {
	Url         string  `json:"url"`
	DisplayUrl  *string `json:"display_url"`
	ExpandedUrl *string `json:"expanded_url"`
	Indices     []int   `json:"indices"`
}

type Hashtag struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
}

type UserMention struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	Indices    []int  `json:"indices"`
}

type ImageSizes struct {
	Large  ImageSize `json:"large"`
	Medium ImageSize `json:"medium"`
	Small  ImageSize `json:"small"`
	Thumb  ImageSize `json:"thumb"`
}

type ImageSize struct {
	W      int    `json:"w"`
	H      int    `json:"h"`
	Resize string `json:"resize"`
}

type GeoPlace struct {
	Id              string                `json:"id"`
	Name            string                `json:"name"`
	FullName        string                `json:"full_name"`
	Url             string                `json:"url"`
	PlaceType       string                `json:"place_type"`
	Polylines       []string              `json:"polylines"`
	Country         string                `json:"country"`
	CountryCode     string                `json:"country_code"`
	Attributes      jsonhelper.JSONObject `json:"attributes"`
	BoundingBox     BoundingBox           `json:"bounding_box"`
	Geometry        BoundingBox           `json:"geometry"`
	ContainedWithin []GeoPlace            `json:"contained_within"`
}

type BoundingBox struct {
	Coordinates [][]int `json:"coordinates"`
	Type        string  `json:"type"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Request string `json:"request"`
}

func (p *ErrorResponse) String() string { return p.Error }
