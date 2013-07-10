package spotify

import "testing"

func TestNewSession(t *testing.T) {
	session, err := NewSession("appkey")
	if session.session == nil {
		t.Error("Session.session should not be nil")
	}

	if err != nil {
		t.Error("err should be nil")
	}
}
