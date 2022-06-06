package storage

type DB interface {
	Connect() bool
}
