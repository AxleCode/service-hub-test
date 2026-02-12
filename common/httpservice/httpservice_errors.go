package httpservice

import (
	"errors"

	"gitlab.com/wit-id/service-hub-test/common/constants"
)

// error message.
var (
	ErrInternalServerError = errors.New("internal server error")

	ErrBadRequest                     = errors.New("bad request payload")
	ErrInvalidAppKey                  = errors.New("invalid app key")
	ErrUnknownSource                  = errors.New("unknown error")
	ErrEmailProviderError             = errors.New("email provider error")
	ErrPaymentNotFound                = errors.New("payment not found")
	ErrUnknownSourceWithMessage       = errors.New("unknown error with message")
	ErrBadRequestWithMessage          = errors.New("error with message")
	ErrInternalServerErrorWithMessage = errors.New("internal server error with message")

	ErrMissingHeaderData = errors.New("missing header data")

	ErrInvalidToken            = errors.New("invalid token")
	ErrUnauthorizedTokenData   = errors.New("unauthorized token data")
	ErrInvalidOTP              = errors.New("invalid otp")
	ErrInvalidOTPToken         = errors.New("invalid otp token")
	ErrInvalidPhoneNumberOTP   = errors.New("invalid phone number for this otp")
	ErrPasswordNotMatch        = errors.New("password not match")
	ErrReferralCodeNotMatch    = errors.New("referral code not match")
	ErrConfirmPasswordNotMatch = errors.New("confirm password not match")
	ErrNoResultData            = errors.New("no result data")
	ErrNoResultDataWithMessage = errors.New("no result data with message")

	ErrUserAlreadyRegistered = errors.New("user is already registered")
	ErrUserNotFound          = errors.New("user not found")
	ErrUnauthorizedUser      = errors.New("unauthorized user")
	ErrInActiveUser          = errors.New("user not active")
	SuccessChangePassword    = errors.New("successfully changed password")

	ErrRoleNotFound = errors.New("role not found")
	ErrDataNotFound = errors.New("data not found")

	ErrInvalidromotionCode         = errors.New("invalid promotion code")
	ErrInsufficientQuantityVoucher = errors.New("insufficient quantities of voucher")
	ErrVoucherIsNotActive          = errors.New("voucher is not active")
	ErrVoucherIsExpired            = errors.New("voucher is expired")

	ErrNoDoctor    = errors.New("doctor is not found")
	ErrNoSchedules = errors.New("schedule is not found")

	ErrUserAlreadyCheckIn       = errors.New("this user is checked in already")
	ErrAttemptIsReachedMax      = errors.New("attempt is reached maximum")
	ErrAttemptIsAlreadyRecorded = errors.New("attempt is taken already")
	ErrAssesmentCantUpdate      = errors.New("assesment cant update while already taken")

	ErrInvalidPaymentID = errors.New("invalid payment id")

	ErrNIKAlreadyExist         = errors.New("nik already exist")
	ErrIDCardAlreadyExist      = errors.New("id card already exist")
	ErrNPWPAlreadyExist        = errors.New("npwp already exist")
	ErrEmailAlreadyExist       = errors.New("email already exist")
	ErrEmailNotFound           = errors.New("email not found")
	ErrPhoneNumberAlreadyExist = errors.New("phone number already exist")
	ErrOTPDoesntMatch          = errors.New("OTP doesn't match")

	//Reservation
	ErrArrivalDateNil      = errors.New("please select a check-in date to proceed with your booking")
	ErrArrivalDateFormat   = errors.New("the check-in date format is incorrect")
	ErrArrivalDatePast     = errors.New("check-in date cannot be in the past")
	ErrDepartureDateNil    = errors.New("please select a check-out date to continue")
	ErrDepartureDateFormat = errors.New("the check-out date format is incorrect")
	ErrDepartureDateAfter  = errors.New("the check-out date must be after the check-in date")
	ErrAddressNil          = errors.New("address cannot be blank. Please enter your address to proceed")
	ErrEmailNil            = errors.New("email address cannot be empty. Please enter a valid email address")
	ErrEmailFormat         = errors.New("the email format is incorrect. Please enter a valid email address")
	ErrFirstnameNil        = errors.New("first name cannot be blank. Please enter your first name")
	ErrLastnameNil         = errors.New("last name cannot be blank. Please enter your last name")
	ErrTelephoneNil        = errors.New("telephone number cannot be blank. Please enter your telephone number")
	ErrTelephoneFormat     = errors.New("the telephone number format is incorrect. Please enter a valid phone number")
	ErrRoomOutOfStock      = errors.New("we apologize, the room you selected is out of stock.")
	ErrSalutationsNil      = errors.New("please select a salutation to proceed")

	//List Room
	ErrPropertyNotFound = errors.New("we couldn't find the property you're looking for. Please check the property details or try searching for a different one")
	ErrDateToNil        = errors.New("please select a check-out date to continue")
	ErrSameDate         = errors.New("your check-in and check-out dates cannot be the same please select different dates")
	ErrDateToFalse      = errors.New("the check-in date must be before the check-out date. Please adjust the dates accordingly")
	ErrDateRangeTooLong = errors.New("bookings can be made for a maximum of 31 days. Please shorten your stay duration")
	ErrDateFromPast     = errors.New("you cannot select a past date")
	ErrCountryNil       = errors.New("please select a country to proceed with your booking")
	ErrAdultMax         = errors.New("adult cannot be more than 2")
	ErrAdultMin         = errors.New("adults should not be less than the number of rooms")
	ErrChildMax         = errors.New("child cannot be more than 1")
	ErrChildrenMax      = errors.New("child cannot be more than 1")
	ErrCountryNull      = errors.New("please select a country to proceed with your booking")

	//Checkout
	ErrPropertyNil      = errors.New("please select a property to proceed with your booking")
	ErrPropertyStaahNil = errors.New("please select a property staah to proceed with your booking")
	ErrBookingDetailNil = errors.New("please select a booking detail to proceed with your booking")
	ErrRoomNil          = errors.New("please select a room to proceed with your booking")
	ErrRoomNotFound     = errors.New("we couldn't find the room id you're looking for. Please check the room id or try searching for a different one")
	ErrRateNotFound     = errors.New("we couldn't find the rate id you're looking for. Please check the rate id or try searching for a different one")
	ErrRatesNil         = errors.New("please select a rate id to proceed with your booking")
	ErrRatesNull        = errors.New("please select a rate to proceed with your booking")

	//Powerpro sync
	ErrPowerproServerDown      = errors.New("Couldn't access to powerpro server at the moment, please try again later")
	ErrPowerproContextExceeded = errors.New("Context exceeded for your request, please contact administrator")
)

