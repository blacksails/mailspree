package mailspree

// NewSession is what the api receives in order to login a user
type NewSession struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// Session is returned on a successful login
type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
