package postgres

import (
	"appDB/package/storage"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

const(
	databaseUrl string = "postgres://postgres:postgres@127.0.0.1:8081/posts_db"
	)
type postgresDB struct{
	ctx context.Context
	pool *pgxpool.Pool
}

func New()*postgresDB{
	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil{
		log.Fatal(err)
	}
	err = pool.Ping(context.Background())
	if err != nil{
		log.Fatal(err)
	}
	return &postgresDB{ctx : context.Background(), pool : pool}
}

func(p *postgresDB) Posts()([]storage.Post, error){  // получение всех публикаций
	rows, err := p.pool.Query(p.ctx, `SELECT posts.id, posts.author_id,
posts.title, posts.content, posts.created_at, authors.name FROM posts, authors WHERE posts.author_id = authors.id;`)
	defer rows.Close()
	if err == pgx.ErrNoRows{
		return []storage.Post{}, nil

	}
	if err != nil{
		return []storage.Post{}, err
	}
	var result []storage.Post
	for rows.Next(){
		var p storage.Post
		err = rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt, &p.AuthorName)

		if err != nil{
			return []storage.Post{}, err
		}
		result = append(result, p)
	}
	err = rows.Err()
	if  err != nil{
		return []storage.Post{}, err
	}
	return result, nil
}



func(p *postgresDB) AddPost(post storage.Post) error { // создание новой публикации
	_, err := p.pool.Exec(p.ctx, `INSERT INTO posts (author_id,
title, content, created_at) VALUES ($1, $2, $3, $4);`,
post.AuthorID, post.Title, post.Content, post.CreatedAt )
	if err != nil{
		return err
	}
	return nil
}

func(p *postgresDB) UpdatePost(post storage.Post) error { // обновление публикации
	_, err := p.pool.Exec(p.ctx, `UPDATE posts SET author_id = $1,
title = $2, content = $3, created_at = $4 WHERE posts.id = $5;`, post.AuthorID,
post.Title, post.Content, post.CreatedAt, post.ID )
	if err != nil{
		return err
	}
	return nil
}
func (p *postgresDB) DeletePost(post storage.Post) error { // удаление публикации по ID
	_, err := p.pool.Exec(p.ctx, `DELETE FROM posts WHERE posts.id = $1;`, post.ID)
	if err != nil{
		return err
	}
	return nil
}

func(p *postgresDB) Close(){ // закрытие соединения, общее для интерфейса

}