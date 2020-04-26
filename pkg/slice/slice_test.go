package slice

import (
	"reflect"
	"testing"
)

func TestFilterInt(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	type args struct {
		xs []int
		f  func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Should return same value if all condition is fullfiled", args{xs, func(v int) bool { return true }}, xs},
		{"Should return empty slice", args{xs, func(v int) bool { return false }}, []int{}},
		{"Should return expected slice ", args{xs, func(v int) bool { return v > 3 }}, []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInt(tt.args.xs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
