// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hanwen/go-fuse/v2/fuse/api.go

// Package mock_fuse is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fuse "github.com/hanwen/go-fuse/v2/fuse"
)

// MockReadResult is a mock of ReadResult interface.
type MockReadResult struct {
	ctrl     *gomock.Controller
	recorder *MockReadResultMockRecorder
}

// MockReadResultMockRecorder is the mock recorder for MockReadResult.
type MockReadResultMockRecorder struct {
	mock *MockReadResult
}

// NewMockReadResult creates a new mock instance.
func NewMockReadResult(ctrl *gomock.Controller) *MockReadResult {
	mock := &MockReadResult{ctrl: ctrl}
	mock.recorder = &MockReadResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadResult) EXPECT() *MockReadResultMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockReadResult) Bytes(buf []byte) ([]byte, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes", buf)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// Bytes indicates an expected call of Bytes.
func (mr *MockReadResultMockRecorder) Bytes(buf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockReadResult)(nil).Bytes), buf)
}

// Done mocks base method.
func (m *MockReadResult) Done() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Done")
}

// Done indicates an expected call of Done.
func (mr *MockReadResultMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockReadResult)(nil).Done))
}

// Size mocks base method.
func (m *MockReadResult) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockReadResultMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockReadResult)(nil).Size))
}

// MockRawFileSystem is a mock of RawFileSystem interface.
type MockRawFileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockRawFileSystemMockRecorder
}

// MockRawFileSystemMockRecorder is the mock recorder for MockRawFileSystem.
type MockRawFileSystemMockRecorder struct {
	mock *MockRawFileSystem
}

// NewMockRawFileSystem creates a new mock instance.
func NewMockRawFileSystem(ctrl *gomock.Controller) *MockRawFileSystem {
	mock := &MockRawFileSystem{ctrl: ctrl}
	mock.recorder = &MockRawFileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRawFileSystem) EXPECT() *MockRawFileSystemMockRecorder {
	return m.recorder
}

// Access mocks base method.
func (m *MockRawFileSystem) Access(cancel <-chan struct{}, input *fuse.AccessIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Access", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Access indicates an expected call of Access.
func (mr *MockRawFileSystemMockRecorder) Access(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Access", reflect.TypeOf((*MockRawFileSystem)(nil).Access), cancel, input)
}

// CopyFileRange mocks base method.
func (m *MockRawFileSystem) CopyFileRange(cancel <-chan struct{}, input *fuse.CopyFileRangeIn) (uint32, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFileRange", cancel, input)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// CopyFileRange indicates an expected call of CopyFileRange.
func (mr *MockRawFileSystemMockRecorder) CopyFileRange(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFileRange", reflect.TypeOf((*MockRawFileSystem)(nil).CopyFileRange), cancel, input)
}

// Create mocks base method.
func (m *MockRawFileSystem) Create(cancel <-chan struct{}, input *fuse.CreateIn, name string, out *fuse.CreateOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cancel, input, name, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRawFileSystemMockRecorder) Create(cancel, input, name, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRawFileSystem)(nil).Create), cancel, input, name, out)
}

// Fallocate mocks base method.
func (m *MockRawFileSystem) Fallocate(cancel <-chan struct{}, input *fuse.FallocateIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fallocate", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Fallocate indicates an expected call of Fallocate.
func (mr *MockRawFileSystemMockRecorder) Fallocate(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fallocate", reflect.TypeOf((*MockRawFileSystem)(nil).Fallocate), cancel, input)
}

// Flush mocks base method.
func (m *MockRawFileSystem) Flush(cancel <-chan struct{}, input *fuse.FlushIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockRawFileSystemMockRecorder) Flush(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockRawFileSystem)(nil).Flush), cancel, input)
}

// Forget mocks base method.
func (m *MockRawFileSystem) Forget(nodeid, nlookup uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Forget", nodeid, nlookup)
}

// Forget indicates an expected call of Forget.
func (mr *MockRawFileSystemMockRecorder) Forget(nodeid, nlookup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forget", reflect.TypeOf((*MockRawFileSystem)(nil).Forget), nodeid, nlookup)
}

