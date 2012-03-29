package googleplus

import (
    "encoding/json"
    "errors"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "io/ioutil"
    "net/url"
)

func retrieveInfo(client oauth2_client.OAuth2Client, uri string, m url.Values, value interface{}) (err error) {
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, nil, uri, m)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            b, _ := ioutil.ReadAll(resp.Body)
            err = errors.New(string(b))
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveSelfContact(client oauth2_client.OAuth2Client, m url.Values) (*Person, error) {
    return RetrieveContact(client, "", m)
}

func RetrieveContact(client oauth2_client.OAuth2Client, userId string, m url.Values) (*Person, error) {
    if userId == "" {
        userId = "me"
    }
    p := new(Person)
    uri := PERSON_ENDPOINT + url.QueryEscape(userId)
    err := retrieveInfo(client, uri, m, p)
    return p, err
}

func RetrieveSelfPublicActivityList(client oauth2_client.OAuth2Client, m url.Values) (*ActivityListResponse, error) {
    return RetrieveActivityList(client, "", "", m)
}

func RetrievePublicActivityList(client oauth2_client.OAuth2Client, userId string, m url.Values) (*ActivityListResponse, error) {
    return RetrieveActivityList(client, userId, "", m)
}

func RetrieveActivityList(client oauth2_client.OAuth2Client, userId, collection string, m url.Values) (*ActivityListResponse, error) {
    if userId == "" {
        userId = "me"
    }
    if collection == "" {
        collection = "public"
    }
    p := new(ActivityListResponse)
    uri := ACTIVITY_LIST_ENDPOINT + url.QueryEscape(userId) + "/activities/" + url.QueryEscape(collection)
    err := retrieveInfo(client, uri, m, p)
    return p, err
}

func RetrieveActivity(client oauth2_client.OAuth2Client, activityId string, m url.Values) (*Activity, error) {
    if activityId == "" {
        return nil, errors.New("activityId cannot be empty")
    }
    p := new(Activity)
    uri := ACTIVITY_ENDPOINT + url.QueryEscape(activityId)
    err := retrieveInfo(client, uri, m, p)
    return p, err
}
