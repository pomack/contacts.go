package linkedin

const (
    LINKEDIN_DEFAULT_USER_ID = "~"

    LINKEDIN_API_ENDPOINT           = "http://api.linkedin.com/v1/"
    LINKEDIN_DEFAULT_PROFILE_FIELDS = `picture-url,id,first-name,last-name,educations,positions,skills,certifications,publications,honors,summary,specialties,twitter-accounts,interests,languages,location,num-connections,num-connections-capped,phone-numbers,main-address,date-of-birth,headline,industry,associations,member-url-resources,public-profile-url,patents,api-standard-profile-request,im-accounts`

    // im account types
    IM_ACCOUNT_TYPE_AIM   = "aim"
    IM_ACCOUNT_TYPE_GTALK = "gtalk"
    IM_ACCOUNT_TYPE_ICQ   = "icq"
    IM_ACCOUNT_TYPE_MSN   = "msn"
    IM_ACCOUNT_TYPE_SKYPE = "skype"
    IM_ACCOUNT_TYPE_YAHOO = "yahoo"

    // phone types
    PHONE_TYPE_HOME   = "home"
    PHONE_TYPE_WORK   = "work"
    PHONE_TYPE_MOBILE = "mobile"
)
