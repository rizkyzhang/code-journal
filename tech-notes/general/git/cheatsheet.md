## Branch

- Checkout `git checkout <branch-name>`
- Create `git branch <branch-name>`
- Current `git branch --show-current`
- Delete `git branch -d <branch-name>`
- Delete remote `git push --delete origin <branch-name>`
- List `git branch`
- List remote `git branch -r`
- Rename `git branch -m <new-branch-name>`

### Merge, rebase and squash

- Abort merge `git merge --abort`
- Merge specific commit from a branch `git cherry-pick <hash>`
- Merge squash `git merge --squash <branch-name>`
- Rebase `git rebase <branch-name>`

---

## Diff and log

- Show difference between two commits `git diff <hash1> <hash2>`
- Show unstaged changes between your index and working directory `git diff --color-words`
- Show difference between working directory and last commit `git diff HEAD`
- Show difference between staged changes and last commit `git diff --cached`
- Show commits that occur between <since> and <until>. Args can be a commit ID, branch name, HEAD, or any other kind of revision reference `git log <since>..<until>`
- Show commits as a graph and limit commit messages to a single line `git log --oneline --graph`
- Search for commits by a particular author `git log --author="<pattern>"`
- Search for commits with a commit message that matches <pattern> `git log --grep=”<pattern>”`

---

## Remote

- Create a remote branch with the same name as the local branch and push changes to the remote branch, used for initial push on empty remote repository `git push -u origin <branch-name>`
- Fetch changes from a remote branch but will not perform a merge on your local branch `git fetch origin <branch-name>`
- Fetch + merge `git pull origin <branch-name>`
- Push changes from a branch `git push origin <branch-name>`
- Push changes from all branches `git push origin --all`
- Remove remote file `git rm <file>`

---

## Stash

### Note

By default, running git stash will stash:

- changes that have been added to your index (staged changes)
- changes made to files that are currently tracked by Git (unstaged changes)

But it will not stash:
By default, git push will not push tags. Tags have to be explicitly passed to git push.

- new files in your working copy that have not yet been staged
- files that have been ignored

Adding the -u option (or --include-untracked) tells git stash to also stash your untracked files.

### Commands

- Clear all `git stash clear`
- Drop `git stash drop <stash-id>`
- List `git stash list`
- Keep stash item and reapplies it to your working copy `git stash apply <stash-id>`
- Removes the latest stash item and reapplies it to your working copy `git stash pop`
- Show `git stash show <stash-id>`
- Stash working directory and staging area `git stash push -u -m <message>`

---

## Tag

### Note

- A best practice is to consider Annotated tags as public, and Lightweight tags as private. Annotated tags store extra meta data such as: the tagger name, email, and date. This is important data for a public release. Lightweight tags are essentially 'bookmarks' to a commit, they are just a name and a pointer to a commit, useful for creating quick links to relevant commits.
- By default, git tag will create a tag on the commit that HEAD is referencing.
- By default, git push will not push tags. Tags have to be explicitly passed to git push.

### Commands

- Checkout `git checkout <tag-name>`
- Create annotated tag for a specific commit `git tag -a <tag-name> <hash>`
- Create lightweight tag for a specific commit `git tag <tag-name> <hash>`
- Delete tag `git tag -d <tag-name>`
- Delete remote tag `git push --delete origin <tag-name>`
- List `git tag --list`
- Push a tag to remote `git push origin <tag-name>`
- Push tags to remote `git push origin --tags`

---

## Undo changes

### Working dir and staging area

- New (untracked) files in working directory `git clean -f`
- Modified files (tracked) in working directory `git restore`
- Modified files (untracked) in staging area `git rm --cached`
- Modified files (tracked) in staging area `git restore --staged`

### Commit history

#### Amend

Best for undoing local private changes

- Amend latest commit `git commit --amend --no-edit`

#### Reset

Best for undoing local private changes

- The default invocation of git reset has implicit arguments of --mixed and HEAD

- Keep changes in working directory and staging area `git reset --soft <hash>`
- Keep changes in working directory `git reset --mixed <hash>`
- Completely destroy any changes `git reset --hard <hash>`

#### Revert

Best for undoing shared public changes

- `git revert <hash>`
