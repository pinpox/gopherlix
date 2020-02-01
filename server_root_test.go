package main

import (
	"reflect"
	"testing"
)

func TestNewGopherServerRoot(t *testing.T) {
	tests := []struct {
		name      string
		root      string
		templates string
		want      *GopherServerRoot
		wantErr   bool
	}{
		{
			name:      "Create a valid GopherServerRoot",
			root:      "testdata/content",
			templates: "testdata/templates",
			want: &GopherServerRoot{
				"testdata/content",
				"testdata/templates",
			},
			wantErr: false,
		},
		{
			name:      "Try to create GopherServerRoot with invalid root",
			root:      "testdata/contentinvalid",
			templates: "testdata/templates",
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "Try to create GopherServerRoot with invalid templates dir",
			root:      "testdata/content",
			templates: "testdata/templatesinvalid",
			want:      nil,
			wantErr:   true,
		},
		{
			name:      "Try to create GopherServerRoot with both invalid dirs",
			root:      "testdata/contentinvalid",
			templates: "testdata/templatesinvalid",
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGopherServerRoot(tt.root, tt.templates)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGopherServerRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGopherServerRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_FileExists(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			if got := sr.FileExists(tt.args.path); got != tt.want {
				t.Errorf("GopherServerRoot.FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_DirExists(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			if got := sr.DirExists(tt.args.path); got != tt.want {
				t.Errorf("GopherServerRoot.DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_getSavePath(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
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
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			got, err := sr.getSavePath(tt.args.subPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServerRoot.getSavePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GopherServerRoot.getSavePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_Type(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	type args struct {
		reqPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			if got := sr.Type(tt.args.reqPath); got != tt.want {
				t.Errorf("GopherServerRoot.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_HeaderTemplate(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			if got := sr.HeaderTemplate(); got != tt.want {
				t.Errorf("GopherServerRoot.HeaderTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_FooterTemplate(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			if got := sr.FooterTemplate(); got != tt.want {
				t.Errorf("GopherServerRoot.FooterTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_GetServerFile(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	type args struct {
		subpath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			got, err := sr.GetServerFile(tt.args.subpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServerRoot.GetServerFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GopherServerRoot.GetServerFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGopherServerRoot_GetServerDir(t *testing.T) {
	type fields struct {
		ServerRootDir string
		TemplatesDir  string
	}
	type args struct {
		subpath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &GopherServerRoot{
				ServerRootDir: tt.fields.ServerRootDir,
				TemplatesDir:  tt.fields.TemplatesDir,
			}
			got, err := sr.GetServerDir(tt.args.subpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GopherServerRoot.GetServerDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GopherServerRoot.GetServerDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
