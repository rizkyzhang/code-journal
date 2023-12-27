1. Perform login with the heroku login command. You will be prompted with your default browser window for completing the login at the Heroku portal.

2. Assuming you already have your app set up at Heroku, you just need to add a new remote for your Git repository with Heroku CLI. Run heroku `git:remote -a example-app`, substitute "example-app" with your app name.

3. git remote -v to check if the remote has been set successfully. You should see something like this appear as a response:

```
heroku  https://git.heroku.com/your-app-name.git (fetch)
heroku  https://git.heroku.com/your-app-name.git (push)
```

4. Push your branch to the new heroku remote
   `git push heroku main` / `git push heroku branch_name:main`

Reference: https://stackoverflow.com/questions/71892543/heroku-and-github-items-could-not-be-retrieved-internal-server-error
