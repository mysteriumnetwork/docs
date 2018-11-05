## Welcome to Mysterium Network

**Mysterium Node** is VPN management software written in Go. Our goal is to facilitate the growth of
a VPN network which could provide better privacy and anonymity. We strive to develop a toolkit which would allow programmers to incorporate VPN services into their dApps.

We want to aid the development of easy and secure access to the distributed Mysterium Network, providing security and anonymity for their users.

Mysterium Node is just one element in a toolkit. Currently, we develop and maintain a VPN API called **TequilAPI**, which is a part of the distributed Mysterium Node. This TequilAPI can be used in different VPN clients that need access to Mysterium Network, see documentation [here](http://tequilapi.mysterium.network/).

We are also developing a proof of concept client. We call it **MysteriumVpn** and it is also completely open-sourced. You can get it [here](https://github.com/mysteriumnetwork/mysterium-vpn/releases).

We encourage everyone to download the node software, try it out and give us your feedback. We looking forward to your contributions.

[![Go Report Card](https://goreportcard.com/badge/github.com/mysteriumnetwork/node)](https://goreportcard.com/report/github.com/mysteriumnetwork/node)
[![Build Status](https://travis-ci.org/mysteriumnetwork/node.svg?branch=master)](https://travis-ci.org/mysteriumnetwork/node)

Cross-platform software to run a node in Mysterium Network. It contains Mysterium server (node),
client API (tequila API) and client-cli (console client) for Mysterium Network.

Currently node supports OpenVPN as its underlying VPN transport.

## Getting Started

- [Homepage](https://mysterium.network)
- [Whitepaper](https://mysterium.network/whitepaper.pdf)
- [Latest](https://github.com/mysteriumnetwork/node/releases/latest) release
- [Installation guide](/user-guide/installation/)

### Prerequisites

To run a node as docker container You will need [docker](https://www.docker.com/).

You should be able to run a node on any OS that supports docker.

Tested on these OSes so far: _Debian 9_, _Ubuntu 16.04_ and _CentOS 7_.

You can check latest docker node versions here: [https://hub.docker.com/r/mysteriumnetwork/myst/](https://hub.docker.com/r/mysteriumnetwork/myst/)

### Installation

Go to [docker](https://www.docker.com/) on how to get a recent docker version for Your OS.

### Running
```bash
sudo docker run --cap-add NET_ADMIN --net host --name myst -d mysteriumnetwork/myst service --agreed-terms-and-conditions
```

### Debugging
```bash
sudo docker logs -f myst
```

More detailed installation options described [here](/user-guide/installation/).

For possible issues while running a node refer to our [FAQ](/user-guide/faq/) and [Troubleshooting](/user-guide/troubleshooting/) sections.

## Built With

* [go](https://golang.org/) - The Go Programming Language
* [travis](https://travis-ci.org/) - Travis continuous integration tool
* [docker](https://www.docker.com/what-docker) - Containerize applications
* [openvpn](https://openvpn.net) - Solid VPN solution
