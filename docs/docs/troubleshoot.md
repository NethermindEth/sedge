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

### Permissions issues

If you found a permission issue while installing some dependencies, and you got a error like this:

```
docker.service - Docker Application Container Engine
     Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
     Active: activating (auto-restart) (Result: exit-code) since Wed 2023-03-22 00:52:34 UTC; 6ms ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
    Process: 1720819 ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock (code=exited, status=1/FAILURE)
   Main PID: 1720819 (code=exited, status=1/FAILURE)

XXX 00 00:00:00 localhost systemd[1]: docker.service: Failed with result 'exit-code'.
XXX 00 00:00:00 localhost systemd[1]: Failed to start Docker Application Container Engine.
dpkg: error processing package docker-ce (--configure):
 installed docker-ce package post-installation script subprocess returned error exit status 1
```

or 

```
permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock
```

You can run the command using sudo, something like:
```shell
sudo sedge deps install
```