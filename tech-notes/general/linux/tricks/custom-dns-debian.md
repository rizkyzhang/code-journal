1. `sudo apt install resolvconf`
2. `sudo systemctl enable resolvconf`
3. `sudo systemctl start resolvconf`
4. `sudo systemctl status resolvconf`
5. `sudo nano /etc/resolvconf/resolv.conf.d/head`
6. Add `nameserver 8.8.8.8`
7. `sudo resolvconf --enable-updates`
8. `sudo resolvconf -u`

Everytime we want to change the dns, edit this file `/etc/resolvconf/resolv.conf.d/head` and then run `sudo resolvconf -u`
