package grule

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"time"
)

func TestSomething(t *testing.T) {
	str := `{"name":"ban11111", "fail":{"name":"ban2222", "pass":{"name":"ban33333"}}}`
	var r rule
	assert.NoError(t, json.UnmarshalFromString(str, &r))
	fmt.Println(r.Name, r.Fail, r.Fail.Pass)
}

func TestTestingComparators(t *testing.T) {
	assert.EqualValues(t, 12, float64(12.000))
	assert.Equal(t, 1, 1)
	assert.NotEqual(t, 1, 2)
	assert.True(t, true)
	assert.False(t, false)
	assert.Contains(t, "asdf", "s")
	assert.NotContains(t, "aaa", "n")
	assert.Len(t, "Asd", 3)
	assert.Zero(t, nil)
	assert.Empty(t, nil)
	// 范围类
	assert.InDelta(t, 6, 7, 1)
	assert.WithinDuration(t, time.Now().Add(time.Second), time.Now().AddDate(0,0,1), time.Hour*24)
	//assert.InEpsilon()
	//// http 类
	//r, err := resty.R().Get("www.baidu.com")
	//assert.NoError(t, err)
	//assert.HTTPSuccess(t, r.StatusCode())
	assert.JSONEq(t, `[{"1":1, "1":1}]`, `[{"1":1}]`)
	assert.Nil(t, nil)
	assert.Regexp(t, `$`, "sdf$fdf")
	assert.Panics(t, func() {panic("111")})
}