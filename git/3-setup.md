# Setup

## Username & email

```bash
> git config --global user.name "username"
> git config --global user.email "email@gmailcom"
```

## VS Code as default editor & diff tool

```bash
> git config --global core.editor "codium --wait"
> git config --global diff.tool "default-difftool"
> git config --global difftool.default-difftool.cmd "codium --wait --diff \$LOCAL \$REMOTE"
```

## Show git config

```bash
> git config --list --show-origin
```
