package bimap

import (
	"errors"
	"reflect"
)

type biMap struct {
	BiMap
	keys        map[interface{}]interface{}
	values      map[interface{}]interface{}
	orderedKeys []interface{}
}

func (bm *biMap) Set(key interface{}, value interface{}) error {
	if bm.keys[key] != nil {
		return errors.New("biMap can\\'t have duplicate key")
	}

	if bm.values[value] != nil {
		return errors.New("biMap can\\'t have duplicate value")
	}

	bm.keys[key] = value
	bm.values[value] = key
	bm.orderedKeys = append(bm.orderedKeys, key)

	return nil
}

func (bm *biMap) Get(key interface{}) interface{} {
	return bm.keys[key]
}

func (bm *biMap) GetKey(value interface{}) interface{} {
	return bm.values[value]
}

func (bm *biMap) getIndexOfKey(key interface{}) int {
	for i, k := range bm.orderedKeys {
		if key == k {
			return i
		}
	}

	return -1
}

func (bm *biMap) deletePair(key interface{}, value interface{}) {
	delete(bm.keys, key)
	delete(bm.values, value)
	keyIndex := bm.getIndexOfKey(key)
	bm.orderedKeys = append(bm.orderedKeys[:keyIndex], bm.orderedKeys[:keyIndex+1])
}

func (bm *biMap) DeleteValue(value interface{}) error {
	key := bm.GetKey(value)
	if key == nil {
		return errors.New("Key does not exist in biMap")
	}

	bm.deletePair(key, value)
	return nil
}

func (bm *biMap) Delete(key interface{}) error {
	value := bm.Get(key)
	if value == nil {
		return errors.New("Value does not exist in biMap")
	}

	bm.deletePair(key, value)
	return nil
}

func (bm *biMap) Size() int {
	return len(bm.keys)
}

func (bm *biMap) Left() map[interface{}]interface{} {
	return bm.keys
}

func (bm *biMap) Right() map[interface{}]interface{} {
	return bm.values
}

func (bm *biMap) Keys() []interface{} {
	return bm.orderedKeys
}

func (bm *biMap) Values() []interface{} {
	slice := make([]interface{}, 0, len(bm.orderedKeys))
	for _, key := range bm.orderedKeys {
		slice = append(slice, bm.Get(key))
	}
	return slice
}

func (bm *biMap) IsEqual(otherBm BiMap) bool {
	return reflect.DeepEqual(bm.values, otherBm.Right()) && reflect.DeepEqual(bm.keys, otherBm.Left())
}

// Tuple is a key-value pair used to initialized a new BiMap with values
type Tuple struct {
	key, value interface{}
}

// NewBiMap creates a new bi-directional map
func NewBiMap(initialValues ...Tuple) (BiMap, error) {
	keys := make(map[interface{}]interface{}, len(initialValues))
	values := make(map[interface{}]interface{}, len(initialValues))
	orderedKeys := make([]interface{}, 0, len(initialValues))

	for _, tuple := range initialValues {
		if keys[tuple.key] != nil {
			return nil, errors.New("Initial values contain duplicated keys")
		}

		if values[tuple.value] != nil {
			return nil, errors.New("Initial values contain duplicated values")
		}

		keys[tuple.key] = tuple.value
		values[tuple.value] = tuple.key
		orderedKeys = append(orderedKeys, tuple.key)
	}

	bm := &biMap{
		keys:        keys,
		values:      values,
		orderedKeys: orderedKeys,
	}
	return bm, nil
}
