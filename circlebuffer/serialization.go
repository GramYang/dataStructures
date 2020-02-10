package circlebuffer

import "encoding/json"

//json加密
func (buffer *Buffer) ToJSON() ([]byte, error) {
	return json.Marshal(buffer.elements[:buffer.Length])
}

//json解密
func (buffer *Buffer) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &buffer.elements)
	if err == nil {
		buffer.Length = len(buffer.elements)
	}
	return err
}