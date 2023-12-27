---
tags:
  - linux-setup
---
- `sudo apt-get install mokutil && mokutil --sb-state`, if it output SecureBoot disabled you can proceed to the next step, otherwise you need to turn off the secure boot option from the BIOS because it can possibly interfere with the driver installation
- `git clone https://github.com/aircrack-ng/rtl8188eus`
- `cd rtl8188eus`
- `make && sudo make install`
- `echo "blacklist r8188eu" | sudo tee /etc/modprobe.d/r8188eu.conf`
- Restart computer
