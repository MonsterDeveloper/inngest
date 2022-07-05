// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/execution/state/state.go

// Package state is a generated GoMock package.
package state

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	inngest "github.com/inngest/inngest-cli/inngest"
	v2 "github.com/oklog/ulid/v2"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// ActionComplete mocks base method.
func (m *MockState) ActionComplete(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionComplete", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ActionComplete indicates an expected call of ActionComplete.
func (mr *MockStateMockRecorder) ActionComplete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionComplete", reflect.TypeOf((*MockState)(nil).ActionComplete), id)
}

// ActionID mocks base method.
func (m *MockState) ActionID(id string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionID", id)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionID indicates an expected call of ActionID.
func (mr *MockStateMockRecorder) ActionID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionID", reflect.TypeOf((*MockState)(nil).ActionID), id)
}

// Actions mocks base method.
func (m *MockState) Actions() map[string]map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(map[string]map[string]interface{})
	return ret0
}

// Actions indicates an expected call of Actions.
func (mr *MockStateMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockState)(nil).Actions))
}

// Errors mocks base method.
func (m *MockState) Errors() map[string]error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].(map[string]error)
	return ret0
}

// Errors indicates an expected call of Errors.
func (mr *MockStateMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*MockState)(nil).Errors))
}

// Event mocks base method.
func (m *MockState) Event() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Event")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Event indicates an expected call of Event.
func (mr *MockStateMockRecorder) Event() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockState)(nil).Event))
}

// Identifier mocks base method.
func (m *MockState) Identifier() Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(Identifier)
	return ret0
}

// Identifier indicates an expected call of Identifier.
func (mr *MockStateMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockState)(nil).Identifier))
}

// Metadata mocks base method.
func (m *MockState) Metadata() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Metadata indicates an expected call of Metadata.
func (mr *MockStateMockRecorder) Metadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockState)(nil).Metadata))
}

// RunID mocks base method.
func (m *MockState) RunID() v2.ULID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunID")
	ret0, _ := ret[0].(v2.ULID)
	return ret0
}

// RunID indicates an expected call of RunID.
func (mr *MockStateMockRecorder) RunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunID", reflect.TypeOf((*MockState)(nil).RunID))
}

// Workflow mocks base method.
func (m *MockState) Workflow() inngest.Workflow {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Workflow")
	ret0, _ := ret[0].(inngest.Workflow)
	return ret0
}

// Workflow indicates an expected call of Workflow.
func (mr *MockStateMockRecorder) Workflow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Workflow", reflect.TypeOf((*MockState)(nil).Workflow))
}

// WorkflowID mocks base method.
func (m *MockState) WorkflowID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkflowID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// WorkflowID indicates an expected call of WorkflowID.
func (mr *MockStateMockRecorder) WorkflowID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkflowID", reflect.TypeOf((*MockState)(nil).WorkflowID))
}

// MockLoader is a mock of Loader interface.
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader.
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance.
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockLoader) Load(ctx context.Context, i Identifier) (State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", ctx, i)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockLoaderMockRecorder) Load(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockLoader)(nil).Load), ctx, i)
}

// MockMutater is a mock of Mutater interface.
type MockMutater struct {
	ctrl     *gomock.Controller
	recorder *MockMutaterMockRecorder
}

// MockMutaterMockRecorder is the mock recorder for MockMutater.
type MockMutaterMockRecorder struct {
	mock *MockMutater
}

// NewMockMutater creates a new mock instance.
func NewMockMutater(ctrl *gomock.Controller) *MockMutater {
	mock := &MockMutater{ctrl: ctrl}
	mock.recorder = &MockMutaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutater) EXPECT() *MockMutaterMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockMutater) New(ctx context.Context, workflow inngest.Workflow, runID v2.ULID, input map[string]any) (State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", ctx, workflow, runID, input)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockMutaterMockRecorder) New(ctx, workflow, runID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockMutater)(nil).New), ctx, workflow, runID, input)
}

