package singleton

import "sync"

/**
懒汉式单例模式会在第一次访问时才初始化实例，
在并发场景下，如果多个协程同时访问 GetInstance()，可能导致实例化多次。为了确保只创建一个实例，我们使用 sync.Once 来保证线程安全。
*/

// 单例对象的定义
// 使用sync.Once 确保只创建一次实例
var (
	lazySinleton *singleton
	once         = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *singleton {
	// once 内的方法只会执行一次，所以不需要再次判断
	once.Do(func() {
		lazySinleton = &singleton{}
	})
	return lazySinleton
}
