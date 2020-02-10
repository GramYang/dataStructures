package hashset

import "encoding/json"

// ToJSON outputs the JSON representation of the set.
func (set *Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}