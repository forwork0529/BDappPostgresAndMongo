package main

import (
	"appDB/package/storage"
	"appDB/package/storage/memDB"
	"reflect"
	"testing"
)

func TestPostID(t *testing.T){
	type args struct {
		db storage.InstanceDB
		id int
	}
	tests := []struct {
		name string
		args args
		want storage.Post
	}{
		{ 	"simple",
			args{memDB.New(),1},
			storage.Post{ID : 1, Title : "first",Content: "1 content", AuthorID : 1, AuthorName : "Ivan", CreatedAt: 10000000, PublishedAt: 10000001}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storage.PostID(tt.args.db, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostID() = %v, want %v", got, tt.want)
			}
		})
	}
}