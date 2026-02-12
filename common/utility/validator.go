package utility

import (
	"context"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/service-hub-test/common/httpservice"
	"gitlab.com/wit-id/service-hub-test/toolkit/log"
)

func ValidateStruct(ctx context.Context, req interface{}) (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(req); err != nil {
		log.FromCtx(ctx).Error(err, "failed parsing payload")
		err = errors.WithStack(httpservice.ErrBadRequest)
		return
	}

	return
}

func ValidateExpiredCard(ctx context.Context, month, year string) error {
	currentYear, currentMonth := time.Now().Year(), int(time.Now().Month())

	// Convert month and year to integers
	expiredMonth, err := strconv.Atoi(month)
	if err != nil {
		return errors.New("invalid month format")
	}
	expiredYear, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("invalid year format")
	}

	// Check if the card is expired
	if expiredYear < currentYear || (expiredYear == currentYear && expiredMonth < currentMonth) {
		return errors.New("card is expired")
	}

	return nil
}
