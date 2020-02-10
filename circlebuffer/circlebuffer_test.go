package circlebuffer

import (
	"dataStructures/utils"
	"fmt"
	"testing"
)

func TestBufferNew(t *testing.T){
	buffer1:=New(10)
	if actualValue := buffer1.Empty(); actualValue == true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	buffer2:=New(5)
	buffer2.Add(1, "b")
	if actualValue := buffer2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := buffer2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, ok := buffer2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := buffer2.Get(2); actualValue != nil || !ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

func TestListAdd(t *testing.T) {
	buffer:= New(3)
	buffer.Add("a")
	buffer.Add("b", "c")
	if actualValue := buffer.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := buffer.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := buffer.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	buffer.Add("d")
	if actualValue, ok := buffer.Get(0); actualValue != "d" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "d")
	}
}

func TestListIndexOf(t *testing.T) {
	buffer := New(5)
	expectedIndex := -1
	if index := buffer.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
	buffer.Add("a")
	buffer.Add("b", "c")
	expectedIndex = 0
	if index := buffer.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
	expectedIndex = 1
	if index := buffer.IndexOf("b"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
	expectedIndex = 2
	if index := buffer.IndexOf("c"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
}

func TestListGet(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a")
	buffer.Add("b", "c")
	if actualValue, ok := buffer.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := buffer.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := buffer.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := buffer.Get(3); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

func TestListSwap(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a")
	buffer.Add("b", "c")
	buffer.Swap(0, 1)
	if actualValue, ok := buffer.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListSort(t *testing.T) {
	buffer:= New(5)
	buffer.Sort(utils.StringComparator)
	buffer.Add("e", "f", "g", "a", "b", "c", "d")
	buffer.Sort(utils.StringComparator)
	for i := 1; i < buffer.Size(); i++ {
		a, _ := buffer.Get(i - 1)
		b, _ := buffer.Get(i)
		if a.(string) > b.(string) {
			t.Errorf("Not sorted! %s > %s", a, b)
		}
	}
}

func TestListClear(t *testing.T) {
	buffer:= New(5)
	buffer.Add("e", "f", "g", "a", "b", "c", "d")
	buffer.Clear()
	if actualValue := buffer.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := buffer.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListContains(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a")
	buffer.Add("b", "c")
	if actualValue := buffer.Contains("a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := buffer.Contains("a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := buffer.Contains("a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	buffer.Clear()
	if actualValue := buffer.Contains("a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := buffer.Contains("a", "b", "c"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestListValues(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a")
	buffer.Add("b", "c")
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", buffer.Values()...), "abc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListSet(t *testing.T) {
	buffer:= New(5)
	buffer.Set(0, "a")
	buffer.Set(1, "b")
	if actualValue := buffer.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	buffer.Set(2, "c") // append
	if actualValue := buffer.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	buffer.Set(4, "d")  // ignore
	buffer.Set(1, "bb") // update
	if actualValue := buffer.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", buffer.Values()...), "abbc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListEach(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	buffer.Each(func(index int, value interface{}) {
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	})
}

func TestListMap(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	mappedList := buffer.Map(func(index int, value interface{}) interface{} {
		return "mapped: " + value.(string)
	})
	if actualValue, _ := mappedList.Get(0); actualValue != "mapped: a" {
		t.Errorf("Got %v expected %v", actualValue, "mapped: a")
	}
	if actualValue, _ := mappedList.Get(1); actualValue != "mapped: b" {
		t.Errorf("Got %v expected %v", actualValue, "mapped: b")
	}
	if actualValue, _ := mappedList.Get(2); actualValue != "mapped: c" {
		t.Errorf("Got %v expected %v", actualValue, "mapped: c")
	}
	if mappedList.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedList.Size(), 3)
	}
}

func TestListSelect(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	selectedList := buffer.Select(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "b"
	})
	if actualValue, _ := selectedList.Get(0); actualValue != "a" {
		t.Errorf("Got %v expected %v", actualValue, "value: a")
	}
	if actualValue, _ := selectedList.Get(1); actualValue != "b" {
		t.Errorf("Got %v expected %v", actualValue, "value: b")
	}
	if selectedList.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedList.Size(), 3)
	}
}

func TestListAny(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	any := buffer.Any(func(index int, value interface{}) bool {
		return value.(string) == "c"
	})
	if any != true {
		t.Errorf("Got %v expected %v", any, true)
	}
	any = buffer.Any(func(index int, value interface{}) bool {
		return value.(string) == "x"
	})
	if any != false {
		t.Errorf("Got %v expected %v", any, false)
	}
}
func TestListAll(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	all := buffer.All(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "c"
	})
	if all != true {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = buffer.All(func(index int, value interface{}) bool {
		return value.(string) >= "a" && value.(string) <= "b"
	})
	if all != false {
		t.Errorf("Got %v expected %v", all, false)
	}
}
func TestListFind(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	foundIndex, foundValue := buffer.Find(func(index int, value interface{}) bool {
		return value.(string) == "c"
	})
	if foundValue != "c" || foundIndex != 2 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, "c", 2)
	}
	foundIndex, foundValue = buffer.Find(func(index int, value interface{}) bool {
		return value.(string) == "x"
	})
	if foundValue != nil || foundIndex != -1 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, nil, nil)
	}
}
func TestListChaining(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	chainedList := buffer.Select(func(index int, value interface{}) bool {
		return value.(string) > "a"
	}).Map(func(index int, value interface{}) interface{} {
		return value.(string) + value.(string)
	})
	if chainedList.Size() != 2 {
		t.Errorf("Got %v expected %v", chainedList.Size(), 2)
	}
	if actualValue, ok := chainedList.Get(0); actualValue != "bb" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := chainedList.Get(1); actualValue != "cc" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListIteratorNextOnEmpty(t *testing.T) {
	buffer:= New(5)
	it := buffer.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty lists")
	}
}

