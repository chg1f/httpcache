package httpcache

import (
	netHttp "net/http"
	"time"
)

type CacheOption struct {
	Key        string
	GetTimeout time.Duration
	SetTTL     time.Duration
}
type Cache interface {
	Get(*CacheOption) (*netHttp.Response, bool)
	Set(*CacheOption, *netHttp.Response) error
}

type LFUCache struct {
}
type LRUCache struct {
}
type RedisCache struct {
}
