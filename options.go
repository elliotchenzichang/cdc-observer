package cdcobserver

type Options struct {
	EnableDocker  bool
	ContainerName string
	DSN           string
	Port          int
	Username      string
	Password      string
}
