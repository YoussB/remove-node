# 1. remove-node aka `main.go`
GO Playground - simple code to remove all node_modules & bower_components folders from a given directory.


## Usage: 

### TODO: update to run with go modules.
inside code directory ```go install```

run code from `GOBIN` folder with folder name as argument.


License Apache v2.0


 # 2. Forever test runner aka `testerr`
 
 use as `testerr path_to_watch command_to_run` 
 
 This little tool can be used for running unit tests in a go environment 
 
 Currently only supporting MacOs

 # 3. Amend for multiple commits
 
`git filter-branch -env-filter "$( git duet )" HEAD~2..HEAD`

This will run `git duet` on `HEAD~2` up to `HEAD`
