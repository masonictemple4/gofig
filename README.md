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
- [ ] Refactor the Load layout function in tmux.
- [ ] Add an overall project config to house the tmux sessions as children so we can do it by project.
- [ ] Round out tests.
- [ ] Overall error handling..
- [ ] Make sure to allow loading new sessions from within tmux.
