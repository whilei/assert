package assert

import (
	"testing"
)

type MyStruct struct {
	Sub *MyStruct
}

func TestEqual(t *testing.T) {
	Equal(t, "foo", "foo")
	Equal(t, true, true)

	myStructA := MyStruct{}
	myStructB := MyStruct{}
	Equal(t, myStructA, myStructB)

	// Equal(t, "foo", "bar", "this should blow up")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "foo", "bar", "msg!")
	NotEqual(t, nil, false)

	myStructA := MyStruct{}
	myStructB := MyStruct{&myStructA}
	NotEqual(t, myStructA, myStructB)
	NotEqual(t, &myStructA, myStructA)

	NotEqual(t, "foo", "foo", "this should blow up")
}

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}

func TestNil(t *testing.T) {
	Nil(t, nil)

	var nilChan chan int
	Nil(t, nilChan)

	var nilFunc func(int) int
	Nil(t, nilFunc)

	var nilInterface interface{}
	Nil(t, nilInterface)

	var nilMap map[string]string
	Nil(t, nilMap)

	var myStruct MyStruct
	Nil(t, myStruct.Sub) // nil pointer

	var nilSlice []string
	Nil(t, nilSlice)

	// Nil(t, "foo", "this should blow up")
}

func TestNotNil(t *testing.T) {
	NotNil(t, "foo")

	myStruct := MyStruct{}
	NotNil(t, myStruct)
	NotNil(t, &myStruct)

	// NotNil(t, nil, "this should blow up")
	// var myNilStruct MyStruct
	// NotNil(t, myNilStruct, "this should blow up")
}
