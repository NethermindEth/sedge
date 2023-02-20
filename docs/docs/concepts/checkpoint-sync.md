---
id: checkpoint-sync
---

# Checkpoint Sync endpoint

Checkpoint sync is a feature that significantly speeds up the initial sync between your beacon node and the Beacon Chain. 
With checkpoint sync configured, your beacon node will begin syncing from a recently finalized checkpoint instead of 
syncing from genesis. This can make installations, validator migrations, recoveries, and testnet deployments way faster.

A Checkpoint Sync endpoint is a consensus node endpoint that can be used by other consensus nodes for Checkpoint sync.