package com

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// json编码
func JsonEncode(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	return string(jsonBytes), err
}

// json解码
func JsonDecode(data string) (interface{}, error) {
	dataByte := []byte(data)
	var decoded interface{}

	err := json.Unmarshal(dataByte, &decoded)
	return decoded, err
}

func Encode(data interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(to)
}
