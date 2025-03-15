package distance

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type DistanceService interface {
	CalculateDistanceStockCity(ctx context.Context, req models.CalculateDistanceStockCityReq) (res models.CalculateDistanceStockCityRes, errRes httpcommon.ErrorDTO)
}
