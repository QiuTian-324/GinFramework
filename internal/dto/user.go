package dto

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Bio       string `json:"bio"`
	Online    int    `json:"online"`
}

// 添加好友请求
type AddFriendRequest struct {
	FriendID string `json:"friend_id"`
}

type AddFriendResponse struct {
	FriendID uint64 `json:"friend_id"`
}

// 好友列表响应
type FriendListResponse struct {
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Online    int    `json:"online"`
}

type GithubAuthURLRequest struct {
	State string `json:"state"`
}

// github授权
type GithubAuthURLResponse struct {
	URL string `json:"url"`
}

// github 登录请求
type GithubLoginRequest struct {
	Code string `json:"code"`
}

// github 登录响应
type GithubLoginResponse struct {
	Token string `json:"token"`
}

// User 结构体用于存储从 GitHub 获取的用户信息
type GithubUser struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	Avatar  string `json:"avatar_url"`
	URL     string `json:"html_url"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Blog    string `json:"blog"`
	Email   string `json:"email"`
}

// Token 结构体用于存储 OAuth2 令牌
type GithubToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// 查询好友响应
type QueryUserResponse struct {
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
}
