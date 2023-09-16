# CLI for Countdown Solver

```
usage: ./countdown [options] <letters>
options:
  -d, --dict string      Location of dictionary file (default "words.txt")
  -h, --help             Prints this message
  -s, --max-size int     Maximum size of word list
  -l, --min-length int   Minimum length answers need to be (default 1)

examples:
  $ ./countdown 'sraneimax'
  $ ./countdown --max-size 10 'mhseoudma'
  $ ./countdown --dict /usr/share/dict/words --min-length 4 'eoiadnnzt'
```

Includes a sorted list of the top 1000 words in English in `words.txt`.
