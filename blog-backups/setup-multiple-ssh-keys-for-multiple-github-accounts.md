Imagine this scenario, you have 3 Github accounts, each have their own purposes. Account A for full time work, account B for portfolio projects and account C for personal projects. Whenever you are going to switch between repo in different accounts, you will need to use a different SSH keys for each account.

Changing SSH config every time you are switching is not good use of your time. Now you might start to think it is possible to have multiple SSH keys for each account and let SSH automatically manage it? Yeah, it is possible, but it can be quite tricky because there is a gotcha which I will explain further in this article.

I am assuming you are running Linux and already have openssh installed. For Mac you should be fine. For Windows you need to have git bash installed or WSL.

## Creating SSH keys

The first thing we need to do is to create a SSH key for each account.
Go to your terminal and run each command.

```bash
ssh-keygen -t rsa -f ~/.ssh/id_rsa_fulltime -C "fulltime@mail.com"
ssh-keygen -t rsa -f ~/.ssh/id_rsa_portfolio -C "portfolio@mail.com"
ssh-keygen -t rsa -f ~/.ssh/id_rsa_personal -C "personal@mail.com"
```

The first command is going to create a SSH key for full time account that are connected with the specified Github email: fulltime@mail.com and save it into the specified path: `~/.ssh/id_rsa_fulltime`.
The second and third commands do the same thing respectively for portfolio account and personal account.

## Adding SSH Keys Into SSH Authentication Agent

```bash
ssh-add ~/.ssh/id_rsa_fulltime
ssh-add ~/.ssh/id_rsa_portfolio
ssh-add ~/.ssh/id_rsa_personal
```

## Creating SSH Config File

Now you need to go to `~/.ssh` and create a file named `config` inside it.
Open it with vim or other text editor and add the following configurations.

```bash
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

`AddKeysToAgent` tells SSH to remember your password so you don't need to enter it every time.

`IdentitiesOnly` is the gotcha I mean earlier. Let's say you are currently authenticated with fulltime account and you want to switch to personal account. If you did not specify it, **SSH will continue to use the first authenticated SSH key even though we have already specified different SSH key for each account**, in this case it will continue to use fulltime SSH key for each account.

The postfix (eg -personal) after `github.com` is used to identify different SSH keys, you can name it whatever you like but it should be unique, usually I named it the context in which the Github account is used for.

## Final Setup

If you already have an existing cloned repo from Github, you need to go the repo directory and edit `.git/config` file to append identifier postfix to the github.com hostname. For example you already clone a repo named `personal-project` from your personal account, to connect the correct SSH key to the repo you have to append this postfix `-personal` to the `github.com` hostname.

```bash
[remote "origin"]
url = git@github.com-personal:personal-account/personal-project.git
```

If you want to clone a new repo. Simply append the identifier postfix to the `github.com` hostname of the clone url.
`git clone git@github.com-personal:personal-account/personal-project.git`

After that, you will need to setup Github username and email for the repo.

```bash
git config user.name "personal-account"
git config user.email "personal@mail.com"
```

Congratulations, you have successfully setup multiple SSH keys.
See you in the next article!
