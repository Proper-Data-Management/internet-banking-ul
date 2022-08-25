package entities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mak-alex/al_hilal_core/modules/logger"
	"go.uber.org/zap"
)

type BasePaginationFilters struct {
	BaseFilter
}

type BaseFilter struct {
	Sort       string `json:"sort"`
	Order      string `json:"order"`
	SearchText string `json:"searchText"`
	Size       uint64 `json:"size"`
	Page       uint64 `json:"page"`
}

func (f *BaseFilter) GetSort() string {
	return f.Sort
}

func (f *BaseFilter) GetOrder() string {
	return f.Order
}

func (f *BaseFilter) GetSearchText() string {
	return f.SearchText
}

func (f *BaseFilter) GetOffset() uint64 {
	return f.Page * f.Size
}

func (f *BaseFilter) GetSize() uint64 {
	return f.Size
}

func NewBaseFilterFromQuery(ctx *fiber.Ctx) (*BasePaginationFilters, error) {
	baseFilter := new(BasePaginationFilters)

	if err := ctx.QueryParser(baseFilter); err != nil {
		logger.WorkLogger.Error("Error parse BasePaginationFilters", zap.Error(err))

		return nil, err
	}

	if baseFilter.Page <= 0 {
		baseFilter.Page = 0
	}

	if baseFilter.Size <= 0 {
		baseFilter.Size = 10
	}

	return baseFilter, nil
}
