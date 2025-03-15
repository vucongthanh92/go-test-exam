package distance

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
	"github.com/vucongthanh92/go-test-exam/internal/repository/external"
)

type DistanceImpl struct {
}

func NewDistanceService() DistanceService {
	return &DistanceImpl{}
}

func (s *DistanceImpl) CalculateDistanceStockCity(ctx context.Context, req models.CalculateDistanceStockCityReq) (
	res models.CalculateDistanceStockCityRes, errRes httpcommon.ErrorDTO) {

	src, err := external.GetCoordinatesFromIP(req.ClientID)
	if err != nil {
		return res, errRes
	}

	dst, err := external.GetCoordinatesFromCity(req.StockCity)
	if err != nil {
		return res, errRes
	}

	distance := external.CalculateDistance(src, dst)

	res.StockCity = req.StockCity
	res.Distance = distance
	res.Unit = "km"

	return res, errRes
}
