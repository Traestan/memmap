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
// 	service NewMemoryCache
// }

//	func (s *endpointSuite) SetupTest() {
//		s.service = NewMemoryCache
//	}
func TestNewMemoryCache(t *testing.T) {
	t.Log("NewMemoryCache init")
	mn := NewMemoryCache()

	//assert.Equal(t, Cache, mn, "The two words should be the same.")
	//assert.Nil(t, mn)
	assert.NotNil(t, mn)
}
func TestPut(t *testing.T) {
	t.Log("NewMemoryCache Put")
	mn := NewMemoryCache()
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
	t.Log("NewMemoryCache IsExist")
	mn := NewMemoryCache()
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
	t.Log("NewMemoryCache Incr")
	mn := NewMemoryCache()
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
func TestDecr(t *testing.T) {
	t.Log("NewMemoryCache Decr")
	mn := NewMemoryCache()
	testCase := []struct {
		name string
		keys string
		args interface{}
		want interface{}
	}{
		{
			"DecrPositive",
			"testMe",
			1,
			0,
		},
		{
			"DecrPositiveNegativeNumber",
			"testMe",
			-1,
			-2,
		},
		{
			"DecrPositiveNegativeNumber",
			"testMe0",
			-10,
			-11,
		},
		{
			"DecrPositiveZeroNumber",
			"testMe0",
			0,
			-1,
		},
		{
			"DecrPositiveZeroNumber",
			"testMe01",
			999999999999999,
			999999999999998,
		},
		{
			"DecrPositiveZeroNumber",
			"testMe01",
			-999999999999999,
			-1000000000000000,
		},
	}
	timeout := 10 * time.Second
	for _, tc := range testCase {
		t.Log(tc.name)
		//put to cache
		err := mn.Put(tc.keys, tc.args, timeout)
		assert.Nil(t, err)
		// incr
		respErr := mn.Decr(tc.keys)
		assert.Nil(t, respErr)
		//get
		respGet := mn.Get(tc.keys)

		assert.Equal(t, tc.want, respGet, "The two words should be the same.")
	}
}
