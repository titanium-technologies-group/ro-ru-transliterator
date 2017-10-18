Ro-ru transliterator
--------------------
Library is used to transliterate from romanian to russian.
Installation:
```
import (
    "github.com/titanium-codes/ro-ru-transliterator/transliteration"
)
```

Simple usage:
```
russianString := transliteration.TransliterateInRussian(romanianString)
```

Some limitations:
* May have unexpected behaviour for not not existing words
* Makes all text lowercase