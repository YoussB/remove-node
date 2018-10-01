# 1. remove-node
GO Playground - simple code to remove all node_modules & bower_components folders from a given directory.


## Usage: 

inside code directory ```go install```

run code from `GOBIN` folder with folder name as argument.


License Apache v2.0


 # 2. Amend for multiple commits
 
`git filter-branch -env-filter "$( git duet )" HEAD~2..HEAD`

This will run `git duet` on `HEAD~2` up to `HEAD`
