# Gofig
gofig is a cli tool that parses user defined configuration files to automatically spin up your terminal, tmux, or ssh session just how you want. 

I've had to resetart my terminal or kill my tmux server one too many times after it was configured just how I wanted it
in order to apply changes or updates. Annoyed after having to set everything up by hand again, I want something
I can use to spin up with just a command.

#### TODOS:
Simplfiy instead of using this to start the ssh session etc.. just use it on the remote server.
Limit scope to tmux related environments.


Gofig also wraps the new dev tool `gonew` to generage new go projects from local or remote templates.

### What is a config?
A config represents a specific tmux/ssh/terminal/template environment. You can choose to define them
however you'd like, whether it's per project, work vs personal, language, remote vs local etc.. All you need to do
is define your desired configuration for gofig to find and parse.

### Defining Configs
Configs can be created with either json or yaml, and you can set a default location (local or remote) for gofind
to search for configs. 

Configs can be identified by their name or an alias you define. By default the Name is used, however if there is ever
a need to abbreviate/alias a config's name you can use the command to set it. 

##### Config Objects
- layout


- type: is either terminal, tmux, go (new go project from template)

- For remote sessions give optional flag to specify
which environment to run tmux in (remote or local). Default is remote

**Warning:** If you're using tmux locally instead of on the remote server
you will need a ssh connection for each window/pane/session. 

##### Example Config

{
    "type": "tmux",
    "name": "default",
    "alias": "d"
    "description": "Default personal development environment with my standard sessions for notes, 
    and work",
    "sessions": [
        {
    ],

}

##### Layouts
- singlepane: Default layout. No splits.
- singlehsplit: 2 horizontal rows in a single window.
- singlevsplit: 2 vertical colums in a single window.
- quadrant: 2x2 grid in a single window.
- quadhsplit: 4 horizontal rows inside a single window.
- quadvsplit: 4 veritical columns inside of a single window.


Layouts are temporarily limited to the basic setups listed above.
These should cover most common use cases, for more complex layouts
you can either manually configure your session or use one of the following
workarounds:

 - You can write setup script(s) to run as a startcommand at either the session, window,
   or window pane level.

 - You can also specify [tmux cli](https://github.com/tmux/tmux/wiki/Getting-Started) commands to run as well.

 - Or if you want to make things difficult you can define your session(s), session window(s),
   and window pane(s) sequentially in your **gofig** json or yaml files. The order in
   which you do things will matter as it would when manually configuring your layout
   with keybindings. 

A much more fitting solution is on the way, so even the most heinous layouts
can be defined with ease.
