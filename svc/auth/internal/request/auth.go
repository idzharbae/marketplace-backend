package request

type GetToken struct {
	UsernameOrEmail string
	Password        string
}

type RefreshToken struct {
	CurrentToken string
}
