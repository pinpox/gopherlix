package main

import "testing"

func Test_fileExists(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"Test existing file", "testdata/file1", true},
		{"Test non-existing file", "testdata/file2", false},
		{"Test existing directory", "testdata/subdir1", false},
		{"Test non-existing directory", "testdata/subdir3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExists(tt.filename); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dirExists(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"Test existing directory", "testdata/subdir1", true},
		{"Test non-existing directory", "testdata/subdir3", false},
		{"Test existing file", "testdata/file1", false},
		{"Test non-existing file", "testdata/file2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirExists(tt.filename); got != tt.want {
				t.Errorf("dirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServer_getSavePath(t *testing.T) {
	tests := []struct {
		name    string
		server  GopherServer
		subPath string
		want    string
		wantErr bool
	}{
		{name: "Test path inside server root",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			subPath: "subdir1",
			want:    "testdata/subdir1",
			wantErr: false,
		},

		{name: "Test absolute path outside server root",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			subPath: "/etc/passwd",
			want:    "testdata/etc/passwd",
			wantErr: false,
		},
		{name: "Test path containing ../../",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			subPath: "subdir1/../../../",
			want:    "",
			wantErr: true,
		},
		{name: "Test path starting with ../../",
			server:  NewGopherServer("8000", "localhost", "localhost", "testdata"),
			subPath: "../../../subdir",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.server.getSavePath(tt.subPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServer.getSavePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GopherServer.getSavePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
