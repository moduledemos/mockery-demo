// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Story is an autogenerated mock type for the Story type
type Story struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Story) Close() {
	_m.Called()
}

// HasChapter provides a mock function with given fields: chapter
func (_m *Story) HasChapter(chapter int) bool {
	ret := _m.Called(chapter)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(chapter)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Open provides a mock function with given fields:
func (_m *Story) Open() {
	_m.Called()
}

// ReadChapter provides a mock function with given fields: chapter
func (_m *Story) ReadChapter(chapter int) ([]byte, error) {
	ret := _m.Called(chapter)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(chapter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(chapter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}