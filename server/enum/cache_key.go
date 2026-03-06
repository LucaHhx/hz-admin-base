package enum

import "fmt"

const Head_RepeatCount = "rabbit-count"

const AppKey = "hz-admin-base"

var CacheKey = fmt.Sprintf("{%s}", AppKey)
