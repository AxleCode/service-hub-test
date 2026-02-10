package payload

import (

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
)

type CheckWeatherPayload struct {
	Latitude  string   `json:"latitude" validate:"required"`
	Longitude string   `json:"longitude" validate:"required"`
}

type GetCheckWeatherPayload struct {
	Key string `json:"key" validate:"required"`
	Q   string `json:"q" validate:"required"`
}

func (p *CheckWeatherPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(p); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}

	return
}

