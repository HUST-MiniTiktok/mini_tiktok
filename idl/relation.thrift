namespace go relation

include "common.thrift"

struct RelationActionRequest {
    1: string token     // 用户token
    2: i64 to_user_id   // 对方用户id
    3: i32 action_type  // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
}

struct RelationFollowListRequest {
    1: i64 user_id      // 用户id
    2: string token     // 用户鉴权token
}

struct RelationFollowListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.User> user_list                             // 用户列表
}

struct RelationFollowerListRequest {
    1: i64 user_id      // 用户id
    2: string token     // 用户鉴权token
}

struct RelationFollowerListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.User> user_list                             // 用户列表
}

struct RelationFriendListRequest {
    1: i64 user_id      // 用户id
    2: string token     // 用户鉴权token
}

struct RelationFriendListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<common.FriendUser> user_list                             // 用户列表
}

struct GetFollowInfoRequest {
    1: string token     // 用户token
    2: i64 to_user_id   // 对方用户id
}

struct GetFollowInfoResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: i64 follow_count                                 // 关注总数
    4: i64 follower_count                               // 粉丝总数
    5: bool is_follow                                   // 是否关注
}

struct GetFollowInfoListRequest {
    1: list<string> token_list     // 用户token
    2: list<i64> to_user_id_list   // 对方用户id
}

struct GetFollowInfoListResponse {
    1: i32 status_code (go.tag="json:\"status_code\"")  // 状态码，0-成功，其他值-失败
    2: optional string status_msg                       // 返回状态描述
    3: list<i64> follow_count_list                                 // 关注总数
    4: list<i64> follower_count_list                               // 粉丝总数
    5: list<bool> is_follow_list                                   // 是否关注
}

service RelationService {
    // 关注操作
    RelationActionResponse RelationAction(RelationActionRequest request) (api.post = "/douyin/relation/action/")
    // 关注列表
    RelationFollowListResponse RelationFollowList(RelationFollowListRequest request) (api.get = "/douyin/relation/follow/list/")
    // 粉丝列表
    RelationFollowerListResponse RelationFollowerList(RelationFollowerListRequest request) (api.get = "/douyin/relation/follower/list/")
    // 好友列表
    RelationFriendListResponse RelationFriendList(RelationFriendListRequest request) (api.get = "/douyin/relation/friend/list/")
    // 获取关注信息
    GetFollowInfoResponse GetFollowInfo(GetFollowInfoRequest request)
    // 批量根据用户token和目标用户ID列表获取关注信息
    GetFollowInfoListResponse GetFollowInfoList(GetFollowInfolistRequest request)
}