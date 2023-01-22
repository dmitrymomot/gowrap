// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/elasticapm
// gowrap: http://github.com/hexdigest/gowrap

package templatestests

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i TestInterface -t ../templates/elasticapm -o interface_with_elasticapm.go -l ""

import (
	"context"

	"go.elastic.co/apm/v2"
)

// TestInterfaceAPMTracing implements TestInterface interface with all methods wrapped
// with go.elastic.co/apm/v2
type TestInterfaceAPMTracing struct {
	base         TestInterface
	startSpan    func(ctx context.Context, name, spanType string) (*apm.Span, context.Context)
	endSpan      func(span *apm.Span)
	setLabel     func(span *apm.Span, key string, value interface{})
	captureError func(ctx context.Context, err error)
}

type TestInterfaceAPMTracingOption func(v *TestInterfaceAPMTracing)

func TestInterfaceAPMTracingWithUsingSetLabel() TestInterfaceAPMTracingOption {
	return func(v *TestInterfaceAPMTracing) {
		v.setLabel = func(span *apm.Span, key string, value interface{}) {
			span.SpanData.Context.SetLabel(key, value)
		}
	}
}

// NewTestInterfaceAPMTracing returns an instance of the TestInterface decorated with go.elastic.co/apm/v2
func NewTestInterfaceAPMTracing(base TestInterface, opts ...TestInterfaceAPMTracingOption) TestInterfaceAPMTracing {
	r := TestInterfaceAPMTracing{
		base:      base,
		startSpan: apm.StartSpan,
		endSpan: func(span *apm.Span) {
			span.End()
		},
		setLabel: func(span *apm.Span, key string, value interface{}) {
		},
		captureError: func(ctx context.Context, err error) {
			apm.CaptureError(ctx, err).Send()
		},
	}

	for _, fn := range opts {
		fn(&r)
	}
	return r
}

// ContextNoError implements TestInterface
func (_d TestInterfaceAPMTracing) ContextNoError(ctx context.Context, a1 string, a2 string) {
	span, ctx := _d.startSpan(ctx, "testinterface.ContextNoError", "testinterface")
	defer func() {
		_d.endSpan(span)
	}()
	_d.setLabel(span, "a1", a1)
	_d.setLabel(span, "a2", a2)

	_d.base.ContextNoError(ctx, a1, a2)
	return
}

// F implements TestInterface
func (_d TestInterfaceAPMTracing) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	span, ctx := _d.startSpan(ctx, "testinterface.F", "testinterface")
	defer func() {
		if err != nil {
			_d.captureError(ctx, err)
		}
		_d.endSpan(span)
	}()
	_d.setLabel(span, "a1", a1)
	_d.setLabel(span, "a2", a2)

	return _d.base.F(ctx, a1, a2...)
}
