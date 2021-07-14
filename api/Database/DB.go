package Database


type DB interface {
	Init() error
	Close() error
}