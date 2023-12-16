package utils

func TableName(m any) string {
	return GetModelName(m) + "s"
}
