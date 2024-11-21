package utils

var SetMapTable = map[string]struct {
	Index   int
	Comment string
}{
	"enableMfa":   {0, "启用MFA验证"},
	"mfaUseEmail": {1, "开启使用邮箱作为MFA方式，关闭则使用SMS"},
}

func GetBitByIndex(data int64, index int) bool {
	return data&(1<<(index%64)) == 1
}

func GetBitByName(data int64, name string) bool {
	return GetBitByIndex(data, SetMapTable[name].Index)
}

func UpdateBitByIndex(data *int64, index int, value bool) {
	if value {
		*data |= (1 << (index % 64))
	} else {
		*data &= ^(1 << (index % 64))
	}
}

func UpdateBitByName(data *int64, name string, value bool) {
	UpdateBitByIndex(data, SetMapTable[name].Index, value)
}
