package googleplus

const (
    GOOGLEPLUS_DATETIME_FORMAT = "2006-01-02T15:04:05.000Z"
    GOOGLEPLUS_OAUTH_SCOPE = "https://www.googleapis.com/auth/plus.me"
    GOOGLEPLUS_DEFAULT_USER_ID = "me"
    
    ACTIVITIES_LIST_ALT_PARAM = "alt"
    ACTIVITIES_LIST_MAXRESULTS_PARAM = "maxResults"
    ACTIVITIES_LIST_PAGETOKEN_PARAM = "pageToken"
    ACTIVITIES_LIST_USERID_PARAM = "userId"
    
    ACTIVITY_GET_ALT_PARAM = "alt"
    
    PERSON_KIND = "plus#person"
    ACTIVITY_LIST_KIND = "plus#activityFeed"
    ACTIVITY_KIND = "plus#activity"
    ACL_KIND = "plus#acl"
    
    PERSON_ENDPOINT = "https://www.googleapis.com/plus/v1/people/"
    ACTIVITY_LIST_ENDPOINT = "https://www.googleapis.com/plus/v1/people/"
    ACTIVITY_ENDPOINT = "https://www.googleapis.com/plus/v1/activities/"
)