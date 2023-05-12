package storage

import (
	"reflect"
	"testing"
)

func initStorage() ItemType {
	s := ItemType{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	return s
}

func Test_storage_Add(t *testing.T) {
	type fields struct {
		s ItemType
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "1",
			fields: fields{s: initStorage()},
			args:   args{key: "key5", value: "value5"},
			want:   "key5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				s: tt.fields.s,
			}
			if got := s.Add(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("storage.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storage_Get(t *testing.T) {
	type fields struct {
		s ItemType
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "1",
			fields: fields{s: initStorage()},
			args: args{
				key: "key1",
			},
			want:    "value1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				s: tt.fields.s,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("storage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storage_List(t *testing.T) {
	type fields struct {
		s ItemType
	}
	tests := []struct {
		name   string
		fields fields
		want   ItemType
	}{
		{
			name:   "1",
			fields: fields{s: initStorage()},
			want:   initStorage(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				s: tt.fields.s,
			}
			if got := s.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storage.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
