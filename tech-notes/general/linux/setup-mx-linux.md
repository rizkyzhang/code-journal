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
2. `sudo apt install openssh-server adb ncdu zsh build-essential python3-pip kdenlive fonts-firacode`
3. `chsh -s /usr/bin/zsh`
4. Copy linux dotfiles to root https://github.com/rizkyzhang/dotfiles
5. Copy .ssh to root and `chmod 600 ~/.ssh/*`
6. Install zgenom https://github.com/jandamm/zgenom
7. Install neovim with deb https://github.com/neovim/neovim/releases/
8. Install chrome, beekeeper studio, mongodb compass, slack, vscode, jetbrain ide, anki, tilix
9. Autostart chrome, slack and tilix
10. Install nvm https://github.com/nvm-sh/nvm -> `nvm install --lts` -> `npm i -g ni pnpm yarn localtunnel`
11. Download go linux-amd64 tar.gz https://go.dev/dl/ -> `tar -C /usr/local -xzf go*.linux-amd64.tar.gz`
12. Install docker engine https://docs.docker.com/engine/install/debian/ -> `sudo usermod -aG docker $USER`
13. `echo fs.inotify. max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p`
14. Install doppler cli
15. Setup work/dev rabbitmq, postgres and mongodb containers
16. Download firacode nerd font https://www.nerdfonts.com/font-downloads -> unzip and copy to ~/.fonts -> `fc-cache -fv` to manually rebuild the font cache
17. Enable systemd, open mx tools and add `init=lib/systemd/systemd` to the end of kernel parameters
18. Install rust `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`

## CLI apps setup

- Install air https://github.com/cosmtrek/air#via-go-install
- Install bat via amd64 deb https://github.com/sharkdp/bat/releases
- Install exa https://the.exa.website/install/linux#manual
- Install fzf via git https://github.com/junegunn/fzf#installation, `[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh` is already set in .zshrc
- Install golang migrate https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Install go global update https://github.com/Gelio/go-global-update#installation
- Install lazydocker https://github.com/jesseduffield/lazydocker#go
- Install lazygit https://github.com/jesseduffield/lazygit#go
- Install rip via binary https://github.com/nivekuil/rip#-installation
- Install scc https://github.com/boyter/scc#go-get
- Install zoxide via install script https://github.com/ajeetdsouza/zoxide#installation, `eval "$(zoxide init zsh)"` is already set in .zshrc
- Install miniconda
  - `pip install s3cmd`
  - `s3cmd --configure`

## Devops setup

- Install virtualbox via deb https://www.virtualbox.org/wiki/Linux_Downloads
- Install maven https://maven.apache.org/install.html
- Install vagrant https://developer.hashicorp.com/vagrant/downloads
- Install aws-cli https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html#cliv2-linux-install

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
3. `vim redshift.conf`

```bash
[redshift]
location-provider=manual

[manual]
lat=<current-lat>
lon=<current-lon>
```

4. `redshift` or `redshift -c .config/redshift/redshift.conf` (if the first doesn't works)
