namespace go publish

include "common.thrift"

struct PublishActionRequest {
    1: string token // 用户鉴权token
    2: binary data  // 视频数据
    3: string title // 视频标题
}

struct PublishActionResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
}

struct PublishListRequest {
    1: i64 user_id  // 用户id
    2: string token // 用户鉴权token
}

struct PublishListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.Video> video_list                           // 用户发布的视频列表
}

struct GetVideoByIdRequest {
    1: i64 id               // 视频唯一标识
    2: string token         // 用户鉴权token
}

struct GetVideoByIdResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: optional common.Video video                      // 视频
}

struct GetVideoByIdListRequest {
    1: list<i64> id                 // 视频唯一标识
    2: string token                 // 用户鉴权token
}

struct GetVideoByIdListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.Video> video_list                           // 视频列表
}

struct GetPublishInfoByUserIdRequest {
    1: i64 user_id  // 用户id
}

struct GetPublishInfoByUserIdResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 work_count                                   // 视频数量
    4: list<i64> video_ids                              // 视频id列表
}

struct GetPublishWorkCountListByUserIdRequest {
    1: list<i64> user_id_list  // 用户id
}

struct GetPublishWorkCountListByUserIdResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<i64> work_count                                   // 视频数量
}


service PublishService {
    // 投稿接口
    PublishActionResponse PublishAction (1: PublishActionRequest request) (api.post = "/douyin/publish/action/")
    // 发布列表
    PublishListResponse PublishList (1: PublishListRequest request) (api.get = "/douyin/publish/list/")
    // 查询视频
    GetVideoByIdResponse GetVideoById (1: GetVideoByIdRequest request)
    // 批量查询视频
    GetVideoByIdListResponse GetVideoByIdList (1: GetVideoByIdListRequest request)
    // 查询发布信息
    GetPublishInfoByUserIdResponse GetPublishInfoByUserId (1: GetPublishInfoByUserIdRequest request)
    // 批量查询用户发布信息(仅作品数)
    GetPublishWorkCountListByUserIdResponse GetPublishWorkCountListByUserId (1: GetPublishWorkCountListByUserIdRequest request)
}