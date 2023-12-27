# Catch all invalid subdomains

## Create default zone

`sudo vim /etc/nginx/sites-available/default`

```
server {
    listen 80 default_server;
    listen [::]:80 default_server;
    listen 443 default_server ssl;
    listen [::]:443 default_server ssl;
    ssl_certificate /etc/nginx/ssl/nginx.crt;
    ssl_certificate_key /etc/nginx/ssl/nginx.key;
    return 444; # silently drop the connection
    # or you can define some landing page here
}
```

`sudo ln -s /etc/nginx/sites-available/default /etc/nginx/sites-enabled/`

## Create ssl cert

```
sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/nginx.key -out /etc/nginx/ssl/nginx.crt
# Test conf
sudo nginx -T
# Reload
sudo service nginx reload
```

Reference: https://stackoverflow.com/questions/72600765/catch-all-nginx-server-blocks-for-invalid-subdomains, https://jonnev.se/nginx-default-server-with-https/
