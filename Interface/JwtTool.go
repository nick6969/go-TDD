package Interface

type JwtTool interface {
	GenerateUserToken(id uint) (token string, err error)
	VerifyUserToken(token string) (id uint, err error)
}