// Fsync mocks base method.
func (m *MockRawFileSystem) Fsync(cancel <-chan struct{}, input *fuse.FsyncIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fsync", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Fsync indicates an expected call of Fsync.
func (mr *MockRawFileSystemMockRecorder) Fsync(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fsync", reflect.TypeOf((*MockRawFileSystem)(nil).Fsync), cancel, input)
}

// FsyncDir mocks base method.
func (m *MockRawFileSystem) FsyncDir(cancel <-chan struct{}, input *fuse.FsyncIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FsyncDir", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// FsyncDir indicates an expected call of FsyncDir.
func (mr *MockRawFileSystemMockRecorder) FsyncDir(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FsyncDir", reflect.TypeOf((*MockRawFileSystem)(nil).FsyncDir), cancel, input)
}

// GetAttr mocks base method.
func (m *MockRawFileSystem) GetAttr(cancel <-chan struct{}, input *fuse.GetAttrIn, out *fuse.AttrOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttr", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// GetAttr indicates an expected call of GetAttr.
func (mr *MockRawFileSystemMockRecorder) GetAttr(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttr", reflect.TypeOf((*MockRawFileSystem)(nil).GetAttr), cancel, input, out)
}

// GetLk mocks base method.
func (m *MockRawFileSystem) GetLk(cancel <-chan struct{}, input *fuse.LkIn, out *fuse.LkOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLk", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// GetLk indicates an expected call of GetLk.
func (mr *MockRawFileSystemMockRecorder) GetLk(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLk", reflect.TypeOf((*MockRawFileSystem)(nil).GetLk), cancel, input, out)
}

// GetXAttr mocks base method.
func (m *MockRawFileSystem) GetXAttr(cancel <-chan struct{}, header *fuse.InHeader, attr string, dest []byte) (uint32, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetXAttr", cancel, header, attr, dest)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// GetXAttr indicates an expected call of GetXAttr.
func (mr *MockRawFileSystemMockRecorder) GetXAttr(cancel, header, attr, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXAttr", reflect.TypeOf((*MockRawFileSystem)(nil).GetXAttr), cancel, header, attr, dest)
}

// Init mocks base method.
func (m *MockRawFileSystem) Init(arg0 *fuse.Server) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0)
}

// Init indicates an expected call of Init.
func (mr *MockRawFileSystemMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRawFileSystem)(nil).Init), arg0)
}

// Link mocks base method.
func (m *MockRawFileSystem) Link(cancel <-chan struct{}, input *fuse.LinkIn, filename string, out *fuse.EntryOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link", cancel, input, filename, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Link indicates an expected call of Link.
func (mr *MockRawFileSystemMockRecorder) Link(cancel, input, filename, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockRawFileSystem)(nil).Link), cancel, input, filename, out)
}

// ListXAttr mocks base method.
func (m *MockRawFileSystem) ListXAttr(cancel <-chan struct{}, header *fuse.InHeader, dest []byte) (uint32, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListXAttr", cancel, header, dest)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// ListXAttr indicates an expected call of ListXAttr.
func (mr *MockRawFileSystemMockRecorder) ListXAttr(cancel, header, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListXAttr", reflect.TypeOf((*MockRawFileSystem)(nil).ListXAttr), cancel, header, dest)
}

// Lookup mocks base method.
func (m *MockRawFileSystem) Lookup(cancel <-chan struct{}, header *fuse.InHeader, name string, out *fuse.EntryOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", cancel, header, name, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Lookup indicates an expected call of Lookup.
func (mr *MockRawFileSystemMockRecorder) Lookup(cancel, header, name, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockRawFileSystem)(nil).Lookup), cancel, header, name, out)
}

// Lseek mocks base method.
func (m *MockRawFileSystem) Lseek(cancel <-chan struct{}, in *fuse.LseekIn, out *fuse.LseekOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lseek", cancel, in, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Lseek indicates an expected call of Lseek.
func (mr *MockRawFileSystemMockRecorder) Lseek(cancel, in, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lseek", reflect.TypeOf((*MockRawFileSystem)(nil).Lseek), cancel, in, out)
}

// Mkdir mocks base method.
func (m *MockRawFileSystem) Mkdir(cancel <-chan struct{}, input *fuse.MkdirIn, name string, out *fuse.EntryOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", cancel, input, name, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockRawFileSystemMockRecorder) Mkdir(cancel, input, name, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockRawFileSystem)(nil).Mkdir), cancel, input, name, out)
}

// Mknod mocks base method.
func (m *MockRawFileSystem) Mknod(cancel <-chan struct{}, input *fuse.MknodIn, name string, out *fuse.EntryOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mknod", cancel, input, name, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Mknod indicates an expected call of Mknod.
func (mr *MockRawFileSystemMockRecorder) Mknod(cancel, input, name, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mknod", reflect.TypeOf((*MockRawFileSystem)(nil).Mknod), cancel, input, name, out)
}

// Open mocks base method.
func (m *MockRawFileSystem) Open(cancel <-chan struct{}, input *fuse.OpenIn, out *fuse.OpenOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockRawFileSystemMockRecorder) Open(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockRawFileSystem)(nil).Open), cancel, input, out)
}

