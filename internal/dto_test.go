package internal

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "Test New User",
			args: args{userName: "instructor"},
			want: User{userName: "instructor"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.userName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsername(t *testing.T) {
	type fields struct {
		userName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test GetUsername()",
			fields: fields{userName: "instructor"},
			want:   "instructor",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				userName: tt.fields.userName,
			}
			if got := u.GetUsername(); got != tt.want {
				t.Errorf("GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
