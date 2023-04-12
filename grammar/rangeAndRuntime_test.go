package main

import (
	"reflect"
	"testing"
)

/*
	func Test_structArray_deepAdd(t *testing.T) {
		type fields struct {
			Value     int
			ValueArr  []int
			ArrayItem []structArrayItem
		}
		tests := []struct {
			name   string
			fields fields
		}{
			{
				name: "",
				fields: fields{
					Value:     1,
					ValueArr:  make([]int, 10),
					ArrayItem: []structArrayItem{{Value: 1, ValueArr: []int{1, 1, 1}}, {Value: 1, ValueArr: []int{1, 1, 1}}},
				},
			},
			// TODO: Add test cases.
		}

		for _, tt := range tests {
			s := &structArray{
				Value:     tt.fields.Value,
				ValueArr:  tt.fields.ValueArr,
				ArrayItem: tt.fields.ArrayItem,
			}
			s.deepAdd()
			fp.Fp(s)
		}

}
*/

func Test_addByRuntineAndRange(t *testing.T) {
	type args struct {
		s *structArray
	}
	tests := []struct {
		name string
		args args
		want *structArray
	}{
		{
			name: "",
			args: args{
				&structArray{
					Value:     1,
					ValueArr:  make([]int, 10),
					ArrayItem: []structArrayItem{{Value: 1, ValueArr: []int{1, 1, 1}}, {Value: 1, ValueArr: []int{2, 2, 2}}},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := addByRuntineAndRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. addByRuntineAndRange() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
