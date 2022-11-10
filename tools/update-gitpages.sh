#!/bin/bash

cd ../

git branch -d gh-pages
git subtree push --prefix dist origin gh-pages