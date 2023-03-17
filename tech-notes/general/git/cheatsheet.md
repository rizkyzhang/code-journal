## Branch

- Checkout `git checkout <branch-name>`
- Create `git branch <branch-name>`
- Current `git branch --show-current`
- Delete `git branch -d <branch-name>`
- List `git branch`
- Rename `git branch -m <new-branch-name>`

### Merge, rebase and squash

- Abort merge `git merge --abort`
- Merge specific commit from a branch `git cherry-pick <hash>`
- Merge squash `git merge --squash <branch-name>`
- Rebase `git rebase <branch-name>`

---

## Diff and log

- `git diff --color-words`
- `git diff <hash1> <hash2>`
- `git log --oneline --graph`

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
- List `git tag --list`
- Pushing tags to remote `git push origin --tags`

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

- `git reset --soft <hash>`
- `git reset --mixed <hash>`
- `git reset --hard <hash>`

#### Revert

Best for undoing shared public changes

- `git revert <hash>`
