package database

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db  *gorm.DB
	dns string
	ctx context.Context
}

func NewClient(ctx context.Context, dns string) *Client {
	return &Client{
		dns: dns,
		ctx: ctx,
	}
}

func (c *Client) Connect() error {
	fmt.Println("connecting to database" + c.dns)
	db, err := gorm.Open(postgres.Open(c.dns), &gorm.Config{})
	if err != nil {
		return errors.New("cannot connect to database")
	}

	c.db = db
	return nil
}

func (c *Client) DB() *gorm.DB {
	return c.db
}
