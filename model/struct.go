package model

type User struct {
	ID            int64  `json:"id"`   //用户id
	Name          string `json:"name"` //用户名称
	Password      string
	FollowCount   int64 `json:"follow_count"`   //关注总数
	FollowerCount int64 `json:"follower_count"` //粉丝总数
	IsFollow      bool  `json:"is_follow"`      //true-已关注，false-未关注
	CreateTime    int64 //注册时间戳
}
type Video struct {
	ID           int64  `json:"id"`             //视频唯一标识
	Author       User   `json:"author"`         //视频作者信息
	PlayURL      string `json:"play_url"`       //视频播放地址
	CoverURL     string `json:"cover_url"`      //视频封面地址
	LikeCount    int64  `json:"favorite_count"` //视频的点赞总数
	CommentCount int64  `json:"comment_count"`  //视频的评论总数
	IsLike       bool   `json:"is_favorite"`    //true-已点赞，false-未点赞
	Title        string `json:"title"`          //视频标题
	PublishTime  int64  //发布时间戳
}
type FeedRequest struct {
	LatestTime int64  `json:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒
	Token      string `json:"token"`       //可选参数，登录用户设置
}
type FeedResponse struct {
	StatusCode int64  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMSG  string `json:"status_msg"`  //返回状态描述
	VideoList  Video  `json:"video_list"`  //视频列表
	NextTime   int64  `json:"next_time"`   //本次返回的视频中，发布最早的时间，作为下次请求时间的Lasted_time
}
type RegisterRequest struct {
	Username string `json:"username"` //注册用户名，最长32个字符
	Password string `json:"password"` //密码，最长32个字符
}
type RegisterResponse struct {
	StatusCode int64  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMSG  string `json:"status_msg"`  //返回状态描述
	UserID     int64  `json:"user_id"`     //用户id
	Token      string `json:"token"`       //用户鉴权token
}
type LoginRequest struct {
	Username string `json:"username"` //登录用户名
	Password string `json:"password"` //登录密码
}
type LoginResponse struct {
	StatusCode int64  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMSG  string `json:"status_msg"`  //返回状态描述
	UserID     int64  `json:"user_id"`     //用户id
	Token      string `json:"token"`       //用户鉴权token
}
type UserRequest struct {
	UserID int64  `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户鉴权token
}
type UserResponse struct {
	StatusCode int64  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMSG  string `json:"status_msg"`  //返回状态描述
	User       User   `json:"user"`        //用户信息
}
type PublishRequest struct {
	Token string `json:"token"` //用户鉴权token
	Data  []byte `json:"data"`  //视频数据
	Title string `json:"title"` //视频标题
}
type PublishResponse struct {
	StatusCode int64  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMSG  string `json:"status_msg"`  //返回状态描述
}
type PublishListRequest struct {
	UserID int64  `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户鉴权token
}
type PublishListResponse struct {
	StatusCode int64  `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMSG  string `json:"status_msg"`  //返回状态描述
	VideoList  Video  `json:"video_list"`  //用户发布的视频列表
}
