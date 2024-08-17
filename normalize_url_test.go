package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "remove scheme",
			args: args{url: "https://blog.boot.dev/path"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove trailing slash",
			args: args{url: "https://blog.boot.dev/path/"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove query params",
			args: args{url: "https://blog.boot.dev/path?q=value&sort=asc"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove query params and trailing slash",
			args: args{url: "https://blog.boot.dev/path/?q=value&sort=asc"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove fragments",
			args: args{url: "https://blog.boot.dev/path#hash"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove fragments and trailing slash",
			args: args{url: "https://blog.boot.dev/path/#hash"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove username and password",
			args: args{url: "https://username:password@blog.boot.dev/path"},
			want: "blog.boot.dev/path",
		}, {
			name: "remove port",
			args: args{url: "https://blog.boot.dev:443/path"},
			want: "blog.boot.dev/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeURL(tt.args.url); got != tt.want {
				t.Errorf("normalizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
