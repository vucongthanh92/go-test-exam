package constants

import "time"

const (
	Time_Cache_5_minutes = 5 * time.Minute
	Time_Cache_1_day     = 24 * time.Hour
	Time_Cache_5_seconds = 5 * time.Second
)

const (
	InvalidValue       = "InvalidValue"
	InvalidLength      = "InvalidLength"
	InvalidEmailFormat = "InvalidEmailFormat"
)

const (
	Yaml               = "yaml"
	Gzip               = "gzip"
	Redis              = "redis"
	ReadDatabase       = "read-database"
	WriteDatabase      = "write-database"
	GoroutineThreshold = "goroutine-threshold"
	Kafka              = "kafka"
)

type ProductType[T any] struct {
	Available  T
	OutOfStock T
	OnOrder    T
}

var ProductTypeName ProductType[string] = ProductType[string]{
	Available:  "Available",
	OutOfStock: "Out of Stock",
	OnOrder:    "On Order",
}
