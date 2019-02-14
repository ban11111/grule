package grule

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewEngine(t *testing.T) {
	newEngine := NewEngine()
	assert.NoError(t, newEngine.AddJson(`{"name":"new1", "value":51, "comparator":"eq"}`))
	result := newEngine.RockNRoll(map[string]interface{}{"new1": 51})
	fmt.Println(result.RespSimple["new1"])

	fmt.Println(result.GetResultOf("new1"))
	newEngine.AddPassJson("new1", `{"name": "new1_pass", "value":"100", "comparator":"eq"}`)
	result = newEngine.RockNRoll(map[string]interface{}{"new1": 51})
	fmt.Println(result.RespSimple["new1"])
	fmt.Println(result.GetResultOf("new1"))
}
