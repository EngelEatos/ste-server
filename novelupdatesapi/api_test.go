package novelupdatesapi

import (
	"reflect"
	"testing"
)

func TestGetBool(t *testing.T) {
	var cases = map[string]bool{
		"No":  false,
		"Yes": true,
		"no":  false,
		"yes": true,
		"nO":  false,
		"yEs": true,
	}
	for o, result := range cases {
		b, _ := getBool(o)
		if b != result {
			t.Errorf("expected %t, but got %t", result, b)
		}
	}
}

func TestParseIdx(t *testing.T) {
	cases := map[string][]string{
		"v4c6": []string{"6"}, "v7c6 part1": []string{"6", "1"}, "c1.2": []string{"1", "2"}, "v6c29 part9": []string{"29", "9"}, "c389 part2": []string{"389", "2"}, "c99": []string{"99"}, "prologue": []string{"0"}, "c991-994": []string{"991", "994"},
		"ex3 c7": []string{"7"}, "ex1 c3 part5": []string{"3", "5"},
	}

	for tcase, solution := range cases {
		i, a := parseIdx(tcase)
		if i == -1 {
			t.Errorf("error %d at case %s", i, tcase)
		}
		if !reflect.DeepEqual(a, solution) {
			t.Errorf("expected %s, got %s", solution, a)
		}
	}
}
