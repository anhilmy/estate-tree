package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var treeData = []repository.TreeModel{
	{"76826522-c32b-4c72-8a2b-8c0231f80500", 23, 45, 15, "471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6"},
	{"bc069b9b-92e8-4782-9aa7-c20449db2b17", 78, 12, 20, "471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6"},
	{"8fc01fb8-27b7-4d5b-9676-8eb04152f2f2", 35, 67, 25, "471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6"},
	{"2ce1b501-b835-4058-a343-09124cd35f5b", 91, 31, 30, "471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6"},
	{"8a6ecc08-70d4-4529-9b8a-6b5321c96e41", 10, 99, 12, "471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6"},
	{"0907aec0-a256-4cc2-9a2b-aefe308d3296", 47, 53, 18, "29e365f7-1bb0-4b99-b09a-1571eb049dbd"},
	{"c96e3908-2577-4032-a637-5ab45ec9d96d", 89, 76, 22, "29e365f7-1bb0-4b99-b09a-1571eb049dbd"},
	{"0d41bec9-5eba-4cc4-a4e6-0b50bd11716c", 54, 33, 28, "29e365f7-1bb0-4b99-b09a-1571eb049dbd"},
	{"dd0d15fe-a18b-496f-86bb-7bed9f1d6950", 19, 82, 35, "29e365f7-1bb0-4b99-b09a-1571eb049dbd"},
	{"09a7c015-3a8b-4ef0-99c8-a09cc0c0749f", 63, 25, 40, "29e365f7-1bb0-4b99-b09a-1571eb049dbd"},
}

func Setup(t *testing.T) (repo repository.Repository, mock sqlmock.Sqlmock, err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return
	}

	repo = repository.Repository{
		Db: db,
	}
	return
}

