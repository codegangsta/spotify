package spotify

/*
#include <libspotify/api.h>
*/
import "C"

type Error struct {
  code C.sp_error
}

func (e *Error) Error() string {
  return "foo"
}
