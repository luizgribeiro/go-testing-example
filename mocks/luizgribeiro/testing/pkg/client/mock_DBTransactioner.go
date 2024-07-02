// Code generated by mockery v2.43.2. DO NOT EDIT.

package client

import (
	context "context"

	pgconn "github.com/jackc/pgx/v5/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"
)

// MockDBTransactioner is an autogenerated mock type for the DBTransactioner type
type MockDBTransactioner struct {
	mock.Mock
}

type MockDBTransactioner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDBTransactioner) EXPECT() *MockDBTransactioner_Expecter {
	return &MockDBTransactioner_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: ctx
func (_m *MockDBTransactioner) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBTransactioner_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type MockDBTransactioner_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDBTransactioner_Expecter) Begin(ctx interface{}) *MockDBTransactioner_Begin_Call {
	return &MockDBTransactioner_Begin_Call{Call: _e.mock.On("Begin", ctx)}
}

func (_c *MockDBTransactioner_Begin_Call) Run(run func(ctx context.Context)) *MockDBTransactioner_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDBTransactioner_Begin_Call) Return(_a0 pgx.Tx, _a1 error) *MockDBTransactioner_Begin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBTransactioner_Begin_Call) RunAndReturn(run func(context.Context) (pgx.Tx, error)) *MockDBTransactioner_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields: ctx
func (_m *MockDBTransactioner) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDBTransactioner_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type MockDBTransactioner_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDBTransactioner_Expecter) Commit(ctx interface{}) *MockDBTransactioner_Commit_Call {
	return &MockDBTransactioner_Commit_Call{Call: _e.mock.On("Commit", ctx)}
}

func (_c *MockDBTransactioner_Commit_Call) Run(run func(ctx context.Context)) *MockDBTransactioner_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDBTransactioner_Commit_Call) Return(_a0 error) *MockDBTransactioner_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactioner_Commit_Call) RunAndReturn(run func(context.Context) error) *MockDBTransactioner_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// Conn provides a mock function with given fields:
func (_m *MockDBTransactioner) Conn() *pgx.Conn {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Conn")
	}

	var r0 *pgx.Conn
	if rf, ok := ret.Get(0).(func() *pgx.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pgx.Conn)
		}
	}

	return r0
}

// MockDBTransactioner_Conn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Conn'
type MockDBTransactioner_Conn_Call struct {
	*mock.Call
}

// Conn is a helper method to define mock.On call
func (_e *MockDBTransactioner_Expecter) Conn() *MockDBTransactioner_Conn_Call {
	return &MockDBTransactioner_Conn_Call{Call: _e.mock.On("Conn")}
}

func (_c *MockDBTransactioner_Conn_Call) Run(run func()) *MockDBTransactioner_Conn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBTransactioner_Conn_Call) Return(_a0 *pgx.Conn) *MockDBTransactioner_Conn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactioner_Conn_Call) RunAndReturn(run func() *pgx.Conn) *MockDBTransactioner_Conn_Call {
	_c.Call.Return(run)
	return _c
}

// CopyFrom provides a mock function with given fields: ctx, tableName, columnNames, rowSrc
func (_m *MockDBTransactioner) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	ret := _m.Called(ctx, tableName, columnNames, rowSrc)

	if len(ret) == 0 {
		panic("no return value specified for CopyFrom")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)); ok {
		return rf(ctx, tableName, columnNames, rowSrc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) int64); ok {
		r0 = rf(ctx, tableName, columnNames, rowSrc)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) error); ok {
		r1 = rf(ctx, tableName, columnNames, rowSrc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBTransactioner_CopyFrom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyFrom'
type MockDBTransactioner_CopyFrom_Call struct {
	*mock.Call
}

// CopyFrom is a helper method to define mock.On call
//   - ctx context.Context
//   - tableName pgx.Identifier
//   - columnNames []string
//   - rowSrc pgx.CopyFromSource
func (_e *MockDBTransactioner_Expecter) CopyFrom(ctx interface{}, tableName interface{}, columnNames interface{}, rowSrc interface{}) *MockDBTransactioner_CopyFrom_Call {
	return &MockDBTransactioner_CopyFrom_Call{Call: _e.mock.On("CopyFrom", ctx, tableName, columnNames, rowSrc)}
}

func (_c *MockDBTransactioner_CopyFrom_Call) Run(run func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource)) *MockDBTransactioner_CopyFrom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Identifier), args[2].([]string), args[3].(pgx.CopyFromSource))
	})
	return _c
}

