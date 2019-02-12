package describe

import "github.com/chengyumeng/khadijah/pkg/utils/resource"

// Option is the option of describe proxy
type Option struct {
	resource  string
	Namespace string
	Output    string
	Cluster   string
	Option    *resource.DescribeOption
}
