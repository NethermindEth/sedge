package templates

const (
	LighthouseConsensus = `
{{ define "consensus" }}
  consensus:
    stop_grace_period: 1m
    container_name: lighthouse-consensus-client
    restart: unless-stopped
    image: ${CC_IMAGE_VERSION}
    depends_on: 
      - execution
    volumes:
      - ./lhconsensus-data:/var/lib/lighthouse
    ports:
      - 9000:9000/tcp
      - 9000:9000/udp
      - 4000:4000/tcp
    command:
      - lighthouse
      - bn
      - --disable-upnp
      - --datadir=/var/lib/lighthouse
      - --port=9000
      - --http
      - --http-address=0.0.0.0
      - --http-port=4000
      - --network=${NETWORK}
      - --target-peers=${PEER_COUNT}
      - --eth1-endpoints=${EC_NODE}
      - --eth1-blocks-per-log-query=150
      - --debug-level=${CC_LOG_LEVEL}
      - --validator-monitor-auto
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"      
{{ end }}
`
	LodestarConsensus = `
{{ define "consensus" }}
  consensus:
    stop_grace_period: 1m
    container_name: lodestar-consensus-client
    restart: unless-stopped
    image: ${CC_IMAGE_VERSION}
    depends_on: 
      - execution
    volumes:
      - ./lsconsensus-data:/var/lib/lodestar/consensus
    ports:
      - 9000:9000/tcp
      - 9000:9000/udp
      - 4000:4000/tcp
    environment:
      NODE_OPTIONS: --max_old_space_size=6144
    command: 
      - beacon 
      - --rootDir=/var/lib/lodestar/consensus
      - --network=${NETWORK}
      - --logFile=/var/lib/lodestart/consensus/logs/beacon.log 
      - --logLevelFile=${CC_LOG_LEVEL}
      - --api.rest.enabled 
      - --api.rest.host=0.0.0.0
      - --api.rest.port=4000
      - --network.discv5.bindAddr="/ip4/0.0.0.0/udp/9000"
      - --network.localMultiaddrs="/ip4/0.0.0.0/tcp/9000"
      - --eth1.providerUrl=${EC_NODE}
      - --weakSubjectivitySyncLatest=${LS_RAPID_SYNC}
      - --network.targetPeers=${PEERS_COUNT}
    logging:
      driver: "json-file"
      options:
      max-size: "10m"
      max-file: "10"
{{ end }}
`
	PrysmConsensus = `
{{ define "consensus" }}
  consensus:
    stop_grace_period: 1m
    container_name: prysm-consesus-client
    restart: unless-stopped
    image: ${CC_IMAGE_VERSION}
    depends_on: 
      - execution
    volumes:
      - ./prysmconsensus-data:/var/lib/prysm
    ports:
      - 13000:13000/tcp
      - 12000:12000/udp
      - 3500:3500/tcp
      - 4000:4000/tcp
    command:
      - --datadir=/var/lib/prysm/
      - --verbosity=${CC_LOG_LEVEL}
      - --${NETWORK}
      - --p2p-tcp-port=13000
      - --p2p-udp-port=12000
      - --p2p-max-peers=${PEER_COUNT}
      - --rpc-host=0.0.0.0
      - --grpc-gateway-host=0.0.0.0
      - --grpc-gateway-port=4000
      - --http-web3provider=${EC_NODE}
      - --fallback-web3provider=${EC_FALLBACK_NODE1}
      - --fallback-web3provider=${EC_FALLBACK_NODE2}
      - --eth1-header-req-limit=150
      - --accept-terms-of-use
    logging:
      driver: "json-file"
      options:
  	    max-size: "10m"
  	    max-file: "10"
{{ end }}    
`
	TekuConsensus = `
{{ define "consensus" }}
  consensus:
    stop_grace_period: 1m
    container_name: teku-consensus-client
    restart: unless-stopped
    image: ${CC_IMAGE_VERSION}
    depends_on: 
      - execution
    user: root
    volumes:
      - ./tekuconsensus-data:/var/lib/teku
    ports:
      - 9000:9000/tcp
      - 9000:9000/udp
      - 4000:4000/tcp
    environment:
      - JAVA_OPTS=-XX:SoftMaxHeapSize=2g -Xmx4g
    command:
      - --data-path=/var/lib/teku
      - --log-destination=CONSOLE
      - --logging=${CC_LOG_LEVEL}
      - --network=${NETWORK}
      - --p2p-port=9000
      - --p2p-peer-upper-bound=${PEER_COUNT}
      - --rest-api-host-allowlist=*
      - --rest-api-enabled=true
      - --rest-api-port=4000
      - --eth1-endpoints=${EC_NODE}
      - --eth1-deposit-contract-max-request-size=150
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"
{{ end }} 
 `
)
