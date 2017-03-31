package encode

import "testing"

func TestMD5(t *testing.T) {
	type args struct {
		rawString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{rawString: "1"},
			want: "c4ca4238a0b923820dcc509a6f75849b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.rawString); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMD5File(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{filePath: "/Users/rogee/Develop/projects/gocode/packages/crawler/encode/md5/file.md5.test"},
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MD5File(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5File() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MD5File() = %v, want %v", got, tt.want)
			}
		})
	}
}
