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
			server: NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
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
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			reqPath: "",
			want: "0file1	file1	localhost	8000\r\n1sub1	sub1	localhost	8000\r\n1subdir1	subdir1	localhost	8000\r\n1subdir2	subdir2	localhost	8000",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			reqPath: "subdir1",
			want: "0file4	subdir1/file4	localhost	8000\r\n0file5	subdir1/file5	localhost	8000",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			reqPath: "subdir2",
			want: "0file2	subdir2/file2	localhost	8000\r\n0file3	subdir2/file3	localhost	8000\r\n0index.gph	subdir2/index.gph	localhost	8000",
			wantErr: false,
		},
		{
			name:    "Test invalid path",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
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

func TestGopherServer_parseTemplate(t *testing.T) {
	tests := []struct {
		name    string
		server  GopherServer
		data    map[string]string
		templ   string
		want    string
		wantErr bool
	}{

		{name: "Empty template",
			server: NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			data: map[string]string{
				"Directory":  "/here/is/a/path",
				"ServerName": "servername.com",
			},
			templ:   "",
			want:    "",
			wantErr: false,
		},

		{name: "Template without variables",
			server: NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			data: map[string]string{
				"Directory":  "/here/is/a/path",
				"ServerName": "servername.com",
			},
			templ:   "templatecontent",
			want:    "templatecontent",
			wantErr: false,
		},

		{name: "Template with variables",
			server: NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			data: map[string]string{
				"Directory":  "/here/is/a/path",
				"ServerName": "servername.com",
			},
			templ:   "templatestart{{.Directory}}{{.ServerName}}templateend",
			want:    "templatestart/here/is/a/pathservername.comtemplateend",
			wantErr: false,
		},

		{name: "Template with missing variables",
			server: NewGopherServer("8000", "localhost", "localhost", "testdata/content", "testdata/templates"),
			data: map[string]string{
				"Directory":  "/here/is/a/path",
				"ServerName": "servername.com",
			},
			templ:   "templatestart{{.Directory}}{{.Somevar}}{{.ServerName}}templateend",
			want:    "templatestart/here/is/a/path<no value>servername.comtemplateend",
			wantErr: false,
		},
		// TODO: Add test more cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.parseTemplate(tt.templ, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServer.parseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GopherServer.parseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
