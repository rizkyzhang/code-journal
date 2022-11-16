1. Create an app.desktop file with the following content

```
[Desktop Entry]
Name=Test
Type=Application
Exec=bash -c './app-cli;$SHELL'
Terminal=true
```

2. Put app-cli in the same folder

Note: `$SHELL` is needed to prevent the command line to exit immediately.
