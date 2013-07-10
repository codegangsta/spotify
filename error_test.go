package spotify

import "testing"

func TestError(t *testing.T) {
	err := Error{SP_ERROR_OK}
	if err.Error() != "sp_error: 0" {
		t.Error("Error equals " + err.Error())
	}
}
