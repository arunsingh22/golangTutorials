package main

import (
	"context"
	"fmt"
	"sync"
)

var (
	dbConn *database
	once   sync.Once
)

type database struct {
	ctx  context.Context
	name string
}

func getDBInstance(ctx context.Context, name string) {
	once.Do(func() {
		dbConn = &database{
			ctx:  ctx,
			name: name,
		}
		fmt.Println("New DB connection established...")
		fmt.Println("this automatically returnes")
	})
	fmt.Println("Already created...")
}
