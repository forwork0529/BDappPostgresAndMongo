package memDB

import (
	"appDB/package/storage"
	"context"
)

type memoryDB struct{
	data []storage.Post
}

/*
type Post struct{
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}
*/


var data  = []storage.Post{
	{ID : 1, Title : "first",Content: "1 content", AuthorID : 1, AuthorName : "Ivan", CreatedAt: 10000000, PublishedAt: 10000001},
	{ID : 2, Title : "second",Content: "2 content", AuthorID : 2, AuthorName : "Vladimir", CreatedAt: 20000000, PublishedAt: 20000002},
}

func New()*memoryDB{
	return &memoryDB{
		data : data,
	}
}

func(m *memoryDB) Posts()([]storage.Post, error){  // получение всех публикаций
	return m.data, nil
}
func(m *memoryDB) AddPost(storage.Post) error { // создание новой публикации
	return nil
}
func(m *memoryDB) UpdatePost(storage.Post) error { // обновление публикации
	return nil
}
func(m *memoryDB) DeletePost(storage.Post) error { // удаление публикации по ID
	return nil
}


func(m *memoryDB) Close(ctx context.Context){ // удаление публикации по ID

}
