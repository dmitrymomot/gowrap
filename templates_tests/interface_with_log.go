package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../templates/log template

//go:generate gowrap gen -d . -i TestInterface -t ../templates/log -o interface_with_log.go

import (
	"context"
	"io"
	"log"
)

// TestInterfaceWithLog implements TestInterface that is instrumented with logging
type TestInterfaceWithLog struct {
	_stdlog, _errlog *log.Logger
	_base            TestInterface
}

// NewTestInterfaceWithLog instruments an implementation of the TestInterface with simple logging
func NewTestInterfaceWithLog(base TestInterface, stdout, stderr io.Writer) TestInterfaceWithLog {
	return TestInterfaceWithLog{
		_base:   base,
		_stdlog: log.New(stdout, "", log.LstdFlags),
		_errlog: log.New(stderr, "", log.LstdFlags),
	}
}

// F implements TestInterface
func (_d TestInterfaceWithLog) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	_params := []interface{}{"TestInterfaceWithLog: calling F with params:", ctx, a1, a2}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"TestInterfaceWithLog: F returned results:", result1, result2, err}
		if err != nil {
			_d._errlog.Println(_results...)
		} else {
			_d._stdlog.Println(_results...)
		}
	}()
	return _d._base.F(ctx, a1, a2...)
}

// NoError implements TestInterface
func (_d TestInterfaceWithLog) NoError(s1 string) (s2 string) {
	_params := []interface{}{"TestInterfaceWithLog: calling NoError with params:", s1}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"TestInterfaceWithLog: NoError returned results:", s2}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.NoError(s1)
}

// NoParamsOrResults implements TestInterface
func (_d TestInterfaceWithLog) NoParamsOrResults() {
	_d._stdlog.Println("TestInterfaceWithLog: calling NoParamsOrResults")
	defer func() {
		_d._stdlog.Println("TestInterfaceWithLog: NoParamsOrResults finished")
	}()
	_d._base.NoParamsOrResults()
	return
}
