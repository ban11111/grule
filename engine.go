package grule

import (
	"github.com/json-iterator/go"
)

type RuleEngine struct {
	r            ruler
	count        int
	respType     int // 默认返回简单版本
	RespSimple   map[string]*bool
	RespComplete map[string]interface{}
}

const (
	respTypeSimple   = iota
	respTypeComplete
)

func NewEngine() *RuleEngine {
	return &RuleEngine{}
}

func (engine *RuleEngine) AddJson(ruleJson string) (err error) {
	var rule rule
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.UnmarshalFromString(ruleJson, &rule); err != nil {
		return
	}
	engine.r.add(rule)
	engine.count ++
	return
}

func (engine *RuleEngine) AddPassJson(name, ruleJson string) (err error) {
	var rule rule
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.UnmarshalFromString(ruleJson, &rule); err != nil {
		return
	}
	engine.r.addSub(name, "pass", rule)
	return
}

func (engine *RuleEngine) AddFailJson(name, ruleJson string) (err error) {
	var rule rule
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.UnmarshalFromString(ruleJson, &rule); err != nil {
		return
	}
	engine.r.addSub(name, "fail", rule)
	return
}

func (engine *RuleEngine) AddRules(rules []rule) {
	// todo, 转换成rules
	engine.r.adds(rules)
	engine.count += len(rules)
}

func (engine *RuleEngine) RockNRoll(data map[string]interface{}) (map[string]*bool) {
	if engine.count <= 0 {
		panic("empty rules, please configure rules before running")
	}
	if engine.respType == respTypeSimple {
		engine.RespSimple = make(map[string]*bool, engine.count)
	}
	for ruleName, value := range data {
		if engine.r[ruleName] == nil {
			print(ruleName + " not configured yet, ignore this rule")
			continue
		}
		engine.r[ruleName].DataReceived = value
		engine.RespSimple[ruleName] = engine.r[ruleName].compare()
	}
	return engine.RespSimple
}

func (engine *RuleEngine) GetResultOf(name string) (result string) {
	if engine.RespSimple[name] != nil {
		r := *engine.RespSimple[name]
		if r {
			return "pass"
		}
		return "fail"
	}
	return "pending or error"
}