package database

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructBuilder(t *testing.T) {
	pe := NewStructBuilder().
		AddString("Name", `json:"json_name"`).
		AddInt64("Age", `json:"json_age"`).
		Build()
	p := pe.New()
	p.SetString("Name", "你好")
	p.SetInt64("Age", 32)
	jsonStr := `{"json_name":"你好","json_age":32}`
	assert.NotNil(t, pe)
	assert.NotNil(t, p)
	bytes, err := json.Marshal(p.Interface())
	assert.Nil(t, err)
	fmt.Println(string(bytes))
	assert.Equal(t, jsonStr, string(bytes))
}
