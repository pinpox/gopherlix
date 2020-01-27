package main

import "testing"

func Test_createLink(t *testing.T) {
	tests := []struct {
		name     string
		itemType string
		text     string
		path     string
		want     string
	}{
		// TODO: Add test cases.
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