func TestListIteratorNext(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	it := buffer.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListIteratorPrevOnEmpty(t *testing.T) {
	buffer:= New(5)
	it := buffer.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty lists")
	}
}

func TestListIteratorPrev(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	it := buffer.Iterator()
	for it.Next() {
	}
	count := 0
	for it.Prev() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListIteratorBegin(t *testing.T) {
	buffer:= New(5)
	it := buffer.Iterator()
	it.Begin()
	buffer.Add("a", "b", "c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestListIteratorEnd(t *testing.T) {
	buffer:= New(5)
	it := buffer.Iterator()
	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}
	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}
	buffer.Add("a", "b", "c")
	it.End()
	if index := it.Index(); index != buffer.Size() {
		t.Errorf("Got %v expected %v", index, buffer.Size())
	}
	it.Prev()
	if index, value := it.Index(), it.Value(); index != buffer.Size()-1 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, buffer.Size()-1, "c")
	}
}

func TestListIteratorFirst(t *testing.T) {
	buffer:= New(5)
	it := buffer.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	buffer.Add("a", "b", "c")
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestListIteratorLast(t *testing.T) {
	buffer:= New(5)
	it := buffer.Iterator()
	if actualValue, expectedValue := it.Last(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	buffer.Add("a", "b", "c")
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "c")
	}
}

func TestListSerialization(t *testing.T) {
	buffer:= New(5)
	buffer.Add("a", "b", "c")
	var err error
	assert := func() {
		if actualValue, expectedValue := fmt.Sprintf("%s%s%s", buffer.Values()...), "abc"; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := buffer.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}
	assert()
	json, err := buffer.ToJSON()
	assert()
	err = buffer.FromJSON(json)
	assert()
}

func benchmarkGet(b *testing.B, buffer *Buffer, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			buffer.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, buffer *Buffer, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			buffer.Add(n)
		}
	}
}

func BenchmarkArrayListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	buffer:= New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, buffer, size)
}

func BenchmarkArrayListGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	buffer := New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, buffer, size)
}

func BenchmarkArrayListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	buffer := New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, buffer, size)
}

func BenchmarkArrayListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	buffer:= New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, buffer, size)
}

func BenchmarkArrayListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	buffer := New(5)
	b.StartTimer()
	benchmarkAdd(b, buffer, size)
}

func BenchmarkArrayListAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	buffer := New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, buffer, size)
}

func BenchmarkArrayListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	buffer:= New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, buffer, size)
}

func BenchmarkArrayListAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	buffer:= New(5)
	for n := 0; n < size; n++ {
		buffer.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, buffer, size)
}