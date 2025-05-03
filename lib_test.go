package main

import (
	"reflect"
	"testing"
)

func Test_lcs_multiple_strings(t *testing.T) {
	type args struct {
		arr1 []string
		arr2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				arr1: []string{"a", "b", "c", "d"},
				arr2: []string{"a", "b", "c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "2",
			args: args{
				arr1: []string{"a", "b", "f", "d"},
				arr2: []string{"a", "g", "d", "f"},
			},
			want: []string{"a", "f"},
		},
		{
			name: "3",
			args: args{
				arr1: []string{"v", "b", "f", "d"},
				arr2: []string{"b", "g", "d", "f"},
			},
			want: []string{"b", "f"},
		},
		{
			name: "4",
			args: args{
				arr1: []string{"i", "b", "f", "d"},
				arr2: []string{"o", "b", "d", "f"},
			},
			want: []string{"b", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcs_multiple_strings(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("lcs_multiple_strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want Diffs
	}{
		{
			name: "1",
			args: args{
				s1: []string{"a", "b", "c", "d"},
				s2: []string{"a", "b", "c", "d"},
			},
			want: []*DiffResult{},
		},
		{
			name: "2",
			args: args{
				s1: []string{"a", "b", "f", "d"},
				s2: []string{"a", "g", "f", "d"},
			},
			want: []*DiffResult{
				{
					Operation: OpReplace,
					OldValues: []string{"b"},
					NewValues: []string{"g"},
				},
			},
		},
		{
			name: "3",
			args: args{
				s1: []string{"v", "b", "o", "d"},
				s2: []string{"b", "g", "d", "f"},
			},
			want: []*DiffResult{
				{
					Operation: OpDelete,
					OldValues: []string{"v"},
				},
				{
					Operation: OpReplace,
					OldValues: []string{"o"},
					NewValues: []string{"g"},
				},
				{
					Operation: OpInsert,
					NewValues: []string{"f"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"diff() = %v\n, want %v",
					got.String(),
					tt.want.String(),
				)
			}
		})
	}
}
