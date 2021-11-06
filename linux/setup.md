## Installing Linux Mint

Note: Step 6, 9, 10 and 11 is important to do if you use Nvidia graphic card because it doesn't support free driver provided by Linux Mint, otherwise you will boot into black screen.

1. Burn iso into usb
2. Reboot
3. Boot into bios (hold F2 after intel logo appear)
4. Reorder the boot order so usb is in the highest order
5. Save setting and reboot
6. After Linux Mint menu appear, select `start in compability mode`
7. Install Linux Mint
8. Repeat step 3-6 again
   <!-- 9. Edit `/boot/grub/grub.cfg`, replace every line `quiet splash` with `noapic noacpi nosplash irqpoll nomodeset` -->
   <!-- 10. Edit `/etc/default/grub`, replace `quiet splash` with `noapic noacpi nosplash irqpoll nomodeset` -->
9. `xed admin:///etc/default/grub`
10. Replace `GRUB_CMDLINE_LINUX=""` with `GRUB_CMDLINE_LINUX="nomodeset"`
11. `sudo update-grub` and reboot
12. Boot into bios and change the boot order back to normal and reboot
13. Open driver manager, install recommended Nvidia driver and reboot

## Setup Linux Mint

1. Open Update Manager -> Edit -> Software Sources -> Mirrors -> Select the closest mirror to your location (for both, main & base) -> OK
2. `sudo apt update & apt upgrade`
3. Nvidia X Server Settings -> X Server Display Configuration -> Advanced -> Force Full Composition Pipeline (fix screen tearing) -> Apply
4. Open System Settings -> Keyboard -> Shortcuts -> Add custom shortcut -> Set name to Take a screenshot of area and command to `mate-screenshot -a` -> Add -> Click on Keyboard bindings entry -> Type your shortcut (Shift + Print)
5. Open System Settings -> Keyboard -> Shortcuts -> Add custom shortcut -> Set name to Xkill and command to xkill -> Add -> Click on Keyboard bindings entry -> Type your shortcut (CTRL + Esc)
6. Open .config directory, create redshift directory and open it, create redshift.conf:

```bash
[redshift]
location-provider=manual

[manual]
lat=0.943520
lon=108.999298
```

5. Enable firewall
6. `sudo apt install chromium git neovim zsh fonts-firacode`
7. Copy .zshenv into ~
8. Copy .zshrc, aliases.zsh into ~/.config/zsh
9. Clone 3 zsh plugins into ~/.config/plugins
10. Install vim plug
11. Copy init.vim plugins.vim into ~/.config/nvim
12. Install nvm
13. Intsall LTS node.js

## Setup git

1. `git config --global user.name "rizkyzhang"`
2. `git config --global user.email "rizkyzhangdev@gmail.com"`
3. `git config --global core.editor "codium --wait"`
4. `git config --global diff.tool "default-difftool"`
5. `git config --global difftool.default-difftool.cmd "codium --wait --diff \$LOCAL \$REMOTE"`
6. Copy .ssh directory into home and chmod 600
7. `eval "$(ssh-agent -s)"`
8. `ssh-add ~/.ssh/id_ed25519`

## Setup VS Codium

1. Install VS code
2. Install settings sync, login Github, choose VS Code gist
3. Command pallete -> Sync: Download Settings

## Decrease swappiness value to 20

1. `xed admin:///etc/sysctl.conf`
2. Add `vm.swappiness=20` to the bottom of the file
3. Save and reboot

## Set a reasonable maximum log size for systemd

1. `sudo journalctl --vacuum-size=50M`
2. `sudo sed -i 's/#SystemMaxFiles=100/SystemMaxFiles=7/g' /etc/systemd/journald.conf`
