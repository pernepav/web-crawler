package main

import (
	"reflect"
	"testing"
)

func Test_sortedByVisits(t *testing.T) {
	type args struct {
		pages map[string]int
	}
	tests := []struct {
		name string
		args args
		want []page
	}{
		{
			name: "pages sorted by visits desc",
			args: args{
				pages: map[string]int{
					"https://example.com/path/one":   2,
					"https://example.com/path/two":   10,
					"https://example.com/path/three": 1,
				},
			},
			want: []page{
				{url: "https://example.com/path/two", visits: 10},
				{url: "https://example.com/path/one", visits: 2},
				{url: "https://example.com/path/three", visits: 1},
			},
		}, {
			name: "empty pages",
			args: args{
				pages: map[string]int{},
			},
			want: []page{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedByVisits(tt.args.pages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedByVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}
