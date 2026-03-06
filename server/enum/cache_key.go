package enum

import "fmt"

const Head_RepeatCount = "rabbit-count"

const AppKey = "hab"

var CacheKey = fmt.Sprintf("{%s}", AppKey)
