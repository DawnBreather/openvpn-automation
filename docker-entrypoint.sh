#!/bin/bash
# Needed for openvpn
mkdir -p /run/openvpn
mkdir -p /dev/net
mknod /dev/net/tun c 10 200

# Firewall
# Allow traffic initiated from VPN to access "the world"
iptables -F
# Allow traffic initiated from VPN to access "the world"
iptables -I FORWARD -i vpninterface -o eth0 -m conntrack --ctstate NEW -j ACCEPT
# Masquerade traffic from VPN to "the world" -- done in the nat table
iptables -t nat -I POSTROUTING -o eth0 -j MASQUERADE

iptables -A PREROUTING -d 1.1.1.1/32 -i vpninterface -j DNAT --to-destination 10.253.3.234
iptables -A PREROUTING -d 10.252.186.254/32 -i vpninterface -j DNAT --to-destination 10.253.3.234
iptables -A PREROUTING -d 8.8.8.8/32 -i vpninterface -j DNAT --to-destination 10.253.3.234

iptables -A FORWARD -m state --state RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -p icmp -j ACCEPT
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE

#iptables -A FORWARD -p udp -m udp --dport 53 -j ACCEPT
#iptables -A INPUT -p udp -m udp --dport 53 -j ACCEPT
#iptables -t nat -A POSTROUTING -p udp -m udp --dport 53 -j ACCEPT

#iptables -A FORWARD -p tcp -m tcp --dport 80 -j ACCEPT
#iptables -A INPUT -p tcp -m tcp --dport 80 -j ACCEPT
#iptables -t nat -A POSTROUTING -p tcp -m tcp --dport 80 -j ACCEPT

#iptables -A FORWARD -p tcp -m tcp --dport 443 -j ACCEPT
#iptables -A INPUT -p tcp -m tcp --dport 443 -j ACCEPT
#iptables -t nat -A POSTROUTING -p tcp -m tcp --dport 443 -j ACCEPT

iptables -A FORWARD -p udp -j ACCEPT
iptables -A INPUT -p udp -j ACCEPT
iptables -t nat -A POSTROUTING -p udp -j ACCEPT

iptables -A FORWARD -p tcp -j ACCEPT
iptables -A INPUT -p tcp -j ACCEPT
iptables -t nat -A POSTROUTING -p tcp -j ACCEPT


# Needed permissions - especially for stupid stuff like CRL
chown -R nobody:nogroup /etc/openvpn
chmod -R 700 /etc/openvpn

# Run actual OpenVPN
exec /usr/sbin/openvpn --writepid /run/openvpn/server.pid --cd /etc/openvpn --config /etc/openvpn/server.conf --script-security 2