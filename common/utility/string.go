package utility

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"gitlab.com/wit-id/service-hub-test/common/constants"

	"github.com/dustin/go-humanize"
)

// PrettyPrint will transform struct data as json string for nicer log ...
func PrettyPrint(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(JSON)
}

// PrettyPrintWithoutIndent will transform struct data as json string for nicer log ...
func PrettyPrintWithoutIndent(data interface{}) string {
	JSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(JSON)
}

// FormatPhoneNumber default for prefix ID.
func FormatPhoneNumber(phoneNumber string) (response string) {
	prefix := phoneNumber[0:3]
	if prefix != "+62" {
		phoneNumber = strings.Replace(phoneNumber, "0", "+62", 1)
	}

	response = phoneNumber

	return
}

// FormatRupiah to convert number 10000 to Rp 10.000.
func FormatRupiah(amount int) string {
	// convert to positive if number negative
	amountConv := math.Abs(float64(amount))
	humanizeValue := humanize.Comma(int64(amountConv))
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)

	return "Rp " + stringValue
}

// GenerateOrderNumber will generate number order ex: 000000001.
func GenerateOrderNumber(id, length int) (idString string) {
	idString = strconv.Itoa(id)

	numOfZero := length - len(idString)

	for i := 1; i <= numOfZero; i++ {
		idString = "0" + idString
	}

	return
}

func GenerateInvoices(prefix string, id, length int) (invoices string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc)

	invoices = prefix + "/" + timeNow.Format("20060102") + "/" + GenerateOrderNumber(id, length)

	return
}

func RandomNumber(low, hi int) int {
	return low + rand.Intn(hi-low)
}

// GenerateSlug will change title to slug format
func GenerateSlug(title string) string {
	// Convert to lowercase and replace all non-alphanumeric characters with hyphen
	slug := regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(strings.ToLower(title), "-")

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	slug = slug + "-" + strconv.Itoa(RandomNumber(1, 999))

	return slug
}

func GenerateSlugARMS(number int32) string {

	str := fmt.Sprintf("ARMS%05d", number)

	return str
}

func PropertyLevel(levelString string) (levelNumber string) {
	if levelString == constants.PropertyLevelCollection {
		levelNumber = constants.EnumPropertyLevelCollection
	}
	if levelString == constants.PropertyLevelEconomy {
		levelNumber = constants.EnumPropertyLevelEconomy
	}
	if levelString == constants.PropertyLevelMidScale {
		levelNumber = constants.EnumPropertyLevelMidScale
	}
	if levelString == constants.PropertyLevelLifestyle {
		levelNumber = constants.EnumPropertyLevelLifestyle
	}
	if levelString == constants.PropertyLevelUpScale {
		levelNumber = constants.EnumPropertyLevelUpScale
	}
	if levelString == constants.PropertyLevelFoodBeverages {
		levelNumber = constants.EnumPropertyLevelFoodBeverages
	}

	return
}

// ActivityBodyMessage Create activity log body message
func ActivityBodyMessage(activityType string, mapData map[string]string) string {
	var message string

	switch activityType {
	case constants.ActivityAddBoardTask:
		message = fmt.Sprintf("%s added %s to %s", mapData[constants.ActivityLoggingFullName],
			mapData[constants.ActivityLoggingTaskTitle],
			mapData[constants.ActivityLoggingColumnTitle])
	case constants.ActivityArchiveBoardTask:
		message = fmt.Sprintf("%s archived %s", mapData[constants.ActivityLoggingFullName],
			mapData[constants.ActivityLoggingTaskTitle])
	case constants.ActivityMoveBoardTask:
		message = fmt.Sprintf("%s moved %s from %s to %s", mapData[constants.ActivityLoggingFullName],
			mapData[constants.ActivityLoggingTaskTitle],
			mapData[constants.ActivityLoggingColumnTitle],
			mapData[constants.ActivityLoggingColumnTitle2])
	}

	return message
}

func stringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return stringWithCharset(length, constants.Charset)
}

// FormatPhoneNumber default for prefix ID.
func FormatAllPhoneNumber(phoneNumber string) (response string) {
	// Remove any non-digit characters (except '+')
	cleanedPhoneNumber := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '+' {
			return r
		}
		return -1
	}, phoneNumber)

	// Check if the number starts with a '+' (international format)
	if strings.HasPrefix(cleanedPhoneNumber, "+") {
		response = cleanedPhoneNumber
	} else {
		// Default to adding Indonesia's country code (+62) if no international prefix
		if strings.HasPrefix(cleanedPhoneNumber, "0") {
			// Replace leading '0' with '+62'
			cleanedPhoneNumber = strings.Replace(cleanedPhoneNumber, "0", "+62", 1)
		}
		response = cleanedPhoneNumber
	}

	return response
}

func DetermineIdentifierType(value string) string {
	var identifierType string
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex := `^(?:\+628|08)[0-9]{8,}$`

	if match, _ := regexp.MatchString(emailRegex, value); match {
		identifierType = constants.IdentifierTypeEmail
	} else if match, _ := regexp.MatchString(phoneRegex, value); match {
		identifierType = constants.IdentifierTypePhone
	} else {
		identifierType = constants.IdentifierTypeEmployeeNumber
	}

	return identifierType
}

func FormatRupiahString(value string) string {
	floatValue, _ := strconv.ParseFloat(value, 64)

	intValue := int64(floatValue)

	formatted := fmt.Sprintf("%d", intValue)

	var result strings.Builder
	length := len(formatted)
	for i, char := range formatted {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(char)
	}

	return "Rp " + result.String()
}

func FormatRupiahFromFloat(value float64) string {
	intValue := int64(value)

	formatted := fmt.Sprintf("%d", intValue)

	var result strings.Builder
	length := len(formatted)
	for i, char := range formatted {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(char)
	}

	return "Rp " + result.String()
}

func FormatThousandString(value string) string {
	floatValue, _ := strconv.ParseFloat(value, 64)

	intValue := int64(floatValue)

	formatted := fmt.Sprintf("%d", intValue)

	var result strings.Builder
	length := len(formatted)
	for i, char := range formatted {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(char)
	}

	return result.String()
}

func ToFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return f
		}
	}
	return 0
}
