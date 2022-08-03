## Installing MX Linux

1. Burn iso into usb
2. Reboot
3. Boot into bios (hold F5)
4. Reorder the boot order so usb is in the first position
5. Save setting and reboot
6. After boot is completed, click intall on the installer
7. Follow the installer instruction
8. Reboot to the installed MX Linux

## General setup

1. `sudo apt update && sudo apt upgrade`
2. `sudo apt install zsh`
3. `chsh -s zsh`
4. Install zgenom https://github.com/jandamm/zgenom
5. Copy linux dotfiles to root https://github.com/rizkyzhang/dotfiles
6. Copy .ssh to root and `chmod 600 ~/.ssh`
7. Install chrome, beekeeper studio, slack, telegram, vscode
8. Autostart chrome, redshift, and terminal
9. Install nvm https://github.com/nvm-sh/nvm
10. `nvm install --lts`
11. Download go linux-amd64 tar.gz https://go.dev/dl/
12. `tar -C /usr/local -xzf go*.linux-amd64.tar.gz`
13. Install docker https://docs.docker.com/engine/install/debian/
14. `sudo usermod -aG docker $USER`
15. `echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p`

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
