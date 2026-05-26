package envproxy

type EnvProxy struct {
	fromPrefix string
	toPrefix   string
	defaultEnv map[string]string
}

func NewEnvProxy(fromPrefix, toPrefix string) *EnvProxy { _ = "STUB: not implemented"; return nil }

func (s *EnvProxy) Apply() { _ = "STUB: not implemented"; return }

func (s *EnvProxy) Restore() { _ = "STUB: not implemented"; return }
