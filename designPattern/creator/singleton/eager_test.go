package singleton

import "testing"

func TestGetDoubleCheckSingletonInstance(t *testing.T) {
 if GetDoubleCheckSingletonInstance() != GetDoubleCheckSingletonInstance() {
  t.Error("双重检测法 获取不同的对象")
 }
}

func BenchmarkGetDoubleCheckSingletonInstance(b *testing.B) {
 for i := 0; i < b.N; i++ {
  if GetDoubleCheckSingletonInstance() != GetDoubleCheckSingletonInstance() {
   b.Error("双重检测法 获取不同的对象")
  }
 }
}