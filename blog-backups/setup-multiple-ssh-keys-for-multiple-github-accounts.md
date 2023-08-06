Imagine this scenario, you have 3 GitHub accounts, each have their own purposes. Account A for full time work, account B for portfolio projects and account C for personal projects. Whenever you are going to switch between repo in different accounts, you will need to use a different SSH keys for each account.

Changing SSH config every time you are switching is not good use of your time. Now you might start to think it is possible to have multiple SSH keys for each account and let SSH automatically manage it? Yeah, it is possible, but it can be quite tricky because there is a gotcha which I will explain further in this article.

I am assuming you are running Linux and already have OpenSSH installed. For Mac, you should be fine. For Windows, you need to have git bash installed or WSL.

## Creating SSH keys

The first thing we need to do is to create a SSH key for each account.
Go to your terminal and run each command.

```conf
ssh-keygen -t rsa -f ~/.ssh/id_rsa_fulltime -C "fulltime@mail.com"
ssh-keygen -t rsa -f ~/.ssh/id_rsa_portfolio -C "portfolio@mail.com"
ssh-keygen -t rsa -f ~/.ssh/id_rsa_personal -C "personal@mail.com"
```

The first command is going to create a SSH key for full time account that are connected with the specified Github email: fulltime@mail.com and save it into the specified path: `~/.ssh/id_rsa_fulltime`.
The second and third commands do the same thing respectively for portfolio account and personal account.

## Adding SSH Keys Into SSH Authentication Agent

```conf
ssh-add ~/.ssh/id_rsa_fulltime
ssh-add ~/.ssh/id_rsa_portfolio
ssh-add ~/.ssh/id_rsa_personal
```

## Creating SSH Config File

Now you need to go to `~/.ssh` and create a file named `config` inside it. Open it with vim or other text editor and add the following configurations.

```conf
Host *
AddKeysToAgent yes
IdentitiesOnly yes

Host github.com-fulltime
HostName github.com
IdentityFile ~/.ssh/id_rsa_fulltime

Host github.com-portfolio
HostName github.com
IdentityFile ~/.ssh/id_rsa_portfolio

Host github.com-personal
HostName github.com
IdentityFile ~/.ssh/id_rsa_personal
```

`Host *` means the config will be applied to every host, not only `github.com-prefix`.

`AddKeysToAgent` tells SSH to remember your passphrase, so you don't need to enter it every time, only on every new login session.

`IdentitiesOnly` is the gotcha I mean earlier. Let's say you are currently authenticated with fulltime account, and you want to switch to personal account. If you did not specify it, **SSH will continue to use the first authenticated SSH key even though we have already specified different SSH key for each account**, in this case it will continue to use fulltime SSH key for each account.

The postfix ( eg -personal) after `github.com` is used to identify different SSH keys, you can name it whatever you like but it should be unique, usually I named it the context in which the GitHub account is used for.

## Final Setup

If you already have an existing cloned repo from GitHub, you need to go the repo directory and edit `.git/config` file to append identifier postfix to the github.com hostname.

For example, you already clone a repo named `personal-project` from your personal account, to connect the correct SSH key to the repo you have to append this postfix `-personal` to the `github.com` hostname.

```conf
[remote "origin"]
url = git@github.com-personal:personal-account/personal-project.git
```

If you want to clone a new repo. Simply append the identifier postfix to the `github.com` hostname of the clone url.
`git clone git@github.com-personal:personal-account/personal-project.git`

After that, you will need to set up Github username and email for the repo by running these commands:

```bash
git config user.name "personal-account"
git config user.email "personal@mail.com"
```

## Additional config if SSH still keep asking for passphrase (optional)

I encounter this problem recently and was so confused at first, luckily I found a [blog post](https://junyonglee.me/ssh/How-to-permanently-add-private-key-passphrase-to-ssh-agent/) which explains how to fix this issue using `keychain` library. `keychain` library is used to
store SSH passphrase for one login session, so you only need to reenter passphrase per login session.

First install the library if you haven't installed it:

```bash
sudo apt-get install keychain
```

Add this to the start of your shell config, `.zshrc` if you are using zsh or `.bashrc` if you are using bash.

```bash
if [[ `uname` == Linux ]] then
  /usr/bin/keychain $HOME/.ssh/id_rsa_personal
  /usr/bin/keychain $HOME/.ssh/id_rsa_portfolio
  /usr/bin/keychain $HOME/.ssh/id_rsa_fulltime
source $HOME/.keychain/$HOST-sh
fi
```

`/usr/bin/keychain` command will prompt for ssh passphrase for every login session and save it to `$HOME/.keychain/$HOST-sh`, in this case because we have 3 keys, it will be prompted 3 times.

## Bonus: complete SSH keys automation!

Currently we still need to append identifier postfix to the github.com hostname every time we clone a new repo so SSH know which key to be used.

This can be very repetitive and prone to typo. Is it possible to only specify the postfix only once? Yes, it is possible with the help `.gitconfig`. With this solution you also don't need to set email and username anymore.

For example, you have a personal SSH key and a directory named personal which contains repositories, you want it to use personal SSH key automatically. Here's the step to make it automatic:

- Create a `.gitconfig` in the root of personal directory.
- Specify email, username and identifier postfix in below format.

```conf
[user]
  email = test@gmail.com
  name = test

[url "git@github.com-personal"]
  insteadOf = git@github.com
```

- Create a global `.gitconfig` in your home directory to tell SSH which local `.gitconfig` to use for personal directory by specifying the path.

```conf
[includeif "gitdir:~/Projects/personal/"]
  path = ~/Projects/personal/.gitconfig
```

Now try to clone a repository, SSH will automatically choose the correct SSH key.

Congratulations, you have successfully setup multiple SSH keys.

See you in the next article!
