## I am getting "403 Forbidden" error

If you see the message "identity is not registered" for the 403 error, that means that you are trying to connect with an unregistered identity.

You need to finish registration process for your identity to be accepted.

This error might also mean that you are connecting to discovery service which is closed for Your IP, either because someone from your IP was abusing the service or due to load issues.

## Can You whitelist my node's IP for Testnet?

Discovery is already open for all nodes willing to test the service.

## What client software to use with Mysterium Network node?

Mysterium team created client software called **Mysterium VPN** and it is available [here](https://github.com/mysteriumnetwork/mysterium-vpn/releases).

You can fill out a [form](https://docs.google.com/forms/d/e/1FAIpQLSfrqlGA2CwEzl24jjHsWC1b6IFcp-H5P6sx5d3j226PYNCgDQ/viewform) for our alpha test client.

You can also use lightweight CLI client available from Mysterium Node repository. If you have node sources you can try running:
```shell
# cd github.com/mysterium/node
# ./bin/client_build; ./bin/client_run_cli
```

## I launched a node, how to know if it is working?

Best is to test with a client. Still there are other ways to be reasonably sure.
If you run your node via docker, check its status and logs:
```shell
# docker ps
# docker logs -f myst
```
You should see `myst` container running. If not, check the logs. You should see something like this:
```
...
[Mysterium.api] Identity registered: 0x4cd126119cd14e38c90e34dd8b6e0e2174b71123
[Mysterium.api] Proposal pinged for node: 0x4cd126119cd14e38c90e34dd8b6e0e2174b71123
...
```
It means that node successfully registered to discovery and its service proposal is available for mysterium network clients.

## In logs I see client's attempts, but it does not fully connect

There might be many things, but most frequent is firewall. If You run node via docker image, check that _ip_forwarding_ is enabled on a host and that UDP service port (1194 by default) is allowed from outside.

check ip_forward status:
```shell
# cat /proc/sys/net/ipv4/ip_forward
```
enable ip_forward if disabled:
```shell
# sysctl -w net.ipv4.ip_forward=1
```
It also might be that default firewall forward policy is set to `DROP`. In that case try setting it to `ACCEPT`.
Generic way to do it, provided there are no other interfering rules:

```shell
# iptables -P FORWARD ACCEPT
```

## I suspect that my hosting provider blocks ports, what to do?

For some hosting providers it is common to open just some pre-defined ports that are commonly used. Try setting any random service to test if incoming UDP port (i.e. 1194) is open. You can try using custom port using `openvpn.port=1234` option.

Sometimes hosting providers block most UDP ports altogether, even outgoing ones. In that case You can still run a node using TCP protocol. You can do that using `-openvpn.proto=TCP` option.

## I suspect that my firewall blocks access, what to do?

You can check all firewall rules with these commands:
```shell
# iptables -L -n
# iptables -L -n -t nat
```

To completely strip host of firewall rules and chains You can do:
```shell
# iptables -F
# iptables -t nat -F
# iptables -X
```

## Every time I reboot a host, I see unneeded firewall rules

Depending on Linux distribution You run, there might be different default firewall policies.
Sometimes You might need to change / disable certain default policies. See respective firewall documentation of Your OS.

In some cases, firewall rules that are being introduced by docker package might interfere with Mysterium Network node rules (such as in Centos7 for example).  In that case try changing conflicting rules or disable extended firewall rules by docker altogether. This might be achieved passing `--iptables=false` option before starting docker service.

## I want to backup / restore my identity, how should I do that?

On the first run node generates its identity automatically. Later this identity is reused each time node is started. If You run Mysterium Network node from sources, binaries or deb packages You can find Your identity in keystore directory in Your home directory (i.e. `~/.keystore`).

If You use docker image, it is strongly recommended **not** to store identity inside docker container (since You might need to remove container in order to upgrade), but to mount Your keystore from Your host to container using `-v host_volume:container_volume` option  as follows:
```shell
# docker run --cap-add NET_ADMIN --net host --publish "1194:1194" -v /home/mysterium-node:/var/lib/mysterium-node --name myst -d mysteriumnetwork/myst service --identity.passphrase=your_passphrase_here
```
Now You can copy and safely store `/home/mysterium-node/keystore` directory to backup all Your identities. Don't forget to save your passphrase that was used with generated identity, otherwise You will not be able to use that identity or access its wallet.

## How should I set a passphrase for node identity?

By default, generated identity is not protected with any password, that is password is an empty string. If You want to generate a password protected identity You can add `--identity.passphrase` to running command:
```shell
# docker run --cap-add NET_ADMIN --net host --publish "1194:1194" --name myst -d mysteriumnetwork/myst service --identity.passphrase=your_passphrase_here
```
If You have multiple identities, You can choose exact identity using `--identity` option. If no identity exists, new one will be created automatically.

After node restart, if no `--identity` option is present, last one used will be reused.

## VPN speed through my node is really low, what can I do?

Speed is affected by many things. Things to check:
* What is the speed between Your desktop and server host without VPN?
* What is delay from Your desktop to host?
* Does Your hosting provider guarantees You a certain minimal throughput?
* Does the speed depends on the time of day?
* Do You have free resources on Your host? (idle CPU / free memory)

To achieve best results, client (desktop computer) has to be as close (low delay) as possible.
Try a different hosting if speed is constantly low.
