package repository

import "context"

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

// GetAllTree implements RepositoryInterface.
func (*Repository) GetAllTree(ctx context.Context, input UuidInput) (output []TreeModel, err error) {
	panic("unimplemented")
}

// InsertEstate implements RepositoryInterface.
func (r *Repository) InsertEstate(ctx context.Context, input CreateEstateInput) (output UuidOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO estate (length, width) VALUES ($1, $2) RETURNING uuid", input.Length, input.Width).Scan(&output.Uuid)
	if err != nil {
		return
	}
	return
}

// InsertTree implements RepositoryInterface.
func (*Repository) InsertTree(ctx context.Context, input CreateTreeInput) (output UuidOutput, err error) {
	panic("unimplemented")
}
