package bimap

import (
	"reflect"
	"testing"
)

func TestBiMap(t *testing.T) {
	initialValues := []Tuple{
		{"a", 1},
		{"b", 2},
		{true, false},
	}
	expectedKeysMap := map[interface{}]interface{}{
		"a":  1,
		"b":  2,
		true: false,
	}
	expectedValuesMap := map[interface{}]interface{}{
		1:     "a",
		2:     "b",
		false: true,
	}
	bm, _ := NewBiMap(initialValues...)

	keys := bm.Left()
	if !reflect.DeepEqual(keys, expectedKeysMap) {
		t.Fatalf("Internal keys map %v is different from expected %v", keys, expectedKeysMap)
	}

	values := bm.Right()
	if !reflect.DeepEqual(values, expectedValuesMap) {
		t.Fatalf("Internal values map %v is different from expected %v", values, expectedValuesMap)
	}
}

func TestBiMapEmpty(t *testing.T) {
	expectedMaps := map[interface{}]interface{}{}
	bm, _ := NewBiMap()

	keys := bm.Left()
	if !reflect.DeepEqual(keys, expectedMaps) {
		t.Fatalf("Internal keys map %v is different from expected %v", keys, expectedMaps)
	}

	values := bm.Right()
	if !reflect.DeepEqual(values, expectedMaps) {
		t.Fatalf("Internal values map %v is different from expected %v", values, expectedMaps)
	}
}

func TestBiMapDuplicate(t *testing.T) {
	testCases := []struct {
		input  []Tuple
		errMsg string
	}{
		{
			input: []Tuple{
				{"a", 1},
				{"a", 2},
			},
			errMsg: "Initial values contain duplicated keys",
		},
		{
			input: []Tuple{
				{"a", 1},
				{"b", 1},
			},
			errMsg: "Initial values contain duplicated values",
		},
	}

	for _, testCase := range testCases {
		_, err := NewBiMap(testCase.input...)
		if err.Error() != testCase.errMsg {
			t.Fatalf("Duplicate error message \"%v\" does not match expected \"%v\"", err.Error(), testCase.errMsg)
		}
	}
}

func TestBiMap_set(t *testing.T) {
	testCases := []struct {
		key          interface{}
		value        interface{}
		nonDuplicate interface{}
	}{
		{"a", 1, 2},
		{"a", "b", 2},
		{1, 2, 10},
		{true, false, 10},
	}

	for _, testCase := range testCases {
		bm, _ := NewBiMap()
		bm.Set(testCase.key, testCase.value)

		mappedValue := bm.Get(testCase.key)
		if mappedValue != testCase.value {
			t.Fatalf("Mapped value %v for key %v doesn't mactch expected %v", mappedValue, testCase.key, testCase.value)
		}

		mappedKey := bm.GetKey(testCase.value)
		if mappedKey != testCase.key {
			t.Fatalf("Mapped key %v for value %v doesn't mactch expected %v", mappedKey, testCase.value, testCase.key)
		}

		duplicateKeyErr := bm.Set(testCase.key, testCase.nonDuplicate)
		if duplicateKeyErr == nil {
			t.Fatalf("BiMap cannot have duplicate keys")
		}

		duplicateValueErr := bm.Set(testCase.nonDuplicate, testCase.value)
		if duplicateValueErr == nil {
			t.Fatalf("BiMap cannot have duplicate values")
		}
	}
}

func TestBiMap_get(t *testing.T) {
	testCases := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"a", "b"},
		{1, 2},
		{true, false},
	}

	for _, testCase := range testCases {
		bm, _ := NewBiMap()
		bm.Set(testCase.key, testCase.value)

		if bm.Get(testCase.key) != testCase.value {
			t.Fatalf("Correct value is not returned from key")
		}
	}
}

func TestBiMap_getKey(t *testing.T) {
	testCases := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	for _, testCase := range testCases {
		bm, _ := NewBiMap()
		bm.Set(testCase.key, testCase.value)

		if bm.GetKey(testCase.value) != testCase.key {
			t.Fatalf("Correct key is not returned from value")
		}
	}
}

