package get

import "github.com/chengyumeng/khadijah/pkg/utils/resource"

// Option is the option of get proxy
type Option struct {
	resource string
	Option   *resource.Option
	NS       string
	App      string
	Output   string
}
