package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	type args struct {
		htmlBody   string
		rawBaseURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "absolute and relative URLs",
			args: args{
				htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			wantErr: false,
		}, {
			name: "single absolute anchor",
			args: args{
				htmlBody: `<a href="https://other.com/path/one"><span>Boot.dev</span></a>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    []string{"https://other.com/path/one"},
			wantErr: false,
		}, {
			name: "single relative anchor",
			args: args{
				htmlBody: `<a href="/articles"><span>Boot.dev</span></a>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    []string{"https://blog.boot.dev/articles"},
			wantErr: false,
		}, {
			name: "anchor without href",
			args: args{
				htmlBody: `<a><span>Boot.dev</span></a>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    nil,
			wantErr: false,
		}, {
			name: "anchor with fragment only",
			args: args{
				htmlBody: `<a href="#header"><span>Boot.dev</span></a>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    []string{"https://blog.boot.dev/#header"},
			wantErr: false,
		}, {
			name: "anchor with query only",
			args: args{
				htmlBody: `<a href="?s=search"><span>Boot.dev</span></a>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    []string{"https://blog.boot.dev/?s=search"},
			wantErr: false,
		}, {
			name: "relative without leading slash",
			args: args{
				htmlBody: `<a href="posts"><span>Boot.dev</span></a>
			`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want:    []string{"https://blog.boot.dev/posts"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLsFromHTML(tt.args.htmlBody, tt.args.rawBaseURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLsFromHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getURLsFromHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRelative(t *testing.T) {
	type args struct {
		href string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "relative url",
			args: args{href: "/relative/path"},
			want: true,
		}, {
			name: "absolute url",
			args: args{href: "https://domain.com/relative/path"},
			want: false,
		}, {
			name: "absolute url without scheme",
			args: args{href: "domain.com/path"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRelative(tt.args.href)
			if got != tt.want {
				t.Errorf("isRelative() = %v, want %v", got, tt.want)
			}
		})
	}
}
