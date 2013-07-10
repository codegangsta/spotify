package spotify

/*
#include <stdlib.h>
#include <libspotify/api.h>
*/
import "C"
import "unsafe"

const (
	SP_ERROR_OK                        = 0
	SP_ERROR_BAD_API_VERSION           = 1
	SP_ERROR_API_INITIALIZATION_FAILED = 2
	SP_ERROR_TRACK_NOT_PLAYABLE        = 3
	SP_ERROR_BAD_APPLICATION_KEY       = 5
	SP_ERROR_BAD_USERNAME_OR_PASSWORD  = 6
	SP_ERROR_USER_BANNED               = 7
	SP_ERROR_UNABLE_TO_CONTACT_SERVER  = 8
	SP_ERROR_CLIENT_TOO_OLD            = 9
	SP_ERROR_OTHER_PERMANENT           = 10
	SP_ERROR_BAD_USER_AGENT            = 11
	SP_ERROR_MISSING_CALLBACK          = 12
	SP_ERROR_INVALID_INDATA            = 13
	SP_ERROR_INDEX_OUT_OF_RANGE        = 14
	SP_ERROR_USER_NEEDS_PREMIUM        = 15
	SP_ERROR_OTHER_TRANSIENT           = 16
	SP_ERROR_IS_LOADING                = 17
	SP_ERROR_NO_STREAM_AVAILABLE       = 18
	SP_ERROR_PERMISSION_DENIED         = 19
	SP_ERROR_INBOX_IS_FULL             = 20
	SP_ERROR_NO_CACHE                  = 21
	SP_ERROR_NO_SUCH_USER              = 22
	SP_ERROR_NO_CREDENTIALS            = 23
	SP_ERROR_NETWORK_DISABLED          = 24
	SP_ERROR_INVALID_DEVICE_ID         = 25
	SP_ERROR_CANT_OPEN_TRACE_FILE      = 26
	SP_ERROR_APPLICATION_BANNED        = 27
	SP_ERROR_OFFLINE_TOO_MANY_TRACKS   = 31
	SP_ERROR_OFFLINE_DISK_CACHE        = 32
	SP_ERROR_OFFLINE_EXPIRED           = 33
	SP_ERROR_OFFLINE_NOT_ALLOWED       = 34
	SP_ERROR_OFFLINE_LICENSE_LOST      = 35
	SP_ERROR_OFFLINE_LICENSE_ERROR     = 36
	SP_ERROR_LASTFM_AUTH_ERROR         = 39
	SP_ERROR_INVALID_ARGUMENT          = 40
	SP_ERROR_SYSTEM_FAILURE            = 41
)

type Error struct {
	code C.sp_error
}

func (e *Error) Error() string {
  cmessage := C.sp_error_message(e.code)
  message := C.GoString(cmessage)
  C.free(unsafe.Pointer(cmessage))
  return message
}
