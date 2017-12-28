package common

import (
	"reflect"
	"strings"
)

// Event represent internal or protocol
type Event interface{}

// Name return event name
func Name(e Event) string {
	typeName := reflect.TypeOf(e).String()
	names := strings.Split(typeName, ".")
	return names[len(names)-1]
}
