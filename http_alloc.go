package main

import (
	"net/http"
	"strconv"
	"fmt"
)

func main() {
	http.HandleFunc("/test/alloc", TestAlloc)
	http.ListenAndServe(":18080", nil)
}

type MomentsAdminRsp struct {
	Id              int64                `json:"mid"`
	Uid             string               `json:"uid"`
	Nickname        string               `json:"nickname"`
	Avatar          string               `json:"avatar"`
	Text            string               `json:"text"`
	TextTp          int8                 `json:"type"`
	FileCnt         int8                 `json:"cnt"`
	State           int8                 `json:"state"`
	StateDesc       string               `json:"state_desc"`
	Fids            string               `json:"fids"`
	PostTime        int64                `json:"post_time"`
	Samples         string               `json:"samples"`
	MediaTime       int                  `json:"media_time"`
	ForwardCnt      int64                `json:"forward_cnt"`
	ClientTime      string               `json:"client_time"`
	CommentCnt      int64                `json:"comment_cnt"`
	LikeCnt         int64                `json:"like_cnt"`
	LikeFlag        int8                 `json:"like_flag"`
	Fee             string               `json:"fee"`
	UidFriends      []string             `json:"-"`
	LocationGPS     string               `json:"location_gps"`
	LocationName    string               `json:"location_name"`
	ForwardInfo     string               `json:"-"`
	ForwardMid      int64                `json:"forward_mid"`
	ForwardText     string               `json:"forward_text"`
	ForwardUid      string               `json:"forward_uid"`
	ForwardNickname string               `json:"forward_nickname"`
	ForwardAvatar   string               `json:"forward_avatar"`
	Flag            int8                 `json:"-"`
	ReportCnt       int64                `json:"report_cnt"`
}
1
func TestAlloc(w http.ResponseWriter, r *http.Request) {
	var ss []*MomentsAdminRsp
	var i int
	iStr := r.FormValue("i")
	if iStr != "" {
		i, _ = strconv.Atoi(iStr)
	}
	for a := 0; a < i; a++{
		ss = append(ss, new(MomentsAdminRsp))
	}
	w.Write([]byte(fmt.Sprintf("alloc object num: [%d]", len(ss))))
}
