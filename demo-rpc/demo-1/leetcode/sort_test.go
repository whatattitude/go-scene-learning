package leetcode

import (
	"reflect"
	"testing"
)

func Test_createRandomArr(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name          string
		args          args
		wantArrRandom []int
		wantArrBase   []int
	}{
		// TODO: Add test cases.
		{
			name: "createArr",
			args: args{
				len: 10,
			},
			wantArrRandom: nil,
			wantArrBase:   nil,
		},
	}
	for _, tt := range tests {
		gotArrRandom, gotArrBase := createRandomArr(tt.args.len)
		if !reflect.DeepEqual(gotArrRandom, tt.wantArrRandom) {
			t.Errorf("%q. createRandomArr() gotArrRandom = %v, want %v", tt.name, gotArrRandom, tt.wantArrRandom)
		}
		if !reflect.DeepEqual(gotArrBase, tt.wantArrBase) {
			t.Errorf("%q. createRandomArr() gotArrBase = %v, want %v", tt.name, gotArrBase, tt.wantArrBase)
		}
	}
}

func Test_sortString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
		// TODO: Add test cases.
	}
	for range tests {
		sortString()
	}
}
