package usecase

import (
	"context"
	"errors"

	"github.com/firerplayer/roda-belem-service/internal/infra/filters"
)

type SaveFilterUseCase struct {
	Filter *filters.CuckooFilter
}

func NewSaveFilterUseCase(filter *filters.CuckooFilter) *SaveFilterUseCase {
	return &SaveFilterUseCase{
		Filter: filter,
	}
}

func (uc *SaveFilterUseCase) Execute(ctx context.Context) error {
	err := uc.Filter.SaveCuckooFilter()
	if err != nil {
		return errors.New("Failed to save filter -> " + err.Error())
	}
	return nil
}
