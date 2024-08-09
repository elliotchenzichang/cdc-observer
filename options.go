package cdcobserver

type Options struct {
	EnableDocker  bool
	ContainerName string
	DSN           string
	ContainerPort int
	Username      string
	Password      string
	DatabaseName  string
}

func (opt *Options) validates() error {
	return nil
}
