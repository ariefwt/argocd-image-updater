// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	plumbing "github.com/go-git/go-git/v5/plumbing"
	mock "github.com/stretchr/testify/mock"
)

// gitRefCache is an autogenerated mock type for the gitRefCache type
type gitRefCache struct {
	mock.Mock
}

// GetGitReferences provides a mock function with given fields: repo, references
func (_m *gitRefCache) GetGitReferences(repo string, references *[]*plumbing.Reference) error {
	ret := _m.Called(repo, references)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *[]*plumbing.Reference) error); ok {
		r0 = rf(repo, references)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGitReferences provides a mock function with given fields: repo, references
func (_m *gitRefCache) SetGitReferences(repo string, references []*plumbing.Reference) error {
	ret := _m.Called(repo, references)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []*plumbing.Reference) error); ok {
		r0 = rf(repo, references)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}