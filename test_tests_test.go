package tests

import (
	"testing"
)

type myT struct {
    testing.T
    FatalfCalled,
    ErrorfCalled bool
 }

func (t *myT) Fatalf(format string, args ...interface{}) {
    t.FatalfCalled = true
    t.T.Fatalf(format, args...)
}

func (t *myT) Errorf(format string, args ...interface{}) {
    t.ErrorfCalled = true
    t.T.Errorf(format, args...)
}

func TestTestSenderMockReplyChannel(t *testing.T) {
    mt := &myT{}
    TestSenderMockReplyChannel(mt)
    if mt.FatalfCalled || mt.ErrorfCalled {
        t.FailNow()
    }
}

func TestSenderMockReplyNick(t *testing.T) {
    mt := &myT{}
    TestSenderMockReplyNick(mt)
    if mt.FatalfCalled || mt.ErrorfCalled {
        t.FailNow()
    }
}

func TestSenderMockMessage(t *testing.T) {
    mt := &myT{}
    TestSenderMockMessage(mt)
    if mt.FatalfCalled || mt.ErrorfCalled {
        t.FailNow()
    }
}
