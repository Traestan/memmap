package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func TestMapMemSuite(t *testing.T) {
// 	suite.Run(t, new(endpointSuite))
// }

// type endpointSuite struct {
// 	suite.Suite
// 	service NewMemCache
// }

// func (s *endpointSuite) SetupTest() {
// 	s.service = NewMemCache
// }
func TestNewMemCache(t *testing.T) {
	t.Log("NewMemCache init")
	mn := NewMemCache()

	//assert.Equal(t, Cache, mn, "The two words should be the same.")
	//assert.Nil(t, mn)
	assert.NotNil(t, mn)
}
func TestPut(t *testing.T) {
	t.Log("NewMemCache Put")
	mn := NewMemCache()
	//testKey := "testMe"
	testCase := []struct {
		name string
		keys string
		args interface{}
		want interface{}
	}{
		{
			"PositiveInt",
			"testMe",
			1,
			1,
		},
		{
			"PositiveString",
			"testMe",
			"match string",
			"match string",
		},
	}
	for _, tc := range testCase {
		t.Log(tc.name)
		timeout := 10 * time.Second
		err := mn.Put(tc.keys, tc.args, timeout)
		assert.Nil(t, err)
		//get
		resp := mn.Get(tc.keys)
		assert.NotNil(t, resp)
		assert.Equal(t, tc.want, resp, "The two words should be the same.")
	}
}

func TestIsExist(t *testing.T) {
	t.Log("NewMemCache IsExist")
	mn := NewMemCache()
	testCase := []struct {
		name string
		keys string
		args interface{}
		want interface{}
	}{
		{
			"Exist",
			"testMe",
			1,
			true,
		},
		{
			"NotExist",
			"failMe",
			nil,
			false,
		},
	}
	for _, tc := range testCase {
		t.Log(tc.name)

		if tc.name == "NotExist" { //fail case
			resp := mn.IsExist(tc.keys)
			assert.NotNil(t, resp)
			assert.Equal(t, tc.want, resp, "The two words should be the same.")
		} else {
			timeout := 10 * time.Second
			err := mn.Put(tc.keys, tc.args, timeout)
			assert.Nil(t, err)
			//get
			resp := mn.IsExist(tc.keys)
			assert.NotNil(t, resp)
			assert.Equal(t, tc.want, resp, "The two words should be the same.")
		}
	}
}
func TestIncr(t *testing.T) {
	t.Log("NewMemCache Incr")
	mn := NewMemCache()
	testCase := []struct {
		name string
		keys string
		args interface{}
		want interface{}
	}{
		{
			"IncPositive",
			"testMe",
			1,
			2,
		},
		{
			"IncPositiveNegativeNumber",
			"testMe",
			-1,
			0,
		},
		{
			"IncPositiveNegativeNumber",
			"testMe0",
			-10,
			-9,
		},
		{
			"IncPositiveZeroNumber",
			"testMe0",
			0,
			1,
		},
	}
	timeout := 10 * time.Second
	for _, tc := range testCase {
		t.Log(tc.name)
		//put to cache
		err := mn.Put(tc.keys, tc.args, timeout)
		assert.Nil(t, err)
		// incr
		respErr := mn.Incr(tc.keys)
		assert.Nil(t, respErr)
		//get
		respGet := mn.Get(tc.keys)

		assert.Equal(t, tc.want, respGet, "The two words should be the same.")
	}
}

//

//Decr(key string) error
