package singleton

import "sync"

var (
 lazySingletonInstance *lazySingleton
 once                  = &sync.Once{}
)

// LazySingleton 懒汉式
type lazySingleton struct {
}

func GetLazySingletonInstance() *lazySingleton {
 if lazySingletonInstance == nil {
  once.Do(func() {
   lazySingletonInstance = &lazySingleton{}
  })
 }
 return lazySingletonInstance
}