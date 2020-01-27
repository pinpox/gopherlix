package main

import (
	"fmt"
	"testing"
)

func Test_createLink(t *testing.T) {
	tests := []struct {
		name     string
		itemType string
		text     string
		path     string
		want     string
	}{
		// TODO: Add more test cases.
		{"Menu entry", "MENU", "mytext1", "path/to/some/file", "1mytext1	path/to/some/file	localhost	8000\r\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createLink(tt.itemType, tt.text, tt.path); got != tt.want {
				t.Errorf("createLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createListing(t *testing.T) {
	tests := []struct {
		name    string
		reqPath string
		want    string
		wantErr bool
	}{
		{
			name:    "Test root",
			reqPath: "testdata",
			want: "0file1	file1	localhost	8000\r\n1subdir1	subdir1	localhost	8000\r\n1subdir2	subdir2	localhost	8000\r\n.",
			wantErr: false,
		},
		{
			name:    "Test root/file1",
			reqPath: "testdata/file1",
			want:    "file1content\r\n",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1",
			reqPath: "testdata/subdir1",
			want:    "TODO",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2",
			reqPath: "testdata/subdir2",
			want:    "TODO",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1/file4 dir request",
			reqPath: "testdata/subdir1/file4",
			want:    "file4content\r\n",
			wantErr: false,
		},
		{
			name:    "Test root/subdir1/file5 dir request",
			reqPath: "testdata/subdir1/file5",
			want:    "file5content\r\n",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2/file2 dir request",
			reqPath: "testdata/subdir2/file2",
			want:    "file2content\r\n",
			wantErr: false,
		},
		{
			name:    "Test root/subdir2/file3 dir request",
			reqPath: "testdata/subdir2/file3",
			want:    "file3content\r\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createListing(tt.reqPath)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("createListing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createListing() = %v, want %v", got, tt.want)
			}
		})
	}
}
