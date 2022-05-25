## Problem

React onClick not working in any of my browsers

## Causes

HtmlWebpackPlugin, HtmlWebpack Plugin is adding additional script tags of bundle in head tag of index.html. So if you already define a script tag in index.html, HtmlWebpackPlugin will replace it with bundled script tags when webpack runs.

## Solutions

1. Delete defined script tags, let HtmlWebpackPlugin inject the bundled script tags automatically.
2. Configure htmlwebpackplugin so that it will not inject any additional script tags.

```js
...
    new HtmlWebpackPlugin({
      name: "index.html",
      inject: false,
      template: path.resolve(__dirname, "public/index.html"),
    }),
...
```

## Reference

https://stackoverflow.com/questions/65770449/react-onclick-not-working-in-any-of-my-browsers-but-for-colleagues-it-does
