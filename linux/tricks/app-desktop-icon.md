# Create a desktop icon on Ubuntu or Debian based systems

Steps:
1. Create a [application-name].desktop file in /usr/share/applications/
2. Copy the following into the .desktop file and update it with the details you want.
```bash
[Desktop Entry]
Version=1.0
Name=Name of the application
Comment=text you want to show when hovering above the icon
Exec=Path of the executable file %U
Icon=Path of the image
Terminal=false
StartupWMClass=name of application
Type=Application
Categories=category of the application
MimeType=Type of application it should open
```
3. Logout and login.

Example:
```bash
[Desktop Entry]
Name=Firefox Developer 
GenericName=Firefox Developer Edition
Exec=/opt/firefox_dev/firefox %U
Terminal=false
Icon=/opt/firefox_dev/browser/chrome/icons/default/default128.png
Type=Application
Categories=Application;Network;X-Developer;
Comment=Firefox Developer Edition Web Browser.
StartupWMClass=Firefox Developer Edition
```
