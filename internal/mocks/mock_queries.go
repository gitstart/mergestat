// Code generated by MockGen. DO NOT EDIT.
// Source: ../../queries/queries.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	pgx "github.com/jackc/pgx/v4"
	db "github.com/mergestat/mergestat/internal/db"
	queries "github.com/mergestat/mergestat/queries"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CheckRunningImps mocks base method.
func (m *MockQuerier) CheckRunningImps(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRunningImps", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRunningImps indicates an expected call of CheckRunningImps.
func (mr *MockQuerierMockRecorder) CheckRunningImps(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRunningImps", reflect.TypeOf((*MockQuerier)(nil).CheckRunningImps), ctx)
}

// CleanOldRepoSyncQueue mocks base method.
func (m *MockQuerier) CleanOldRepoSyncQueue(ctx context.Context, dollar_1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanOldRepoSyncQueue", ctx, dollar_1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanOldRepoSyncQueue indicates an expected call of CleanOldRepoSyncQueue.
func (mr *MockQuerierMockRecorder) CleanOldRepoSyncQueue(ctx, dollar_1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanOldRepoSyncQueue", reflect.TypeOf((*MockQuerier)(nil).CleanOldRepoSyncQueue), ctx, dollar_1)
}

// DeleteGitHubRepoInfo mocks base method.
func (m *MockQuerier) DeleteGitHubRepoInfo(ctx context.Context, repoID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGitHubRepoInfo", ctx, repoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGitHubRepoInfo indicates an expected call of DeleteGitHubRepoInfo.
func (mr *MockQuerierMockRecorder) DeleteGitHubRepoInfo(ctx, repoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGitHubRepoInfo", reflect.TypeOf((*MockQuerier)(nil).DeleteGitHubRepoInfo), ctx, repoID)
}

// DeleteRemovedRepos mocks base method.
func (m *MockQuerier) DeleteRemovedRepos(ctx context.Context, arg db.DeleteRemovedReposParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRemovedRepos", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRemovedRepos indicates an expected call of DeleteRemovedRepos.
func (mr *MockQuerierMockRecorder) DeleteRemovedRepos(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRemovedRepos", reflect.TypeOf((*MockQuerier)(nil).DeleteRemovedRepos), ctx, arg)
}

// DequeueSyncJob mocks base method.
func (m *MockQuerier) DequeueSyncJob(ctx context.Context) (db.DequeueSyncJobRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DequeueSyncJob", ctx)
	ret0, _ := ret[0].(db.DequeueSyncJobRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DequeueSyncJob indicates an expected call of DequeueSyncJob.
func (mr *MockQuerierMockRecorder) DequeueSyncJob(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DequeueSyncJob", reflect.TypeOf((*MockQuerier)(nil).DequeueSyncJob), ctx)
}

// EnqueueAllSyncs mocks base method.
func (m *MockQuerier) EnqueueAllSyncs(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqueueAllSyncs", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnqueueAllSyncs indicates an expected call of EnqueueAllSyncs.
func (mr *MockQuerierMockRecorder) EnqueueAllSyncs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueAllSyncs", reflect.TypeOf((*MockQuerier)(nil).EnqueueAllSyncs), ctx)
}

// FetchContainerSync mocks base method.
func (m *MockQuerier) FetchContainerSync(ctx context.Context, id uuid.UUID) (db.FetchContainerSyncRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchContainerSync", ctx, id)
	ret0, _ := ret[0].(db.FetchContainerSyncRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchContainerSync indicates an expected call of FetchContainerSync.
func (mr *MockQuerierMockRecorder) FetchContainerSync(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchContainerSync", reflect.TypeOf((*MockQuerier)(nil).FetchContainerSync), ctx, id)
}

// FetchGitHubToken mocks base method.
func (m *MockQuerier) FetchGitHubToken(ctx context.Context, pgpSymDecrypt string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchGitHubToken", ctx, pgpSymDecrypt)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchGitHubToken indicates an expected call of FetchGitHubToken.
func (mr *MockQuerierMockRecorder) FetchGitHubToken(ctx, pgpSymDecrypt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchGitHubToken", reflect.TypeOf((*MockQuerier)(nil).FetchGitHubToken), ctx, pgpSymDecrypt)
}

// GetRepoById mocks base method.
func (m *MockQuerier) GetRepoById(ctx context.Context, id uuid.UUID) (db.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoById", ctx, id)
	ret0, _ := ret[0].(db.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepoById indicates an expected call of GetRepoById.
func (mr *MockQuerierMockRecorder) GetRepoById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoById", reflect.TypeOf((*MockQuerier)(nil).GetRepoById), ctx, id)
}

// GetRepoIDsFromRepoImport mocks base method.
func (m *MockQuerier) GetRepoIDsFromRepoImport(ctx context.Context, arg db.GetRepoIDsFromRepoImportParams) ([]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoIDsFromRepoImport", ctx, arg)
	ret0, _ := ret[0].([]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepoIDsFromRepoImport indicates an expected call of GetRepoIDsFromRepoImport.
func (mr *MockQuerierMockRecorder) GetRepoIDsFromRepoImport(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoIDsFromRepoImport", reflect.TypeOf((*MockQuerier)(nil).GetRepoIDsFromRepoImport), ctx, arg)
}

// GetRepoImportByID mocks base method.
func (m *MockQuerier) GetRepoImportByID(ctx context.Context, id uuid.UUID) (db.MergestatRepoImport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoImportByID", ctx, id)
	ret0, _ := ret[0].(db.MergestatRepoImport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepoImportByID indicates an expected call of GetRepoImportByID.
func (mr *MockQuerierMockRecorder) GetRepoImportByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoImportByID", reflect.TypeOf((*MockQuerier)(nil).GetRepoImportByID), ctx, id)
}

// GetRepoUrlFromImport mocks base method.
func (m *MockQuerier) GetRepoUrlFromImport(ctx context.Context, importid uuid.UUID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoUrlFromImport", ctx, importid)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepoUrlFromImport indicates an expected call of GetRepoUrlFromImport.
func (mr *MockQuerierMockRecorder) GetRepoUrlFromImport(ctx, importid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoUrlFromImport", reflect.TypeOf((*MockQuerier)(nil).GetRepoUrlFromImport), ctx, importid)
}

// InsertGitHubRepoInfo mocks base method.
func (m *MockQuerier) InsertGitHubRepoInfo(ctx context.Context, arg db.InsertGitHubRepoInfoParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertGitHubRepoInfo", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertGitHubRepoInfo indicates an expected call of InsertGitHubRepoInfo.
func (mr *MockQuerierMockRecorder) InsertGitHubRepoInfo(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertGitHubRepoInfo", reflect.TypeOf((*MockQuerier)(nil).InsertGitHubRepoInfo), ctx, arg)
}

// InsertNewDefaultSync mocks base method.
func (m *MockQuerier) InsertNewDefaultSync(ctx context.Context, arg db.InsertNewDefaultSyncParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewDefaultSync", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewDefaultSync indicates an expected call of InsertNewDefaultSync.
func (mr *MockQuerierMockRecorder) InsertNewDefaultSync(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewDefaultSync", reflect.TypeOf((*MockQuerier)(nil).InsertNewDefaultSync), ctx, arg)
}

// InsertSyncJobLog mocks base method.
func (m *MockQuerier) InsertSyncJobLog(ctx context.Context, arg db.InsertSyncJobLogParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSyncJobLog", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertSyncJobLog indicates an expected call of InsertSyncJobLog.
func (mr *MockQuerierMockRecorder) InsertSyncJobLog(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSyncJobLog", reflect.TypeOf((*MockQuerier)(nil).InsertSyncJobLog), ctx, arg)
}

// ListRepoImportsDueForImport mocks base method.
func (m *MockQuerier) ListRepoImportsDueForImport(ctx context.Context) ([]db.ListRepoImportsDueForImportRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepoImportsDueForImport", ctx)
	ret0, _ := ret[0].([]db.ListRepoImportsDueForImportRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepoImportsDueForImport indicates an expected call of ListRepoImportsDueForImport.
func (mr *MockQuerierMockRecorder) ListRepoImportsDueForImport(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepoImportsDueForImport", reflect.TypeOf((*MockQuerier)(nil).ListRepoImportsDueForImport), ctx)
}

// MarkRepoImportAsUpdated mocks base method.
func (m *MockQuerier) MarkRepoImportAsUpdated(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkRepoImportAsUpdated", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkRepoImportAsUpdated indicates an expected call of MarkRepoImportAsUpdated.
func (mr *MockQuerierMockRecorder) MarkRepoImportAsUpdated(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkRepoImportAsUpdated", reflect.TypeOf((*MockQuerier)(nil).MarkRepoImportAsUpdated), ctx, id)
}

// MarkSyncsAsTimedOut mocks base method.
func (m *MockQuerier) MarkSyncsAsTimedOut(ctx context.Context) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkSyncsAsTimedOut", ctx)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkSyncsAsTimedOut indicates an expected call of MarkSyncsAsTimedOut.
func (mr *MockQuerierMockRecorder) MarkSyncsAsTimedOut(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkSyncsAsTimedOut", reflect.TypeOf((*MockQuerier)(nil).MarkSyncsAsTimedOut), ctx)
}

// SetLatestKeepAliveForJob mocks base method.
func (m *MockQuerier) SetLatestKeepAliveForJob(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLatestKeepAliveForJob", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLatestKeepAliveForJob indicates an expected call of SetLatestKeepAliveForJob.
func (mr *MockQuerierMockRecorder) SetLatestKeepAliveForJob(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLatestKeepAliveForJob", reflect.TypeOf((*MockQuerier)(nil).SetLatestKeepAliveForJob), ctx, id)
}

// SetSyncJobStatus mocks base method.
func (m *MockQuerier) SetSyncJobStatus(ctx context.Context, arg db.SetSyncJobStatusParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSyncJobStatus", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSyncJobStatus indicates an expected call of SetSyncJobStatus.
func (mr *MockQuerierMockRecorder) SetSyncJobStatus(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSyncJobStatus", reflect.TypeOf((*MockQuerier)(nil).SetSyncJobStatus), ctx, arg)
}

// UpdateImportStatus mocks base method.
func (m *MockQuerier) UpdateImportStatus(ctx context.Context, arg db.UpdateImportStatusParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImportStatus", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImportStatus indicates an expected call of UpdateImportStatus.
func (mr *MockQuerierMockRecorder) UpdateImportStatus(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImportStatus", reflect.TypeOf((*MockQuerier)(nil).UpdateImportStatus), ctx, arg)
}

// UpsertRepo mocks base method.
func (m *MockQuerier) UpsertRepo(ctx context.Context, arg db.UpsertRepoParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRepo", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRepo indicates an expected call of UpsertRepo.
func (mr *MockQuerierMockRecorder) UpsertRepo(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRepo", reflect.TypeOf((*MockQuerier)(nil).UpsertRepo), ctx, arg)
}

// UpsertWorkflowRunJobs mocks base method.
func (m *MockQuerier) UpsertWorkflowRunJobs(ctx context.Context, arg db.UpsertWorkflowRunJobsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkflowRunJobs", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkflowRunJobs indicates an expected call of UpsertWorkflowRunJobs.
func (mr *MockQuerierMockRecorder) UpsertWorkflowRunJobs(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkflowRunJobs", reflect.TypeOf((*MockQuerier)(nil).UpsertWorkflowRunJobs), ctx, arg)
}

// UpsertWorkflowRuns mocks base method.
func (m *MockQuerier) UpsertWorkflowRuns(ctx context.Context, arg db.UpsertWorkflowRunsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkflowRuns", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkflowRuns indicates an expected call of UpsertWorkflowRuns.
func (mr *MockQuerierMockRecorder) UpsertWorkflowRuns(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkflowRuns", reflect.TypeOf((*MockQuerier)(nil).UpsertWorkflowRuns), ctx, arg)
}

// UpsertWorkflowsInPublic mocks base method.
func (m *MockQuerier) UpsertWorkflowsInPublic(ctx context.Context, arg db.UpsertWorkflowsInPublicParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkflowsInPublic", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkflowsInPublic indicates an expected call of UpsertWorkflowsInPublic.
func (mr *MockQuerierMockRecorder) UpsertWorkflowsInPublic(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkflowsInPublic", reflect.TypeOf((*MockQuerier)(nil).UpsertWorkflowsInPublic), ctx, arg)
}

// WithTx mocks base method.
func (m *MockQuerier) WithTx(tx pgx.Tx) queries.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", tx)
	ret0, _ := ret[0].(queries.Querier)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockQuerierMockRecorder) WithTx(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockQuerier)(nil).WithTx), tx)
}
