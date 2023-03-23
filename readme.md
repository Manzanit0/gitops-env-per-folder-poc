# GitOps POC with environments per directory 

Showcases 3 things:

* Organising application charts across different directories per environment
* A workflow to release a new changeset to _staging_
* A workflow to promote a change from _staging_ to _production_


## Benefits

Main benefits of this approach are:
* The repository's default branch, i.e. `master`, is the source of truth for
  what's running where.
* Promotion of changes from one env to another are dead simple: copy the
  content of one chart to the other.
* Reverts are simply a `git revert`

## Disadvantages

Some of the drawbacks are:
* Charts can potentially be modified manually thus provoking a situation where
  any given version of a chart has different content in different environments.
