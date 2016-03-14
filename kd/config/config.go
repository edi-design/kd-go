// Author: andre@freshest.me
// Date: 23.04.2015
// Version: 1

// configuration structs and global const.
package config

// Api configuration struct
type Config struct {
	Service struct {
		Username string
		Password string
		Listen   string
	}
}

// struct to login to api
type SignIn struct {
	DomainID    int
	SiteGuid    string
	LoginStatus int
	UserData    interface{}
}

type ChannelList []Channel

// channellist
// ToDo http://freshest.me/how-to-reverse-engineer-the-kabeldeutschland-tv-streaming-api/
type Channel struct {
	MediaID               interface{}
	MediaName             string
	MediaTypeID           interface{}
	MediaTypeName         interface{}
	Rating                interface{}
	ViewCounter           interface{}
	Description           interface{}
	CreationDate          interface{}
	LastWatchDate         interface{}
	StartDate             interface{}
	CatalogStartDate      interface{}
	PicURL                interface{}
	URL                   interface{}
	MediaWebLink          interface{}
	Duration              interface{}
	FileID                interface{}
	MediaDynamicData      interface{}
	SubDuration           interface{}
	SubFileFormat         interface{}
	SubFileID             interface{}
	SubURL                interface{}
	GeoBlock              interface{}
	TotalItems            interface{}
	like_counter          interface{}
	Tags                  interface{}
	AdvertisingParameters interface{}
	Files                 []struct {
		FileID          string
		URL             string
		Duration        string
		Format          string
		PreProvider     string
		PostProvider    string
		BreakProvider   string
		OverlayProvider string
		BreakPoints     string
		OverlayPoints   string
		Language        string
		IsDefaultLang   bool
		CoGuid          string
	}
	Pictures    interface{}
	ExternalIDs interface{}
}

// response of a licensed link
type LicensedLink struct {
	MainUrl string
	AltUrl  string
}

// global const to be used in code
const (
	GATEWAY              = "https://api-live.iptv.kabel-deutschland.de/v2_9/gateways/jsonpostgw.aspx"
	IOS_VERSION          = "8.1.2"
	APP_VERSION          = "1.2.3"
	METHOD_SIGNIN        = "SSOSignIn"
	METHOD_CHANNELLIST   = "GetChannelMediaList"
	METHOD_LICENSED_LINK = "GetLicensedLinks"
	INIT_OBJECT          = "eyJBcGlVc2VyIjoidHZwYXBpXzE4MSIsIlVESUQiOiJEMkFDNjMzQUZCNjQ0Q0YwQTY3NTA1MzcwNTc4Q0RFNSIsIkRvbWFpbklEIjozMTUzODQsIlNpdGVHdWlkIjo2Nzk4NzAsIlBsYXRmb3JtIjoiaVBhZCIsIkFwaVBhc3MiOiJhek5ETHpzbktER3RBclZXMlNIUiIsIkxvY2FsZSI6eyJMb2NhbGVEZXZpY2UiOiJudWxsIiwiTG9jYWxlVXNlclN0YXRlIjoiVW5rbm93biIsIkxvY2FsZUNvdW50cnkiOiJudWxsIiwiTG9jYWxlTGFuZ3VhZ2UiOiJudWxsIn19"
	CHANNEL_OBJECT       = "\"orderBy\":\"None\",\"pageSize\":1000,\"picSize\":\"100X100\",\"ChannelID\":340758"
	M3U_HEAD             = "#EXTM3U\n"
	M3U_LINE             = "#EXTINF:-1,%s\n%s\n"

	QUALITY_LOW    = "CCURstream564000.m3u8"
	QUALITY_MEDIUM = "CCURstream1064000.m3u8"
	QUALITY_HIGH   = "CCURstream1664000.m3u8"

	CACHE_FILE     = "playlist_%s.m3u"
	CACHE_LIFETIME = 86400
)
