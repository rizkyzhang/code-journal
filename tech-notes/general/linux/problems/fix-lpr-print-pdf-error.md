---
tags:
  - linux-problem
---
## Problem

`lpr` command can't print some images possibly because of missing orientation `exif`

## Possible solutions

- Add rotation `exif` with `exiftool -Orientation=0 -n img.jpg`
- Add auto rotation `exif` with `exiftran -a img.jpg -o img2.jpg`
