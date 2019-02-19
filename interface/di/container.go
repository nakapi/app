package di

type DiContainer interface {
	Register(...interface{}) error
	Resolve(interface{}) (interface{}, error)
}
