package repository

import (
	"context"
	"database/sql"
)

// GetAllTree implements RepositoryInterface.
func (r *Repository) GetAllTree(ctx context.Context, input UuidInput) (output []TreeModel, err error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT uuid, x_axis, y_axis, height, estate_uuid FROM tree where estate_uuid = $1", input.Uuid)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tree TreeModel
		err = rows.Scan(&tree.Uuid, &tree.X, &tree.Y, &tree.Height, &tree.EstateUuid)
		if err != nil {
			return
		}

		output = append(output, tree)
	}
	return
}

func (r *Repository) GetEstateStats(ctx context.Context, input UuidInput) (output EstateStatsOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT COUNT(*), MAX(height), MIN(height), percentile_cont(0.5) WITHIN GROUP (ORDER BY height) as median FROM tree where estate_uuid = $1", input.Uuid).Scan(&output.Count, &output.Max, &output.Min, &output.Median)
	if err == sql.ErrNoRows {
		output = EstateStatsOutput{
			Count:  0,
			Max:    0,
			Min:    0,
			Median: 0,
		}
		err = nil
	}
	return
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
func (r *Repository) InsertTree(ctx context.Context, input CreateTreeInput) (output UuidOutput, err error) {

	err = r.Db.QueryRowContext(ctx, "INSERT INTO tree (x_axis, y_axis, height, estate_uuid) VALUES ($1, $2, $3, $4) RETURNING uuid", input.X, input.Y, input.Height, input.EstateId).Scan(&output.Uuid)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetEstate(ctx context.Context, input UuidInput) (output EstateModel, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT uuid, width, length FROM estate where uuid = $1", input.Uuid).Scan(&output.Uuid, &output.Width, &output.Length)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetTree(ctx context.Context, input GetTreeByCoordinateInput) (output TreeModel, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT uuid, x_axis, y_axis, height, estate_uuid FROM tree where x_axis = $1 and y_axis = $2 and estate_uuid = $3 order by y_axis, x_axis", input.X, input.Y, input.EstateUuid).Scan(&output.Uuid, &output.X, &output.Y, &output.Height, &output.EstateUuid)
	return
}
