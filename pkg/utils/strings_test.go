package utils

import (
	"reflect"
	"testing"
)

func TestSplitMultiString(t *testing.T) {
	testCases := map[string]struct {
		s        string
		wantList []string
	}{
		"test single case 1": {
			s:        "a",
			wantList: []string{"a"},
		},
		"test single case 2": {
			s:        "a,",
			wantList: []string{"a"},
		},
		"test single case 3": {
			s:        "a,,",
			wantList: []string{"a"},
		},
		"test multi string case 1": {
			s:        "a,b,c",
			wantList: []string{"a", "b", "c"},
		},
		"test multi string case 2": {
			s:        "a,b,c,",
			wantList: []string{"a", "b", "c"},
		},
	}

	for k, item := range testCases {
		got := SplitMultiString(item.s)
		if !reflect.DeepEqual(got, item.wantList) {
			t.Errorf("%s failure, want:%v,got:%v", k, item.wantList, got)
		}
	}

}

/*
func TestGetNormalizedGroupNameVersion(t *testing.T) {
	g := "app"
	n := "label-x"
	v := "1.0.0"
	want := fmt.Sprintf("%s-%s-%s", g, n, v)
	got := GetNormalizedGroupNameVersion(g, n, v)
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
*/

func TestShuffleStrings(t *testing.T) {
	check := []string{
		"a",
		"b",
		"c",
		"d",
	}

	dummy := []string{
		"a",
		"b",
		"c",
		"d",
	}

	ShuffleStrings(check)

	if reflect.DeepEqual(check, dummy) {
		t.Errorf("check failure, got same result, want shuffle")
	}

}
