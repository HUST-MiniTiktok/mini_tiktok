namespace go favorite

include "common.thrift"

struct FavoriteActionRequest {
    1: string token     // 用户鉴权token
    2: i64 video_id     // 视频id
    3: i32 action_type  // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
}

struct FavoriteListRequest {
    1: i64 user_id  // 用户id
    2: string token // 用户鉴权token
}

struct FavoriteListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.Video> video_list                           // 用户点赞视频列表
}

struct GetVideoFavoriteInfoRequest {
    1: i64 user_id                                      // 用户id
    2: i64 video_id                                     // 视频id
}

struct GetVideoFavoriteInfoResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 favorite_count                               // 视频点赞总数
    4: bool is_favorite                                 // true-已点赞，false-未点赞
}

struct GetVideoFavoriteInfoListRequest {
    1: i64 user_id                                      // 用户id
    2: list<i64> video_id                               // 视频id
}

struct GetVideoFavoriteInfoListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<i64> favorite_count                         // 视频点赞总数
    4: list<bool> is_favorite                           // true-已点赞，false-未点赞
}

struct GetUserFavoriteInfoRequest {
    1: i64 user_id  // 用户id
}

struct GetUserFavoriteInfoResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 total_favorited                              // 获赞数量
    4: i64 favorite_count                               // 点赞数量
}

struct GetUserFavoriteInfoListRequest {
    1: list<i64> user_id_list  // 用户id
}

struct GetUserFavoriteInfoListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<i64> total_favorited_list                   // 获赞数量
    4: list<i64> favorite_count_list                    // 点赞数量
}

service FavoriteService {
    // 赞操作
    FavoriteActionResponse FavoriteAction (1: FavoriteActionRequest request) (api.post = "/douyin/favorite/action/")
    // 喜欢列表
    FavoriteListResponse FavoriteList (1: FavoriteListRequest request) (api.get = "/douyin/favorite/list/")
    // 获取视频点赞信息
    GetVideoFavoriteInfoResponse GetVideoFavoriteInfo (1: GetVideoFavoriteInfoRequest request)
    // 批量获取视频点赞信息
    GetVideoFavoriteInfoListResponse GetVideoFavoriteInfoList (1: GetVideoFavoriteInfoListRequest request)
    // 获取用户点赞信息
    GetUserFavoriteInfoResponse GetUserFavoriteInfo (1: GetUserFavoriteInfoRequest request)
    // 批量获取用户点赞信息
    GetUserFavoriteInfoListResponse GetUserFavoriteInfoList (1: GetUserFavoriteInfoListRequest request)
}

