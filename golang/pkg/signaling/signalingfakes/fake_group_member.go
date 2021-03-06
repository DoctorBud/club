// Code generated by counterfeiter. DO NOT EDIT.
package signalingfakes

import (
	"sync"

	"github.com/ryanrolds/club/pkg/signaling"
)

type FakeGroupMember struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	IDStub        func() signaling.PeerID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 signaling.PeerID
	}
	iDReturnsOnCall map[int]struct {
		result1 signaling.PeerID
	}
	SendMessageStub        func(signaling.Message) error
	sendMessageMutex       sync.RWMutex
	sendMessageArgsForCall []struct {
		arg1 signaling.Message
	}
	sendMessageReturns struct {
		result1 error
	}
	sendMessageReturnsOnCall map[int]struct {
		result1 error
	}
	TimedoutStub        func() bool
	timedoutMutex       sync.RWMutex
	timedoutArgsForCall []struct {
	}
	timedoutReturns struct {
		result1 bool
	}
	timedoutReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGroupMember) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeGroupMember) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeGroupMember) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeGroupMember) ID() signaling.PeerID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeGroupMember) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeGroupMember) IDCalls(stub func() signaling.PeerID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeGroupMember) IDReturns(result1 signaling.PeerID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 signaling.PeerID
	}{result1}
}

func (fake *FakeGroupMember) IDReturnsOnCall(i int, result1 signaling.PeerID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 signaling.PeerID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 signaling.PeerID
	}{result1}
}

func (fake *FakeGroupMember) SendMessage(arg1 signaling.Message) error {
	fake.sendMessageMutex.Lock()
	ret, specificReturn := fake.sendMessageReturnsOnCall[len(fake.sendMessageArgsForCall)]
	fake.sendMessageArgsForCall = append(fake.sendMessageArgsForCall, struct {
		arg1 signaling.Message
	}{arg1})
	fake.recordInvocation("SendMessage", []interface{}{arg1})
	fake.sendMessageMutex.Unlock()
	if fake.SendMessageStub != nil {
		return fake.SendMessageStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendMessageReturns
	return fakeReturns.result1
}

func (fake *FakeGroupMember) SendMessageCallCount() int {
	fake.sendMessageMutex.RLock()
	defer fake.sendMessageMutex.RUnlock()
	return len(fake.sendMessageArgsForCall)
}

func (fake *FakeGroupMember) SendMessageCalls(stub func(signaling.Message) error) {
	fake.sendMessageMutex.Lock()
	defer fake.sendMessageMutex.Unlock()
	fake.SendMessageStub = stub
}

func (fake *FakeGroupMember) SendMessageArgsForCall(i int) signaling.Message {
	fake.sendMessageMutex.RLock()
	defer fake.sendMessageMutex.RUnlock()
	argsForCall := fake.sendMessageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGroupMember) SendMessageReturns(result1 error) {
	fake.sendMessageMutex.Lock()
	defer fake.sendMessageMutex.Unlock()
	fake.SendMessageStub = nil
	fake.sendMessageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGroupMember) SendMessageReturnsOnCall(i int, result1 error) {
	fake.sendMessageMutex.Lock()
	defer fake.sendMessageMutex.Unlock()
	fake.SendMessageStub = nil
	if fake.sendMessageReturnsOnCall == nil {
		fake.sendMessageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMessageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGroupMember) Timedout() bool {
	fake.timedoutMutex.Lock()
	ret, specificReturn := fake.timedoutReturnsOnCall[len(fake.timedoutArgsForCall)]
	fake.timedoutArgsForCall = append(fake.timedoutArgsForCall, struct {
	}{})
	fake.recordInvocation("Timedout", []interface{}{})
	fake.timedoutMutex.Unlock()
	if fake.TimedoutStub != nil {
		return fake.TimedoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.timedoutReturns
	return fakeReturns.result1
}

func (fake *FakeGroupMember) TimedoutCallCount() int {
	fake.timedoutMutex.RLock()
	defer fake.timedoutMutex.RUnlock()
	return len(fake.timedoutArgsForCall)
}

func (fake *FakeGroupMember) TimedoutCalls(stub func() bool) {
	fake.timedoutMutex.Lock()
	defer fake.timedoutMutex.Unlock()
	fake.TimedoutStub = stub
}

func (fake *FakeGroupMember) TimedoutReturns(result1 bool) {
	fake.timedoutMutex.Lock()
	defer fake.timedoutMutex.Unlock()
	fake.TimedoutStub = nil
	fake.timedoutReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGroupMember) TimedoutReturnsOnCall(i int, result1 bool) {
	fake.timedoutMutex.Lock()
	defer fake.timedoutMutex.Unlock()
	fake.TimedoutStub = nil
	if fake.timedoutReturnsOnCall == nil {
		fake.timedoutReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.timedoutReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGroupMember) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.sendMessageMutex.RLock()
	defer fake.sendMessageMutex.RUnlock()
	fake.timedoutMutex.RLock()
	defer fake.timedoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGroupMember) recordInvocation(key string, args []interface{}) {
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

var _ signaling.GroupMember = new(FakeGroupMember)
