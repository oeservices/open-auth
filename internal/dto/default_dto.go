package dto

type DefaultDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PaginationMetaDataDto struct {
	Total int64 `json:"total"`
}

type PaginationDto[T any] struct {
	Data     []T                   `json:"data"`
	MetaData PaginationMetaDataDto `json:"metaData"`
}
