package compress

import (
	"testing"
)

func TestCompress(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "baidu.com",
			args: args{
				link: "https://qu.la/book/123_123/",
			},
			want: "/book/123_123/",
		},
		{
			name: "baidu.com",
			args: args{
				link: "http://qu.la/book/123_123/?id=123",
			},
			want: "/book/123_123/?id=123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compress(tt.args.link); got != tt.want {
				t.Errorf("Compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
