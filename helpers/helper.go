package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	databases "github.com/AkapongAlone/validate-helper/database"
	"github.com/AkapongAlone/validate-helper/responses"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gopkg.in/validator.v2"
)

func Validate(v interface{}, c *gin.Context) *responses.Fail {
	validate := InitValidator(c)
	mapError := map[string]string{}

	if errs := validate.Validate(v); errs != nil {
		errorArray := errs.Error() //message จะบอกมี pattern ว่า FieldName: message error,...
		parseErrors(errorArray, mapError)
		errResponse := responses.ValidateResponse(mapError)
		return &errResponse
	}

	return &responses.Fail{}
}

func parseErrors(errorArray string, mapError map[string]string) {
	re := regexp.MustCompile(`\b\w+?\[.*?\]?\.\w+?:|\b\w+?:`)
	listDatas := re.FindAllString(errorArray, -1)

	for index, fieldName := range listDatas {
		var listErr []string
		key := parseFieldName(strings.TrimSuffix(fieldName, ":"))
		lastFieldName := getLastFieldName(strings.TrimSuffix(fieldName, ":"))

		indexFirstElement := strings.Index(errorArray, fieldName) + len(fieldName)
		indexNextElement := len(errorArray)
		if index != len(listDatas)-1 {
			indexNextElement = strings.Index(errorArray, listDatas[index+1])
		}

		errMessages := strings.Split(strings.TrimSpace(errorArray[indexFirstElement:indexNextElement]), ",")
		for _, errMsg := range errMessages {
			if errMsg = strings.TrimSpace(errMsg); errMsg != "" {
				listErr = append(listErr, HandleErrMesssage(lo.SnakeCase(lastFieldName), errMsg))
			}
		}

		mapError[key] = strings.Join(listErr, "|")
	}
}

func getLastFieldName(fieldName string) string {
	re := regexp.MustCompile(`\w+$`)
	return re.FindString(fieldName)
}

func parseFieldName(fieldName string) string {
	re := regexp.MustCompile(`(\w+|\[\d+\])`)
	parts := re.FindAllString(fieldName, -1)

	for i, part := range parts {
		if strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]") {
			parts[i] = part // Keep array indices as they are
		} else {
			if i == 0 {
				parts[i] = lo.SnakeCase(part)
			} else {
				parts[i] = "._" + lo.SnakeCase(part)
			}
		}
	}

	return strings.Join(parts, "")
}

// Custom
// example `validate:"acceptlist=asc|dec"`
func acceptList(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return nil
	}
	paramContains := strings.ToUpper("|" + param + "|")
	value := "|" + strings.ToUpper(st.String()+"|")
	if exists := strings.Contains(paramContains, value); !exists {
		return errors.New("ONLY_SUPPORT&v=" + param)
	}

	return nil
}

// Custom
// example `validate:"date"` check format yyyy-mm-dd
// optional `validate:"date=lt"` lt = <= date now
func dateFormat(v interface{}, param string) error {
	date, ok := v.(string)
	if !ok {
		return errors.New("VALID_DATE_FORMAT")
	}
	if date == "" {
		return nil
	}

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return errors.New("VALID_DATE_FORMAT")
	}

	if strings.TrimSpace(param) == "lt" {
		now := time.Now()
		if dateTime.After(now) {
			return errors.New("VALID_DATE_OVERNOW")
		}
	}

	return nil
}

// first param is table name
// second param is column name
// `validate:"unique=contract_infos,contract_number"`
func validateUniqueValue(c *gin.Context) validator.ValidationFunc {
	return func(v interface{}, param string) error {
		db := databases.NewPostgres()
		value, ok := v.(string)
		if !ok {
			return fmt.Errorf("validateUniqueValue only validates strings")
		}

		params := strings.Split(param, "|")
		if len(params) != 2 {
			return fmt.Errorf("validateUniqueValue requires 2 parameters, got %d", len(params))
		}

		tableName, columnName := params[0], params[1]

		var count int64
		db.Table(tableName).Where(columnName+" = ? AND status = 'A'", value).Count(&count)

		if count > 0 {
			if c.Request.Method == "POST" {
				return errors.New("already used")
			}
			if c.Request.Method == "PUT" {
				parentID := c.Param("id")
				var existingCount int64
				db.Table(tableName).Where(columnName+" = ? AND status = 'A' AND id = ?", value, parentID).Count(&existingCount)
				if existingCount == 0 {
					return errors.New("already used")
				}
			}
		}

		return nil
	}
}

// HandleErrMesssage maps error messages to human-readable text
func HandleErrMesssage(errField, err string) string {
	errorMessages := initErrorMessages()

	//for accept list validate
	data := strings.Split(err, "&v=")
	if len(data) > 1 {
		sub := data[1]
		return "รองรับเฉพาะ" + " " + sub
	}
	///
	if fieldErrors, exists := errorMessages[err]; exists {
		if msg, ok := fieldErrors[errField]; ok {
			return msg
		}
		if msg, ok := fieldErrors["default"]; ok {
			return msg
		}
	}
	return "เกิดข้อผิดพลาด โปรดลองอีกครั้ง"
}
