package get

import "github.com/chengyumeng/khadijah/pkg/utils/resource"

type Option struct {
	resource string
	Option   *resource.Option
	NS       string
	App      string
	Output   string
}
