package repository

import "github.com/google/wire"

// Set RepositorySet is a Wire provider set for all repositories
var Set = wire.NewSet(
	NewUserRepository,
	NewContactRepository,
	NewAddressRepository,
)
