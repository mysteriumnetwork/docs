[What is current state of the project?](#what-is-current-state-of-the-project)

[Where I can get Mysterium Network node software?](#where-i-can-get-mysterium-network-node-software)

[How should I run node software?](#how-should-i-run-node-software)

## What is current state of the project?

Currently project is in alpha state. Meaning that it is still under heavy development.
At the moment node discovery service is centralised and open only to development team.
Anyone can try out running Mystermium Network node, but they will also need their own discovery service for this. Current discovery service falls under api project name and is available [here](https://github.com/MysteriumNetwork/api).

As with node itself, latest api version can be downloaded from [releases](https://github.com/MysteriumNetwork/api) page, or from [DockerHub](https://hub.docker.com/r/mysteriumnetwork/mysterium-api/) directly:

```shell
# docker pull mysteriumnetwork/mysterium-api
```

## Where I can get Mysterium Network node software?

### Docker image

You can download Mysterium image from [DockerHub](https://hub.docker.com/r/mysteriumnetwork/mysterium-node/) directly:
```shell
# docker pull mysteriumnetwork/mysterium-node
```
It will download latest version of node.

### Packages and binaries

Latest stable versions can be found in [releases](https://github.com/MysteriumNetwork/node/releases) section of Github.

### Sources
If you are willing to try latest and greatest, you can fetch sources:
```shell
# git clone https://github.com/MysteriumNetwork/node.git
```

## How should I run node software?

Go to dedicated FAQ section on how to [[operate|node-operation]] a node.
