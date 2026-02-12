package payload

import (
	// "database/sql"
	"fmt"
	"regexp"
	"strings"

	// "github.com/labstack/echo/v4"
	// "gitlab.com/wit-id/service-hub-test/common/constants"
	// sqlc "gitlab.com/wit-id/service-hub-test/src/repository/pgbo_sqlc"
)

const (
	defaultLimit      = 10
	defaultOrderValue = "created_at DESC"
)

// type MddwUserData struct {
// 	ReadAuthenticationFCM
// }

type CommonLanguagePayload struct {
	EN string `json:"en"`
	ID string `json:"id"`
}

func limitWithDefault(limit int32) int32 {
	if limit <= 0 {
		return defaultLimit
	}

	return limit
}

func makeOffset(limit, offset int32) int32 {

	if offset == 0 {
		return (1 * limit) - limit
	} else {
		return (offset * limit) - limit
	}
}

func makeOrderParam(orderBy, sort string) string {
	if orderBy == "" || sort == "" {
		return defaultOrderValue
	}

	return fmt.Sprintf(strings.ToLower("%s %s"), orderBy, sort)
}

func queryStringLike(param string) string {
	return "%" + param + "%"
}

func isNumberAndDot(str string) bool {
	re := regexp.MustCompile(`^[0-9.]+$`)
	return re.MatchString(str)
}

func containsAlphabet(str string) bool {
	re := regexp.MustCompile(`[a-zA-Z]`)
	return re.MatchString(str)
}

type CommonSubResponsePayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Pagination struct {
	Limit  int32  `json:"limit" valid:"required"`
	Offset int32  `json:"page" valid:"required"`
	Order  string `json:"order" valid:"required"`
	Sort   string `json:"sort" valid:"required"` // ASC, DESC
}

type ResponseData struct {
	Code      string      `json:"code"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
	MessageEn string      `json:"message_en"`
	MessageID string      `json:"message_id"`
}

type CommonPayload struct {
	GuID string `json:"guid" param:"guid"`
	Name string `json:"name" param:"name"`
}

// func GetActor(ctx echo.Context) MddwUserData {
// 	actor, _ := ctx.Get(constants.MddwUserBackoffice).(sqlc.GetAuthenticationByIDRow)

// 	return MddwUserData{
// 		ReadAuthenticationFCM: ReadAuthenticationFCM{
// 			Guid:             actor.Guid,
// 			EmployeeGuid:     actor.EmployeeGuid.String,
// 			Username:         actor.Username,
// 			EmployeeFullname: actor.EmployeeFullname.String,
// 			Email:            actor.Email.String,
// 			UserFCMToken:     actor.FcmToken.String,
// 			Status:           actor.Status,
// 			CreatedAt:        actor.CreatedAt.Time,
// 			CreatedBy:        actor.CreatedBy,
// 		},
// 	}
// }

// func ToTypeDeductionPayment(s string) (sqlc.TypeDeductionPayment, error) {
// 	switch s {
// 	case string(sqlc.TypeDeductionPaymentNONPAYMENT):
// 		return sqlc.TypeDeductionPaymentNONPAYMENT, nil
// 	case string(sqlc.TypeDeductionPaymentVIRTUALACCOUNT):
// 		return sqlc.TypeDeductionPaymentVIRTUALACCOUNT, nil
// 	case string(sqlc.TypeDeductionPaymentQRIS):
// 		return sqlc.TypeDeductionPaymentQRIS, nil
// 	case string(sqlc.TypeDeductionPaymentEWALLET):
// 		return sqlc.TypeDeductionPaymentEWALLET, nil
// 	case string(sqlc.TypeDeductionPaymentCARDS):
// 		return sqlc.TypeDeductionPaymentCARDS, nil
// 	default:
// 		return "", fmt.Errorf("invalid TypeDeductionPayment: %s", s)
// 	}
// }

// func ToEntityDeleteDeduction(key string) (data sqlc.UpdateDeductionStatusParams) {
// 	data = sqlc.UpdateDeductionStatusParams{
// 		Status:    constants.StatusDeleted,
// 		UpdatedBy: sql.NullString{String: constants.CreatedByTemporaryBySystem, Valid: true},
// 		Guid:      key,
// 	}

// 	return
// }