func (_c *MockDBTransactioner_CopyFrom_Call) Return(_a0 int64, _a1 error) *MockDBTransactioner_CopyFrom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBTransactioner_CopyFrom_Call) RunAndReturn(run func(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)) *MockDBTransactioner_CopyFrom_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: ctx, sql, arguments
func (_m *MockDBTransactioner) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, arguments...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, sql, arguments...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, sql, arguments...)
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, arguments...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBTransactioner_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockDBTransactioner_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - sql string
//   - arguments ...interface{}
func (_e *MockDBTransactioner_Expecter) Exec(ctx interface{}, sql interface{}, arguments ...interface{}) *MockDBTransactioner_Exec_Call {
	return &MockDBTransactioner_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, sql}, arguments...)...)}
}

func (_c *MockDBTransactioner_Exec_Call) Run(run func(ctx context.Context, sql string, arguments ...interface{})) *MockDBTransactioner_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDBTransactioner_Exec_Call) Return(commandTag pgconn.CommandTag, err error) *MockDBTransactioner_Exec_Call {
	_c.Call.Return(commandTag, err)
	return _c
}

func (_c *MockDBTransactioner_Exec_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)) *MockDBTransactioner_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// LargeObjects provides a mock function with given fields:
func (_m *MockDBTransactioner) LargeObjects() pgx.LargeObjects {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LargeObjects")
	}

	var r0 pgx.LargeObjects
	if rf, ok := ret.Get(0).(func() pgx.LargeObjects); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pgx.LargeObjects)
	}

	return r0
}

// MockDBTransactioner_LargeObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LargeObjects'
type MockDBTransactioner_LargeObjects_Call struct {
	*mock.Call
}

// LargeObjects is a helper method to define mock.On call
func (_e *MockDBTransactioner_Expecter) LargeObjects() *MockDBTransactioner_LargeObjects_Call {
	return &MockDBTransactioner_LargeObjects_Call{Call: _e.mock.On("LargeObjects")}
}

func (_c *MockDBTransactioner_LargeObjects_Call) Run(run func()) *MockDBTransactioner_LargeObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBTransactioner_LargeObjects_Call) Return(_a0 pgx.LargeObjects) *MockDBTransactioner_LargeObjects_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactioner_LargeObjects_Call) RunAndReturn(run func() pgx.LargeObjects) *MockDBTransactioner_LargeObjects_Call {
	_c.Call.Return(run)
	return _c
}

// Prepare provides a mock function with given fields: ctx, name, sql
func (_m *MockDBTransactioner) Prepare(ctx context.Context, name string, sql string) (*pgconn.StatementDescription, error) {
	ret := _m.Called(ctx, name, sql)

	if len(ret) == 0 {
		panic("no return value specified for Prepare")
	}

	var r0 *pgconn.StatementDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*pgconn.StatementDescription, error)); ok {
		return rf(ctx, name, sql)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *pgconn.StatementDescription); ok {
		r0 = rf(ctx, name, sql)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pgconn.StatementDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, sql)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBTransactioner_Prepare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prepare'
type MockDBTransactioner_Prepare_Call struct {
	*mock.Call
}

// Prepare is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - sql string
func (_e *MockDBTransactioner_Expecter) Prepare(ctx interface{}, name interface{}, sql interface{}) *MockDBTransactioner_Prepare_Call {
	return &MockDBTransactioner_Prepare_Call{Call: _e.mock.On("Prepare", ctx, name, sql)}
}

