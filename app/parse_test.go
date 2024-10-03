package main

import (
	"fmt"
	"testing"
)

func TestParseArgsRESP(t *testing.T) {
	data := []byte("*2\r\n$4\r\nECHO\r\n$5\r\nmykey\r\n")
	got, err := parseRequest(data)
	if err != nil {
		t.Errorf("parseRequest() error = %v", err)
	}
	want := []string{"echo", "mykey"}
	if !equalString(got, want) {
		t.Errorf("parseRequest() = %v, want %v", got, want)
	}
}

func TestParseArgsRESP2(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []string
		wantErr bool
	}{
		{
			name:    "Valid data",
			data:    []byte("*2\r\n$3\r\nSET\r\n$5\r\nmykey\r\n"),
			want:    []string{"set", "mykey"},
			wantErr: false,
		},
		{
			name:    "Empty data",
			data:    []byte(""),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid data",
			data:    []byte("invalid"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid array length",
			data:    []byte("*a\r\n$3\r\nSET\r\n$5\r\nmykey\r\n"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid command length",
			data:    []byte("*2\r\n$a\r\nSET\r\n$5\r\nmykey\r\n"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid command args length",
			data:    []byte("*2\r\n$3\r\nSET\r\n$a\r\nmykey\r\n"),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRequest(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalString(got, tt.want) {
				t.Errorf("parseRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseResponse(t *testing.T) {
	data := []byte("*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n")
	command, err := parseRequest(data)
	fmt.Println(command)
	got, err := parseResponseCommand(command)
	if err != nil {
		t.Errorf("parseResponseCommand() error = %v", err)
	}
	want := []byte("$3\r\nhey\r\n")
	if !equalByte(got, want) {
		t.Errorf("parseRequest() = %s, want %s", got, want)
	}
}
func equalString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalByte(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
