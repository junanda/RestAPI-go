package database

type Database interface {
	Connect()
	Close()
}
