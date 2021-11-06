Problem: Create react app, reload not working.

Causes: Insufficient `max_user_wathces` value and React bug regarding fast refresh which replace hot reload but fast refresh did not work.

Solution:

1.  Increase `max_user_wathces` value temporarily.

```bash
> sudo -i
> echo 1048576 > /proc/sys/fs/inotify/max_user_watches
> exit
```

2. Increase `max_user_wathces` value permanently (persist after power off).

```bash
echo "fs.inotify.max_user_watches=1048576" >> /etc/sysctl.conf
sudo sysctl -p
```

3. Create a .env file in the root directory of your project and add FAST_REFRESH=false in the file.

Reference: https://stackoverflow.com/questions/42189575/create-react-app-reload-not-working
