package helpers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

// Initialize error messages mapping

// Set up the custom validation functions
func InitValidator(c *gin.Context) *validator.Validator {
	validate := validator.NewValidator()
	validate.SetValidationFunc("acceptlist", acceptList)
	validate.SetValidationFunc("date", dateFormat)
	validate.SetValidationFunc("unique", validateUniqueValue(c))
	return validate
}

func initErrorMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"less than min": {
			"tel":      "โปรดระบุหมายเลขโทรศัพท์ให้ครบ 10 หลัก",
			"password": "โปรดระบุรหัสผ่านไม่ต่ำกว่า 12 ตัวอักษร",
			"default":  "ค่าไม่สามารถติดลบได้",
		},
		"zero value":                           {"default": "โปรดระบุ"},
		"regular expression mismatch":          {"default": "รูปแบบอีเมลไม่ถูกต้อง"},
		"incorrect":                            {"default": "ข้อมูลไม่ถูกต้อง"},
		"companyEstablishment value not equal": {"default": "ข้อมูลทุนจัดตั้งที่นำเข้ากับข้อมูลหน้าตั้งค่าไม่เท่ากัน"},
		"mismatch regis":                       {"default": "รหัสผ่านไม่ตรงกัน"},
		"mismatch":                             {"old_password": "รหัสผ่านไม่ถูกต้อง", "default": "รหัสผ่านไม่ตรงกัน"},
		"mismatch with old password":           {"default": "ยืนยันรหัสผ่านไม่ถูกต้อง"},
		"already used":                         {"default": "ข้อมูลนี้ถูกใช้ในระบบแล้ว"},
		"not found":                            {"default": "ไม่พบข้อมูล"},
		"inactive":                             {"default": "ข้อมูลไม่ได้ใช้งาน"},
		"not verify":                           {"default": "ข้อมูลไม่ได้ยืนยัน"},
		"not equal":                            {"default": "ข้อมูลไม่ครบถ้วน"},
		"duplicate":                            {"default": "ข้อมูลนี้ถูกใช้แล้ว"},
		"have space":                           {"default": "โปรดระบุข้อมูลที่ไม่มีช่องว่าง"},
		"password must contain at least one uppercase letter":  {"default": "โปรดระบุตัวพิมพ์ใหญ่อย่างน้อย 1 ตัว"},
		"password must contain at least one lowercase letter":  {"default": "โปรดระบุตัวพิมพ์เล็กอย่างน้อย 1 ตัว"},
		"password must contain at least one digit letter":      {"default": "โปรดระบุตัวเลข 0-9 อย่างน้อย 1 ตัว"},
		"password must contain at least one special character": {"default": "โปรดระบุสัญลักษณ์อย่างน้อย 1 ตัว"},
		"duplicate input":                  {"default": "Code หรือ Name ซ้ำกัน"},
		"companyEstablishment value error": {"default": "รหัสทุนจัดตั้งบริษัทไม่ถูกต้อง"},
		"invalid email address":            {"default": "รูปแบบอีเมลไม่ถูกต้อง"},
		"greater than max": {
			"code":             "โปรดระบุข้อมูลไม่เกิน 100 ตัวอักษร",
			"peak_code":        "โปรดระบุข้อมูลไม่เกิน 100 ตัวอักษร",
			"express_code":     "โปรดระบุข้อมูลไม่เกิน 100 ตัวอักษร",
			"description":      "โปรดระบุข้อมูลไม่เกิน 125 ตัวอักษร",
			"discount_percent": "โปรดระบุข้อมูลเปอร์เซ็นต์ไม่เกิน 100 เปอร์เซ็นต์",
			"name":             "โปรดระบุข้อมูลไม่เกิน 255 ตัวอักษร",
			"detail":           "โปรดระบุข้อมูลไม่เกิน 255 ตัวอักษร",
			"tel":              "โปรดระบุหมายเลขโทรศัพท์ไม่เกิน 10 หลัก",
			"round":            "โปรดระบุค่าเป็นบวก",
			"first_name":       "โปรดระบุไม่เกิน 50 ตัวอักษร",
			"last_name":        "โปรดระบุไม่เกิน 50 ตัวอักษร",
			"contract_number":  "โปรดระบุไม่เกิน 50 ตัวอักษร",
		},
		"summary debit and credit not equal": {"default": "ผลรวมเดบิตไม่เท่ากับผลรวมเครดิต"},
		"less than max 255 characters":       {"default": "โปรดระบุข้อมูลไม่เกิน 255 ตัวอักษร"},
		"less than min 10 number":            {"default": "โปรดระบุหมายเลขโทรศัพท์ให้ครบ 10 หลัก"},
		"must be integer":                    {"default": "โปรดระบุเป็นตัวเลข"},
		"out of package":                     {"default": "กรุณาเลือก Package ใหม่เนื่องจาก Package ไม่เพียงพอไม่สามารถทำรายการต่อได้"},
		"invalid length":                     {"default": "โปรดระบุตัวเลขให้ครบ 13 หลัก"},
		"invalid Thai National ID":           {"default": "รูปแบบรหัสประชาชนไม่ถูกต้อง"},
	}
}
