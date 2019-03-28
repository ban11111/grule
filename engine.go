package grule

import (
	"github.com/json-iterator/go"
)

type RuleEngine struct {
	r            ruler
	count        int
	respType     int                    // default RespSimple
	data         map[string]interface{} // 全部输入参数
	RespSimple   map[string]*bool       // 简版结果
	RespComplete map[string]interface{} // 祥版结果
}

// config of rule
type RuleConfig struct {
	// rule id
	Name string `json:"name"`
	// rule param
	Param string `json:"param"`
	// value of param
	Value interface{} `json:"value"`
	// comparator
	Comparator string `json:"cmp"`
}

const (
	respTypeSimple = iota
	respTypeComplete
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewEngine() *RuleEngine {
	return new(RuleEngine)
}

func (engine *RuleEngine) AddRule(rc *RuleConfig) (err error) {
	engine.r.add(&rule{
		Name:       rc.Name,
		Param:      rc.Param,
		Value:      rc.Value,
		Comparator: rc.Comparator,
	})
	engine.count ++
	return
}

func (engine *RuleEngine) AddJSON(ruleJson string) (err error) {
	var rule rule
	if err = json.UnmarshalFromString(ruleJson, &rule); err != nil {
		return
	}
	engine.r.add(&rule)
	engine.count ++
	return
}

func (engine *RuleEngine) AddPassJSON(name, ruleJson string) (err error) {
	var rule rule
	if err = json.UnmarshalFromString(ruleJson, &rule); err != nil {
		return
	}
	//paths := strings.Split(name, ".")
	// todo, 优化
	engine.r.addSub(name, "pass", &rule)
	return
}

func (engine *RuleEngine) AddFailJSON(name, ruleJson string) (err error) {
	var rule rule
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err = json.UnmarshalFromString(ruleJson, &rule); err != nil {
		return
	}
	engine.r.addSub(name, "fail", &rule)
	return
}

func (engine *RuleEngine) AddJSONs(rulesJson string) (err error) {
	var rules []rule
	if err = json.UnmarshalFromString(rulesJson, &rules); err != nil {
		return
	}
	engine.r.adds(rules)
	engine.count += len(rules)
	return
}

func (engine *RuleEngine) AddRules(rules []rule) {
	engine.r.adds(rules)
	engine.count += len(rules)
}

// one way of running the engine
func (engine *RuleEngine) RockNRoll(rulesToRun []string, data map[string]interface{}, doParallel ...bool) *RuleEngine {
	resultEngine := engine.clean()
	resultEngine.data = data
	if resultEngine.count <= 0 {
		panic("empty rules, please configure rules before running")
	}
	if resultEngine.respType == respTypeSimple {
		resultEngine.RespSimple = make(map[string]*bool, resultEngine.count)
	}
	//if len(doParallel)>0 && doParallel[0] {
	//	finish := make(chan bool)
	//	for ruleName, value := range data {
	//		if resultEngine.r[ruleName] == nil {
	//			print(ruleName + " not configured yet, ignore this rule")
	//			continue
	//		}
	//		go func(f chan bool) {
	//			resultEngine.r[ruleName].Data = value
	//			resultEngine.RespSimple[ruleName] = resultEngine.r[ruleName].compare()
	//			finish <- true
	//		}(finish)
	//	}
	//	for {
	//		if _, ok := <- finish; !ok {
	//			close(finish)
	//			break
	//		}
	//	}
	//	return resultEngine
	//}
	for _, ruleName := range rulesToRun {
		if resultEngine.r[ruleName] == nil {
			print(ruleName + " not configured yet, ignore this rule")
			continue
		}
		resultEngine.RespSimple[ruleName] = resultEngine.doCompare(ruleName)
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
		r:        engine.r.clone(),
		count:    engine.count,
		respType: engine.respType,
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
		r:        engine.r.clone(),
		count:    engine.count,
		respType: engine.respType,
	}
	return e
}

func (engine *RuleEngine) loadData(dataStr string) (err error) {
	if err = json.UnmarshalFromString(dataStr, &engine.data); err != nil {
		return
	}
	return
}

func (engine *RuleEngine) doCompare(ruleName string) *bool {
	return engine.r[ruleName].compare(engine.data)
}
