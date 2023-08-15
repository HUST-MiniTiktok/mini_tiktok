package utils

import (
	hertz_user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/user"
	kitex_user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user"
)

func ToHertzUser(user *kitex_user.User) *hertz_user.User {
	return &hertz_user.User{
		ID:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}

func ToKitexUserRequest(req *hertz_user.UserRequest) *kitex_user.UserRequest {
	return &kitex_user.UserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
}

func ToHertzUserResponse(resp *kitex_user.UserResponse) *hertz_user.UserResponse {
	return &hertz_user.UserResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		User:       ToHertzUser(resp.User),
	}
}

func ToKitexUserRegisterRequest(req *hertz_user.UserRegisterRequest) *kitex_user.UserRegisterRequest {
	return &kitex_user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func ToHertzUserRegisterResponse(resp *kitex_user.UserRegisterResponse) *hertz_user.UserRegisterResponse {
	return &hertz_user.UserRegisterResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserID:     resp.UserId,
		Token:      resp.Token,
	}
}

func ToKitexUserLoginRequest(req *hertz_user.UserLoginRequest) *kitex_user.UserLoginRequest {
	return &kitex_user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func ToHertzUserLoginResponse(resp *kitex_user.UserLoginResponse) *hertz_user.UserLoginResponse {
	return &hertz_user.UserLoginResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		UserID:     resp.UserId,
		Token:      resp.Token,
	}
}
