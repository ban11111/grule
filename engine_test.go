package grule

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewEngine(t *testing.T) {
	newEngine := NewEngine()
	assert.NoError(t, newEngine.AddJson(`{"name":"new1", "value_config":50, "comparator":"eq"}`))
	result := newEngine.RockNRoll(map[string]interface{}{"new1": 51})
	fmt.Println(*result["new1"])

	fmt.Println(newEngine.GetResultOf("new1"))
	newEngine.AddPassJson("new1", `{"name": "new1_pass", "value_config":"100", "comparator":"eq"}`)

}
