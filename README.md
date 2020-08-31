# Cliconfig
> Cliconfig is a utility package that helps setup and manage cli configurations

``` golang
cliconfig.New("path/to/config/folder")
cliconfig.Init(configStruct) // Creates the yaml config.yml file inside the config folder
cliconfig.FromFile(configStruct) // Updates the configStruct from the config.yml file
```
