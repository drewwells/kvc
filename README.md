This will not work if you have a read/write connection to the bolt DB file already. You must close that before using the viewer.


Usage:

`go get github.com/drewwells/kvc`

List all boltDB buckets
```bash
-> % kvc --db /path/to/storelib.db list
users
products
```

Get key/vals from a bucket. It assumes JSON encoding of the values, but will dump the raw string if unmarshalling fails.

```bash
-> % kvc --db /var/lib/stackengine/storelib.db get users
|  ------------ Key  -----------  | ---------------- Value ---------------------
                              100 : {
                                      "first_name": "first",
                                      "last_name": "last",
                                      "password": "cLfgjFzg3H1H.Jm6oUiWtWmSi4Rf",
                                      "token": "ajIln2jxk",
                                      "username": "admin"
                                      }

```
