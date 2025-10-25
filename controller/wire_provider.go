package controller

import "github.com/google/wire"

// Set ControllerSet is a Wire provider set for all controllers
var Set = wire.NewSet(
	NewUserController,
	NewContactController,
	NewAddressController,
)
