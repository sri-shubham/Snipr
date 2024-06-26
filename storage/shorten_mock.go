// Code generated by MockGen. DO NOT EDIT.
// Source: shortenedURL.go

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sri-shubham/snipr/storage/models"
)

// MockURLStorage is a mock of URLStorage interface.
type MockURLStorage struct {
	ctrl     *gomock.Controller
	recorder *MockURLStorageMockRecorder
}

// MockURLStorageMockRecorder is the mock recorder for MockURLStorage.
type MockURLStorageMockRecorder struct {
	mock *MockURLStorage
}

// NewMockURLStorage creates a new mock instance.
func NewMockURLStorage(ctrl *gomock.Controller) *MockURLStorage {
	mock := &MockURLStorage{ctrl: ctrl}
	mock.recorder = &MockURLStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLStorage) EXPECT() *MockURLStorageMockRecorder {
	return m.recorder
}

// GetOriginalURL mocks base method.
func (m *MockURLStorage) GetOriginalURL(ctx context.Context, shortURL string) (*models.ShortenedURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginalURL", ctx, shortURL)
	ret0, _ := ret[0].(*models.ShortenedURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginalURL indicates an expected call of GetOriginalURL.
func (mr *MockURLStorageMockRecorder) GetOriginalURL(ctx, shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginalURL", reflect.TypeOf((*MockURLStorage)(nil).GetOriginalURL), ctx, shortURL)
}

// StoreShortURL mocks base method.
func (m *MockURLStorage) StoreShortURL(ctx context.Context, shortUrl *models.ShortenedURL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreShortURL", ctx, shortUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreShortURL indicates an expected call of StoreShortURL.
func (mr *MockURLStorageMockRecorder) StoreShortURL(ctx, shortUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreShortURL", reflect.TypeOf((*MockURLStorage)(nil).StoreShortURL), ctx, shortUrl)
}
