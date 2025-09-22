Simple scripts for simple go repositories.

## GGIT IS NOT MEANT FOR COLLABORATIVE PROJECTS

I've noticed i spend a lot of time doing the same
commands for creating repositories and uploading changes.
GGIT aims to simplify the process.

Available commands:
- ggit init :: create go repository
```bash
ggit init <optional:project>
# mkdir -p <project>
# go mod init github.com/<user>/<project>
# git init
# git remote add origin git@github.com:<user>/<project>
```

- ggit :: upload changes
```bash
ggit <optional:message>
# git add --all
# git commit -m <message>
# git push --set-upstream origin master
```

- ggit redo (or ggit fuck) :: amend last commit
```bash
ggit redo
# git add .
# git commit --amend --no-edit
# git push --force
```

### Installation
run the makefile to install into your /usr/bin
```bash
doas make install
```
