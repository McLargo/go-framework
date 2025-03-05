package db

const (
	DBTypeInMemory string = "memory"
	DBTypeMongo    string = "mongo"
)

// DatabaseAdapter is the interface used by the controllers to interact with the database.
type DatabaseAdapter interface {
	Connect() error
	Disconnect() error
}
