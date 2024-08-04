package cdcobserver

func NewCDCObserver(*Options) {
	StartMySQLContainer()
	select {}
}