// OpenDir mocks base method.
func (m *MockRawFileSystem) OpenDir(cancel <-chan struct{}, input *fuse.OpenIn, out *fuse.OpenOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenDir", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// OpenDir indicates an expected call of OpenDir.
func (mr *MockRawFileSystemMockRecorder) OpenDir(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenDir", reflect.TypeOf((*MockRawFileSystem)(nil).OpenDir), cancel, input, out)
}

// Read mocks base method.
func (m *MockRawFileSystem) Read(cancel <-chan struct{}, input *fuse.ReadIn, buf []byte) (fuse.ReadResult, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", cancel, input, buf)
	ret0, _ := ret[0].(fuse.ReadResult)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRawFileSystemMockRecorder) Read(cancel, input, buf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRawFileSystem)(nil).Read), cancel, input, buf)
}

// ReadDir mocks base method.
func (m *MockRawFileSystem) ReadDir(cancel <-chan struct{}, input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockRawFileSystemMockRecorder) ReadDir(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockRawFileSystem)(nil).ReadDir), cancel, input, out)
}

// ReadDirPlus mocks base method.
func (m *MockRawFileSystem) ReadDirPlus(cancel <-chan struct{}, input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDirPlus", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// ReadDirPlus indicates an expected call of ReadDirPlus.
func (mr *MockRawFileSystemMockRecorder) ReadDirPlus(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDirPlus", reflect.TypeOf((*MockRawFileSystem)(nil).ReadDirPlus), cancel, input, out)
}

// Readlink mocks base method.
func (m *MockRawFileSystem) Readlink(cancel <-chan struct{}, header *fuse.InHeader) ([]byte, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readlink", cancel, header)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// Readlink indicates an expected call of Readlink.
func (mr *MockRawFileSystemMockRecorder) Readlink(cancel, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readlink", reflect.TypeOf((*MockRawFileSystem)(nil).Readlink), cancel, header)
}

// Release mocks base method.
func (m *MockRawFileSystem) Release(cancel <-chan struct{}, input *fuse.ReleaseIn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release", cancel, input)
}

// Release indicates an expected call of Release.
func (mr *MockRawFileSystemMockRecorder) Release(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockRawFileSystem)(nil).Release), cancel, input)
}

// ReleaseDir mocks base method.
func (m *MockRawFileSystem) ReleaseDir(input *fuse.ReleaseIn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseDir", input)
}

// ReleaseDir indicates an expected call of ReleaseDir.
func (mr *MockRawFileSystemMockRecorder) ReleaseDir(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseDir", reflect.TypeOf((*MockRawFileSystem)(nil).ReleaseDir), input)
}

// RemoveXAttr mocks base method.
func (m *MockRawFileSystem) RemoveXAttr(cancel <-chan struct{}, header *fuse.InHeader, attr string) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveXAttr", cancel, header, attr)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// RemoveXAttr indicates an expected call of RemoveXAttr.
func (mr *MockRawFileSystemMockRecorder) RemoveXAttr(cancel, header, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveXAttr", reflect.TypeOf((*MockRawFileSystem)(nil).RemoveXAttr), cancel, header, attr)
}

// Rename mocks base method.
func (m *MockRawFileSystem) Rename(cancel <-chan struct{}, input *fuse.RenameIn, oldName, newName string) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", cancel, input, oldName, newName)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockRawFileSystemMockRecorder) Rename(cancel, input, oldName, newName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockRawFileSystem)(nil).Rename), cancel, input, oldName, newName)
}

