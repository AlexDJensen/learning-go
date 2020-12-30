# Alex is learning GO

Directory structure is, while overkill, meant to look like the kubernetes github repository https://github.com/kubernetes/kubernetes.

Only some folders have been included at this point though.

* cmd is main function folder, and contains a subfolder per sub-application (not important now, but might be later).
* internal is for secret code, that shouldn't be shared outside this project.
* pkg is for general purpose and non-secret reusable code. One subfolder per sub-application (as with cmd)

## Notes to self:
---
* Each repo has a git-named `go mod init github.com/AlexDJensen/project` call at the root.
* Create a folder for each thing I want to do (or put them in a cmd folder?, and create a `go mod init <whatever>` for them. Add a replace statement for each package it needs (to point to local))
* run `go mod init <whatever, not github>`