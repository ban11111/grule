package grule

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewEngine(t *testing.T) {
	newEngine := NewEngine()
	assert.NoError(t, newEngine.AddJSON(`{"name":"new1", "param":"age", "value":51, "cmp":"eq"}`))
	result := newEngine.RockNRoll([]string{"new1"}, map[string]interface{}{"age": 51})
	fmt.Println(result.GetResultOf("new1"))

	newEngine.AddPassJSON("new1", `{"param":"age2", "value":50, "cmp":"eq"}`)
	result = newEngine.RockNRoll([]string{"new1"}, map[string]interface{}{"age": 51, "age2": 52})
	fmt.Println(result.GetResultOf("new1"))

	//newEngine.AddPassJSON("p", `{"name": "p.p", "param":"age", "value":"100", "cmp":"eq"}`)
	//result = newEngine.RockNRoll([]string{"new1"}, map[string]interface{}{"age": 51})
	//fmt.Println(result.RespSimple["new1"])
	//fmt.Println(result.GetResultOf("new1"))
}
