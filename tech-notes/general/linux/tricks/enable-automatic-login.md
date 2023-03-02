1. `sudo vim /etc/lightdm/lightdm.conf`
2. Find `autologin-user` line and uncomment it
3. Next to the `autologin-user` line, write in your username. Be sure to write it exactly as it appears on the system. One missed character could mess up your entire LightDM configuration. Example `autologin-user=john`
4. Reboot

Reference: https://www.addictivetips.com/ubuntu-linux-tips/enable-automatic-login-on-linux/
