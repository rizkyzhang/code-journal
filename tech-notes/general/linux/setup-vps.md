## First steps

### 1. `apt update && apt upgrade`

### 2. `adduser username`

### 3. `adduser username sudo`

### 4. Add ssh key to your vps user

If you have have password-based SSH access available, run `ssh-copy-id /path/to/ssh-private-key username@ip-address` and the setup is done, otherwise follow below steps.

If you do not have password-based SSH access available, you must add your public key to the remote server manually.

On your local machine, output and copy the contents of your public key. `cat ~/.ssh/id_rsa.pub`.

On your vps, create the `~/.ssh` directory if it does not already exist:

`mkdir -p ~/.ssh`

The public keys listed in `~/.ssh/authorized_keys` are the ones that you can use to log in to the server as this user, so you need to add the public key you copied into this file.

`echo "ssh-rsa EXAMPLEzaC1yc2E...GvaQ== username@ip-address" >> ~/.ssh/authorized_keys`

The `~/.ssh` directory and `authorized_keys` file must have specific restricted permissions (700 for `~/.ssh` and 600 for `authorized_keys`). If they don’t, you won’t be able to log in.

Once the authorized_keys file contains the public key, set the permissions and ownership of the files:

`chmod -R go= ~/.ssh`
`chown -R $USER:$USER ~/.ssh`

## Secure SSH

1. Setup firewall to only allow authorized IPs (you can also VPN) to connect to the SSH port
2. Change the default SSH Port (port 22) to random port ranging from 1025 to 65535, for example port 5723 `Port 5723`
3. Disable root login, `PermitRootLogin no`
4. Disable password login, `PasswordAuthentication no`
5. Disable empty password login, `PermitEmptyPasswords no`
6. Disable X11/TCP port forwarding, `AllowTcpForwarding no` and `X11Forwarding no`
7. Limit user's SSH access `AllowUsers user1 user2`
8. Only use SSH protocol 2, change `#Protocol 2, 1` to `Protocol 2`

---

## Setup ufw firewall

### 1. Setup openvpn

This setup is needed in order for client to able to connect to the openvpn, ufw firewall will route all traffic from the client through the OpenVPN Server.

#### 1.1 Forwarding policy

Change default forward policy, edit `/etc/sysctl.conf` to permanently enable ipv4 packet forwarding. (Note: This will take effect at next boot).

```
# Enable packet forwarding
net.ipv4.ip_forward=1
```

#### 1.2 UFW config

Edit /etc/default/ufw

`DEFAULT_FORWARD_POLICY="ACCEPT"`

#### 1.3 UFW before rules

Edit `/etc/ufw/before.rules` to add the following code after the header and before the “filter” line. Match the IP/subnet mask to the same one as in `/etc/openvpn/server.conf`.

```conf
# START OPENVPN RULES
# NAT table rules
*nat
:POSTROUTING ACCEPT [0:0]
# Allow traffic from OpenVPN client to eth0
-A POSTROUTING -s 10.8.0.0/8 -o eth0 -j MASQUERADE
COMMIT
# END OPENVPN RULES
```

#### 1.4 Allow incoming connection to openvpn port (default is 1194 unless you changed it)

`sudo ufw allow 1194`

### 2. Only allow specific IP to access SSH port

`sudo ufw allow from 11.111.1.11 to 11.111.1.11 port 1111 proto tcp`

### 3. Script to only allow cloudflare ips to port 80 and 443

```bash
#!/bin/bash

set -euo pipefail

# Get the Cloudflare IPs
curl -s https://www.cloudflare.com/ips-v4 -o /tmp/cloudflare_ips_$$
echo "" >> /tmp/cloudflare_ips_$$
curl -s https://www.cloudflare.com/ips-v6 >> /tmp/cloudflare_ips_$$

# Allow traffic from Cloudflare IPs on all ports.
for ip in $(cat /tmp/cloudflare_ips_$$)
do
  echo "add cloudflare ip $ip"
  ufw allow from $ip to any port 80,443 proto tcp comment 'Cloudflare'
done

# Reject HTTP/HTTPS, for good measure.
ufw reject 80
ufw reject 443

# Reload ufw.
ufw reload > /dev/null

# Show the rules to verify it worked.
ufw status numbered
```

### 4. UFW hide published docker port

Modify the UFW configuration file `/etc/ufw/after.rules` and add the following rules at the end of the file:

```conf
# BEGIN UFW AND DOCKER
*filter
:ufw-user-forward - [0:0]
:ufw-docker-logging-deny - [0:0]
:DOCKER-USER - [0:0]
-A DOCKER-USER -j ufw-user-forward

-A DOCKER-USER -j RETURN -s 10.0.0.0/8
-A DOCKER-USER -j RETURN -s 172.16.0.0/12
-A DOCKER-USER -j RETURN -s 192.168.0.0/16

-A DOCKER-USER -p udp -m udp --sport 53 --dport 1024:65535 -j RETURN

-A DOCKER-USER -j ufw-docker-logging-deny -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -d 192.168.0.0/16
-A DOCKER-USER -j ufw-docker-logging-deny -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -d 10.0.0.0/8
-A DOCKER-USER -j ufw-docker-logging-deny -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -d 172.16.0.0/12
-A DOCKER-USER -j ufw-docker-logging-deny -p udp -m udp --dport 0:32767 -d 192.168.0.0/16
-A DOCKER-USER -j ufw-docker-logging-deny -p udp -m udp --dport 0:32767 -d 10.0.0.0/8
-A DOCKER-USER -j ufw-docker-logging-deny -p udp -m udp --dport 0:32767 -d 172.16.0.0/12

-A DOCKER-USER -j RETURN

-A ufw-docker-logging-deny -m limit --limit 3/min --limit-burst 10 -j LOG --log-prefix "[UFW DOCKER BLOCK] "
-A ufw-docker-logging-deny -j DROP

COMMIT
# END UFW AND DOCKER
```

Using command sudo systemctl restart ufw or sudo ufw reload to restart UFW after changing the file. Now the public network can't access any published docker ports, the container and the private network can visit each other normally, and the containers can also access the external network from inside.

If you want to allow public networks to access the services provided by the Docker container, for example, the service port of a container is `80`. Run the following command to allow the public networks to access this service:

`ufw route allow proto tcp from any to any port 80`

This allows the public network to access all published ports whose container port is `80`.

Note: If we publish a port by using option `-p 8080:80`, we should use the container port `80`, not the host port `8080`.

---

Reference:

- https://docs.digitalocean.com/products/droplets/how-to/add-ssh-keys/to-existing-droplet/
- https://www.gaggl.com/2013/04/openvpn-forward-all-client-traffic-through-tunnel-using-ufw/
- https://www.stavros.io/posts/block-non-cloudflare-ips-with-ufw/
- https://github.com/chaifeng/ufw-docker
