// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	InsertEstate(ctx context.Context, input CreateEstateInput) (output UuidOutput, err error)
	InsertTree(ctx context.Context, input CreateTreeInput) (output UuidOutput, err error)
	GetAllTree(ctx context.Context, input UuidInput) (output []TreeModel, err error)
	GetTree(ctx context.Context, input GetTreeByCoordinateInput) (output TreeModel, err error)
	GetEstateStats(ctx context.Context, input UuidInput) (output EstateStatsOutput, err error)
	GetEstate(ctx context.Context, input UuidInput) (output EstateModel, err error)
}