// Rmdir mocks base method.
func (m *MockRawFileSystem) Rmdir(cancel <-chan struct{}, header *fuse.InHeader, name string) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rmdir", cancel, header, name)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Rmdir indicates an expected call of Rmdir.
func (mr *MockRawFileSystemMockRecorder) Rmdir(cancel, header, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rmdir", reflect.TypeOf((*MockRawFileSystem)(nil).Rmdir), cancel, header, name)
}

// SetAttr mocks base method.
func (m *MockRawFileSystem) SetAttr(cancel <-chan struct{}, input *fuse.SetAttrIn, out *fuse.AttrOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAttr", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// SetAttr indicates an expected call of SetAttr.
func (mr *MockRawFileSystemMockRecorder) SetAttr(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAttr", reflect.TypeOf((*MockRawFileSystem)(nil).SetAttr), cancel, input, out)
}

// SetDebug mocks base method.
func (m *MockRawFileSystem) SetDebug(debug bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDebug", debug)
}

// SetDebug indicates an expected call of SetDebug.
func (mr *MockRawFileSystemMockRecorder) SetDebug(debug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDebug", reflect.TypeOf((*MockRawFileSystem)(nil).SetDebug), debug)
}

// SetLk mocks base method.
func (m *MockRawFileSystem) SetLk(cancel <-chan struct{}, input *fuse.LkIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLk", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// SetLk indicates an expected call of SetLk.
func (mr *MockRawFileSystemMockRecorder) SetLk(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLk", reflect.TypeOf((*MockRawFileSystem)(nil).SetLk), cancel, input)
}

// SetLkw mocks base method.
func (m *MockRawFileSystem) SetLkw(cancel <-chan struct{}, input *fuse.LkIn) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLkw", cancel, input)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// SetLkw indicates an expected call of SetLkw.
func (mr *MockRawFileSystemMockRecorder) SetLkw(cancel, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLkw", reflect.TypeOf((*MockRawFileSystem)(nil).SetLkw), cancel, input)
}

// SetXAttr mocks base method.
func (m *MockRawFileSystem) SetXAttr(cancel <-chan struct{}, input *fuse.SetXAttrIn, attr string, data []byte) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetXAttr", cancel, input, attr, data)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// SetXAttr indicates an expected call of SetXAttr.
func (mr *MockRawFileSystemMockRecorder) SetXAttr(cancel, input, attr, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetXAttr", reflect.TypeOf((*MockRawFileSystem)(nil).SetXAttr), cancel, input, attr, data)
}

// StatFs mocks base method.
func (m *MockRawFileSystem) StatFs(cancel <-chan struct{}, input *fuse.InHeader, out *fuse.StatfsOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatFs", cancel, input, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// StatFs indicates an expected call of StatFs.
func (mr *MockRawFileSystemMockRecorder) StatFs(cancel, input, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatFs", reflect.TypeOf((*MockRawFileSystem)(nil).StatFs), cancel, input, out)
}

// String mocks base method.
func (m *MockRawFileSystem) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockRawFileSystemMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockRawFileSystem)(nil).String))
}

// Symlink mocks base method.
func (m *MockRawFileSystem) Symlink(cancel <-chan struct{}, header *fuse.InHeader, pointedTo, linkName string, out *fuse.EntryOut) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Symlink", cancel, header, pointedTo, linkName, out)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Symlink indicates an expected call of Symlink.
func (mr *MockRawFileSystemMockRecorder) Symlink(cancel, header, pointedTo, linkName, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Symlink", reflect.TypeOf((*MockRawFileSystem)(nil).Symlink), cancel, header, pointedTo, linkName, out)
}

// Unlink mocks base method.
func (m *MockRawFileSystem) Unlink(cancel <-chan struct{}, header *fuse.InHeader, name string) fuse.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlink", cancel, header, name)
	ret0, _ := ret[0].(fuse.Status)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockRawFileSystemMockRecorder) Unlink(cancel, header, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockRawFileSystem)(nil).Unlink), cancel, header, name)
}

// Write mocks base method.
func (m *MockRawFileSystem) Write(cancel <-chan struct{}, input *fuse.WriteIn, data []byte) (uint32, fuse.Status) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", cancel, input, data)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(fuse.Status)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockRawFileSystemMockRecorder) Write(cancel, input, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockRawFileSystem)(nil).Write), cancel, input, data)
}
