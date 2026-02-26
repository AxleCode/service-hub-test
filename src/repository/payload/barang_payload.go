package payload

import (
	"database/sql"
	"time"
	
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
	"gitlab.com/wit-id/service-hub-test/common/utility"
)

type ReadBarang struct {
	Guid 	   string `json:"guid" validate:"required,uuid"`
	KodeBarang string `json:"kode_barang"`
	NamaBarang string `json:"nama_barang"`
	Deskripsi  string `json:"deskripsi"`
	Kategori   string `json:"kategori"`
	Harga      int32 `json:"harga"`
	CreateAt   time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"updated_at"`
}

type InsertBarangPayload struct {
	KodeBarang string `json:"kode_barang"`
	NamaBarang string `json:"nama_barang"`
	Kategori   string `json:"kategori"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int32    `json:"harga"`
}

type BarangResponse struct {
	Guid       string `json:"guid"`
	KodeBarang string `json:"kode_barang"`
	NamaBarang string `json:"nama_barang"`
	Kategori   string `json:"kategori"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int32 `json:"harga"`
}

type ListFilterBarangPayload struct {
	SetGuid bool   `json:"set_guid"`
	Guid    string `json:"guid"`
	SetKodeBarang bool   `json:"set_kode_barang"`
	KodeBarang string `json:"kode_barang"`
	SetNamaBarang bool   `json:"set_nama_barang"`
	NamaBarang string `json:"nama_barang"`
    SetKategori bool   `json:"set_kategori"`
    Kategori    string `json:"kategori"`

}

type ListBarangPayload struct {
    Filter   ListFilterBarangPayload `json:"filter"`
    Pagination
}

type ListBarangFilterPayload struct {
	SetGuid bool   `json:"set_guid"`
	Guid    string `json:"guid"`
	SetKodeBarang bool   `json:"set_kode_barang"`
	KodeBarang string `json:"kode_barang"`
	SetNamaBarang bool   `json:"set_nama_barang"`
	NamaBarang string `json:"nama_barang"`
    SetKategori bool   `json:"set_kategori"`
    Kategori    string `json:"kategori"`
}

type CountListBarangParams struct {
	Kategori sql.NullString
	Harga    sql.NullFloat64
}

type UpdateBarangPayload struct {
	KodeBarang string `json:"kode_barang"`
	NamaBarang string `json:"nama_barang"`
	Kategori   string `json:"kategori"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int32    `json:"harga"`
}

func (p *InsertBarangPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (p *ListBarangPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (p *ListBarangFilterPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (p *UpdateBarangPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func(payload *InsertBarangPayload) ToEntity() (data sqlc.InsertBarangParams) {
	data = sqlc.InsertBarangParams{
		Guid:           utility.GenerateGoogleUUID(),
		KodeBarang: payload.KodeBarang,
		NamaBarang: payload.NamaBarang,
		Deskripsi: sql.NullString{
			String: payload.Deskripsi,
			Valid:  payload.Deskripsi != "",
		},
		Kategori: payload.Kategori,
		Harga: payload.Harga,
	}
	return
}

func (params *ListBarangPayload) ToEntity() sqlc.ListBarangParams {
	return sqlc.ListBarangParams{
		SetGuid: params.Filter.SetGuid,
		Guid:    params.Filter.Guid,
		SetKodeBarang: params.Filter.SetKodeBarang,
		KodeBarang: sql.NullString{
			String: params.Filter.KodeBarang,
			Valid:  params.Filter.SetKodeBarang,
		},
		SetNamaBarang: params.Filter.SetNamaBarang,
		NamaBarang: sql.NullString{
			String: params.Filter.NamaBarang,
			Valid:  params.Filter.SetNamaBarang,
		},
		SetKategori:  params.Filter.SetKategori,
		Kategori: sql.NullString{
			String: params.Filter.Kategori,
			Valid:  params.Filter.SetKategori,
		},
		OrderParam:   makeOrderParam(params.Order, params.Sort),
		OffsetPages:  makeOffset(params.Limit, params.Offset),
		LimitData:    limitWithDefault(params.Limit),
	}
}

func ToPayloadListBarang(listData []sqlc.ListBarangRow) (payload []*ReadBarang) {
	payload = make([]*ReadBarang, len(listData))
	for i := range listData {
		payload[i] = new(ReadBarang)

		data := ToPayloadBarang (sqlc.GetOneBarangRow{
			Guid:       listData[i].Guid,
			KodeBarang: listData[i].KodeBarang,
			NamaBarang: listData[i].NamaBarang,
			Deskripsi:  listData[i].Deskripsi,
			Kategori:   listData[i].Kategori,
			Harga:      listData[i].Harga,
			CreatedAt:  listData[i].CreatedAt,
			UpdatedAt:  listData[i].UpdatedAt,
		})
		payload[i] = &data
	}
	return
}

func ToPayloadBarang(barang sqlc.GetOneBarangRow) (payload ReadBarang) {
	payload = ReadBarang{
		Guid:       barang.Guid,
		KodeBarang: barang.KodeBarang,
		NamaBarang: barang.NamaBarang,
		Deskripsi:  barang.Deskripsi.String,
		Kategori:   barang.Kategori,
		Harga:      barang.Harga,
		CreateAt:   barang.CreatedAt,
		UpdateAt:   barang.UpdatedAt,
	}
	return
}

func (p *UpdateBarangPayload) ToEntity(guid string) (data sqlc.UpdateBarangParams) {
	data = sqlc.UpdateBarangParams{
		Guid:           guid,
		KodeBarang: p.KodeBarang,
		NamaBarang: p.NamaBarang,
		Deskripsi: sql.NullString{
			String: p.Deskripsi,
			Valid:  p.Deskripsi != "",
		},
		Kategori: p.Kategori,
		Harga: p.Harga,
	}
	return
}
