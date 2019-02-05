package describe

import "github.com/chengyumeng/khadijah/pkg/utils/resource"

type Option struct {
	resource  string
	Namespace string
	Output    string
	Cluster   string
	Option    *resource.DescribeOption
}
