## Installing MX Linux

1. Burn iso into usb
2. Reboot
3. Boot into bios (hold F2/F5/F11)
4. Reorder the boot order so usb is in the first position
5. Save setting and reboot
6. After boot is completed, click install on the installer
7. Follow the installer instruction
8. Reboot to the installed MX Linux

## General setup

1. `sudo apt update && sudo apt upgrade`
2. `sudo apt install openssh-server adb scrcpy ncdu zsh bat exa build-essential clang python3-pip kdenlive wireshark`
3. `chsh -s /usr/bin/zsh`
4. Install rust `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`
5. Install neovim with deb https://github.com/neovim/neovim/releases/
6. Install zgenom https://github.com/jandamm/zgenom
7. Copy linux dotfiles to root https://github.com/rizkyzhang/dotfiles
8. Copy .ssh to root and `chmod 600 ~/.ssh/*`
9. Install chrome, beekeeper studio, mongodb compass, slack, telegram, vscode
10. Autostart chrome, slack and terminal
11. Install nvm https://github.com/nvm-sh/nvm
12. `nvm install --lts`
13. Download go linux-amd64 tar.gz https://go.dev/dl/
14. `tar -C /usr/local -xzf go*.linux-amd64.tar.gz`
15. Install docker https://docs.docker.com/engine/install/debian/
16. `sudo usermod -aG docker $USER`
17. `echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p`
18. Download firacode nerd font https://www.nerdfonts.com/font-downloads
19. Unzip and copy to ~/.fonts
20. `fc-cache -fv` to manually rebuild the font cache
21. `sudo usermod -a -G wireshark coderz`

## Postman setup

1. Download postman tar.gz https://www.postman.com/downloads/
2. `tar -xzf postman*.tar.gz`
3. If any postman version is installed before, `sudo rm -rf /opt/Postman`
4. `sudo mv Postman /opt/Postman`
5. `sudo ln -s /opt/Postman/Postman /usr/bin/postman`
6. Create desktop icon for Postman

```
cat > ~/.local/share/applications/postman.desktop <<EOL
[Desktop Entry]
Encoding=UTF-8
Name=Postman
Exec=postman
Icon=/opt/Postman/app/resources/app/assets/icon.png
Terminal=false
Type=Application
Categories=Development;
EOL
```

## Redshift setup

1. `cd .config`
2. `mkdir redshift && cd redshift`
3. `vim redshift.config`

```bash
[redshift]
location-provider=manual

[manual]
lat=<current-lat>
lon=<current-lon>
```

4. `redshift -c .config/redshift/redshift.conf`
