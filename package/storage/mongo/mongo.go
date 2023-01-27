package mongo

import (
	"appDB/package/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const(
	database string = "news"
	collection string = "posts"
)

type mongoDB struct{
	ctx context.Context
	pool *mongo.Client
}

func New()*mongoDB{
	mongoOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil{
		log.Fatal(err)
	}
	err = client.Ping(context.Background(),nil)
	if err != nil{
		log.Fatal(err)
	}

	return & mongoDB{ctx : context.Background(), pool : client}
}



func(m *mongoDB) Posts()([]storage.Post, error){  // получение всех публикаций
	collection := m.pool.Database(database).Collection(collection)
	filter := bson.D{}
	cur, err := collection.Find(m.ctx, filter)
	defer cur.Close(m.ctx)
	if err != nil{
		return []storage.Post{}, err
	}
	var result []storage.Post
	for cur.Next(m.ctx){
		var p storage.Post
		err = cur.Decode(&p)
		if err != nil{
			return []storage.Post{}, err
		}
		result = append(result, p)

	}
	err = cur.Err()
	if err != nil{
		return []storage.Post{}, err
	}

	return result, nil
}
func(m *mongoDB) AddPost(post storage.Post) error { // создание новой публикации
	collection := m.pool.Database(database).Collection(collection)
	_, err := collection.InsertOne(m.ctx, post)
	if err != nil{
		return err
	}
	return nil
}
func(m *mongoDB) UpdatePost(post storage.Post) error { // обновление публикации
	collection := m.pool.Database(database).Collection(collection)
	filter := bson.M{"id" : post.ID}
	set := bson.D{{"$set", bson.D{{"title", post.Title},{"content", post.Content},
		{"authorid", post.AuthorID}, {"authorname", post.AuthorName},
		{"createdat", post.CreatedAt}, {"publishedat", post.PublishedAt}}}}
	_, err := collection.UpdateOne(m.ctx, filter, set )
	if err != nil{
		return err
	}
	return nil
}
func (m *mongoDB) DeletePost(post storage.Post) error { // удаление публикации по ID
	collection := m.pool.Database(database).Collection(collection)
	filter := bson.M{"id" : post.ID}
	_, err := collection.DeleteOne(m.ctx, filter)
	if err != nil{
		return err
	}
	return nil
}


func(m *mongoDB) Close(){ // закрытие соединения, общее для интерфейса
	err := m.pool.Disconnect(m.ctx)
	if err != nil{
		log.Fatal(err)
	}
}
