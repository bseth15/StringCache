package collections

import (
	"fmt"
	"strings"
)

type stringCache struct {
	cache map[string]string
}

// NewStringCache initializes a StringCache struct with provided initial values.
// Returns a pointer to the new StringCache struct.
func NewStringCache(initial_values ...string) *stringCache {
	cache := &stringCache{cache: make(map[string]string, len(initial_values))}
	for _, strValue := range initial_values {
		cache.cache[strValue] = strValue
	}
	return cache
}

// Call processes an op code and dispatches the correct method handler.
// Returns result of operation or an error if an invalid op is called.
func (s *stringCache) Call(op string) (string, error) {
	command := strings.Split(op, "|")

	switch command[0] {
	case "Add":
		return fmt.Sprint(s.add(command[1])), nil
	case "Get":
		return fmt.Sprint(s.get(command[1])), nil
	case "Has":
		return fmt.Sprint(s.has(command[1])), nil
	case "Remove":
		return fmt.Sprint(s.remove(command[1])), nil
	case "Reset":
		return fmt.Sprint(s.reset()), nil
	default:
		return "", fmt.Errorf("invalid command: %q", command[0])
	}
}

// add inserts value into StringCache if it is not already present.
// Returns true if the value was added, otherwise returns false.
func (s *stringCache) add(value string) bool {
	if !s.has(value) {
		s.cache[value] = value
		return true
	}
	return false
}

// get retrieves the value if it is present in the StringCache.
// Returns a string if value is present, otherwise nil.
func (s *stringCache) get(value string) interface{} {
	if !s.has(value) {
		return nil
	}
	return value
}

// has checks to see if value is contained in the StringCache.
// Returns true if value has been cached, otherwise false.
func (s *stringCache) has(value string) bool {
	_, ok := s.cache[value]
	return ok
}

// remove deletes the value from the StringCache.
// Returns true if the value was removed, otherwise returns false.
func (s *stringCache) remove(value string) bool {
	if s.has(value) {
		delete(s.cache, value)
		return true
	}
	return false
}

// reset clears all values from the StringCache.
// Returns the number of strings cleared.
func (s *stringCache) reset() int {
	count := len(s.cache)
	clear(s.cache)
	return count
}
