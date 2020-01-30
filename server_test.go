package main

import "testing"

func TestGopherServer_parseRequest(t *testing.T) {

	tests := []struct {
		name    string
		server  GopherServer
		request string
		want    string
		wantErr bool
	}{

		{
			name:    "Test root",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "\r\n",
			want: "0file1	file1	localhost	8000\r\n1sub1	sub1	localhost	8000\r\n1subdir1	subdir1	localhost	8000\r\n1subdir2	subdir2	localhost	8000\r\n.",
			wantErr: false,
		},
		{
			name:    "Test root/file1",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "file1",
			want:    "file1content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir1\r\n",
			want: "0file4	subdir1/file4	localhost	8000\r\n0file5	subdir1/file5	localhost	8000\r\n.",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir2\r\n",
			want:    "indexcontent",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1/file4 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir1/file4\r\n",
			want:    "file4content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1/file5 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir1/file5\r\n",
			want:    "file5content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2/file2 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir2/file2\r\n",
			want:    "file2content",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2/file3 dir request",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir2/file3\r\n",
			want:    "file3content",
			wantErr: false,
		},
		{
			name:    "Test invalid path",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			request: "subdir2/file48\r\n",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.parseRequest(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServer.parseRequest() error = \n        %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GopherServer.parseRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
