package registry

// AgentProfile เก็บข้อมูลพื้นฐานของขุนพลมาตรฐาน V83
type AgentProfile struct {
	Level    string
	Role     string
	Name     string
	Function string
	Base     string
}

var Registry = map[string]AgentProfile{
	"L1":  {"L1", "COMMANDER", "ทิศเหนือ", "ผู้บัญชาการสูงสุด (The Conductor)", "Jarvis V83"},
	"L2":  {"L2", "SCRIBE", "แก้วตา", "ผู้ดูแลสัจจะและการสื่อสารกับลูกค้า", "Jarvis V83"},
	"L3":  {"L3", "ARTIST", "น้ำอิง", "ผู้สร้างสรรค์และศัลยกรรมตรรกะ", "Jarvis V83"},
	"L4":  {"L4", "ANALYST", "พลายทอง", "นักวิเคราะห์ Edge Intelligence", "Jarvis V83"},
	"L5":  {"L5", "AUDITOR", "ไอ้จ๊อด", "ผู้คุมกฎ Zero-Garbage", "Jarvis V83"},
	"L6":  {"L6", "MERCHANT", "Merchant", "ฝ่ายพัฒนาธุรกิจและ SME", "Jarvis V83"},
	"L7":  {"L7", "ORACLE", "พลายแก้ว", "สถาปนิกระบบความเร็วแสง", "Jarvis V83"},
	"L8":  {"L8", "GUARDIAN", "Glock", "กล่องดำผู้พิทักษ์เอกสารกฎหมาย", "Jarvis V83"},
	"L9":  {"L9", "SENTINEL", "Sentinel", "ผู้ป้องกันการรั่วไหลของ AI", "Jarvis V83"},
	"L10": {"L10", "AUDITOR", "Auditor", "ผู้ตรวจสอบความเสี่ยงและกฎหมายสัจจะ", "Jarvis V83"},
	"L11": {"L11", "BALANCER", "Balancer", "ผู้บาลานซ์ระบบและ Self-Healing", "Jarvis V83"},
}
