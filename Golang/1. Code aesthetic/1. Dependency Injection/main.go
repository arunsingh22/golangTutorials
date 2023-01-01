package main

import (
	"fmt"
	"time"
)

// interfaces are abstract types with only collection of funcs signatures and that is why we cannot create objects (aka instantiate them but we can definately create variables for it.

type storage interface {
	execute() // only func name and return type.
}

type database struct{}

func (db *database) execute() {
	fmt.Println("Storing into DB")
}

type fileSystem struct{}

func (fs *fileSystem) execute() {
	fmt.Println("Storing into fileSystem")
}

type mail interface {
	connect()
	send()
}

type Gmail struct {
	name string
}

func (g *Gmail) connect() {
	fmt.Println(g.name, ":Connected to gmail server..")
	time.Sleep(5 * time.Millisecond)
}

func (g *Gmail) send() {
	fmt.Println(g.name, "Mail sent to the destination.")
}

// // Method 1:
// type Service struct {
// 	name string
// 	db   *database   // tight coupling aka object dependecy
// 	fs   *fileSystem // tight coupling aka object dependecy
// }

// // Public (because it's not having prefix ) Constructor which returns the object for Service
// func NewService(db *database, fs *fileSystem) *Service {
// 	srv := &Service{
// 		name: "service_1",
// 		db:   db,
// 		fs:   fs,
// 	}
// 	return srv
// }

// Method 2:
type Service struct {
	name  string
	sto   storage // Using interface we are decoupling (followed `D` of SOLID principle)
	email mail    // Using interface we are decoupling (followed `D` of SOLID principle)
}

// Public (because it's not having prefix ) Constructor which returns the object for Service
func NewService(str storage, email mail) *Service {
	srv := &Service{
		name:  "service_1",
		sto:   str,
		email: email,
	}
	return srv
}

func main() {
	// tight coupling
	// d := &database{}
	// d.execute()

	// syntax to create a var for interface and not objects as interfaces are abstract
	// var st storage

	// Method 1:
	// fs := &fileSystem{}
	// db := &database{}
	// s := NewService(db, fs)
	// fmt.Println(s.name)
	// s.db.execute()

	// Method 2:
	fs := &fileSystem{}
	// gmail := &Gmail{
	// 	name: "gmail",
	// }
	s := NewService(fs, nil)
	fmt.Println(s.name)
	s.email.connect()
}
