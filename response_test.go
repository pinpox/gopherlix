package main

import (
	"testing"
)

func TestGopherServer_createLink(t *testing.T) {
	type args struct {
		itemType string
		text     string
		path     string
	}
	tests := []struct {
		name   string
		server GopherServer
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "Menu entry",
			server: NewGopherServer("8000", "localhost", "localhost", "testdata"),
			args: args{
				itemType: "MENU",
				text:     "mytext1",
				path:     "path/to/some/file"},
			want: "1mytext1	path/to/some/file	localhost	8000\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.server.createLink(tt.args.itemType, tt.args.text, tt.args.path); got != tt.want {
				t.Errorf("GopherServer.createLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServer_createListing(t *testing.T) {

	tests := []struct {
		name    string
		server  GopherServer
		reqPath string
		want    string
		wantErr bool
	}{
		{
			name:    "Test root",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata",
			want: "0file1	file1	localhost	8000\r\n1sub1	sub1	localhost	8000\r\n1subdir1	subdir1	localhost	8000\r\n1subdir2	subdir2	localhost	8000\r\n.",
			wantErr: false,
		},
		{
			name:    "Test root/file1",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/file1",
			want:    "file1content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/subdir1",
			want: "0file4	file4	localhost	8000\r\n0file5	file5	localhost	8000\r\n.",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/subdir2",
			want:    "indexcontent",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1/file4 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/subdir1/file4",
			want:    "file4content",
			wantErr: false,
		},
		{
			name:    "Test root/sub1/subdir1/subdir2/subdir3/file6  dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/sub1/subdir1/subdir2/subdir3/file6",
			want:    "file6content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1/file5 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/subdir1/file5",
			want:    "file5content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2/file2 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/subdir2/file2",
			want:    "file2content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2/file3 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "testdata/subdir2/file3",
			want:    "file3content",
			wantErr: false,
		},
		{
			name:    "Test invalid path",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			reqPath: "subdir2/file48\r\n",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.createListing(tt.reqPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServer.createListing() error = \n       %v,\n wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createListing() =\n      %v,\n want %v", replaceCRLF(got), replaceCRLF(tt.want))
			}
		})
	}
}
