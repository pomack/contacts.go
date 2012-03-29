package facebook

import (
    "encoding/json"
    "errors"
    "github.com/pomack/oauth2_client.go/oauth2_client"
    "io/ioutil"
)

func retrieveInfo(client oauth2_client.OAuth2Client, id, scope string, value interface{}) (err error) {
    uri := FACEBOOK_GRAPH_API_ENDPOINT
    for _, s := range []string{id, scope} {
        if len(s) > 0 {
            if uri[len(uri)-1] != '/' {
                uri += "/"
            }
            uri += s
        }
    }
    resp, _, err := oauth2_client.AuthorizedGetRequest(client, nil, uri, nil)
    if err != nil {
        return err
    }
    if resp != nil {
        if resp.StatusCode >= 400 {
            e := new(ErrorResponse)
            b, _ := ioutil.ReadAll(resp.Body)
            if err = json.Unmarshal(b, e); err != nil {
                err = errors.New(string(b))
            } else {
                err = e
            }
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveUserInfo(client oauth2_client.OAuth2Client) (*Contact, error) {
    return RetrieveContact(client, "me")
}

func RetrieveContact(client oauth2_client.OAuth2Client, id string) (*Contact, error) {
    contact := new(Contact)
    err := retrieveInfo(client, id, "", contact)
    return contact, err
}

func RetrieveActivities(client oauth2_client.OAuth2Client, id string) ([]Activity, error) {
    resp := new(ActivitiesResponse)
    err := retrieveInfo(client, id, "activities", resp)
    return resp.Data, err
}

func RetrieveAlbums(client oauth2_client.OAuth2Client, id string) ([]Album, error) {
    resp := new(AlbumsResponse)
    err := retrieveInfo(client, id, "albums", resp)
    return resp.Data, err
}

func RetrieveBooks(client oauth2_client.OAuth2Client, id string) ([]Book, error) {
    resp := new(BooksResponse)
    err := retrieveInfo(client, id, "books", resp)
    return resp.Data, err
}

func RetrieveCheckins(client oauth2_client.OAuth2Client, id string) ([]Checkin, Paging, error) {
    resp := new(CheckinsResponse)
    err := retrieveInfo(client, id, "checkins", resp)
    return resp.Data, resp.Paging, err
}

func RetrieveEvents(client oauth2_client.OAuth2Client, id string) ([]Event, Paging, error) {
    resp := new(EventsResponse)
    err := retrieveInfo(client, id, "events", resp)
    return resp.Data, resp.Paging, err
}

func RetrieveFamily(client oauth2_client.OAuth2Client, id string) ([]Relationship, error) {
    resp := new(FamilyResponse)
    err := retrieveInfo(client, id, "family", resp)
    return resp.Data, err
}

func RetrieveFriends(client oauth2_client.OAuth2Client, id string) ([]ContactReference, error) {
    resp := new(FriendsResponse)
    err := retrieveInfo(client, id, "friends", resp)
    return resp.Data, err
}

func RetrieveGames(client oauth2_client.OAuth2Client, id string) ([]Game, error) {
    resp := new(GamesResponse)
    err := retrieveInfo(client, id, "games", resp)
    return resp.Data, err
}

func RetrieveGroups(client oauth2_client.OAuth2Client, id string) ([]Group, error) {
    resp := new(GroupsResponse)
    err := retrieveInfo(client, id, "groups", resp)
    return resp.Data, err
}

func RetrieveInterests(client oauth2_client.OAuth2Client, id string) ([]Interest, error) {
    resp := new(InterestsResponse)
    err := retrieveInfo(client, id, "interests", resp)
    return resp.Data, err
}

func RetrieveLikes(client oauth2_client.OAuth2Client, id string) ([]LikeInfo, error) {
    resp := new(LikesResponse)
    err := retrieveInfo(client, id, "likes", resp)
    return resp.Data, err
}

func RetrieveMovies(client oauth2_client.OAuth2Client, id string) ([]Movie, error) {
    resp := new(MoviesResponse)
    err := retrieveInfo(client, id, "movies", resp)
    return resp.Data, err
}

func RetrieveMusic(client oauth2_client.OAuth2Client, id string) ([]Music, error) {
    resp := new(MusicResponse)
    err := retrieveInfo(client, id, "music", resp)
    return resp.Data, err
}

func RetrieveNotes(client oauth2_client.OAuth2Client, id string) ([]Note, Paging, error) {
    resp := new(NotesResponse)
    err := retrieveInfo(client, id, "notes", resp)
    return resp.Data, resp.Paging, err
}

func RetrievePhotos(client oauth2_client.OAuth2Client, id string) ([]Photo, Paging, error) {
    resp := new(PhotosResponse)
    err := retrieveInfo(client, id, "photos", resp)
    return resp.Data, resp.Paging, err
}

func RetrieveStatuses(client oauth2_client.OAuth2Client, id string) ([]Status, error) {
    resp := new(StatusesResponse)
    err := retrieveInfo(client, id, "statuses", resp)
    return resp.Data, err
}

func RetrieveTVShows(client oauth2_client.OAuth2Client, id string) ([]TVShow, error) {
    resp := new(TelevisionResponse)
    err := retrieveInfo(client, id, "television", resp)
    return resp.Data, err
}
