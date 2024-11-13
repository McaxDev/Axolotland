package backup

// 获取世界目录在服务器目录的位置
func getWorldPath(game string) string {

	switch game {
	case "bedrock":
		return "worlds/"
	case "dst":
		return "DoNotStarveTogether/"
	case "terraria":
		return "world/"
	case "stardew":
		return "data/"
	default:
		return ""
	}
}
