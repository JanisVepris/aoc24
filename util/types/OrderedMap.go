package types

// A map that maintains the order of its keys
// This is useful for iterating over a map in a predictable order
type OrderedMap[K comparable, T any] struct {
	keys   []K     // Ordered map keys
	values map[K]T // Map values
}

// Initialize a new OrderedMap
func NewOrderedMap[K comparable, T any]() *OrderedMap[K, T] {
	return &OrderedMap[K, T]{
		keys:   []K{},
		values: make(map[K]T, 0),
	}
}

// Set a key/value pair in the map
func (om *OrderedMap[K, T]) Set(key K, value T) {
	if _, ok := om.values[key]; ok {
		om.values[key] = value
		return
	}

	om.values[key] = value
	om.keys = append(om.keys, key)
}

// Get a value from the map by key
// Returns the value and a boolean indicating whether the key was found
// If the key was not found, the value will be the zero value for the type
// and the boolean will be false
func (om *OrderedMap[K, T]) Get(key K) (T, bool) {
	value, ok := om.values[key]
	return value, ok
}

// Returns a slice of the map's keys in the order they were added
func (om *OrderedMap[K, T]) Keys() []K {
	return om.keys
}

// Returns a slice of the map's values in the order they were added
func (om *OrderedMap[K, T]) Values() []T {
	values := make([]T, len(om.keys))

	for i, key := range om.keys {
		values[i] = om.values[key]
	}

	return values
}

// Returns the number of items in the map
func (om *OrderedMap[K, T]) Len() int {
	return len(om.keys)
}

// Delete a key/value pair from the map
// If the key does not exist, this is a no-op
func (om *OrderedMap[K, T]) Delete(key K) {
	delete(om.values, key)

	for i, k := range om.keys {
		if k == key {
			om.keys = append(om.keys[:i], om.keys[i+1:]...)

			break
		}
	}
}
