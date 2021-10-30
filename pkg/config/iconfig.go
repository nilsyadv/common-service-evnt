package config

type IConfig interface {
	GetString(string) string
	GetInt64(string) int64
	GetBool(string) bool
}
