package service

import "github.com/google/wire"

// Set ServiceSet is a Wire provider set for all services
var Set = wire.NewSet(
	NewUserService,
	NewContactService,
	NewAddressService,
)
