package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const FollowTableName = "follow"

type Follow struct {
	ID         int64          `json:"id"`
	UserId     int64          `json:"user_id"`
	FollowerId int64          `json:"follower_id"`
	CreatedAt  time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (Follow) TableName() string {
	return FollowTableName
}

func CreateFollow(ctx context.Context, follow *Follow) (id int64, err error) {
	// 如果关注关系已经存在，那么就不需要再次创建
	ok, err := CheckFollow(ctx, follow.UserId, follow.FollowerId)
	if err != nil {
		return -1, err
	}
	if ok {
		return 0, nil
	}
	err = DB.WithContext(ctx).Create(follow).Error
	if err != nil {
		return -1, err
	}
	return follow.ID, nil
}

func DeleteFollow(ctx context.Context, follow *Follow) (ok bool, err error) {
	err = DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", follow.UserId, follow.FollowerId).Delete(follow).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckFollow(ctx context.Context, user_id int64, follower_id int64) (ok bool, err error) {
	var db_follow Follow
	err = DB.WithContext(ctx).Where("user_id = ? and follower_id = ?", user_id, follower_id).Limit(1).Find(&db_follow).Error
	if err != nil {
		return false, err
	}
	return db_follow != Follow{}, nil
}

func GetFollowUserIdList(ctx context.Context, userId int64) (user_ids []int64, err error) {
	var follows []Follow
	err = DB.WithContext(ctx).Where("follower_id = ?", userId).Order("user_id asc").Find(&follows).Error
	if err != nil {
		return nil, err
	}
	for _, follow := range follows {
		user_ids = append(user_ids, follow.UserId)
	}
	return user_ids, nil
}

func GetFollowerUserIdList(ctx context.Context, userId int64) (user_ids []int64, err error) {
	var follows []Follow
	err = DB.WithContext(ctx).Where("user_id = ?", userId).Order("user_id asc").Find(&follows).Error
	if err != nil {
		return nil, err
	}
	for _, follow := range follows {
		user_ids = append(user_ids, follow.FollowerId)
	}
	return user_ids, nil
}

func GetFriendUserIdList(ctx context.Context, userId int64) (user_ids []int64, err error) {
	var follows []Follow
	// Friend 要求 A关注B 且 B关注A
	// 可以使用表连接查询，或者子查询
	// 这里使用子查询
	// user_id 表示被关注者A, follower_id 表示关注者B
	err = DB.WithContext(ctx).Where("user_id = ?", userId).Where("follower_id IN (SELECT user_id FROM follow WHERE follower_id = ?)", userId).Order("user_id asc").Find(&follows).Error
	if err != nil {
		return nil, err
	}
	for _, follow := range follows {
		user_ids = append(user_ids, follow.FollowerId)
	}
	return user_ids, nil
}

func GetFollowUserCount(ctx context.Context, userId int64) (count int64, err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("follower_id = ?", userId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func GetFollowerUserCount(ctx context.Context, userId int64) (count int64, err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func GetFriendUserCount(ctx context.Context, userId int64) (count int64, err error) {
	// Friend 要求 A关注B 且 B关注A
	// 可以使用表连接查询，或者子查询
	// 这里使用子查询
	// user_id 表示被关注者A, follower_id 表示关注者B
	err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ?", userId).Where("follower_id IN (SELECT user_id FROM follow WHERE follower_id = ?)", userId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}