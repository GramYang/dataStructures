package treemap

// ToJSON outputs the JSON representation of the map.
func (m *Map) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates the map from the input JSON representation.
func (m *Map) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}