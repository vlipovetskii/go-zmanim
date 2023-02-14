package assert

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"log"
	"reflect"
	"testing"
)

/*
<<
Are libraries for testing incorporated in final build ?
So no, dependencies only referred to from tests are not included in the executable binary.

If it isn't imported from your non-test code then it won't be built into the final binary.
>>
*/

func Equal(t *testing.T, tag string, want any, got any) {
	if !reflect.DeepEqual(got, want) {
		log.Println(fmt.Sprintf("%s got = %v, want %v", tag, got, want))
		helper.TraceStack()
		t.FailNow()
	}
}

func False(t *testing.T, tag string, condition bool) {
	Equal(t, tag, false, condition)
}

func True(t *testing.T, tag string, condition bool) {
	Equal(t, tag, true, condition)
}

// Raises --- How to test panics? https://stackoverflow.com/questions/31595791/how-to-test-panics
func Raises(t *testing.T, tag string) func() {
	return func() {
		if r := recover(); r == nil {
			t.Errorf("%s: The code did not panic", tag)
		}
	}
}
