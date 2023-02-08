package main

import (
	"appDB/package/storage"
	"appDB/package/storage/memDB"
	"reflect"
	"testing"
	"time"
)

//    -count=1 флаг тменяющий кеширование во втором варианте запуска тестов
//    go test -run=XXX -bench=. -benchmem - ля бренчмарка и контролем распределения памяти


/*func TestMain(m *testing.M){					// ункция для подготовки к тесту setup tear down
	fmt.Println("I will wait")
	time.Sleep(time.Second * 3)
	fmt.Println("All right")
	time.Sleep(time.Second * 2)
	m.Run()
}*/

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

func BenchmarkTestPostId(b *testing.B){
	var db storage.InstanceDB = memDB.New()
	time.Sleep(time.Second * 3)               // Работает
	b.ResetTimer()							  // Работает
	for i := 0; i < b.N ; i ++ {
		_ = storage.PostID(db, 1)
	}
}