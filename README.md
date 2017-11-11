ehinfo
===

Handy config storing program to read data from stdin and store it in a json
file to be read back and/or amended.

eg:
```sh
UUID=unique
$ echo "arse face" | ehinfo set server $UUID
$ ehinfo info server $UUID
{"arse":"face"}
$ echo "less arse" | ehinfo set server $UUID
$ ehinfo info server $UUID
{"arse":"face","less":"arse"}
```
