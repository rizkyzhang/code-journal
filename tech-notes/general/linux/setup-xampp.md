- Download installer https://www.apachefriends.org/download.html
- `chmod +x ./Downloads/xampp-linux-x64-x.x.x-x-installer.run`
- `sudo ./Downloads/xampp-linux-x64-x.x.x-x-installer.run`
- Add this line into .zshrc `export PATH=/opt/lampp/bin:$PATH` so we can access php-cli
- Download composer.phar from https://getcomposer.org/download/
- Rename to composer
- `chmod +x composer`
- Move it to `/opt/lampp/bin`
- `sudo chown -R username:group /opt/lampp/htdocs`

## php.ini config

- Display error on web page, change line `display_errors=Off` to `display_errors=On`
