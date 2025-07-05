package main


type Config struct {
	ApiPort						string		`yaml:"REST_API_port"`
	OrchestratorTarget 			string		`yaml:"orchestrator_grpc_target"`

}
