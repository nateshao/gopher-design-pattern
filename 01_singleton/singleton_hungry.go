package singleton

/**
饿汉式单例模式
饿汉式模式的特点是实例在程序启动时就初始化，因此它天然是线程安全的。
*/

// 单例对象的定义
type singleton struct {
}

// 在程序启动的时候，单例实例就会被创建
var instance = &singleton{}

// 获取单例对象的函数
func GetHungryInstance() *singleton {
	return instance
}
