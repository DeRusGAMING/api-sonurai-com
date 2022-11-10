// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	store "api/internal/store"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *Storer) Get(ctx context.Context, id string) (*store.WallpaperWithTags, error) {
	ret := _m.Called(ctx, id)

	var r0 *store.WallpaperWithTags
	if rf, ok := ret.Get(0).(func(context.Context, string) *store.WallpaperWithTags); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.WallpaperWithTags)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByOldID provides a mock function with given fields: ctx, id
func (_m *Storer) GetByOldID(ctx context.Context, id int) (*store.WallpaperWithTags, error) {
	ret := _m.Called(ctx, id)

	var r0 *store.WallpaperWithTags
	if rf, ok := ret.Get(0).(func(context.Context, int) *store.WallpaperWithTags); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.WallpaperWithTags)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, q
func (_m *Storer) List(ctx context.Context, q store.ListQuery) ([]store.Wallpaper, error) {
	ret := _m.Called(ctx, q)

	var r0 []store.Wallpaper
	if rf, ok := ret.Get(0).(func(context.Context, store.ListQuery) []store.Wallpaper); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]store.Wallpaper)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.ListQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByTag provides a mock function with given fields: ctx, tag, after
func (_m *Storer) ListByTag(ctx context.Context, tag string, after float64) ([]store.Wallpaper, float64, error) {
	ret := _m.Called(ctx, tag, after)

	var r0 []store.Wallpaper
	if rf, ok := ret.Get(0).(func(context.Context, string, float64) []store.Wallpaper); ok {
		r0 = rf(ctx, tag, after)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]store.Wallpaper)
		}
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(context.Context, string, float64) float64); ok {
		r1 = rf(ctx, tag, after)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, float64) error); ok {
		r2 = rf(ctx, tag, after)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewStorer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorer(t mockConstructorTestingTNewStorer) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
