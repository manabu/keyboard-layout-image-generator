# keyboard-layout-image-generator
Generate Keyboard Image PDF file
 from source code `keymap.c`
 which is for [qmk/qmk_firmware: keyboard controller firmware for Atmel AVR USB family](https://github.com/qmk/qmk_firmware)

Currently design for only ErgoDox


# Background

* [ErgoDox EZ Configurator](http://configure.ergodox-ez.com/keyboard_layouts/new)

* [Keyboard Layout Editor](http://www.keyboard-layout-editor.com/)

These are nice editor and very good looking easy to see my keyboard layouts.

But when I create custom one with:

* [qmk/qmk_firmware: keyboard controller firmware for Atmel AVR USB family](https://github.com/qmk/qmk_firmware/)

No image to keymap.

After the change of keymap with source code `keymap.c` .

No image of keyboard layouts .

This application generate PDF image from `keymap.c` .

# Goal

Goal is no need to maintain double.

# Download

* [Releases · manabu/keyboard-layout-image-generator](https://github.com/manabu/keyboard-layout-image-generator/releases)

# How to build

```
peg keymap_image.peg
```

# How to run

```
go run keymap_image.go keymap_image.peg.go keymap.c > layout.pdf
```

## Get keymap

If you don't have own image.

```
curl -o keymap.c https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/ergodox/keymaps/default/keymap.c
```

# How it works

Using PEG for handling `keymap.c`

Firstly I try to use [peg/c.peg at master · pointlander/peg](https://github.com/pointlander/peg/blob/master/grammars/c/c.peg)
but it can't parse well.

So I create PEG file for very simple and adapt for ErgoDox `keymap.c` .

# TODO

* [ ] Single Binary
* [ ] Long press key
* [ ] Key combinations
* [ ] Color
* [ ] Rotate some keys label
* [ ] More flexible mapping to keys
* [ ] Output JSON for [Keyboard Layout Editor](http://www.keyboard-layout-editor.com/)
* [ ] If it can , more easy to see stable version and current develop version.  

# Reference

* [Go言語 - PEGで構文解析 - 字句解析 - はけの徒然日記](http://d.hatena.ne.jp/hake/20151004/p1)

* [GolangとPEGで作る言語処理系 vol.1 - Qiita](http://qiita.com/erukiti/items/9e9cada94178ed10a1fa)

* [pointlander/peg: Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.](https://github.com/pointlander/peg)

* [peg/c.peg at master · pointlander/peg](https://github.com/pointlander/peg/blob/master/grammars/c/c.peg)

* [Parsing Expression Grammar - Wikipedia](https://ja.wikipedia.org/wiki/Parsing_Expression_Grammar)
