package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Collection *mongo.Collection
}

func (s *UserService) Register(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	var existing models.User
	err := s.Collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		return mongo.ErrNoDocuments
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Role = "user"

	_, err = s.Collection.InsertOne(ctx, user)
	return err
}

func (s *UserService) Login(email, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := s.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, errors.New("invalid credentials")
	}
	
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return user, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.Collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.User{}, err
	}

	var users []models.User
	for cursor.Next(ctx) {
		var u models.User
		cursor.Decode(&u)
		users = append(users, u)
	}
	return users, nil
}
