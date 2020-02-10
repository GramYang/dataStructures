package binaryheap

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}