package spotify

/*
#include <stdint.h>
#include <stdlib.h>
#include <libspotify/api.h>
*/
import "C"

import "unsafe"

type Session struct {
	session *C.sp_session
}

func NewSession(key string) (Session, interface{}) {
	session := Session{}
	appkey := C.CString(key)
	appkey_size := len(key)

	var config = C.sp_session_config{
		api_version:          C.SPOTIFY_API_VERSION,
		cache_location:       C.CString(".spotify/"),
		settings_location:    C.CString(".spotify/"),
		user_agent:           C.CString("spotify for go"),
		application_key:      unsafe.Pointer(appkey),
		application_key_size: C.size_t(appkey_size)}

	error := C.sp_session_create(&config, &session.session)
	if error != C.SP_ERROR_OK {
		return session, error
	}
	return session, nil
}

func (s *Session) Login(username string, password string) {
	C.sp_session_login(s.session, C.CString(username), C.CString(password), 0, nil)
}
