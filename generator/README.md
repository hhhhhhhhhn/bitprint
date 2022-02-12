# How to add a font
This is done through the psf2txt utility (from psftools), so it must be installed.
The fonts are read in psf format, so if you have a different one,
you'll need to convert it.

For example, for bdf files:

```bash
# Example from https://bbs.archlinux.org/viewtopic.php?pid=770649#p770649
./bdf2psf --fb \
	font.bdf \
	/usr/share/bdf2psf/standard.equivalents \
	/usr/share/bdf2psf/ascii.set+/usr/share/bdf2psf/useful.set \
	512 \
	font.psf
```

Once done that, run the program

```bash
go run . font.psf MyCoolFont
```

This will output go code defining the font to stdout.
Output it to a file:

```bash
go run . font.psf MyCoolFont > mycoolfont.go
```