func TestInsertEstateError(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	ent := repository.CreateEstateInput{
		Length: 0,
		Width:  0,
	}
	errExp := errors.New(gomock.Any().String())
	mock.ExpectQuery(`INSERT INTO estate \(length, width\) VALUES \(\$1, \$2\) RETURNING uuid`).
		WithArgs(ent.Length, ent.Width).
		WillReturnError(errors.New(gomock.Any().String()))

	uuid, err := repo.InsertEstate(context.Background(), ent)
	assert.Equal(t, errExp, err)
	assert.Zero(t, uuid)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestInsertEstateNormal(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	ent := repository.CreateEstateInput{
		Length: 20,
		Width:  19,
	}

	uuidExp := "1234"
	row := sqlmock.NewRows([]string{"uuid"}).AddRow(uuidExp)

	mock.ExpectQuery(`INSERT INTO estate \(length, width\) VALUES \(\$1, \$2\) RETURNING uuid`).
		WithArgs(ent.Length, ent.Width).
		WillReturnRows(row)

	uuid, err := repo.InsertEstate(context.Background(), ent)
	assert.Nil(t, err)
	assert.Equal(t, uuidExp, uuid.Uuid)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

}

func TestInsertTreeError(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	ent := repository.CreateTreeInput{
		X:        0,
		Y:        0,
		Height:   0,
		EstateId: "",
	}
	errExp := errors.New(gomock.Any().String())
	mock.ExpectQuery(`INSERT INTO tree \(x_axis, y_axis, height, estate_uuid\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING uuid`).
		WithArgs(ent.X, ent.Y, ent.Height, ent.EstateId).
		WillReturnError(errors.New(gomock.Any().String()))

	uuid, err := repo.InsertTree(context.Background(), ent)
	assert.Equal(t, errExp, err)
	assert.Zero(t, uuid)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestInsertTreeNormal(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	ent := repository.CreateTreeInput{
		X:        0,
		Y:        0,
		Height:   0,
		EstateId: "",
	}
	uuidExp := "1234"
	row := sqlmock.NewRows([]string{"uuid"}).AddRow(uuidExp)

	mock.ExpectQuery(`INSERT INTO tree \(x_axis, y_axis, height, estate_uuid\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING uuid`).
		WithArgs(ent.X, ent.Y, ent.Height, ent.EstateId).
		WillReturnRows(row)

	uuid, err := repo.InsertTree(context.Background(), ent)
	assert.Nil(t, err)
	assert.Equal(t, uuidExp, uuid.Uuid)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

}

func TestGetAllTreeNormal(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	uuidExp := "1234"
	input := repository.UuidInput{
		Uuid: uuidExp,
	}
	rows := sqlmock.NewRows([]string{"uuid", "x_axis", "y_axis", "height", "estate_uuid"})
	for _, tr := range treeData {
		rows.AddRow(tr.Uuid, tr.X, tr.Y, tr.Height, tr.EstateUuid)
	}

	mock.ExpectQuery(`SELECT uuid, x_axis, y_axis, height, estate_uuid FROM tree where estate_uuid = \$1`).
		WithArgs(input.Uuid).
		WillReturnRows(rows)

	allData, err := repo.GetAllTree(context.Background(), input)
	assert.Nil(t, err)
	assert.Equal(t, len(treeData), len(allData))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetAllTreeErrorScan(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	uuidExp := "1234"
	input := repository.UuidInput{
		Uuid: uuidExp,
	}
	rows := mock.NewRows([]string{"uuid", "x_axis", "y_axis", "height", "estate_uuid"}).AddRow(
		"76826522-c32b-4c72-8a2b-8c0231f80500", "FAIL", 45, 15, "471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6",
	)

	mock.ExpectQuery(`SELECT uuid, x_axis, y_axis, height, estate_uuid FROM tree where estate_uuid = \$1`).
		WithArgs(input.Uuid).
		WillReturnRows(rows)

	allData, err := repo.GetAllTree(context.Background(), input)
	assert.Error(t, err)
	assert.Equal(t, 0, len(allData))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetEstateStats(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	rows := sqlmock.NewRows([]string{"count", "max", "min", "median"})
	rows.AddRow(5, 30, 12, 20)

	uuidExp := "1234"
	input := repository.UuidInput{
		Uuid: uuidExp,
	}

	mock.ExpectQuery(`(?i)SELECT.*?COUNT.*?MAX.*?MIN.*?PERCENTILE_CONT.*?tree.*?estate_uuid.*
*
`).
		WithArgs(input.Uuid).
		WillReturnRows(rows)

	data, err := repo.GetEstateStats(context.Background(), input)
	assert.Nil(t, err)
	assert.Positive(t, data.Count)
	assert.Positive(t, data.Max)
	assert.Positive(t, data.Min)
	assert.Positive(t, data.Median)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetEstateStatsErrNoRows(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	rows := sqlmock.NewRows([]string{"count", "max", "min", "median"})
	uuidExp := "1234"
	input := repository.UuidInput{
		Uuid: uuidExp,
	}

	mock.ExpectQuery(`(?i)SELECT.*?COUNT.*?MAX.*?MIN.*?PERCENTILE_CONT.*?tree.*?estate_uuid.*
		*
		`).
		WithArgs(input.Uuid).
		WillReturnRows(rows)

	data, err := repo.GetEstateStats(context.Background(), input)
	assert.Nil(t, err)
	assert.Zero(t, data.Count)
	assert.Zero(t, data.Max)
	assert.Zero(t, data.Min)
	assert.Zero(t, data.Median)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetDetailEstateNormal(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	uuidExp := "1234"
	input := repository.UuidInput{
		Uuid: uuidExp,
	}
	rows := sqlmock.NewRows([]string{"uuid", "width", "length"})
	rows.AddRow("471eb2ad-207f-47f8-ab69-4ab4b6c8a8f6", 50, 100)

	mock.ExpectQuery(`SELECT .*? FROM estate where.*?uuid = \$1`).
		WithArgs(input.Uuid).
		WillReturnRows(rows)

	estate, err := repo.GetEstate(context.Background(), input)
	assert.Nil(t, err)
	assert.Equal(t, 100, estate.Length)
	assert.Equal(t, 50, estate.Width)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

}

func TestGetDetailEstateNotFound(t *testing.T) {
	repo, mock, err := Setup(t)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Db.Close()

	uuidExp := "1234"
	input := repository.UuidInput{
		Uuid: uuidExp,
	}
	rows := sqlmock.NewRows([]string{"uuid", "width", "length"})

	mock.ExpectQuery(`SELECT .*? FROM estate where.*?uuid = \$1`).
		WithArgs(input.Uuid).
		WillReturnRows(rows)

	estate, err := repo.GetEstate(context.Background(), input)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.Zero(t, estate)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}

}
