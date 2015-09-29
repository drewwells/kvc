This will not work if you have a read/write connection to the bolt DB file already. You must close that before using the viewer.


Usage:

`go get github.com/drewwells/kvc`

List all boltDB buckets
```bash
-> % kvc --db /var/lib/stackengine/storelib.db list
users
products
```

Get key/vals from a bucket. It assumes JSON encoding of the values

```bash
-> % kvc --db /var/lib/stackengine/storelib.db get users

```
