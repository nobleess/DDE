# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. 

Please note we have a code of conduct, please follow it in all your interactions with the project.

## Merge Request Process

1. Ensure any install or build dependencies are removed before the end of the layer when doing a 
   build.
2. Update the README.md with details of changes to the interface, this includes new environment 
   variables, exposed ports, useful file locations and container parameters.
3. Increase the version numbers in any examples files and the README.md to the new version that this
   Merge Request would represent. The versioning scheme we use is [SemVer](http://semver.org/).
4. You may merge the Merge Request in once you have the sign-off of two other developers, or if you 
   do not have permission to do that, you may request the second reviewer to merge it for you.

## Code of Conduct

### Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

### Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members
* Code style https://google.github.io/styleguide/go/guide 



### Development flow


## Main Branches  
### 1. `main`
- The main branch containing the stable production-ready version.  
- No direct commits are allowed. Changes are introduced via merge requests (MR) from `release` or `hotfix` branches.  

### 2. `develop`
- The main development branch.  
- Contains the latest stable code with new features and fixes undergoing testing.  

---

## Supporting Branches 

### 1. Feature Branches
- Used for developing new features.  
- **Naming convention:** `feature/<Task|Issue>`  
- Created from `develop` and merged back into it.  
- **Example:** `feature/AddDefaultTTL`

### 2. Release Branches
- Used for preparing a new release.  
- **Naming convention:** `release/<version>`  
- Created from `develop` and merged into both `main` and `develop`.  
- **Naming version** must satisfy `^\d+\.\d+\.\d+$`
- **Example:** `release/1.0.0`

### 3. Hotfix Branches
- For urgent fixes in production.  
- **Naming convention:** `hotfix/<Task|Issue>`  
- Created from `main` and merged into both `main` and `develop`.  
- **Example:** `hotfix/MoreLogs`

### 4. Bugfix Branches
- Used for fixing bugs during development.  
- **Naming convention:** `bugfix/<Task|Issue>`  
- Created from `develop` and merged back into it.  
- **Example:** `bugfix/LostKind`

## Commits

### 1. Small better Big
- Used consistent commits
- Describe the goal, not the action.

---

## Merge Request Template
When opening a new Merge Request, please follow this template to ensure clarity and completeness.

### PR Title
- Use a concise and descriptive title.
- Example: `Add user authentication feature`

### Description
- **Which issue does it fix?** 

### Changes Made
- List major changes introduced in this PR.
- Example:
  - Added login and registration forms.
  - Implemented JWT authentication.

### How to Test
- Changes must be covered by tests, previous tests must not fail





### Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

### Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

### Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at myEmail. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

### 