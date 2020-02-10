package arraylist

import "encoding/json"

//json加密
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

//json解密
func (list *List) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}