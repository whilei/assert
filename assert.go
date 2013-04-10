package assert

import (
	"fmt"
	"github.com/kr/pretty"
	"reflect"
	"runtime"
	"testing"
)

var errorPrefix = "\U0001F4A9 "

// -- Assertion handlers

func assert(t *testing.T, success bool, f func(), callDepth int) {
	if !success {
		_, file, line, _ := runtime.Caller(callDepth + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}

func equal(t *testing.T, expected, got interface{}, callDepth int, args ...interface{}) {
	fn := func() {
		for _, desc := range pretty.Diff(expected, got) {
			t.Error(errorPrefix, desc)
		}
		if len(args) > 0 {
			t.Error(errorPrefix, "-", fmt.Sprint(args...))
		}
	}
	assert(t, isEqual(expected, got), fn, callDepth+1)
}

func notEqual(t *testing.T, expected, got interface{}, callDepth int, args ...interface{}) {
	fn := func() {
		t.Errorf("%s Unexpected: %#v", errorPrefix, got)
		if len(args) > 0 {
			t.Error(errorPrefix, "-", fmt.Sprint(args...))
		}
	}
	assert(t, !isEqual(expected, got), fn, callDepth+1)
}

// -- Matching

func isEqual(expected, got interface{}) bool {
	if expected == nil {
		return isNil(got)
	} else {
		return reflect.DeepEqual(expected, got)
	}
}

func isNil(got interface{}) bool {
	if got == nil {
		return true
	}
	value := reflect.ValueOf(got)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

// -- Public API

func Equal(t *testing.T, expected, got interface{}, args ...interface{}) {
	equal(t, expected, got, 1, args...)
}

func NotEqual(t *testing.T, expected, got interface{}, args ...interface{}) {
	notEqual(t, expected, got, 1, args...)
}

func True(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, true, got, 1, args...)
}

func False(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, false, got, 1, args...)
}

func Nil(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, nil, got, 1, args...)
}

func NotNil(t *testing.T, got interface{}, args ...interface{}) {
	notEqual(t, nil, got, 1, args...)
}
