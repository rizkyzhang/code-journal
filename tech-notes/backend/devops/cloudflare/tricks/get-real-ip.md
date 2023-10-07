1. `HTTP_X_FORWARDED_FOR`, This header is added by the proxy and it contains a comma-separated list of IP addresses, where the first IP address is the IP address of the client and the subsequent IP addresses are the IP addresses of the proxies that the request has passed through.
2. `HTTP_CF_CONNECTING_IP`

Reference: https://saturncloud.io/blog/how-to-get-the-real-ip-address-behind-a-proxy-behind-cloudflare/
