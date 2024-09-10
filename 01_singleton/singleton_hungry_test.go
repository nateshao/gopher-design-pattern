package singleton_test

import (
	"github.com/stretchr/testify/assert"
	singleton "gopher-design-pattern/01_singleton"
	"testing"
)

/*
*
验证懒汉式单例模式 和 饿汉式单例模式是否返回同一个实例对象
1. 懒汉式单例，第一次调用时才创建实例。
2. 饿汉式单例，在程序启动时立即创建实例。
*/
func TestGetInstance(t *testing.T) {
	assert.True(t, singleton.GetLazyInstance() == singleton.GetHungryInstance())
}

/*
*
测试在并发环境下，懒汉式和饿汉式单例是否能够安全工作，并且返回同一个实例。这里使用了 b.RunParallel 来模拟多线程/多协程并发场景
 1. pb.Next()：每次循环都会调用，模拟高频率的并发操作。
 2. singleton.GetLazyInstance() != singleton.GetHungryInstance()：每次迭代时检查两个单例实例是否不同，
    如果不同则报告错误 (b.Errorf("test fail"))，意味着在并发环境下，某些情况可能导致单例模式失效。
*/
func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetHungryInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
