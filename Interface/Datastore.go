package Interface

// Datastore is interface with real database
type Datastore interface {
	DatastoreCustomer
}

// DatastoreCustomer is interface with real database specific table `Customer`
type DatastoreCustomer interface {
	CheckUserNameCanUse(name string) bool
	CreateCustomer(username, password string) (uint, error)
}
