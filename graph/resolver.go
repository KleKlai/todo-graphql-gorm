package graph

import (
	"github.com/kleklai/todoAppv1/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service *service.Service
}
