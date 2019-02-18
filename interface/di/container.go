package di

type DiContainer interface {
	Resolve(interface{}) (interface{}, error)
}
