# Client-side cache busting

Web browser will often cache website resources like html, css, js, images and others in order to save time. When a request of web resources to the server happen, browser will check whether the files is cached or not, if it is, browser will use the cached files instead.  
This could be a problem if the files is updated on the server, but the browser still used the cached files. Examples:

1. JavaScript files is updated on the server to fix some bugs. After pushing the changes to production. Some users still complain that the bug is not fixed althought you have already test it. Turns out, some of the user's browser still use the old buggy cached JavaScript files.
2. Browser cache json file from an API that updates frequently.

Easiest solution is to append timestamp query parameter to endpoint to trick browser as if it is a new request.
