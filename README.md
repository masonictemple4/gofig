# Gofig
Gofig is a CLI tool built with go to quickly manage tmux sessions, and session configs.


[CLI Docs](./docs/gofig.md)

### Install
`$ go install github.com/masonictemple4/gofig@latest`

### Use
With your tmux session open run the following:

`$ gofig export layoutname.json` (you can also use .yaml)

To load your config, make sure tmux is not open:

`$ gofig load layoutname.yaml`

### TODOs
- [ ] Still a little finicky with layout, need to test and fix that on both single and mutlipae windows..
- [ ] Add an overall project config to house the tmux sessions as children so we can do it by project.
- [ ] Round out tests.
- [ ] Overall error handling..
- [ ] Make sure to allow loading new sessions from within tmux.
- [X] Fix bug with same named panes.
    - If naming windows the same thing across sessions.
- [X] Fix bug with split panes.
- [X] Add better start command support for panes and windows.
- [X] Honestly redo the generation of the objects from the existing tmux session.
- [ ] Implement a better verbosity mode.
- [ ] See if we can make the default window the first for all sessions, previously
solved by calling the attach-session command with the window id. However, 
the following sessions are all currently attached to their last window.
- [ ] Add update functionality, currently if you export a layout to an existing path that
exists you will replace that file.
    - [ ] At the very least add a safety check to confirm whether or not they would like to
    overwrite that file.
    - [ ] Add an update feature to take current export and apply changes to existing layout
    file. (Not sure if this is totally necessary, one could just load their existing layout 
    and if they make additions they would like to save for next time just re-export it.)
