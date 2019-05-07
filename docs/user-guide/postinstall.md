## Node installed with a DEB package
Mysterium node can be installed using a DEB package for most of the Debian-like systems.

[Prebuilt image](/user-guide/installation/#running-mysterium-node-on-a-raspberry-pi) for Raspberry Pi is using a node installed with the same DEB package.

### Startup configuration
The installed DEB package will be starting a node service automatically.

You can check details using the following command:

```
sudo systemctl status mysterium-node.service
```

Systemd unit file, that describes service is located at the: `/lib/systemd/system/mysterium-node.service`.

It should look similar to the following:

```
[Unit]
Description=Server for decentralised VPN built on blockchain
Documentation=https://mysterium.network/
After=network-online.target

[Service]
User=mysterium-node
Group=mysterium-node

RuntimeDirectory=mysterium-node
RuntimeDirectoryMode=0750
LogsDirectory=mysterium-node

EnvironmentFile=-/etc/default/mysterium-node
ExecStart=/usr/bin/myst $CONF_DIR $RUN_DIR $DISCOVERY $BROKER $PROTO service --agreed-terms-and-conditions
KillMode=process
TimeoutStopSec=10
SendSIGKILL=yes
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

The line starting with an `ExecStart=` describes how the process starts.
You can change the line to configure the service by your need.

List of available options can be found using a help commands: `myst --help` or `myst service --help`.

For example, if you want to start only OpenVPN service on 1190 port and serve Mysterium verified consumer only, change the `ExecStart=` line to the following:

```
ExecStart=/usr/bin/myst $CONF_DIR $RUN_DIR $DISCOVERY $BROKER $PROTO service openvpn --openvpn.port=1190 --access-policy.list mysterium --agreed-terms-and-conditions
```

To apply the changed service configuration you will need to re-read configuration and restart it.

```
sudo systemctl daemon-reload
```

```
sudo systemctl restart mysterium-node
```

### TequilAPI and CLI
Runtime configuration of the running Mysterium node can be changed using a [TequilAPI](https://tequilapi.mysterium.network/) or CLI commands.

**TequilAPI** is a REST API that allows you to manipulate a Mysterium node multiple ways.

```
curl http://localhost:4050/healthcheck

{
    "uptime": "47m55.616006453s",
    "process": 5490,
    "version": "0.0.0-dev",
    "buildInfo": {
        "commit": "99d37edc952b736e3fad069f56e7d968276634a8",
        "branch": "master",
        "buildNumber": "4755"
    }
}
```

**CLI** is a command line interface that allows you manipulate a Mysterium node using just your terminal.
It can be started using a `myst cli` command:

```
myst@raspberrypi:~ $ myst cli
Mysterium Node
  Version: 0.0.0-dev
  Build info: Branch: master. Build id: 4755. Commit: 99d37edc952b736e3fad069f56e7d968276634a8
...
» healthcheck
[INFO] Uptime: 50m25.039797872s
[INFO] Process: 5490
[INFO] Version: 0.0.0-dev
[INFO] Branch: master. Build id: 4755. Commit: 99d37edc952b736e3fad069f56e7d968276634a8
»
```

#### Identities

To get a list of Mysterium node identities you can use the following commands:

```
curl http://localhost:4050/identities

{
    "identities": [
        {
            "id": "0x041e42becaa3a6f92e155a7a5cab62c49dcdc578"
        }
    ]
}
```

```
myst@raspberrypi:~ $ myst cli
...
» identities list
[+] 0x041e42becaa3a6f92e155a7a5cab62c49dcdc578
»
```

To add a payout ETH address to the identity use the following commands:

```
curl -X PUT -d '{"ethAddress":"0x000000000000000000000000000000000000000a"}' http://localhost:4050/identities/0x041e42becaa3a6f92e155a7a5cab62c49dcdc578/payout
```

```
myst@raspberrypi:~ $ myst cli
...
» payout set 0x041e42becaa3a6f92e155a7a5cab62c49dcdc578 0x000000000000000000000000000000000000000a
[SUCCESS] Payout address 0x000000000000000000000000000000000000000a registered.
»
```

### Logs

To get a detailed log for Mysterium node service, you can use the following command:

```
sudo journalctl -u mysterium-node.service
```
