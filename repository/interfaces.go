// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	InsertEstate(ctx context.Context, input CreateEstateInput) (output UuidOutput, err error)
	InsertTree(ctx context.Context, input CreateTreeInput) (output UuidOutput, err error)
	GetAllTree(ctx context.Context, input UuidInput) (output []TreeModel, err error)
}
