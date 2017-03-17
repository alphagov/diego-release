// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/loggregator_v2"
	"github.com/cloudfoundry/sonde-go/events"
)

type FakeClient struct {
	SendAppLogStub        func(appID, message, sourceType, sourceInstance string) error
	sendAppLogMutex       sync.RWMutex
	sendAppLogArgsForCall []struct {
		appID          string
		message        string
		sourceType     string
		sourceInstance string
	}
	sendAppLogReturns struct {
		result1 error
	}
	SendAppErrorLogStub        func(appID, message, sourceType, sourceInstance string) error
	sendAppErrorLogMutex       sync.RWMutex
	sendAppErrorLogArgsForCall []struct {
		appID          string
		message        string
		sourceType     string
		sourceInstance string
	}
	sendAppErrorLogReturns struct {
		result1 error
	}
	SendAppMetricsStub        func(metrics *events.ContainerMetric) error
	sendAppMetricsMutex       sync.RWMutex
	sendAppMetricsArgsForCall []struct {
		metrics *events.ContainerMetric
	}
	sendAppMetricsReturns struct {
		result1 error
	}
	SendDurationStub        func(name string, value time.Duration) error
	sendDurationMutex       sync.RWMutex
	sendDurationArgsForCall []struct {
		name  string
		value time.Duration
	}
	sendDurationReturns struct {
		result1 error
	}
	SendMebiBytesStub        func(name string, value int) error
	sendMebiBytesMutex       sync.RWMutex
	sendMebiBytesArgsForCall []struct {
		name  string
		value int
	}
	sendMebiBytesReturns struct {
		result1 error
	}
	SendMetricStub        func(name string, value int) error
	sendMetricMutex       sync.RWMutex
	sendMetricArgsForCall []struct {
		name  string
		value int
	}
	sendMetricReturns struct {
		result1 error
	}
	SendBytesPerSecondStub        func(name string, value float64) error
	sendBytesPerSecondMutex       sync.RWMutex
	sendBytesPerSecondArgsForCall []struct {
		name  string
		value float64
	}
	sendBytesPerSecondReturns struct {
		result1 error
	}
	SendRequestsPerSecondStub        func(name string, value float64) error
	sendRequestsPerSecondMutex       sync.RWMutex
	sendRequestsPerSecondArgsForCall []struct {
		name  string
		value float64
	}
	sendRequestsPerSecondReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) SendAppLog(appID string, message string, sourceType string, sourceInstance string) error {
	fake.sendAppLogMutex.Lock()
	fake.sendAppLogArgsForCall = append(fake.sendAppLogArgsForCall, struct {
		appID          string
		message        string
		sourceType     string
		sourceInstance string
	}{appID, message, sourceType, sourceInstance})
	fake.recordInvocation("SendAppLog", []interface{}{appID, message, sourceType, sourceInstance})
	fake.sendAppLogMutex.Unlock()
	if fake.SendAppLogStub != nil {
		return fake.SendAppLogStub(appID, message, sourceType, sourceInstance)
	} else {
		return fake.sendAppLogReturns.result1
	}
}

func (fake *FakeClient) SendAppLogCallCount() int {
	fake.sendAppLogMutex.RLock()
	defer fake.sendAppLogMutex.RUnlock()
	return len(fake.sendAppLogArgsForCall)
}

func (fake *FakeClient) SendAppLogArgsForCall(i int) (string, string, string, string) {
	fake.sendAppLogMutex.RLock()
	defer fake.sendAppLogMutex.RUnlock()
	return fake.sendAppLogArgsForCall[i].appID, fake.sendAppLogArgsForCall[i].message, fake.sendAppLogArgsForCall[i].sourceType, fake.sendAppLogArgsForCall[i].sourceInstance
}

