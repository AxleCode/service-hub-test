package utility

import "encoding/json"

const (
	defaultSuperAdmin = "superadmin"
)

func DefaultActorParser(actorID string) string {
	if actorID == "" {
		return defaultSuperAdmin
	}

	return actorID
}

func IsAllLanguageParser(isAllLanguage string) bool {
	switch isAllLanguage {
	case "1":
		return true
	case "0":
		return false
	default:
		return false
	}
}

func PayloadToJsonRawMessageParser(payload interface{}) json.RawMessage {
	// Serialize the struct to JSON
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Create a json.RawMessage from the JSON byte slice
	rawMessage := json.RawMessage(jsonBytes)

	return rawMessage
}
