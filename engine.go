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
	return new(RuleEngine)
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

// one way of running the engine
func (engine *RuleEngine) RockNRoll(data map[string]interface{}) *RuleEngine {
	resultEngine := engine.clean()
	if resultEngine.count <= 0 {
		panic("empty rules, please configure rules before running")
	}
	if resultEngine.respType == respTypeSimple {
		resultEngine.RespSimple = make(map[string]*bool, resultEngine.count)
	}
	for ruleName, value := range data {
		if resultEngine.r[ruleName] == nil {
			print(ruleName + " not configured yet, ignore this rule")
			continue
		}
		resultEngine.r[ruleName].Data = value
		resultEngine.RespSimple[ruleName] = resultEngine.r[ruleName].compare()
	}
	return resultEngine
}

// get named rule's result
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

// full copy of engine
func (engine *RuleEngine) clone() *RuleEngine {
	e := &RuleEngine{
		r:            engine.r.clone(),
		count:        engine.count,
		respType:     engine.respType,
	}
	for key, value := range engine.RespSimple {
		e.RespSimple[key] = value
	}
	for key, value := range engine.RespComplete {
		e.RespComplete[key] = value
	}
	return e
}

// a clean copy of engine before running
func (engine *RuleEngine) clean() *RuleEngine {
	e := &RuleEngine{
		r:            engine.r.clone(),
		count:        engine.count,
		respType:     engine.respType,
	}
	return e
}
