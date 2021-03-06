// Code generated by counterfeiter. DO NOT EDIT.
package signalingfakes

import (
	"sync"

	"github.com/ryanrolds/club/pkg/signaling"
)

type FakeDispatcher struct {
	DispatchStub        func(signaling.RoomMember, signaling.Message) error
	dispatchMutex       sync.RWMutex
	dispatchArgsForCall []struct {
		arg1 signaling.RoomMember
		arg2 signaling.Message
	}
	dispatchReturns struct {
		result1 error
	}
	dispatchReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDispatcher) Dispatch(arg1 signaling.RoomMember, arg2 signaling.Message) error {
	fake.dispatchMutex.Lock()
	ret, specificReturn := fake.dispatchReturnsOnCall[len(fake.dispatchArgsForCall)]
	fake.dispatchArgsForCall = append(fake.dispatchArgsForCall, struct {
		arg1 signaling.RoomMember
		arg2 signaling.Message
	}{arg1, arg2})
	fake.recordInvocation("Dispatch", []interface{}{arg1, arg2})
	fake.dispatchMutex.Unlock()
	if fake.DispatchStub != nil {
		return fake.DispatchStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dispatchReturns
	return fakeReturns.result1
}

func (fake *FakeDispatcher) DispatchCallCount() int {
	fake.dispatchMutex.RLock()
	defer fake.dispatchMutex.RUnlock()
	return len(fake.dispatchArgsForCall)
}

func (fake *FakeDispatcher) DispatchCalls(stub func(signaling.RoomMember, signaling.Message) error) {
	fake.dispatchMutex.Lock()
	defer fake.dispatchMutex.Unlock()
	fake.DispatchStub = stub
}

func (fake *FakeDispatcher) DispatchArgsForCall(i int) (signaling.RoomMember, signaling.Message) {
	fake.dispatchMutex.RLock()
	defer fake.dispatchMutex.RUnlock()
	argsForCall := fake.dispatchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDispatcher) DispatchReturns(result1 error) {
	fake.dispatchMutex.Lock()
	defer fake.dispatchMutex.Unlock()
	fake.DispatchStub = nil
	fake.dispatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDispatcher) DispatchReturnsOnCall(i int, result1 error) {
	fake.dispatchMutex.Lock()
	defer fake.dispatchMutex.Unlock()
	fake.DispatchStub = nil
	if fake.dispatchReturnsOnCall == nil {
		fake.dispatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.dispatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDispatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dispatchMutex.RLock()
	defer fake.dispatchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDispatcher) recordInvocation(key string, args []interface{}) {
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

var _ signaling.Dispatcher = new(FakeDispatcher)
