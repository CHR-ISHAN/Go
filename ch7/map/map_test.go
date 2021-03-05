package _map

import (
	"fmt"
	"sort"
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	t.Log(m1)
	m1[2] = 14
	t.Log(m1)
	m2 := map[int]int{}
	t.Log(m2)
	m3 := make(map[int]int, 10)
	t.Log(m3)
}

func TestAccess(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1)
	if _, ok := m1[3]; ok {
		t.Log("11111")
	} else {
		t.Log("vbvvvv")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

//func TestMap(t *testing.T) {
//	var mapList map[string]int
//	var mapAss map[string]int
//	mapList = map[string]int{"one": 1}
//	var mapTest map[string]string
//	mapCreated := make(map[string]float32)
//	mapCreated["key1"] = 4.5
//	mapCreated["key2"] = 4.5
//	mapCreated["key4"] = 4.5
//	v, Is := mapTest["key5"]
//	t.Log(v, Is)
//	mapAss = mapList
//	t.Log(mapAss, mapCreated, mapList)
//}

func TestWeek(t *testing.T) {
	week := map[int]string{1: "Mon", 2: "Tues", 3: "Th"}
	if _, Is := week[2]; Is {
		t.Log(week[2])
	}
	if _, Is := week[4]; Is {
		t.Log(week[4])
	}
}

var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87,
	"juliet": 65, "kili": 43, "lima": 98}

func TestMapSort(t *testing.T) {
	for k, v := range barVal {
		t.Logf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		t.Logf("Key: %v, Value: %v / ", k, barVal[k])
	}
}

func TestMap(t *testing.T) {
	var m, n map[int32]int32
	m = n
	n = map[int32]int32{1: 1}
	fmt.Println(m)
	fmt.Println(n)
}

func TestMap1(t *testing.T) {
	m, n := map[int32]int32{1: 2}, map[int32]int32{}
	m = n
	t.Log(m, n)
	m[1] = 3
	t.Log(m, n)
	t.Logf("%p,%p", &m, &n)
}
