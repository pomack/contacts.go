package facebook

import (
    "github.com/pomack/oauth2_client"
    "io/ioutil"
    "json"
    "os"
)

func retrieveInfo(client oauth2_client.OAuth2Client, id, scope string, value interface{}) (err os.Error) {
    uri := FACEBOOK_GRAPH_API_ENDPOINT
    for _, s := range []string{id, scope} {
        if len(s) > 0 {
            if uri[len(uri) - 1] != '/' {
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
            b, _ := ioutil.ReadAll(resp.Body)
            err = os.NewError(string(b))
        } else {
            err = json.NewDecoder(resp.Body).Decode(value)
        }
    }
    return err
}

func RetrieveUserInfo(client oauth2_client.OAuth2Client) (*Contact, os.Error) {
    return RetrieveContact(client, "me")
}

func RetrieveContact(client oauth2_client.OAuth2Client, id string) (*Contact, os.Error) {
    contact := new(Contact)
    err := retrieveInfo(client, id, "", contact)
    return contact, err
}

func RetrieveActivities(client oauth2_client.OAuth2Client, id string) ([]Activity, os.Error) {
    resp := new(ActivitiesResponse)
    err := retrieveInfo(client, id, "activities", resp)
    return resp.Data, err
}

func RetrieveAlbums(client oauth2_client.OAuth2Client, id string) ([]Album, os.Error) {
    resp := new(AlbumsResponse)
    err := retrieveInfo(client, id, "albums", resp)
    return resp.Data, err
}

func RetrieveBooks(client oauth2_client.OAuth2Client, id string) ([]Book, os.Error) {
    resp := new(BooksResponse)
    err := retrieveInfo(client, id, "books", resp)
    return resp.Data, err
}

func RetrieveCheckins(client oauth2_client.OAuth2Client, id string) ([]Checkin, Paging, os.Error) {
    resp := new(CheckinsResponse)
    err := retrieveInfo(client, id, "checkins", resp)
    return resp.Data, resp.Paging, err
}

func RetrieveEvents(client oauth2_client.OAuth2Client, id string) ([]Event, Paging, os.Error) {
    resp := new(EventsResponse)
    err := retrieveInfo(client, id, "events", resp)
    return resp.Data, resp.Paging, err
}

func RetrieveFamily(client oauth2_client.OAuth2Client, id string) ([]Relationship, os.Error) {
    resp := new(FamilyResponse)
    err := retrieveInfo(client, id, "family", resp)
    return resp.Data, err
}

func RetrieveFriends(client oauth2_client.OAuth2Client, id string) ([]ContactReference, os.Error) {
    resp := new(FriendsResponse)
    err := retrieveInfo(client, id, "friends", resp)
    return resp.Data, err
}

func RetrieveGames(client oauth2_client.OAuth2Client, id string) ([]Game, os.Error) {
    resp := new(GamesResponse)
    err := retrieveInfo(client, id, "games", resp)
    return resp.Data, err
}

func RetrieveGroups(client oauth2_client.OAuth2Client, id string) ([]Group, os.Error) {
    resp := new(GroupsResponse)
    err := retrieveInfo(client, id, "groups", resp)
    return resp.Data, err
}

func RetrieveInterests(client oauth2_client.OAuth2Client, id string) ([]Interest, os.Error) {
    resp := new(InterestsResponse)
    err := retrieveInfo(client, id, "interests", resp)
    return resp.Data, err
}

func RetrieveLikes(client oauth2_client.OAuth2Client, id string) ([]LikeInfo, os.Error) {
    resp := new(LikesResponse)
    err := retrieveInfo(client, id, "likes", resp)
    return resp.Data, err
}

func RetrieveMovies(client oauth2_client.OAuth2Client, id string) ([]Movie, os.Error) {
    resp := new(MoviesResponse)
    err := retrieveInfo(client, id, "movies", resp)
    return resp.Data, err
}

func RetrieveMusic(client oauth2_client.OAuth2Client, id string) ([]Music, os.Error) {
    resp := new(MusicResponse)
    err := retrieveInfo(client, id, "music", resp)
    return resp.Data, err
}

func RetrieveNotes(client oauth2_client.OAuth2Client, id string) ([]Note, Paging, os.Error) {
    resp := new(NotesResponse)
    err := retrieveInfo(client, id, "notes", resp)
    return resp.Data, resp.Paging, err
}

func RetrievePhotos(client oauth2_client.OAuth2Client, id string) ([]Photo, Paging, os.Error) {
    resp := new(PhotosResponse)
    err := retrieveInfo(client, id, "photos", resp)
    return resp.Data, resp.Paging, err
}

func RetrieveStatuses(client oauth2_client.OAuth2Client, id string) ([]Status, os.Error) {
    resp := new(StatusesResponse)
    err := retrieveInfo(client, id, "statuses", resp)
    return resp.Data, err
}

func RetrieveTVShows(client oauth2_client.OAuth2Client, id string) ([]TVShow, os.Error) {
    resp := new(TelevisionResponse)
    err := retrieveInfo(client, id, "television", resp)
    return resp.Data, err
}
