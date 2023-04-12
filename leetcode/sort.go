package leetcode

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
	"sort"
)


func createRandomArr(len int) (arrRandom []int, arrBase []int) {
	arrRandom = make([]int, len)
	for i := 0; i < len; i++ {
		arrRandom[i] = i
	}
	arrBase = make([]int, len)
	copy(arrBase, arrRandom)
	for i := 0; i < len; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(128))
		arrRandom[i], arrRandom[result.Int64()%int64(len)] = arrRandom[result.Int64()%int64(len)], arrRandom[i]
	}

	return
}

func sortString() {
	
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}

	fmt.Println(reflect.TypeOf(s))
	s1 := sort.StringSlice(s)
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(s1)
    // 从大到小进行排序
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

}
