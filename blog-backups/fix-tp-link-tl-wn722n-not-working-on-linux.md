This issue has been bugging me since last year. It started when I am still running Linux Mint, suddenly after an update, my wireless adapter stop working suddenly. I try to find a way to fix it on Google, but did not found any working solutions. Because I am quite busy at that time, I simply used USB tethering to get around it, and soon forget about the problem.

Early this year, I decided to try MX Linux because it seems to be more lightweight and stable (once a year release compared to monthly release which mean less crash) than Linux Mint. I did notice the performance is noticeably better, and I never have a crash or bug anymore after update, overall it is very stable, but the wireless adapter problem still persist, so clearly the problem is not the OS.

Tonight I remembered about the issue and because I am still curious about it, I decided to research about the issues again with a fresh approach. The first thing I do is running `lsusb` command on the terminal after I connected my wireless adapter. This is the output:

![ lsusb output](https://xmyforsgwfqphgifjxpe.supabase.co/storage/v1/object/public/strapi-blog/strapi-blog/lsusb.png-16bd855051b18e46f53d34ff8671297c.png)

Gotcha! I found the needed driver version for my wireless adapter which is `RTL8188EUS`. I then run `lsmod | grep 8188` to check whether the driver is loaded or not on the kernel, turn out it is not installed yet.

After a brief google search for the driver, I found it on https://github.com/aircrack-ng/rtl8188eus. I tried the instruction there, but sadly it still doesn't work. Suddenly I got an idea to try this command `echo "blacklist r8188eu" | sudo tee /etc/modprobe.d/r8188eu.conf` and restart the computer. Finally, it's working! To spare you from wasting time like me, I will give you the precise step here:

1. Connect your wireless adapter
2. Open your terminal and run ` lsusb` to find your driver model, if it is not the same as mine, you can modify this tutorial by searching the driver model that match yours on Google and modify the last command to `echo "blacklist <driver-model>" | sudo tee/etc/modprobe.d/<driver-model>.conf`
3. Disconnect your wireless adapter and connect to internet with USB tethering from your phone
4. `sudo apt-get install mokutil && mokutil --sb-state`, if it output ` SecureBoot disabled` you can proceed to the next step, otherwise you need to turn off the secure boot option from the BIOS because it can possibly interfere with the driver installation
5. `git clone https://github.com/aircrack-ng/rtl8188eus`
6. ` cd rtl8188eus`
7. `make && sudo make install`
8. `echo "blacklist r8188eu" | sudo tee /etc/modprobe.d/r8188eu.conf`
9. Restart computer

**Important note: you might need to repeat the steps again after updating your Linux because the driver might got replaced.**

I hope my article is useful to you, see you in the next article!
