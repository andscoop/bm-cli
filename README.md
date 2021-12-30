# cbm
A CLI utility for listing and opening chrome bookmarks from the json file.

This utility came out of a desire for more searchable and discoverable bookmarks. By utilizing the path with fuzzy finder, **where** you store your bookmarks can contain contextual information, making them easier to find.


## Examples

List all bookmarks
`./bm-cli ls`

Pipe all bookmarks to `fzf`
`/bm-cli ls | fzf)`

Use `fzf` to find a bookmark and open bookmark in your default browser
```
open "$(./bm-cli find "$(./bm-cli ls | fzf)")"
```

## TODO
- Allow for piping from std in for finding bookmarks
- Find should take many bookmark IDs. This opens possibility for opening folders/batches of bookmarks
- bring in basic config and chrome bookmark fp discovery


