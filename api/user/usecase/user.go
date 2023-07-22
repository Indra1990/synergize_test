package usecase

type ServiceUser interface {
	UserList(cmd UserQueryParam) (dto []*UserListResponse, err error)
	UserDetail(userId uint) (dto *UserListResponse, err error)
}