// error message.
var (
	MsgHeaderTokenNotFound = constants.ErrorResponse{
		ID: "Header `token` tidak ditemukan",
		EN: "Header `token` not found",
	}

	MsgHeaderRefreshTokenNotFound = constants.ErrorResponse{
		ID: "Header `refresh-token` tidak ditemukan",
		EN: "Header `refresh-token` not found",
	}

	MsgHeaderTokenUnauthorized = constants.ErrorResponse{
		ID: "Token tidak sah",
		EN: "Unauthorized token",
	}

	MsgHeaderRefreshTokenUnauthorized = constants.ErrorResponse{
		ID: "Refresh token tidak sah",
		EN: "Unauthorized refresh token",
	}

	MsgIsNotLogin = constants.ErrorResponse{
		ID: "Silahkan masuk terlebih dahulu",
		EN: "Please login first",
	}

	MsgUnauthorizedUser = constants.ErrorResponse{
		ID: "Pengguna tidak sah",
		EN: "Unauthorized user",
	}

	MsgUserNotActive = constants.ErrorResponse{
		ID: "User tidak aktif",
		EN: "User not active",
	}

	MsgInvalidIDParam = constants.ErrorResponse{
		ID: "parameter id tidak valid",
		EN: "invalid id parameter",
	}

	MsgInvalidCheckIN = constants.ErrorResponse{
		ID: "user sudah check in hari ini",
		EN: "this user is checked in already",
	}

	MsgAttemptReachedMax = constants.ErrorResponse{
		ID: "user untuk attempt module ini sudah melampaui batas",
		EN: "user for this module Attempt Reached Max",
	}

	MsgAttemptIsTaken = constants.ErrorResponse{
		ID: "user untuk attempt module ini sudah tersimpan di database",
		EN: "user for this module Attempt is taken already",
	}

	MsgCantEditAssesmentWhichAlreadyTaken = constants.ErrorResponse{
		ID: "user tidak dapat mengubah assesment ketika assesment sudah di gunakan",
		EN: "user cant update this assesment while this assesment already taken",
	}

	MsgPowerproServerDown = constants.ErrorResponse{
		ID: "Saat ini server Powerpro Sedang tidak bisa di akses, mohon di coba lagi beberapa saat",
		EN: "Couldn't access to powerpro server at the moment, please try again later",
	}
)
