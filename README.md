# Gofig
gofig is a cli tool that parses user defined configuration files to automatically spin up your terminal, tmux, or ssh session just how you want. 

I've had to resetart my terminal or kill my tmux server one too many times after it was configured just how I wanted it
in order to apply changes or updates. Annoyed after having to set everything up by hand again, I want something
I can use to spin up with just a command.


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

**Note: Your config aliases live on your local device, you can use the `$ gofig export db` helper to 
generate a copy of your custom gofig settings.`

Although now that I'm thinking about it, it would be smart to add a shared gofig config file that can
be read from a centralized repository so that your aliases and other settings are consistent on all devices.
Like a control node in Ansible.

For a list of available **Gofig** specific settings see the [section placeholder here]().

After defining your config you can run `$ gofig <conf-name> or alias`

##### Config options

- type: is either terminal, tmux, go (new go project from template)

- For remote sessions give optional flag to specify
which environment to run tmux in (remote or local). Default is remote

**Warning:** If you're using tmux locally instead of on the remote server
you will need a ssh connection for each window/pane/session. 

##### Example config
{
    "type": "development",
    "name": "personal",
    "description": "Default personal development environment with my standard sessions for notes, 
    and work",

}

