1. Open `/etc/hosts` file
2. Add `127.0.0.1 your-domain.com` and save the file
3. If you haven't install nginx, `sudo apt install nginx`
4. Configure virtual hosts `sudo nano /usr/local/etc/nginx/nginx.conf`
5. Add the following config:

```nginx
server {
    listen 80;
    listen [::]:80;
    server_name your-domain.com www.your-domain.com;
    access_log /var/log/nginx/your-domain.com.access.log;
    error_log /var/log/nginx/your-domain.com.error.log;
    location / {
        proxy_pass http://127.0.0.1;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection ‘upgrade’;
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

After saving the file, run `sudo systemctl restart nginx.service  
`.
