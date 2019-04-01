package grule

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEngine(t *testing.T) {
	newEngine := NewEngine()
	assert.NoError(t, newEngine.AddJSON(`{"name":"new1", "param":"age", "value":51, "cmp":"eq"}`))
	result := newEngine.RockNRoll([]string{"new1"}, map[string]interface{}{"age": 51})
	fmt.Println(result.GetResultOf("new1"))

	newEngine.AddPassJSON("new1", `{"param":"age2", "value":50, "cmp":"eq"}`)
	result = newEngine.RockNRoll([]string{"new1"}, map[string]interface{}{"age": 51, "age2": 52})
	fmt.Println(result.GetResultOf("new1"))

	err := newEngine.AddRule(&RuleConfig{
		// rule id
		Name: "rule1",
		// rule param
		Param: "address",
		// value of param
		Value: "abc",
		// comparator
		Comparator: "eq",
	})
	assert.NoError(t, err)
	result = newEngine.RockNRoll([]string{"rule1"}, map[string]interface{}{"address": "abc"})
	assert.Equal(t, "pass", result.GetResultOf("rule1"))
	newEngine.AddPassRule("rule1", &RuleConfig{
		// rule param
		Param: "address2",
		// value of param
		Value: "aaa",
		// comparator
		Comparator: "eq",
	})
	result = newEngine.RockNRoll([]string{"rule1"}, map[string]interface{}{"address": "abc", "address2": "aaa"})
	assert.Equal(t, "pass", result.GetResultOf("rule1"))
	newEngine.AddFailRule("rule1", &RuleConfig{
		// rule param
		Param: "address2",
		// value of param
		Value: "aaa",
		// comparator
		Comparator: "eq",
	})
	result = newEngine.RockNRoll([]string{"rule1"}, map[string]interface{}{"address": "abc1", "address2": "aaa"})
	assert.Equal(t, "pass", result.GetResultOf("rule1"))
}
