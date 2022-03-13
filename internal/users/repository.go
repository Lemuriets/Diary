package users

type Repository interface {
	Get()
	Create()
	Update()
	Delete()
}
