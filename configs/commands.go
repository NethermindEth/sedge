package configs

// Commands
const (
	DockerComposeUpCMD         = "docker-compose -f %s up -d execution consensus"
	DockerPsCMD                = "docker ps -a"
	DockerComposePsServicesCMD = "docker-compose -f %s ps --services --filter \"status=running\""
	DockerComposeLogsFollowCMD = "docker-compose -f %s logs --follow %s"
	DockerComposeLogsTailCMD   = "docker-compose -f %s logs --tail=20 %s"
	DepositCLIDockerBuildCMD   = "docker build github.com/ethereum/eth2.0-deposit-cli -t %s"
	DockerInspectCMD           = "docker inspect %s"
)
