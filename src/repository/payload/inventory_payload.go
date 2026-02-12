package payload

import (
	"time"
	"database/sql"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/common/utility"
)

type ReadInventory struct {
	Guid      string `json:"guid" validate:"required,uuid"`
	BarangID  string `json:"barang_id"`
	Jumlah    int32    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	Status    string `json:"status"`
	IsDeleted bool `json:"is_deleted"`
	CreateAt  time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type InsertInventoryPayload struct {
	BarangID string `json:"barang_id"`
	Jumlah     int32    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	Status    string `json:"status"`
}

type InventoryResponse struct {
	Guid      string `json:"guid"`
	BarangID  string `json:"barang_id"`
	Jumlah      int32    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	Status    string `json:"status"`
	CreateAt  time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type ListInventoryPayload struct {
	Filter  ListInventoryFilterPayload `json:"filter"`
	Pagination
}

type ListInventoryFilterPayload struct {
	BarangID string `json:"barang_id"`
	Jumlah   int32    `json:"jumlah"`
	Status   string `json:"status"`
}

func (p *InsertInventoryPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (p *ListInventoryPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func(payload *InsertInventoryPayload) ToEntity() (data sqlc.InsertInventoryParams) {
	data = sqlc.InsertInventoryParams{
		Guid:      utility.GenerateGoogleUUID(),
		BarangID:  payload.BarangID,
		Jumlah:     payload.Jumlah,
		Keterangan: sql.NullString{String: payload.Keterangan, Valid: true},
		Status:    payload.Status,
	}
	return
}

func (p *ListInventoryFilterPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (params *ListInventoryPayload) ToEntity() sqlc.ListInventoryParams {

	return sqlc.ListInventoryParams{
		BarangID:    params.Filter.BarangID,
		Jumlah:    params.Filter.Jumlah,
		Status:    params.Filter.Status,
		OrderParam:   makeOrderParam(params.Order, params.Sort),
		OffsetPages:  makeOffset(params.Limit, params.Offset),
		LimitData:    limitWithDefault(params.Limit),
	}
}

func ToPayloadListInventory(listData []sqlc.ListInventoryRow) (payload []*ReadInventory) {
	payload = make([]*ReadInventory, len(listData))
	for i, data := range listData {
		payload[i] = &ReadInventory{
			Guid:      data.Guid,
			BarangID:  data.BarangID,
			Jumlah:      data.Jumlah,
			Keterangan: data.Keterangan.String,
			Status:    data.Status,
			CreateAt:  data.CreatedAt,
			UpdateAt:  data.UpdatedAt,
		}
	}
	return
}