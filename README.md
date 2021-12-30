# cbm
A CLI utility for listing and opening chrome bookmarks from the json file.

This utility came out of a desire for more searchable and discoverable bookmarks. By utilizing the path with fuzzy finder, **where** you store your bookmarks can contain contextual information, making them easier to find.


## Examples

List all bookmarks
`./cbm ls`

Pipe all bookmarks to `fzf`
`/cbm ls | fzf)`

Use `fzf` to find a bookmark and open bookmark in your default browser
```
./cbm ls | fzf | ./cbm open
```

## TODO
- bring in basic config and chrome bookmark fp discovery
- `cbm open` should probably carry most of the water here
    - `cbm open --fzf` should run eqivilant of `cbm ls | fzf | cbm open`
    - any reason to not make this the default behavior of `cbm` with no args?
- allow for customization of what is output in `cbm ls` this will allow for users to customize for their searching preferences


