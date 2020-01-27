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
		{"Test existing directory", "tesdata/subdir1", true},
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
	type fields struct {
		Port    string
		Domain  string
		Host    string
		RootDir string
		run     bool
		signals chan bool
	}
	type args struct {
		subPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &GopherServer{
				Port:    tt.fields.Port,
				Domain:  tt.fields.Domain,
				Host:    tt.fields.Host,
				RootDir: tt.fields.RootDir,
				run:     tt.fields.run,
				signals: tt.fields.signals,
			}
			got, err := server.getSavePath(tt.args.subPath)
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
