package utils

var SetMapTable = map[string]struct {
	Index   int
	Comment string
}{
	"enableMfa":   {0, "启用MFA验证"},
	"mfaUseEmail": {1, "开启使用邮箱作为MFA方式，关闭则使用SMS"},
}

func GetBitByIndex(value, index int) bool {
	return value&(1<<index) == 1
}

func GetBitByName(value int, name string) bool {
	return GetBitByIndex(value, SetMapTable[name].Index)
}

func UpdateBitByIndex(bitmask, index int, value bool) int {
	if value {
		bitmask |= (1 << index)
	} else {
		bitmask &= ^(1 << index)
	}
	return bitmask
}

func UpdateBitByName(bitmask int, name string, value bool) int {
	return UpdateBitByIndex(bitmask, SetMapTable[name].Index, value)
}