func (_c *MockDBTransactioner_Prepare_Call) Run(run func(ctx context.Context, name string, sql string)) *MockDBTransactioner_Prepare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockDBTransactioner_Prepare_Call) Return(_a0 *pgconn.StatementDescription, _a1 error) *MockDBTransactioner_Prepare_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBTransactioner_Prepare_Call) RunAndReturn(run func(context.Context, string, string) (*pgconn.StatementDescription, error)) *MockDBTransactioner_Prepare_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, sql, args
func (_m *MockDBTransactioner) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, sql, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, sql, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBTransactioner_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockDBTransactioner_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - sql string
//   - args ...interface{}
func (_e *MockDBTransactioner_Expecter) Query(ctx interface{}, sql interface{}, args ...interface{}) *MockDBTransactioner_Query_Call {
	return &MockDBTransactioner_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, sql}, args...)...)}
}

func (_c *MockDBTransactioner_Query_Call) Run(run func(ctx context.Context, sql string, args ...interface{})) *MockDBTransactioner_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDBTransactioner_Query_Call) Return(_a0 pgx.Rows, _a1 error) *MockDBTransactioner_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBTransactioner_Query_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgx.Rows, error)) *MockDBTransactioner_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields: ctx, sql, args
func (_m *MockDBTransactioner) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, sql)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, sql, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// MockDBTransactioner_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type MockDBTransactioner_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
//   - ctx context.Context
//   - sql string
//   - args ...interface{}
func (_e *MockDBTransactioner_Expecter) QueryRow(ctx interface{}, sql interface{}, args ...interface{}) *MockDBTransactioner_QueryRow_Call {
	return &MockDBTransactioner_QueryRow_Call{Call: _e.mock.On("QueryRow",
		append([]interface{}{ctx, sql}, args...)...)}
}

func (_c *MockDBTransactioner_QueryRow_Call) Run(run func(ctx context.Context, sql string, args ...interface{})) *MockDBTransactioner_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDBTransactioner_QueryRow_Call) Return(_a0 pgx.Row) *MockDBTransactioner_QueryRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactioner_QueryRow_Call) RunAndReturn(run func(context.Context, string, ...interface{}) pgx.Row) *MockDBTransactioner_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields: ctx
func (_m *MockDBTransactioner) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDBTransactioner_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type MockDBTransactioner_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDBTransactioner_Expecter) Rollback(ctx interface{}) *MockDBTransactioner_Rollback_Call {
	return &MockDBTransactioner_Rollback_Call{Call: _e.mock.On("Rollback", ctx)}
}

func (_c *MockDBTransactioner_Rollback_Call) Run(run func(ctx context.Context)) *MockDBTransactioner_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDBTransactioner_Rollback_Call) Return(_a0 error) *MockDBTransactioner_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactioner_Rollback_Call) RunAndReturn(run func(context.Context) error) *MockDBTransactioner_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// SendBatch provides a mock function with given fields: ctx, b
func (_m *MockDBTransactioner) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	ret := _m.Called(ctx, b)

	if len(ret) == 0 {
		panic("no return value specified for SendBatch")
	}

	var r0 pgx.BatchResults
	if rf, ok := ret.Get(0).(func(context.Context, *pgx.Batch) pgx.BatchResults); ok {
		r0 = rf(ctx, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.BatchResults)
		}
	}

	return r0
}

// MockDBTransactioner_SendBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendBatch'
type MockDBTransactioner_SendBatch_Call struct {
	*mock.Call
}

// SendBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - b *pgx.Batch
func (_e *MockDBTransactioner_Expecter) SendBatch(ctx interface{}, b interface{}) *MockDBTransactioner_SendBatch_Call {
	return &MockDBTransactioner_SendBatch_Call{Call: _e.mock.On("SendBatch", ctx, b)}
}

func (_c *MockDBTransactioner_SendBatch_Call) Run(run func(ctx context.Context, b *pgx.Batch)) *MockDBTransactioner_SendBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pgx.Batch))
	})
	return _c
}

func (_c *MockDBTransactioner_SendBatch_Call) Return(_a0 pgx.BatchResults) *MockDBTransactioner_SendBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBTransactioner_SendBatch_Call) RunAndReturn(run func(context.Context, *pgx.Batch) pgx.BatchResults) *MockDBTransactioner_SendBatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDBTransactioner creates a new instance of MockDBTransactioner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDBTransactioner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDBTransactioner {
	mock := &MockDBTransactioner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
