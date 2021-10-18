package models

import (
	"reflect"
	"testing"
)

func Test_user_Map(t *testing.T) {
	type fields struct {
		ID         string
		FatherName string
		FirstName  string
		LastName   string
		Gender     string
		CNIC       string
		Email      string
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &user{
				ID:         tt.fields.ID,
				FatherName: tt.fields.FatherName,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				Gender:     tt.fields.Gender,
				CNIC:       tt.fields.CNIC,
				Email:      tt.fields.Email,
				Password:   tt.fields.Password,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_Names(t *testing.T) {
	type fields struct {
		ID         string
		FatherName string
		FirstName  string
		LastName   string
		Gender     string
		CNIC       string
		Email      string
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &user{
				ID:         tt.fields.ID,
				FatherName: tt.fields.FatherName,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				Gender:     tt.fields.Gender,
				CNIC:       tt.fields.CNIC,
				Email:      tt.fields.Email,
				Password:   tt.fields.Password,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
