package config

import (
	"os/user"
	"testing"
)



func Test_validateHomeDir(t *testing.T) {
	usr, _ := user.Current()

	type args struct {
		path string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{path: "~/test"}, want: usr.HomeDir + "/test"},
		{name: "test2", args: args{path: "/test"}, want: "/test"},
		{name: "test3", args: args{path: "test"}, want: usr.HomeDir + "/test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHomeDir(tt.args.path); got != tt.want {
				t.Errorf("validateHomeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
