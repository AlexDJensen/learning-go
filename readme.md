# Alex is learning GO

Directory structure is, while overkill, meant to look like the kubernetes github repository https://github.com/kubernetes/kubernetes.

Only some folders have been included at this point though.

* cmd is main function folder, and contains a subfolder per sub-application (not important now, but might be later).
* internal is for secret code, that shouldn't be shared outside this project.
* pkg is for general purpose and non-secret reusable code. One subfolder per sub-application (as with cmd)