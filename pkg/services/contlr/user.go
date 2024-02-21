package contlr

import (
	"app-controller/pkg/model"
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	//"encoding/json"
)

func (s *service) GetUsers(ctx context.Context) ([]model.User, error) {

	out, err := s.usersRepo.Query(ctx)

	if err != nil {
		return []model.User{}, err
	}

	return out, err
}

func (s *service) GetUser(ctx context.Context,u string) (model.User, error) {

	out, err := s.usersRepo.QueryInfo(ctx,u)

	if err != nil {
		return model.User{}, err
	}

	return out, err
}

func (s *service) AddUser(ctx context.Context, f model.User) (string,error) {
	err := 	s.usersRepo.Create(ctx, model.User{
			Username:   	f.Username,
			Password:   	f.Password,
			Name:   		f.Name,
			Surname: 		f.Surname,
			Email: 			f.Email,
			Github: 		f.Github,
			Phone: 			f.Phone,
			Description: 	f.Description,
			Attachment:     f.Attachment,
	})

	token := jwt.New(jwt.SigningMethodHS256)
	claims := model.UserToken{
		Username: f.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte("notdoingthatnowlol"))
	if err != nil {
		return "" , err
	}

	if err := 	s.usersRepo.CreateToken(ctx,claims); err != nil {
		return "" , errors.New("fail to create token")
	}

	if err != nil {
		return "",err
	}

	return tokenString , nil
}

func (s *service) Login(ctx context.Context, f model.UserLogin) (string,error) {
	usr , err := 	s.usersRepo.QueryByUsername(ctx,f.Username)

	if err != nil {
		return "" , err
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(usr.Password)
	if err != nil {
		return "" , errors.New("error decoding base64 string")
	}

	if err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(f.Password)); err != nil {
		return "" , errors.New("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := model.UserToken{
		Username: f.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token.Claims = claims

	tokenString, err := token.SignedString([]byte("notdoingthatnowlol"))
	if err != nil {
		return "" , err
	}

	if err := 	s.usersRepo.CreateToken(ctx,claims); err != nil {
		return "" , errors.New("fail to create token")
	}

	return tokenString , nil
}

func (s *service) Auth(ctx context.Context, f model.UserLogin) (model.UserToken,error) {
	usr , err := 	s.usersRepo.QueryToken(ctx,f.Username)

	if err != nil {
		return model.UserToken{} , err
	}

	return usr , nil
}

func (s *service) EditUser(ctx context.Context, f model.User, u string) error {
	usr , err := s.usersRepo.QueryByUsername(ctx,u)

	if err != nil {
		return  err
	}

	err = 	s.usersRepo.Update(ctx, model.User{
			ID: 			usr.ID,	
			Name:   		f.Name,
			Surname: 		f.Surname,
			Email: 			f.Email,
			Github: 		f.Github,
			Phone: 			f.Phone,
			Description: 	f.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) EditProfile(ctx context.Context, f model.ProfileAttachment, u string) error {
	usr , err := s.usersRepo.QueryByUsername(ctx,u)

	if err != nil {
		return  err
	}

	err = 	s.usersRepo.UpdateProfile(ctx, model.ProfileAttachment{
			ID: 			usr.Attachment.ID,	
			UserId: 		usr.ID,
			Name:   		f.Name,
			Src:			f.Src,
			Size:			f.Size,
	})

	if err != nil {
		return err
	}

	return nil
}