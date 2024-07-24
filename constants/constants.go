package constants

const (
	//User error
	INVALID_ID_PASSWORD     string = "ชื่อผู้ใช้/รหัสผ่านไม่ถูกต้อง กรุณาระบุใหม่อีกครั้ง"
	INVALID_PASSWORD        string = "รหัสผ่านไม่ถูกต้อง"
	USER_NOT_FOUND          string = "ชื่อผู้ใช้งานไม่ถูกต้อง กรุณาระบุใหม่อีกครั้ง"
	USER_BLOCKED            string = "ผู้ใช้งานถูกระงับการเข้าใช้"
	INVALID_IP_ADDRESS      string = "มีการใช้งานจากที่อื่นอยู่"
	INVALID_USER_PERMISSION string = "ผู้ใช้งานไม่มีสิทธิ์ในการใช้งานในส่วนนี้"
	IS_CANNOT_DELETE        string = "ไม่สามารถลบได้เนื่องจากมีการใช้งานอยู่"

	IS_CANNOT_DELETE_FINANCIAL_STATEMENT string = "ไม่สามารถลบรายการได้ เนื่องจากมีข้อมูล AJE"
	INVALID_VERIFY_EMAIL                 string = "ผู้ใช้งานยังไม่ได้ยืนยันอีเมล"
	INVALID_VERIFY_USER                  string = "ผู้ใช้งานยังไม่ถูกยืนยัน"

	EMAIL_ALREADY_REGISTERED string = "Email นี้มีอยู่แล้วในระบบ"

	MUST_BE_DIGIT string = "โปรดระบุตัวเลขหลักหน่วย"

	INVALID_EMAIL string = "รูปแบบอีเมลไม่ถูกต้อง"

	MORE_THAN_MAX_NUMBER string = "โปรดระบุตัวเลขไม่เกิน "

	MUST_BE_NUMERIC string = "โปรดระบุข้อมูลเป็นตัวเลขเท่านั้น"

	AJE_NOT_FOUND string = "ไม่สามารถพิมพ์ PDF ได้ เนื่องจากไม่พบข้อมูล AJE"

	VALID_ZERO_VALUE     string = "โปรดระบุ"
	VALID_DUPLICATE      string = "ข้อมูลนี้มีอยู่ในระบบแล้ว"
	VALID_ONLY_SUPPORT   string = "รองรับเฉพาะ"
	VALID_MIN_NUMBER     string = "โปรดระบุจำนวนอย่างน้อย"
	VALID_MAX_NUMBER     string = "โปรดระบุจำนวนไม่เกิน"
	VALID_DATA_NOT_FOUND string = "ไม่พบข้อมูล"
	VALID_DATE_FORMAT    string = "รูปแบบวันที่ไม่ถูกต้อง"
	VALID_FILE_TYPE      string = "ประเภทไฟล์ไม่รองรับ"
	VALID_FILE_NOT_FOUND string = "ไม่พบไฟล์"
	VALID_DATA_FORMAT    string = "รูปแบบข้อมูลไม่ถูกต้อง"
	VALID_DATE_OVERNOW   string = "วันที่ต้องไม่เกินวันที่ปัจจุบัน"
)
