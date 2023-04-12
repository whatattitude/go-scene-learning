package singleton

import "testing"

func TestGetLazySingletonInstance(t *testing.T) {
 if GetLazySingletonInstance() != GetLazySingletonInstance() {
  t.Error("懒汉式 单例模式： 出现同一对象")
 }
}

func BenchmarkGetLazySingletonInstance(b *testing.B) {
 for i := 0; i < b.N; i++ {
  if GetLazySingletonInstance() != GetLazySingletonInstance() {
   b.Error("懒汉式 单例模式： 出现同一对象")
  }
 }
}