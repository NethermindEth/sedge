package templates

const (
	Nethermind = `
{{ define "execution" }}
  execution:
    stop_grace_period: 1m
    container_name: nethermind-execution-client
    restart: unless-stopped
    image: ${EXECUTION_IMAGE_VERSION}
    volumes:
      - ./nethermind-execution-data/:/nethermind/data
    ports:
      - 30303:30303/tcp
      - 30303:30303/udp
      - 8545:8545/tcp
    command: 
      - --config=${NETHERMIND_CONFIG}
      - --datadir=/nethermind/data
      - --log=${NETHERMIND_LOG_LEVEL}
      - --Metrics.Enabled=${NETHERMIND_METRICSCONFIG_ENABLED}
      - --Metrics.NodeName=${NETHERMIND_METRICSCONFIG_NODENAME}
      - --Metrics.PushGatewayUrl=${NETHERMIND_METRICSCONFIG_PUSHGATEWAYURL}
      - --Init.WebSocketsEnabled=true
      - --JsonRpc.Enabled=true
      - --JsonRpc.Host=0.0.0.0
      - --JsonRpc.EnabledModules=${NETHERMIND_JSONRPCCONFIG_ENABLEDMODULES  }
      - --HealthChecks.Enabled=${NETHERMIND_HEALTHCHECKSCONFIG_ENABLED}
      - --Pruning.CacheMb=${NETHERMIND_PRUNINGCONFIG_CACHEMB}
      - --EthStats.Enabled=${NETHERMIND_ETHSTATSCONFIG_ENABLED}
      - --EthStats.Server=${NETHERMIND_ETHSTATSCONFIG_SERVER}
      - --EthStats.Name=${NETHERMIND_ETHSTATSCONFIG_NAME}
      - --EthStats.Secret=${NETHERMIND_ETHSTATSCONFIG_SECRET}
      - --EthStats.Contact=${NETHERMIND_ETHSTATSCONFIG_CONTACT}
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"
{{ end }}
	`
)
