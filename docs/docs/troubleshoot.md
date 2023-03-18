---
slug: troubleshooting
---

# Troubleshooting

### Docker Compose not found after Sedge installs it

If you see a log similar to this just before Sedge is going to interact with the docker-compose script:

```
2022-XX-XX 00:00:00 -- [ERRO] it seems docker compose plugin is not installed. Please install it and try again.
```

Then probably is because the `compose` plugin is not installed for the `root` user. Currently Sedge uses `sudo` (root /
admin powers) to run `docker compose` commands. We are working on this issue. To solve it, you can copy the `compose`
plugin from your Home directory to `/root/` (root's Home directory) with the following command:

```
sudo cp ~/.docker/cli-plugins/docker-compose /root/.docker/cli-plugins/
```

### Prysm doesn't find a wallet datadir and throws an error

If Prysm throws an error that indicates that can't find a wallet datadir, that means that the wallet was not created
correctly.

To solve this, you should run `sedge import-key` with the flag `--start-validator` to starts the validator client after
import, regardless of the state the validator was in before.

This will create the wallet datadir and the validator will start correctly.
