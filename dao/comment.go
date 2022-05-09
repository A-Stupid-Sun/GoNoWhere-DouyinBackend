package dao

type commentDAO struct{}

var (
	CommentDAO   = &commentDAO{}
	FavoriteDAO  = &favoriteDAO{}
	FollowDAO    = &followDAO{}
	UserDAO      = &userDAO{}
	UserLoginDAO = &userLoginDAO{}
	VideoDAO     = &videoDAO{}
)
