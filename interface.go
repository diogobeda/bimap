package bimap

// BiMap is a bi-directional map
type BiMap interface {
	// Sets a key-value pair on the map. Returns an error if key or value is duplicate
	Set(key interface{}, value interface{}) error
	// Gets a value from a key
	Get(key interface{}) interface{}
	// Gets a key from a value
	GetKey(value interface{}) interface{}
	// Deletes a key-value pair from a key. Returns an error if provided argument is not a key
	Delete(key interface{}) error
	// Deletes a key-value pair from a value. Returns an error if provided argument is not a value
	DeleteValue(value interface{}) error
	// Returns the size of the map
	Size() int
	// Returns the "key: value" mapping of the BiMap
	Left() map[interface{}]interface{}
	// Returns the "value: key" mapping of the BiMap
	Right() map[interface{}]interface{}
	// Returns a slice with all the BiMap keys
	Keys() []interface{}
	// Returns a slice with all the BiMap values
	Values() []interface{}
	// Checks if a BiMap is equal to another
	IsEqual(BiMap) bool
}
