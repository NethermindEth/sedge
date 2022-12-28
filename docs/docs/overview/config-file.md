---
slug: config-file
---

# Configuration file

When you run Sedge for the first time, it generates a `.sedge.yml` on your HOME directory. This file should look like this:

```yaml
dependencies:
  - docker

executionClients:
  - geth
  - nethermind

consensusClients:
  - lighthouse
  - lodestar
  - prysm
  - teku
 
validatorClients:
  - lighthouse
  - lodestar
  - prysm
  - teku

logs:
  logLevel: info
```

If you want to know what Sedge does in every step, then just read the logs. Sedge is logging every step. Some of the applied steps or commands are shown on the debug logs only. To see these logs you need to replace `info` for `debug` on the `logLevel` field.

You can modify the clients there to customize Sedge's random selection of clients. If you remove a client there, it won't be randomly selected. Be careful not to add an unsupported client, if this client is choosed you will face an error. Sedge knows very well which clients it supports.

You don't need to modify the `dependencies` field at all. If you alter it, Sedge may not work as expected.