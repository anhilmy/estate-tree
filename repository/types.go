// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type CreateEstateInput struct {
	Length int
	Width  int
}

type UuidInput struct {
	Uuid string
}
type UuidOutput struct {
	Uuid string
}

type CreateTreeInput struct {
	X        int
	Y        int
	Height   int
	EstateId string
}

type TreeModel struct {
	Uuid       string
	X          int
	Y          int
	Height     int
	EstateUuid string
}

type EstateStatsOutput struct {
	Count  int
	Max    int
	Min    int
	Median int
}

type EstateModel struct {
	Uuid   string
	Length int
	Width  int
}

type GetTreeByCoordinateInput struct {
	X          int
	Y          int
	EstateUuid string
}
