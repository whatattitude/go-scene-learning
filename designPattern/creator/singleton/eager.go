package singleton

import (
	"fmt"
	"sync"
)

type doubleCheckSingleton struct {
}

var lock = &sync.Mutex{}

var doubleCheckSingletonInstance *doubleCheckSingleton

func GetDoubleCheckSingletonInstance() *doubleCheckSingleton {
 if doubleCheckSingletonInstance == nil {
  lock.Lock()
  defer lock.Unlock()
  if doubleCheckSingletonInstance == nil {
   fmt.Println("创建 单例对象")
   doubleCheckSingletonInstance = &doubleCheckSingleton{}
  } else {
   fmt.Println("单例对象已经创建了")
  }
 } else {
  fmt.Println("单例对象已经创建了")
 }

 return doubleCheckSingletonInstance
}