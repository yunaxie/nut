package iterf

type Map interface {
	ContainsKey(interface{}) bool
	Get(interface{}) interface{}
	Put(interface{}, interface{}) interface{}
	PutIfAbsent(interface{}, interface{}) interface{}
	ComputeIfAbsent(interface{}, func(interface{}) interface{}) interface{}
	ComputeIfPresent(interface{}, func(interface{}, interface{}) interface{}) interface{}
	Remove(interface{}) interface{}
	PutAll(Map)

	Collection
	Iterator
}
