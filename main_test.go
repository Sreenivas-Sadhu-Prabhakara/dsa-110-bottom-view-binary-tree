package main

import (
	"reflect"
	"testing"
)

func TestBottomView(t *testing.T) {
	// Columns: 20:0, 8:-1, 22:1, 5:-2, 3:0, 25:2, 10:-1, 14:1.
	// Deepest per column: -2:5, -1:10, 0:3, 1:14 (deeper than 22), 2:25.
	root := tn(20,
		tn(8, tn(5, nil, nil), tn(3, tn(10, nil, nil), tn(14, nil, nil))),
		tn(22, nil, tn(25, nil, nil)))
	want := []int{5, 10, 3, 14, 25}
	if got := BottomView(root); !reflect.DeepEqual(got, want) {
		t.Errorf("BottomView = %v, want %v", got, want)
	}

	if got := BottomView(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("BottomView(nil) = %v", got)
	}
	if got := BottomView(tn(1, nil, nil)); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("single: %v", got)
	}
	// Left-leaning chain: every node its own column.
	chain := tn(3, tn(2, tn(1, nil, nil), nil), nil)
	if got := BottomView(chain); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("chain: %v", got)
	}
}
