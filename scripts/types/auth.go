package types

type ReqAuthToken struct {
	ID       string `query:"id" validate:"required"`
	Password string `query:"password" validate:"required"`
}
