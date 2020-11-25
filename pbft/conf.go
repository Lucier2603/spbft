package pbft

type Config struct {
	ElectionTick	int						`yaml:"ElectionTick"`

	ID				uint64					`yaml:"ID"`
	IP				string					`yaml:"IP"`
	PBFT_PORT		string					`yaml:"PBFT_PORT"`
	HTTP_PORT		string					`yaml:"HTTP_PORT"`

	Peers			[]uint64				`yaml:"Peers"`
	Routers			map[uint64]string		`yaml:"Routers"`
	HeartbeatTick	uint					`yaml:"HeartbeatTick"`

	StorageDir		string					`yaml:"StorageDir"`
}
