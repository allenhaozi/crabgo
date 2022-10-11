package common

type LogStoreEnd string

var (
	LogStoreLoki LogStoreEnd = "loki"
	LogStoreES   LogStoreEnd = "es"
)

var DefaultLogStore = LogStoreLoki
