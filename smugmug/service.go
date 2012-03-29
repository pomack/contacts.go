package smugmug

import (
    "encoding/json"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
)

func retrieveInfo(client oauth2_client.OAuth2Client, method string, m url.Values, value interface{}) (err error) {
    uri := SMUGMUG_API_ENDPOINT
    headers := make(http.Header)
    headers.Set("Accept", "application/json")
    if m == nil {
        m = make(url.Values)
    }
    m.Set("method", method)
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, headers, uri, m)
    if err != nil {
        return err
    }
    if resp != nil {
        var e ErrorResponse
        b, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(b, &e)
        if e.Stat != "ok" {
            err = &e
        } else {
            err = json.Unmarshal(b, value)
        }
    }
    return err
}

func CheckAccessToken(client oauth2_client.OAuth2Client, m url.Values) (*AuthTokenResponse, error) {
    resp := new(AuthTokenResponse)
    err := retrieveInfo(client, "smugmug.auth.checkAccessToken", m, resp)
    return resp, err
}

func RetrieveAlbums(client oauth2_client.OAuth2Client, heavy bool, m url.Values) (*GetAlbumsResponse, error) {
    resp := new(GetAlbumsResponse)
    if m == nil {
        m = make(url.Values)
    }
    if heavy {
        m.Set("Heavy", "true")
    }
    err := retrieveInfo(client, "smugmug.albums.get", m, resp)
    return resp, err
}

func RetrieveAlbumInfo(client oauth2_client.OAuth2Client, albumId int64, albumKey string, m url.Values) (*GetAlbumInfoResponse, error) {
    resp := new(GetAlbumInfoResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("AlbumID", strconv.FormatInt(albumId, 10))
    m.Set("AlbumKey", albumKey)
    err := retrieveInfo(client, "smugmug.albums.getInfo", m, resp)
    return resp, err
}

func RetrieveAlbumStats(client oauth2_client.OAuth2Client, albumId int64, year, month int, heavy bool, m url.Values) (*AlbumStatsResponse, error) {
    resp := new(AlbumStatsResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("AlbumID", strconv.FormatInt(albumId, 10))
    m.Set("Year", strconv.Itoa(year))
    m.Set("Month", strconv.Itoa(month))
    if heavy {
        m.Set("Heavy", "true")
    }
    err := retrieveInfo(client, "smugmug.albums.getStats", m, resp)
    return resp, err
}

func RetrieveCategories(client oauth2_client.OAuth2Client, m url.Values) (*GetCategoriesResponse, error) {
    resp := new(GetCategoriesResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.categories.get", m, resp)
    return resp, err
}

func RetrieveUserInfo(client oauth2_client.OAuth2Client, nickName string, m url.Values) (*GetUserInfoResponse, error) {
    resp := new(GetUserInfoResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("NickName", nickName)
    err := retrieveInfo(client, "smugmug.users.getInfo", m, resp)
    return resp, err
}

func RetrieveFeaturedAlbums(client oauth2_client.OAuth2Client, m url.Values) (*GetFeaturedAlbumsResponse, error) {
    resp := new(GetFeaturedAlbumsResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.featured.albums.get", m, resp)
    return resp, err
}

func RetrieveFamily(client oauth2_client.OAuth2Client, m url.Values) (*GetFamilyResponse, error) {
    resp := new(GetFamilyResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.family.get", m, resp)
    return resp, err
}

func RetrieveFans(client oauth2_client.OAuth2Client, m url.Values) (*GetFansResponse, error) {
    resp := new(GetFansResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.fans.get", m, resp)
    return resp, err
}

func RetrieveFriends(client oauth2_client.OAuth2Client, m url.Values) (*GetFriendsResponse, error) {
    resp := new(GetFriendsResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.friends.get", m, resp)
    return resp, err
}
