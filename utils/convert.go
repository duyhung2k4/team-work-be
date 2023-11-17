package utils

import (
	"strings"
	"team-work-be/model"
	"unicode"
)

func ConvertNameField(data map[string]interface{}, option model.OPTION_CONVERT_FIELD) map[string]interface{} {
	dataConvert := make(map[string]interface{})

	for key, value := range data {
		switch option {
		case model.MODEL_TO_TABLE:
			fieldTable := ""

			for _, c := range key {
				if unicode.IsUpper(rune(c)) {
					fieldTable = fieldTable + "_" + Lowercase(string(c))
				} else {
					fieldTable += Lowercase(string(c))
				}
			}

			dataConvert[fieldTable] = value
		case model.TABLE_TO_MODEL:
			cutStrings := strings.Split(key, "_")
			fieldModel := cutStrings[0]

			for i := 1; i < len(cutStrings); i++ {
				fieldModel += UppercaseFirstChar(cutStrings[i])
			}

			dataConvert[fieldModel] = value
		}
	}

	return dataConvert
}
