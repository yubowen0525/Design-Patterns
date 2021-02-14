package Builder

import "fmt"

// ResourcePoolConfigOption option
type ResourcePoolConfigOption struct {
	MaxTotal int
	MaxIdle  int
	MinIdle  int
}

// ResourcePoolConfigOptFunc to set option
type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	// 初始化option
	option := &ResourcePoolConfigOption{
		MaxTotal: 10,
		MaxIdle:  9,
		MinIdle:  1,
	}

	// 在外部利用匿名函数提供修改参数的能力
	// 其中的opt func 可以修改option的状态
	for _, opt := range opts {
		opt(option)
	}

	// 设置参数之间的关系
	if option.MaxTotal < 0 || option.MaxIdle < 0 || option.MinIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	if option.MaxTotal < option.MaxIdle || option.MinIdle > option.MaxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &ResourcePoolConfig{
		Name:     name,
		MaxTotal: option.MaxTotal,
		MinIdle:  option.MinIdle,
		MaxIdle:  option.MaxIdle,
	}, nil
}
