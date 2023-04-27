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


## How does it work?

### Releasing to staging

Releasing a brand new app version to staging can be done through the GitHub
Action (GHA):

![image](https://user-images.githubusercontent.com/10437518/227167447-6c341914-0811-43d0-b3f3-0e0cd78cea6b.png)

This will:
1. Always bump the chart version
2. Update the app version if provided
3. Create a pull request with the changes

Keep in mind that if you want to ship further changes to the chart, the
recommended approach is to add more commits to the pull request generated
by the GHA. This way reverting the chart is a matter of reverting a single
commit.

### Promoting to production

As for promoting, it's even simpler, just provide the service name to the
GHA action for promotions and it'll literally pick what's running in stage
and set up a PR so the same is shipped to production.

![image](https://user-images.githubusercontent.com/10437518/227168949-b46cbfba-2bf3-4da2-a28f-608afb5309e4.png)

