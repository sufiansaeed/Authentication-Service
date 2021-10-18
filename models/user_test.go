package models

import (
	"reflect"
	"testing"
)

func TestUser_Map(t *testing.T) {
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
		{
			name: "successful test case",
			fields: fields{
				ID: "123",
				FatherName: "Saeed",
				FirstName: "Sufian",
				LastName: "Saeed",
				Gender: "Male",
				CNIC: "13101-5418171-5",
				Email: "sufiansaeed97@gmail.com",
				Password: "123456789",
			},
			want: map[string]interface{}{
				"id": "123",
				"father_name": "Saeed",
				"first_name": "Sufian",
				"last_name": "Saeed",
				"gender": "Male",
				"cnic": "13101-5418171-5",
				"email": "sufiansaeed97@gmail.com",
				"password": "123456789",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &User{
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

func TestUser_Names(t *testing.T) {
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
		{
			name: "Success!",
			fields: fields{
				ID: "123",
				FatherName: "Saeed",
				FirstName: "Sufian",
				LastName: "Saeed",
				Gender: "Male",
				CNIC: "13101-5418171-5",
				Email: "sufiansaeed97@gmail.com",
				Password: "123456789",
			},
			want: []string{"id", "father_name", "first_name", "last_name", "gender", "cnic", "email", "password"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:         tt.fields.ID,
				FatherName: tt.fields.FatherName,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				Gender:     tt.fields.Gender,
				CNIC:       tt.fields.CNIC,
				Email:      tt.fields.Email,
				Password:   tt.fields.Password,
			}
			if got := u.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
