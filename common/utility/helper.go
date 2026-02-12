package utility

// import (
// 	"crypto/sha256"
// 	"fmt"
// 	"strings"

// 	"github.com/pkg/errors"
// 	"gitlab.com/wit-id/service-hub-test/common/constants"
// )

// func ParsePhoneNumber(phoneNumber string) string {
// 	if strings.HasPrefix(phoneNumber, "0") {
// 		return "+62" + phoneNumber[1:]
// 	}
// 	if strings.HasPrefix(phoneNumber, "+62") {
// 		return phoneNumber
// 	}
// 	return phoneNumber
// }

// func WriteStringTemplate(stringTemplate string, args ...interface{}) string {
// 	return fmt.Sprintf(stringTemplate, args...)
// }

// // SHA256 is a function to hash a string using SHA256
// func SHA256(text string) string {
// 	hash := sha256.New()
// 	hash.Write([]byte(text))
// 	return fmt.Sprintf("%x", hash.Sum(nil))
// }

// // ParseGenderToEnumXendit is a function to parse gender from enums database to enums xendit
// func ParseSalutationToEnumXendit(salutation string) (genderEnum string, err error) {
// 	switch salutation {
// 	case "Mr.":
// 		genderEnum = constants.GenderMale
// 	case "Mrs.":
// 		genderEnum = constants.GenderFemale
// 	case "Ms.":
// 		genderEnum = constants.GenderFemale
// 	default:
// 		return "", errors.New("Gender is not valid")
// 	}

// 	return
// }

// // ParseGenderToEnumXendit is a function to parse gender from enums database to enums xendit
// func ParseGenderToEnumXendit(gender string) (genderEnum string, err error) {
// 	switch gender {
// 	case constants.GenderLakiLaki:
// 		genderEnum = constants.GenderMale
// 	case constants.GenderPerempuan:
// 		genderEnum = constants.GenderFemale
// 	default:
// 		return "", errors.New("Gender is not valid")
// 	}

// 	return
// }

// // ParseCountryToCountryCode is a function to parse country to country code (ISO 3166-2 alpha-2)
// // This function is used in service of xendit to parse country to country code
// func ParseCountryToCountryCode(country string) string {
// 	countryMap := map[string]string{
// 		"UNITED STATES":  "US",
// 		"CANADA":         "CA",
// 		"UNITED KINGDOM": "GB",
// 		"GERMANY":        "DE",
// 		"FRANCE":         "FR",
// 		"ITALY":          "IT",
// 		"SPAIN":          "ES",
// 		"AUSTRALIA":      "AU",
// 		"JAPAN":          "JP",
// 		"CHINA":          "CN",
// 		"INDIA":          "IN",
// 		"INDONESIA":      "ID",
// 		"BRAZIL":         "BR",
// 		"MEXICO":         "MX",
// 		"SOUTH AFRICA":   "ZA",
// 		"NIGERIA":        "NG",
// 	}

// 	if code, exists := countryMap[country]; exists {
// 		return code
// 	}
// 	return ""
// }
