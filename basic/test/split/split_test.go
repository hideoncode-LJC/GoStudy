package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("i love you", "love")
	want := []string{"i ", " you"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v, got:%v", want, got)
	}
}
