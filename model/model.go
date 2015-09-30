package model

type Model interface {
	Save(object interface{}) bool
	Find(key string, value string) interface{}
	Update(object interface{}) bool
	Delete(object interface{}) bool
}
