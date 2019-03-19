# building-blocks-app
App created from building blocks

## How to include a building block
1. Clone to a directory:
e.g. `git clone git@github.com:ggerritsen/postgresql-block repository`
2. Remove the .git directory and .gitignore file:
`rm -rf .git && rm .gitignore`
3. Update the package name of the building block files to reflect the new path
e.g. `sed -i '.bak' 's/package main/package <INSERT NEW NAME>/g' $(git grep package | cut -f1 -d':')`
4. Use the building block's code in the app and edit it where necessary
