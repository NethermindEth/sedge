---
slug: troubleshooting
---

# Troubleshooting

### Docker Compose not found after Sedge installs it

If you see a log similar to this just before Sedge is going to interact with the docker-compose script:
```
2022-XX-XX 00:00:00 -- [ERRO] it seems docker compose plugin is not installed. Please install it and try again.
```

Then probably is because the `compose` plugin is not installed for the `root` user. Currently Sedge uses `sudo` (root / admin powers) to run `docker compose` commands. We are working on this issue. To solve it, you can copy the `compose` plugin from your Home directory to `/root/` (root's Home directory) with the following command:

```
sudo cp ~/.docker/cli-plugins/docker-compose /root/.docker/cli-plugins/
```

### One of more services are not running after running the script

Work in progress

### When using a public execution or consensus node, the JWT token is not valid

Currently, to use public execution or consensus nodes on any network other than mainned, you need to a valid JWT token.

As sedge creates a JWT token for you, you can find the relevant public key in `docker-compose-scripts/jwtsecret`, which you should include in the remote node to enable communication between the two nodes.
