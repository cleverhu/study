package lib

import (
	"errors"
	"strings"
)

var admins = []string{"root", "admin"}

func init() {
	E.AddFunction("splitMatchBySplit", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 2 {
			return splitMatchBySplit(arguments...), nil
		}
		return nil, errors.New("error match")
	})

	E.AddFunction("isSuperAdmin", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 1 {
			return isSuperAdmin(arguments...), nil
		}
		return nil, errors.New("error match")
	})
}

func splitMatchBySplit(arguments ...interface{}) bool {
	keys := strings.Split(arguments[1].(string), " ")
	for _, key := range keys {
		if key == arguments[0].(string) {
			return true
		}
	}
	return false
}

func isSuperAdmin(arguments ...interface{}) bool {
	user := arguments[0].(string)
	for _, key := range admins {
		if key == user {
			return true
		}
	}
	return false
}
