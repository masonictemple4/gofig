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
- [ ] Fix bug with same named panes.
    - If naming windows the same thing across sessions.
- [ ] Fix bug with split panes.
- [ ] Add better start command support for panes and windows.
- [ ] Honestly redo the generation of the objects from the existing tmux session.


### Scratch notes:

To start let's walk through the existing process from start to finish..
1. `$ tmux list-sessions -F "#{session_id}|#{session_name}|#{session_path}"` (might need to add -p when running manually)

    ```zsh
    $0|gofig|/home/mason
    $3|kbd|/home/mason
    ```

    This process will return a list of the sessions on the current running tmux server.
    Then removes trailing newline character, and split into lines each representing a session.
    Iterate over the sessionStrings and call the `sessionFromString()` function.

    Immediately after generating the base session object, we set the session.Windows by calling `session.GetWindows()`
    Which leads us to our second command.

2.  `$ tmux list-windows -t sessionName -F "#{window_id}|#{window_name}|#{window_index}|#{window_height}|#{window_width}|#{window_offset_x}|#{window_offset_y}|#{window_layout}|#{current_pane_path}"`

    This process will return a list of the windows for the target session.
    Removes the trailing newline character, and splits into lines each representing a window in that session.
    Iterate over the windowStrings and call the `windowFromString()` function.
    (NOTE: This is all happening before we reach the `sessions = append(sesssions, session)` line in the previous step.
    
    Immediately after generating the base window object, we set the window.Panes by calling `window.GetPanes()`.
    Which leads us to the last command.

3. `$ tmux list-panes -a -F "#{pane_id}|#{pane_index}|#{pane_title}|#{pane_height}|#{pane_width}|#{pane_current_path}|#{pane_layout}" -f "#{m:window.Name,#{window_name}}"`

    This is potentially problematic, filtering on window name isn't enough because there can be more than 1 window with the
    same name, there can also be windows in other sessions with the same name.

    After calling this command it returns the `PanesFromString()` function output Which is just a slice of panes.



Once the final command to `list-panes` completes technically, we should then resume the `ListSessions()` process,
appending the now "complete" session to the list of sessions and moving onto the next iteration.
