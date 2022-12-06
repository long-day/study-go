package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
   @author:君
   @date:2022-12-06
   @note:
*/

// 原子计数器
func TestAtomicCounter(t *testing.T) {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}

// 互斥锁

type checkinMapStruct struct {
	muLock     sync.Mutex
	checkinMap map[string]uint64
}

func (c *checkinMapStruct) checkinByName(name string) {
	c.muLock.Lock()
	defer c.muLock.Unlock()
	c.checkinMap[name]++
}

func TestMutualExclusionLock(t *testing.T) {
	c := checkinMapStruct{
		checkinMap: map[string]uint64{"qq": 0, "github": 0},
	}
	var wg sync.WaitGroup
	doCheckinInterface := func(checkinName string, checkinCount uint64) {
		wg.Add(1)
		for i := uint64(0); i < checkinCount; i++ {
			c.checkinByName(checkinName)
		}
		wg.Done()
	}
	doCheckinInterface("qq", 10000)
	doCheckinInterface("github", 10000)
	doCheckinInterface("github", 20000)

	wg.Wait()
	fmt.Println(c.checkinMap)
}

//状态协程

func TestStatusCoroutine(t *testing.T) {

}
