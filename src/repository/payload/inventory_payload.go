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

type InventoryCountEachProductResponse struct {
	BarangID  string `json:"barang_id"`
	KodeBarang string `json:"kode_barang"`
	NamaBarang string `json:"nama_barang"`
	Kategori   string `json:"kategori"`
	TotalStok  int32  `json:"total_stok"`
}

type ListInventoryPayload struct {
	Filter  ListInventoryFilterPayload `json:"filter"`
	Pagination
}

type ListInventoryFilterPayload struct {
	SetGuid bool   `json:"set_guid"`
	Guid    string `json:"guid"`
	SetBarangID bool   `json:"set_barang_id"`
	BarangID string `json:"barang_id"`
	SetStatus bool   `json:"set_status"`
	Status string `json:"status"`
}

type UpdateInventoryPayload struct {
	BarangID  string `json:"barang_id"`
	Jumlah     int32    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	Status    string `json:"status"`
}

type ListCountAllInventoryPayload struct {
	Filter ListCountAllInventoryFilterPayload `json:"filter"`
	Pagination
}

type ListCountAllInventoryFilterPayload struct {
	SetNamaBarang bool   `json:"set_nama_barang"`
	NamaBarang string `json:"nama_barang"`
	SetKategori bool   `json:"set_kategori"`
	Kategori   string `json:"kategori"`
}

type StockItemsByCategoryPayload struct {
	Filter StockItemsByCategoryFilterPayload `json:"filter"`
	Pagination
}

type StockItemsByCategoryFilterPayload struct {
	SetKategori bool   `json:"set_kategori"`
	Kategori   string `json:"kategori"`
}

type StockItemsByCategoryResponse struct {
	Kategori   string `json:"kategori"`
	TotalStock int `json:"total_stock"`
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

func (p *UpdateInventoryPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (p *StockItemsByCategoryPayload) Validate() (err error){
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

func (p *ListCountAllInventoryPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (params *ListInventoryPayload) ToEntity() sqlc.ListInventoryParams {

	return sqlc.ListInventoryParams{
		SetGuid: params.Filter.SetGuid,
		Guid:    params.Filter.Guid,
		SetBarangID: params.Filter.SetBarangID,
		BarangID:    params.Filter.BarangID,
		SetStatus:   params.Filter.SetStatus,
		Status: sql.NullString{
			String: params.Filter.Status,
			Valid:  params.Filter.Status != "",
		},
		OrderParam:   makeOrderParam(params.Order, params.Sort),
		OffsetPages:  makeOffset(params.Limit, params.Offset),
		LimitData:    limitWithDefault(params.Limit),
	}
}

func (p *UpdateInventoryPayload) ToEntity(guid string) sqlc.UpdateInventoryParams {
	return sqlc.UpdateInventoryParams{
		Guid:      guid,
		BarangID:  p.BarangID,
		Jumlah:     p.Jumlah,
		Keterangan: sql.NullString{String: p.Keterangan, Valid: true},
		Status:    p.Status,
	}
}

func (p *ListCountAllInventoryPayload) ToEntity() sqlc.ListCountAllInventoryEachProductParams {
	return sqlc.ListCountAllInventoryEachProductParams{
		SetNamaBarang: p.Filter.SetNamaBarang,
		NamaBarang: sql.NullString{
			String: p.Filter.NamaBarang,
			Valid: true,
		},
		SetKategori: p.Filter.SetKategori,
		Kategori: sql.NullString{
			String: p.Filter.Kategori,
			Valid: true,
		},
		OrderParam: makeOrderParam(p.Order, p.Sort),
		OffsetPages: makeOffset(p.Limit, p.Offset),
		LimitData: limitWithDefault(p.Limit),
	}
}

func (p *StockItemsByCategoryPayload) ToEntity() sqlc.StockItemsByCategoryParams {
	return sqlc.StockItemsByCategoryParams{
		SetKategori: p.Filter.SetKategori,
		Kategori: sql.NullString{
			String: p.Filter.Kategori,
			Valid: true,
		},
		OrderParam: makeOrderParam(p.Order, p.Sort),
		OffsetPages: makeOffset(p.Limit, p.Offset),
		LimitData: limitWithDefault(p.Limit),
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

func ToPayloadInventory(data sqlc.GetOneInventoryRow) (payload *InventoryResponse) {
	payload = &InventoryResponse{
		Guid:      data.Guid,
		BarangID:  data.BarangID,
		Jumlah:      data.Jumlah,
		Keterangan: data.Keterangan.String,
		Status:    data.Status,
		CreateAt:  data.CreatedAt,
		UpdateAt:  data.UpdatedAt,
	}
	return
}

func ToPayloadCountAllInventoryEachProduct(listData []sqlc.ListCountAllInventoryEachProductRow,
	) (payload []*InventoryCountEachProductResponse) {

    payload = make([]*InventoryCountEachProductResponse, len(listData))

    for i, data := range listData {
        payload[i] = &InventoryCountEachProductResponse{
            BarangID:  data.BarangID,
			KodeBarang: data.KodeBarang,
			NamaBarang: data.NamaBarang,
			Kategori:   data.Kategori,
			TotalStok:  data.TotalStok,
        }
    }

    return
}

func ToPayloadStockItemsByCategory(listData []sqlc.StockItemsByCategoryRow) []*StockItemsByCategoryResponse {
    payload := make([]*StockItemsByCategoryResponse, len(listData))

    for i, data := range listData {
        payload[i] = &StockItemsByCategoryResponse{
            Kategori: data.Kategori,
			TotalStock: int(data.TotalStok),
        }
    }

    return payload
}




