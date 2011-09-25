package smugmug

import (
    "github.com/pomack/oauth2_client"
    "http"
    "io/ioutil"
    "json"
    "os"
    "strconv"
    "url"
)

func retrieveInfo(client oauth2_client.OAuth2Client, method string, m url.Values, value interface{}) (err os.Error) {
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

func CheckAccessToken(client oauth2_client.OAuth2Client, m url.Values) (*AuthTokenResponse, os.Error) {
    resp := new(AuthTokenResponse)
    err := retrieveInfo(client, "smugmug.auth.checkAccessToken", m, resp)
    return resp, err
}

func RetrieveAlbums(client oauth2_client.OAuth2Client, heavy bool, m url.Values) (*GetAlbumsResponse, os.Error) {
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

func RetrieveAlbumInfo(client oauth2_client.OAuth2Client, albumId int64, albumKey string, m url.Values) (*GetAlbumInfoResponse, os.Error) {
    resp := new(GetAlbumInfoResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("AlbumID", strconv.Itoa64(albumId))
    m.Set("AlbumKey", albumKey)
    err := retrieveInfo(client, "smugmug.albums.getInfo", m, resp)
    return resp, err
}

func RetrieveAlbumStats(client oauth2_client.OAuth2Client, albumId int64, year, month int, heavy bool, m url.Values) (*AlbumStatsResponse, os.Error) {
    resp := new(AlbumStatsResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("AlbumID", strconv.Itoa64(albumId))
    m.Set("Year", strconv.Itoa(year))
    m.Set("Month", strconv.Itoa(month))
    if heavy {
        m.Set("Heavy", "true")
    }
    err := retrieveInfo(client, "smugmug.albums.getStats", m, resp)
    return resp, err
}

func RetrieveCategories(client oauth2_client.OAuth2Client, m url.Values) (*GetCategoriesResponse, os.Error) {
    resp := new(GetCategoriesResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.categories.get", m, resp)
    return resp, err
}

func RetrieveUserInfo(client oauth2_client.OAuth2Client, nickName string, m url.Values) (*GetUserInfoResponse, os.Error) {
    resp := new(GetUserInfoResponse)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("NickName", nickName)
    err := retrieveInfo(client, "smugmug.users.getInfo", m, resp)
    return resp, err
}

func RetrieveFeaturedAlbums(client oauth2_client.OAuth2Client, m url.Values) (*GetFeaturedAlbumsResponse, os.Error) {
    resp := new(GetFeaturedAlbumsResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.featured.albums.get", m, resp)
    return resp, err
}

func RetrieveFamily(client oauth2_client.OAuth2Client, m url.Values) (*GetFamilyResponse, os.Error) {
    resp := new(GetFamilyResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.family.get", m, resp)
    return resp, err
}

func RetrieveFans(client oauth2_client.OAuth2Client, m url.Values) (*GetFansResponse, os.Error) {
    resp := new(GetFansResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.fans.get", m, resp)
    return resp, err
}

func RetrieveFriends(client oauth2_client.OAuth2Client, m url.Values) (*GetFriendsResponse, os.Error) {
    resp := new(GetFriendsResponse)
    if m == nil {
        m = make(url.Values)
    }
    err := retrieveInfo(client, "smugmug.friends.get", m, resp)
    return resp, err
}
