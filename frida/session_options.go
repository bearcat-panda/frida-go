package frida

//#include <frida-core.h>
import "C"

// SessionOptions type is used to configure session
type SessionOptions struct {
	opts *C.FridaSessionOptions
}

// NewSessionOptions create new SessionOptions with the realm and
// timeout to persist provided
func NewSessionOptions(realm Realm, persist_timeout uint) *SessionOptions {
	opts := C.frida_session_options_new()
	C.frida_session_options_set_realm(opts, C.FridaRealm(realm))
	C.frida_session_options_set_persist_timeout(opts, C.guint(persist_timeout))

	return &SessionOptions{opts}
}

// GetRealm returns the realm of the options
func (s *SessionOptions) GetRealm() Realm {
	rlm := C.frida_session_options_get_realm(s.opts)
	return Realm(rlm)
}

// GetPersistTimeout returns the persist timeout of the script.s
func (s *SessionOptions) GetPersistTimeout() int {
	return int(C.frida_session_options_get_persist_timeout(s.opts))
}