func (fake *FakeClient) SendAppLogReturns(result1 error) {
	fake.SendAppLogStub = nil
	fake.sendAppLogReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendAppErrorLog(appID string, message string, sourceType string, sourceInstance string) error {
	fake.sendAppErrorLogMutex.Lock()
	fake.sendAppErrorLogArgsForCall = append(fake.sendAppErrorLogArgsForCall, struct {
		appID          string
		message        string
		sourceType     string
		sourceInstance string
	}{appID, message, sourceType, sourceInstance})
	fake.recordInvocation("SendAppErrorLog", []interface{}{appID, message, sourceType, sourceInstance})
	fake.sendAppErrorLogMutex.Unlock()
	if fake.SendAppErrorLogStub != nil {
		return fake.SendAppErrorLogStub(appID, message, sourceType, sourceInstance)
	} else {
		return fake.sendAppErrorLogReturns.result1
	}
}

func (fake *FakeClient) SendAppErrorLogCallCount() int {
	fake.sendAppErrorLogMutex.RLock()
	defer fake.sendAppErrorLogMutex.RUnlock()
	return len(fake.sendAppErrorLogArgsForCall)
}

func (fake *FakeClient) SendAppErrorLogArgsForCall(i int) (string, string, string, string) {
	fake.sendAppErrorLogMutex.RLock()
	defer fake.sendAppErrorLogMutex.RUnlock()
	return fake.sendAppErrorLogArgsForCall[i].appID, fake.sendAppErrorLogArgsForCall[i].message, fake.sendAppErrorLogArgsForCall[i].sourceType, fake.sendAppErrorLogArgsForCall[i].sourceInstance
}

func (fake *FakeClient) SendAppErrorLogReturns(result1 error) {
	fake.SendAppErrorLogStub = nil
	fake.sendAppErrorLogReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendAppMetrics(metrics *events.ContainerMetric) error {
	fake.sendAppMetricsMutex.Lock()
	fake.sendAppMetricsArgsForCall = append(fake.sendAppMetricsArgsForCall, struct {
		metrics *events.ContainerMetric
	}{metrics})
	fake.recordInvocation("SendAppMetrics", []interface{}{metrics})
	fake.sendAppMetricsMutex.Unlock()
	if fake.SendAppMetricsStub != nil {
		return fake.SendAppMetricsStub(metrics)
	} else {
		return fake.sendAppMetricsReturns.result1
	}
}

func (fake *FakeClient) SendAppMetricsCallCount() int {
	fake.sendAppMetricsMutex.RLock()
	defer fake.sendAppMetricsMutex.RUnlock()
	return len(fake.sendAppMetricsArgsForCall)
}

func (fake *FakeClient) SendAppMetricsArgsForCall(i int) *events.ContainerMetric {
	fake.sendAppMetricsMutex.RLock()
	defer fake.sendAppMetricsMutex.RUnlock()
	return fake.sendAppMetricsArgsForCall[i].metrics
}

func (fake *FakeClient) SendAppMetricsReturns(result1 error) {
	fake.SendAppMetricsStub = nil
	fake.sendAppMetricsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendDuration(name string, value time.Duration) error {
	fake.sendDurationMutex.Lock()
	fake.sendDurationArgsForCall = append(fake.sendDurationArgsForCall, struct {
		name  string
		value time.Duration
	}{name, value})
	fake.recordInvocation("SendDuration", []interface{}{name, value})
	fake.sendDurationMutex.Unlock()
	if fake.SendDurationStub != nil {
		return fake.SendDurationStub(name, value)
	} else {
		return fake.sendDurationReturns.result1
	}
}

func (fake *FakeClient) SendDurationCallCount() int {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return len(fake.sendDurationArgsForCall)
}

func (fake *FakeClient) SendDurationArgsForCall(i int) (string, time.Duration) {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return fake.sendDurationArgsForCall[i].name, fake.sendDurationArgsForCall[i].value
}

func (fake *FakeClient) SendDurationReturns(result1 error) {
	fake.SendDurationStub = nil
	fake.sendDurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendMebiBytes(name string, value int) error {
	fake.sendMebiBytesMutex.Lock()
	fake.sendMebiBytesArgsForCall = append(fake.sendMebiBytesArgsForCall, struct {
		name  string
		value int
	}{name, value})
	fake.recordInvocation("SendMebiBytes", []interface{}{name, value})
	fake.sendMebiBytesMutex.Unlock()
	if fake.SendMebiBytesStub != nil {
		return fake.SendMebiBytesStub(name, value)
	} else {
		return fake.sendMebiBytesReturns.result1
	}
}

func (fake *FakeClient) SendMebiBytesCallCount() int {
	fake.sendMebiBytesMutex.RLock()
	defer fake.sendMebiBytesMutex.RUnlock()
	return len(fake.sendMebiBytesArgsForCall)
}

func (fake *FakeClient) SendMebiBytesArgsForCall(i int) (string, int) {
	fake.sendMebiBytesMutex.RLock()
	defer fake.sendMebiBytesMutex.RUnlock()
	return fake.sendMebiBytesArgsForCall[i].name, fake.sendMebiBytesArgsForCall[i].value
}

func (fake *FakeClient) SendMebiBytesReturns(result1 error) {
	fake.SendMebiBytesStub = nil
	fake.sendMebiBytesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendMetric(name string, value int) error {
	fake.sendMetricMutex.Lock()
	fake.sendMetricArgsForCall = append(fake.sendMetricArgsForCall, struct {
		name  string
		value int
	}{name, value})
	fake.recordInvocation("SendMetric", []interface{}{name, value})
	fake.sendMetricMutex.Unlock()
	if fake.SendMetricStub != nil {
		return fake.SendMetricStub(name, value)
	} else {
		return fake.sendMetricReturns.result1
	}
}

func (fake *FakeClient) SendMetricCallCount() int {
	fake.sendMetricMutex.RLock()
	defer fake.sendMetricMutex.RUnlock()
	return len(fake.sendMetricArgsForCall)
}

func (fake *FakeClient) SendMetricArgsForCall(i int) (string, int) {
	fake.sendMetricMutex.RLock()
	defer fake.sendMetricMutex.RUnlock()
	return fake.sendMetricArgsForCall[i].name, fake.sendMetricArgsForCall[i].value
}

func (fake *FakeClient) SendMetricReturns(result1 error) {
	fake.SendMetricStub = nil
	fake.sendMetricReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendBytesPerSecond(name string, value float64) error {
	fake.sendBytesPerSecondMutex.Lock()
	fake.sendBytesPerSecondArgsForCall = append(fake.sendBytesPerSecondArgsForCall, struct {
		name  string
		value float64
	}{name, value})
	fake.recordInvocation("SendBytesPerSecond", []interface{}{name, value})
	fake.sendBytesPerSecondMutex.Unlock()
	if fake.SendBytesPerSecondStub != nil {
		return fake.SendBytesPerSecondStub(name, value)
	} else {
		return fake.sendBytesPerSecondReturns.result1
	}
}

func (fake *FakeClient) SendBytesPerSecondCallCount() int {
	fake.sendBytesPerSecondMutex.RLock()
	defer fake.sendBytesPerSecondMutex.RUnlock()
	return len(fake.sendBytesPerSecondArgsForCall)
}

func (fake *FakeClient) SendBytesPerSecondArgsForCall(i int) (string, float64) {
	fake.sendBytesPerSecondMutex.RLock()
	defer fake.sendBytesPerSecondMutex.RUnlock()
	return fake.sendBytesPerSecondArgsForCall[i].name, fake.sendBytesPerSecondArgsForCall[i].value
}

func (fake *FakeClient) SendBytesPerSecondReturns(result1 error) {
	fake.SendBytesPerSecondStub = nil
	fake.sendBytesPerSecondReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SendRequestsPerSecond(name string, value float64) error {
	fake.sendRequestsPerSecondMutex.Lock()
	fake.sendRequestsPerSecondArgsForCall = append(fake.sendRequestsPerSecondArgsForCall, struct {
		name  string
		value float64
	}{name, value})
	fake.recordInvocation("SendRequestsPerSecond", []interface{}{name, value})
	fake.sendRequestsPerSecondMutex.Unlock()
	if fake.SendRequestsPerSecondStub != nil {
		return fake.SendRequestsPerSecondStub(name, value)
	} else {
		return fake.sendRequestsPerSecondReturns.result1
	}
}

func (fake *FakeClient) SendRequestsPerSecondCallCount() int {
	fake.sendRequestsPerSecondMutex.RLock()
	defer fake.sendRequestsPerSecondMutex.RUnlock()
	return len(fake.sendRequestsPerSecondArgsForCall)
}

func (fake *FakeClient) SendRequestsPerSecondArgsForCall(i int) (string, float64) {
	fake.sendRequestsPerSecondMutex.RLock()
	defer fake.sendRequestsPerSecondMutex.RUnlock()
	return fake.sendRequestsPerSecondArgsForCall[i].name, fake.sendRequestsPerSecondArgsForCall[i].value
}

func (fake *FakeClient) SendRequestsPerSecondReturns(result1 error) {
	fake.SendRequestsPerSecondStub = nil
	fake.sendRequestsPerSecondReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendAppLogMutex.RLock()
	defer fake.sendAppLogMutex.RUnlock()
	fake.sendAppErrorLogMutex.RLock()
	defer fake.sendAppErrorLogMutex.RUnlock()
	fake.sendAppMetricsMutex.RLock()
	defer fake.sendAppMetricsMutex.RUnlock()
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	fake.sendMebiBytesMutex.RLock()
	defer fake.sendMebiBytesMutex.RUnlock()
	fake.sendMetricMutex.RLock()
	defer fake.sendMetricMutex.RUnlock()
	fake.sendBytesPerSecondMutex.RLock()
	defer fake.sendBytesPerSecondMutex.RUnlock()
	fake.sendRequestsPerSecondMutex.RLock()
	defer fake.sendRequestsPerSecondMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ loggregator_v2.Client = new(FakeClient)
