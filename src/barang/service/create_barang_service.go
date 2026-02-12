package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
	"gitlab.com/wit-id/service-hub-test/common/utility"
	"gitlab.com/wit-id/service-hub-test/src/repository/payload"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"

)

func (s *BarangService) CreateBarang(ctx context.Context, request sqlc.InsertBarangParams, p payload.InsertBarangPayload) (
	 barangGUID string, err error) {
		q := sqlc.New(s.mainDB)

		barangData, err := q.InsertBarang(ctx, request)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed to insert barang", err)
			err = errors.Wrap(httpservice.ErrBadRequestWithMessage,
				utility.ParseSqlError(err))
			return
		}

		barangGUID = barangData.Guid
		return 
}
