package spotify

import "testing"

func TestNewSession(t *testing.T) {
	// good appkey
	session, err := NewSession("appkey_good")
	if session.session == nil {
		t.Error("Session.session should not be nil")
	}

	if err != nil {
		t.Error("Failed to initialize session:", err)
	}

	// bad appkey
	session, err = NewSession("appkey_bad")
	if err.Error() != "sp_error: 5" {
		t.Error("err should be SP_ERROR_BAD_APPLICATION_KEY, was", err, "instead")
	}
}
