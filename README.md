# jwt-golang

Semplice progetto con l'implementazione in [GO](http://www.golang.org) dei Web Token in JSON ([JWT](http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html), partendo dal modulo [jwt-go](https://github.com/dgrijalva/jwt-go).

**TODO LIST:** Permettere la customizzazione di Claims e Metodi di Signing. La versione attuale supporta solo l'algoritmo di firma: SigningMethodHS256.

## How to use it

```golang
import "gitlab.com/fabiosebastiano/jwt-golang"
```
