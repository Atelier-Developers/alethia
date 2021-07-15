package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type AuthInterface interface {
	CreateAuth(uint64, *TokenDetails) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type ClientData struct {
	ctx    context.Context
	client *redis.Client
}

var _ AuthInterface = &ClientData{}

func NewAuth(client *redis.Client, ctx context.Context) *ClientData {
	return &ClientData{client: client, ctx: ctx}
}

type AccessDetails struct {
	TokenUuid string
	UserId    uint64
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	TokenUuid           string
	RefreshUuid         string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

func (tk *ClientData) CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AccessTokenExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RefreshTokenExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(tk.ctx, td.TokenUuid, strconv.Itoa(int(userid)), at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(tk.ctx, td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

func (tk *ClientData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := tk.client.Get(tk.ctx, tokenUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func (tk *ClientData) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)
	fmt.Println(authD)
	//delete access token
	deletedAt, err := tk.client.Del(tk.ctx, authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := tk.client.Del(tk.ctx, refreshUuid).Result()
	if err != nil {
		return err
	}
	fmt.Println(deletedAt, deletedRt)
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *ClientData) DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	deleted, err := tk.client.Del(tk.ctx, refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
