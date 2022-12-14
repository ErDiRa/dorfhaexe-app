#!/bin/bash

cd ../

yarn build

git add .

git commit -m "New build"

git push

git subtree push --prefix dist origin gh-pages
