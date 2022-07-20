# Contributing

## Commit Process

This repo comes with pre-commit hooks. These should be installed before commiting, done with:
```bash
pre-commit install --hook-type pre-commit --hook-type pre-push
```
This will lint and run Go tools on commit, and run unit tests when pushing.
