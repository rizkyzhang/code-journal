## Problem

ESLint Definition for rule 'import/extensions' was not found

## Solution

You missed adding this in your eslint.json file.

`"plugins": ["import"]`

or

`"extends": ["plugin:import/recommended"]`

## Reference

https://stackoverflow.com/questions/68878189/eslint-definition-for-rule-import-extensions-was-not-found
