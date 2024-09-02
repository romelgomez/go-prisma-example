package repository

import (
	"context"
	"errors"
	"fmt"
	"go-prisma-example/helper"
	"go-prisma-example/model"
	"go-prisma-example/prisma/db"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewPostRepository(Db *db.PrismaClient) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// FindAll implements PostRepository.
func (p *PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {
	allPost, err := p.Db.Post.FindMany().Exec(ctx)
	helper.ErrorPanic(err)

	var posts []model.Post

	for _, post := range allPost {
		published := post.Published
		description, _ := post.Description()

		postData := model.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: description,
		}

		posts = append(posts, postData)
	}

	return posts
}

// FindById implements PostRepository.
func (p *PostRepositoryImpl) FindById(ctx context.Context, postId string) (model.Post, error) {
	// Attempt to find the post by ID
	post, err := p.Db.Post.FindFirst(db.Post.ID.Equals(postId)).Exec(ctx)
	if err != nil {
		// Return an error if the query fails
		return model.Post{}, err
	}

	// Check if post is nil to avoid a nil pointer dereference
	if post == nil {
		// Return an error indicating that the post was not found
		return model.Post{}, errors.New("post id not found")
	}

	// Safely access post fields and methods after confirming post is not nil
	published := post.Published
	description, _ := post.Description()

	// Create and return the Post model
	postData := model.Post{
		Id:          post.ID,
		Title:       post.Title,
		Published:   published,
		Description: description,
	}

	return postData, nil
}

// Save implements PostRepository.
func (p *PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	result, err := p.Db.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	result, err := p.Db.Post.FindMany(db.Post.ID.Equals(post.Id)).Update(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}
