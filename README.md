Simple scripts for simple go repositories.

## GGIT IS NOT MEANT FOR PUBLIC USE

I've noticed i spend a lot of time doing the same
commands for creating repositories and uploading changes.
This can save some time on solo projects.

Available commands:
- ggit init (create repository)
```bash
ggit init <optional:project>
# mkdir -p <project>
# go mod init github.com/<user>/<project>
# git init
# git remote add origin git@github.com:<user>/<project>
```

- ggit (upload changes)
```bash
ggit <optional:message>
# git add --all
# git commit -m <message>
# git push --set-upstream origin master
```
