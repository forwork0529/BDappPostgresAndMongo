package storage

func PostID(db InstanceDB, id int)Post{
	posts, err := db.Posts()
	if err != nil{
		return Post{}
	}
	for _, post := range posts{
		if post.ID == id{
			return post
		}
	}
	return Post{}
}
