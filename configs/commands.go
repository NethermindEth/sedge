package configs

// Commands
const (
	DockerComposeUpCMD         = "sudo docker-compose -f %s up -d execution consensus"
	DockerPsCMD                = "sudo docker ps -a"
	DockerComposePsServicesCMD = "sudo docker-compose -f %s ps --services --filter \"status=running\""
	DockerComposeLogsFollowCMD = "sudo docker-compose -f %s logs --follow %s"
	DockerComposeLogsTailCMD   = "sudo docker-compose -f %s logs --tail=20 %s"
	DepositCLIDockerBuildCMD   = "sudo docker build github.com/ethereum/eth2.0-deposit-cli -t %s"
)
