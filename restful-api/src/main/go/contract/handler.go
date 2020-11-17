package contract

type Handler interface {
	SaveOne(c Context)
	GetOne(c Context)
	UpdateOne(c Context)
	DeleteOne(c Context)
	PatchOne(c Context)
	GetAll(c Context)
	GetAllMethod(c Context)
	GetBasicMethod(c Context)
}
