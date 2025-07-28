# folder-organizer v0.9

organize a local folder in subfolders by type, date modified
reads a user config settings.json file to specify operations at run time

v0.9 -> ability to bypass organizeByType through settings.json

v0.8 -> ability to further organize by --time flag, by year or month (specified in settings.json)

```shell
$ folder-organizer --time ./Downloads
Organizing downloads
         2 files => ./Pictures
         3 files => ./Documents
organizing ./Pictures
         2 files => July 2025
organizing ./Documents
         3 files => July 2025
```

v0.1 -> ability to organize by file type, settings.json determines which folder should contain which file type when subfolders by type

```shell
$ folder-organizer ./Downloads
Organizing downloads
         2 files => ./Pictures
         3 files => ./Documents
```

