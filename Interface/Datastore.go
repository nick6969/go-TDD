package Interface

import "tdd/Database/mysql"

// Datastore is interface with real database
type Datastore interface {
	DatastoreCustomer
	DatastoreAuthCustomer
}

// DatastoreCustomer is interface with real database specific table `Customer`
type DatastoreCustomer interface {
	CheckUserNameCanUse(name string) bool
	CreateCustomer(username, password string) (uint, error)
	FindUserWithUserName(name string) (user mysql.Customer, err error)
}

type DatastoreAuthCustomer interface {
}
