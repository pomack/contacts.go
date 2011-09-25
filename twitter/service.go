package twitter

import (
    "github.com/pomack/oauth2_client"
    "http"
    "io/ioutil"
    "json"
    "os"
    "strconv"
    "url"
)

func retrieveInfo(client oauth2_client.OAuth2Client, apiUrl, id string, m url.Values, value interface{}) (err os.Error) {
    uri := TWITTER_API_ENDPOINT
    for _, s := range []string{apiUrl, id} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    headers := make(http.Header)
    headers.Set("Accept", "application/json")
    if m == nil {
        m = make(url.Values)
    }
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, headers, uri, m)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            var e ErrorResponse
            b, _ := ioutil.ReadAll(resp.Body)
            json.Unmarshal(b, &e)
            if len(e.Error) > 0 {
                err = os.NewError(e.Error)
            } else {
                err = os.NewError(string(b))
            }
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveSelfUser(client oauth2_client.OAuth2Client, m url.Values) (*User, os.Error) {
    resp := new(User)
    if m == nil {
        m = make(url.Values)
    }
    m.Set("include_entities", "true")
    err := retrieveInfo(client, "account/verify_credentials.json", "", m, resp)
    return resp, err
}

func RetrieveUser(client oauth2_client.OAuth2Client, screenName string, userId int64, includeEntities bool, m url.Values) (*User, os.Error) {
    resp := new(User)
    if m == nil {
        m = make(url.Values)
    }
    if userId > 0 {
        m.Set("user_id", strconv.Itoa64(userId))
    } else if len(screenName) > 0 {
        m.Set("screen_name", screenName)
    } else {
        return nil, os.NewError("Must specify either userId or screenName in twitter.RetrieveUser()")
    }
    if includeEntities {
        m.Set("include_entities", "true")
    }
    err := retrieveInfo(client, "users/show.json", "", m, resp)
    return resp, err
}

func RetrieveProfileImage(client oauth2_client.OAuth2Client, screenName, size string, userId int64, includeEntities bool, m url.Values) (*http.Response, os.Error) {
    uri := TWITTER_API_ENDPOINT + "users/profile_image"
    if m == nil {
        m = make(url.Values)
    }
    if len(screenName) > 0 {
        m.Set("screen_name", screenName)
    } else {
        return nil, os.NewError("Must specify either screenName in twitter.RetrieveProfileImage()")
    }
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, nil, uri, m)
    return resp, err
}

func RetrieveStatus(client oauth2_client.OAuth2Client, id int64, trimUser, includeEntities bool, m url.Values) (*Status, os.Error) {
    resp := new(Status)
    if m == nil {
        m = make(url.Values)
    }
    if id > 0 {
        m.Set("id", strconv.Itoa64(id))
    } else {
        return nil, os.NewError("Must specify either status Id in twitter.RetrieveStatus()")
    }
    if trimUser {
        m.Set("trim_user", "true")
    }
    if includeEntities {
        m.Set("include_entities", "true")
    }
    err := retrieveInfo(client, "statuses/show.json", "", m, resp)
    return resp, err
}

func RetrieveStatuses(client oauth2_client.OAuth2Client, screenName string, userId int64, trimUser, includeRts, includeEntities bool, m url.Values) (*[]Status, os.Error) {
    return RetrieveStatusesWithOptions(client, screenName, userId, 0, 0, 0, 0, trimUser, includeRts, includeEntities, true, true, m)
}

func RetrieveStatusesWithOptions(client oauth2_client.OAuth2Client, screenName string, userId, sinceId, count, maxId, page int64, trimUser, includeRts, includeEntities, excludeReplies, contributorDetails bool, m url.Values) (*[]Status, os.Error) {
    resp := new([]Status)
    if m == nil {
        m = make(url.Values)
    }
    if userId > 0 {
        m.Set("user_id", strconv.Itoa64(userId))
    } else if len(screenName) > 0 {
        m.Set("screen_name", screenName)
    } else {
        return nil, os.NewError("Must specify either userId or screenName in twitter.RetrieveStatus()")
    }
    if sinceId > 0 {
        m.Set("since_id", strconv.Itoa64(sinceId))
    }
    if count > 0 {
        m.Set("count", strconv.Itoa64(count))
    }
    if maxId > 0 {
        m.Set("max_id", strconv.Itoa64(maxId))
    }
    if page > 0 {
        m.Set("page", strconv.Itoa64(page))
    }
    if trimUser {
        m.Set("trim_user", "true")
    }
    if includeRts {
        m.Set("include_rts", "true")
    }
    if includeEntities {
        m.Set("include_entities", "true")
    }
    if excludeReplies {
        m.Set("exclude_replies", "true")
    }
    if contributorDetails {
        m.Set("contributor_details", "true")
    }
    err := retrieveInfo(client, "statuses/user_timeline.json", "", m, resp)
    return resp, err
}
