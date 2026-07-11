package registry

type AgentProfile struct {
	Level, Role, Name, Function, Base, Link string
}

var Registry = map[string]AgentProfile{
	"L1": {"L1", "COMMANDER", "ทิศเหนือ", "ผู้บัญชาการสูงสุด", "Jarvis V83", "https://arthitsiangwan.github.io/ThitNueaHub-Admin/"},
	"L2": {"L2", "SCRIBE", "แก้วตา", "สื่อสารและสัจจะ", "Jarvis V83", "https://thitnueahub.com/kaewta"},
    // ... เพิ่มให้ครบ 11 ท่าน
}