func TestBiMap_delete(t *testing.T) {
	testCases := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	for _, testCase := range testCases {
		bm, _ := NewBiMap()
		bm.Set(testCase.key, testCase.value)
		bm.DeleteValue(testCase.value)

		if bm.Get(testCase.key) != nil {
			t.Fatalf("Item not deleted from key")
		}
	}
}

func TestBiMap_deleteKey(t *testing.T) {
	testCases := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	for _, testCase := range testCases {
		bm, _ := NewBiMap()
		bm.Set(testCase.key, testCase.value)
		bm.Delete(testCase.key)

		if bm.Get(testCase.key) != nil {
			t.Fatalf("Item not deleted from key")
		}
	}
}

func TestBiMap_size(t *testing.T) {
	valuesToInsert := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	bm, _ := NewBiMap()
	for _, val := range valuesToInsert {
		bm.Set(val.key, val.value)
	}

	if bm.Size() != 4 {
		t.Fatalf("BiMap size is not being calculated correctly")
	}
}

func TestBiMap_left(t *testing.T) {
	valuesToInsert := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	bm, _ := NewBiMap()
	for _, val := range valuesToInsert {
		bm.Set(val.key, val.value)
	}

	expected := map[interface{}]interface{}{
		"a":  1,
		"b":  2,
		"c":  3,
		true: false,
	}

	if !reflect.DeepEqual(bm.Left(), expected) {
		t.Fatalf("Incorrect left value")
	}
}

func TestBiMap_right(t *testing.T) {
	valuesToInsert := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	bm, _ := NewBiMap()
	for _, val := range valuesToInsert {
		bm.Set(val.key, val.value)
	}

	expected := map[interface{}]interface{}{
		1:     "a",
		2:     "b",
		3:     "c",
		false: true,
	}

	if !reflect.DeepEqual(bm.Right(), expected) {
		t.Fatalf("Incorrect right value")
	}
}

func TestBiMap_keys(t *testing.T) {
	valuesToInsert := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	bm, _ := NewBiMap()
	for _, val := range valuesToInsert {
		bm.Set(val.key, val.value)
	}

	expected := []interface{}{
		"a", "b", "c", true,
	}

	returned := bm.Keys()
	if !reflect.DeepEqual(returned, expected) {
		t.Fatalf("Returned slice %v is different from expected %v", returned, expected)
	}
}

func TestBiMap_values(t *testing.T) {
	valuesToInsert := []struct {
		key   interface{}
		value interface{}
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{true, false},
	}

	bm, _ := NewBiMap()
	for _, val := range valuesToInsert {
		bm.Set(val.key, val.value)
	}

	expected := []interface{}{
		1, 2, 3, false,
	}

	returned := bm.Values()
	if !reflect.DeepEqual(returned, expected) {
		t.Fatalf("Returned slice %v is different from expected %v", returned, expected)
	}
}

func TestBiMap_isEqual(t *testing.T) {
	testCases := []struct {
		firstBmValues  []Tuple
		secondBmValues []Tuple
		expected       bool
	}{
		{
			firstBmValues: []Tuple{
				{"a", 1},
				{"b", 2},
			},
			secondBmValues: []Tuple{
				{"a", 1},
				{"b", 2},
			},
			expected: true,
		},
		{
			firstBmValues: []Tuple{
				{"a", 1},
				{"b", 2},
			},
			secondBmValues: []Tuple{
				{"a", 1},
				{"c", 3},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		firstBm, _ := NewBiMap(testCase.firstBmValues...)
		secondBm, _ := NewBiMap(testCase.secondBmValues...)

		if firstBm.IsEqual(secondBm) != testCase.expected {
			t.Fatalf("IsEqual result %v is different from %v for maps %v and %v", firstBm.IsEqual(secondBm), testCase.expected, firstBm, secondBm)
		}
	}
}
