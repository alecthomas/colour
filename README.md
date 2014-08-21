Package colour provides [Quake-style colour formatting][1] for Unix terminals.

eg.

    Printf("^0black ^1red ^2green ^3yellow ^4blue ^5magenta ^6cyan ^7white^R\n")

The following sequences are converted to their equivalent ANSI colours:

| Sequence | Effect |
| -------- | ------ |
| ^0 | Black |
| ^1 | Red |
| ^2 | Green |
| ^3 | Yellow |
| ^4 | Blue |
| ^5 | Cyan (light blue) |
| ^6 | Magenta (purple) |
| ^7 | White |
| ^8 | Black Background |
| ^9 | Red Background |
| ^a | Green Background |
| ^b | Yellow Background |
| ^c | Blue Background |
| ^d | Cyan (light blue) Background |
| ^e | Magenta (purple) Background |
| ^f | White Background |
| ^R | Reset |
| ^U | Underline |
| ^B | Bold |

[1]: http://www.holysh1t.net/quake-live-colors-nickname/