// SaveResponse mocks base method.
func (m *MockMutater) SaveResponse(ctx context.Context, i Identifier, r DriverResponse, attempt int) (State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResponse", ctx, i, r, attempt)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveResponse indicates an expected call of SaveResponse.
func (mr *MockMutaterMockRecorder) SaveResponse(ctx, i, r, attempt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResponse", reflect.TypeOf((*MockMutater)(nil).SaveResponse), ctx, i, r, attempt)
}

// Scheduled mocks base method.
func (m *MockMutater) Scheduled(ctx context.Context, i Identifier, step inngest.Step) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheduled", ctx, i, step)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scheduled indicates an expected call of Scheduled.
func (mr *MockMutaterMockRecorder) Scheduled(ctx, i, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheduled", reflect.TypeOf((*MockMutater)(nil).Scheduled), ctx, i, step)
}

// MockPauseMutater is a mock of PauseMutater interface.
type MockPauseMutater struct {
	ctrl     *gomock.Controller
	recorder *MockPauseMutaterMockRecorder
}

// MockPauseMutaterMockRecorder is the mock recorder for MockPauseMutater.
type MockPauseMutaterMockRecorder struct {
	mock *MockPauseMutater
}

// NewMockPauseMutater creates a new mock instance.
func NewMockPauseMutater(ctrl *gomock.Controller) *MockPauseMutater {
	mock := &MockPauseMutater{ctrl: ctrl}
	mock.recorder = &MockPauseMutaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPauseMutater) EXPECT() *MockPauseMutaterMockRecorder {
	return m.recorder
}

// ConsumePause mocks base method.
func (m *MockPauseMutater) ConsumePause(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumePause", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumePause indicates an expected call of ConsumePause.
func (mr *MockPauseMutaterMockRecorder) ConsumePause(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumePause", reflect.TypeOf((*MockPauseMutater)(nil).ConsumePause), ctx, id)
}

// LeasePause mocks base method.
func (m *MockPauseMutater) LeasePause(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeasePause", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeasePause indicates an expected call of LeasePause.
func (mr *MockPauseMutaterMockRecorder) LeasePause(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeasePause", reflect.TypeOf((*MockPauseMutater)(nil).LeasePause), ctx, id)
}

// SavePause mocks base method.
func (m *MockPauseMutater) SavePause(ctx context.Context, p Pause) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePause", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePause indicates an expected call of SavePause.
func (mr *MockPauseMutaterMockRecorder) SavePause(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePause", reflect.TypeOf((*MockPauseMutater)(nil).SavePause), ctx, p)
}

// MockPauseGetter is a mock of PauseGetter interface.
type MockPauseGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPauseGetterMockRecorder
}

// MockPauseGetterMockRecorder is the mock recorder for MockPauseGetter.
type MockPauseGetterMockRecorder struct {
	mock *MockPauseGetter
}

// NewMockPauseGetter creates a new mock instance.
func NewMockPauseGetter(ctrl *gomock.Controller) *MockPauseGetter {
	mock := &MockPauseGetter{ctrl: ctrl}
	mock.recorder = &MockPauseGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPauseGetter) EXPECT() *MockPauseGetterMockRecorder {
	return m.recorder
}

// PauseByStep mocks base method.
func (m *MockPauseGetter) PauseByStep(ctx context.Context, i Identifier, actionID string) (*Pause, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseByStep", ctx, i, actionID)
	ret0, _ := ret[0].(*Pause)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseByStep indicates an expected call of PauseByStep.
func (mr *MockPauseGetterMockRecorder) PauseByStep(ctx, i, actionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseByStep", reflect.TypeOf((*MockPauseGetter)(nil).PauseByStep), ctx, i, actionID)
}

// PausesByEvent mocks base method.
func (m *MockPauseGetter) PausesByEvent(ctx context.Context, eventName string) (PauseIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PausesByEvent", ctx, eventName)
	ret0, _ := ret[0].(PauseIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PausesByEvent indicates an expected call of PausesByEvent.
func (mr *MockPauseGetterMockRecorder) PausesByEvent(ctx, eventName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PausesByEvent", reflect.TypeOf((*MockPauseGetter)(nil).PausesByEvent), ctx, eventName)
}

// MockPauseIterator is a mock of PauseIterator interface.
type MockPauseIterator struct {
	ctrl     *gomock.Controller
	recorder *MockPauseIteratorMockRecorder
}

// MockPauseIteratorMockRecorder is the mock recorder for MockPauseIterator.
type MockPauseIteratorMockRecorder struct {
	mock *MockPauseIterator
}

// NewMockPauseIterator creates a new mock instance.
func NewMockPauseIterator(ctrl *gomock.Controller) *MockPauseIterator {
	mock := &MockPauseIterator{ctrl: ctrl}
	mock.recorder = &MockPauseIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPauseIterator) EXPECT() *MockPauseIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockPauseIterator) Next(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockPauseIteratorMockRecorder) Next(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockPauseIterator)(nil).Next), ctx)
}

// Val mocks base method.
func (m *MockPauseIterator) Val(arg0 context.Context) *Pause {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Val", arg0)
	ret0, _ := ret[0].(*Pause)
	return ret0
}

// Val indicates an expected call of Val.
func (mr *MockPauseIteratorMockRecorder) Val(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Val", reflect.TypeOf((*MockPauseIterator)(nil).Val), arg0)
}

// MockPauseManager is a mock of PauseManager interface.
type MockPauseManager struct {
	ctrl     *gomock.Controller
	recorder *MockPauseManagerMockRecorder
}

// MockPauseManagerMockRecorder is the mock recorder for MockPauseManager.
type MockPauseManagerMockRecorder struct {
	mock *MockPauseManager
}

// NewMockPauseManager creates a new mock instance.
func NewMockPauseManager(ctrl *gomock.Controller) *MockPauseManager {
	mock := &MockPauseManager{ctrl: ctrl}
	mock.recorder = &MockPauseManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPauseManager) EXPECT() *MockPauseManagerMockRecorder {
	return m.recorder
}

// ConsumePause mocks base method.
func (m *MockPauseManager) ConsumePause(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumePause", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumePause indicates an expected call of ConsumePause.
func (mr *MockPauseManagerMockRecorder) ConsumePause(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumePause", reflect.TypeOf((*MockPauseManager)(nil).ConsumePause), ctx, id)
}

// LeasePause mocks base method.
func (m *MockPauseManager) LeasePause(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeasePause", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeasePause indicates an expected call of LeasePause.
func (mr *MockPauseManagerMockRecorder) LeasePause(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeasePause", reflect.TypeOf((*MockPauseManager)(nil).LeasePause), ctx, id)
}

// PauseByStep mocks base method.
func (m *MockPauseManager) PauseByStep(ctx context.Context, i Identifier, actionID string) (*Pause, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseByStep", ctx, i, actionID)
	ret0, _ := ret[0].(*Pause)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseByStep indicates an expected call of PauseByStep.
func (mr *MockPauseManagerMockRecorder) PauseByStep(ctx, i, actionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseByStep", reflect.TypeOf((*MockPauseManager)(nil).PauseByStep), ctx, i, actionID)
}

// PausesByEvent mocks base method.
func (m *MockPauseManager) PausesByEvent(ctx context.Context, eventName string) (PauseIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PausesByEvent", ctx, eventName)
	ret0, _ := ret[0].(PauseIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PausesByEvent indicates an expected call of PausesByEvent.
func (mr *MockPauseManagerMockRecorder) PausesByEvent(ctx, eventName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PausesByEvent", reflect.TypeOf((*MockPauseManager)(nil).PausesByEvent), ctx, eventName)
}

// SavePause mocks base method.
func (m *MockPauseManager) SavePause(ctx context.Context, p Pause) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePause", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePause indicates an expected call of SavePause.
func (mr *MockPauseManagerMockRecorder) SavePause(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePause", reflect.TypeOf((*MockPauseManager)(nil).SavePause), ctx, p)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// ConsumePause mocks base method.
func (m *MockManager) ConsumePause(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumePause", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumePause indicates an expected call of ConsumePause.
func (mr *MockManagerMockRecorder) ConsumePause(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumePause", reflect.TypeOf((*MockManager)(nil).ConsumePause), ctx, id)
}

// LeasePause mocks base method.
func (m *MockManager) LeasePause(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeasePause", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeasePause indicates an expected call of LeasePause.
func (mr *MockManagerMockRecorder) LeasePause(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeasePause", reflect.TypeOf((*MockManager)(nil).LeasePause), ctx, id)
}

// Load mocks base method.
func (m *MockManager) Load(ctx context.Context, i Identifier) (State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", ctx, i)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load.
func (mr *MockManagerMockRecorder) Load(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockManager)(nil).Load), ctx, i)
}

// New mocks base method.
func (m *MockManager) New(ctx context.Context, workflow inngest.Workflow, runID v2.ULID, input map[string]any) (State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", ctx, workflow, runID, input)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockManagerMockRecorder) New(ctx, workflow, runID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockManager)(nil).New), ctx, workflow, runID, input)
}

// PauseByStep mocks base method.
func (m *MockManager) PauseByStep(ctx context.Context, i Identifier, actionID string) (*Pause, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseByStep", ctx, i, actionID)
	ret0, _ := ret[0].(*Pause)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseByStep indicates an expected call of PauseByStep.
func (mr *MockManagerMockRecorder) PauseByStep(ctx, i, actionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseByStep", reflect.TypeOf((*MockManager)(nil).PauseByStep), ctx, i, actionID)
}

// PausesByEvent mocks base method.
func (m *MockManager) PausesByEvent(ctx context.Context, eventName string) (PauseIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PausesByEvent", ctx, eventName)
	ret0, _ := ret[0].(PauseIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PausesByEvent indicates an expected call of PausesByEvent.
func (mr *MockManagerMockRecorder) PausesByEvent(ctx, eventName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PausesByEvent", reflect.TypeOf((*MockManager)(nil).PausesByEvent), ctx, eventName)
}

// SavePause mocks base method.
func (m *MockManager) SavePause(ctx context.Context, p Pause) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePause", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePause indicates an expected call of SavePause.
func (mr *MockManagerMockRecorder) SavePause(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePause", reflect.TypeOf((*MockManager)(nil).SavePause), ctx, p)
}

// SaveResponse mocks base method.
func (m *MockManager) SaveResponse(ctx context.Context, i Identifier, r DriverResponse, attempt int) (State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResponse", ctx, i, r, attempt)
	ret0, _ := ret[0].(State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveResponse indicates an expected call of SaveResponse.
func (mr *MockManagerMockRecorder) SaveResponse(ctx, i, r, attempt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResponse", reflect.TypeOf((*MockManager)(nil).SaveResponse), ctx, i, r, attempt)
}

// Scheduled mocks base method.
func (m *MockManager) Scheduled(ctx context.Context, i Identifier, step inngest.Step) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheduled", ctx, i, step)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scheduled indicates an expected call of Scheduled.
func (mr *MockManagerMockRecorder) Scheduled(ctx, i, step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheduled", reflect.TypeOf((*MockManager)(nil).Scheduled), ctx, i, step)
}