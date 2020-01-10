## ficklepickle

> it's a simple struct pickling library with compression and encryption
>
> read/write/delete available over file, db and tcp server

* [sample usage](_tests_behavioral_/pickler.go); [sample rpc server](_tests_behavioral_/fickle_server.go)

![FicklePickle](docs/ficklepickle.png) | [![Go Report Card](https://goreportcard.com/badge/github.com/abhishekkr/ficklepickle)](https://goreportcard.com/report/github.com/abhishekkr/ficklepickle)

---

### Usage

* to fetch this go package `go get -u github.com/abhishekkr/ficklepickle`

* import required to use `import "github.com/abhishekkr/ficklepickle"`

* say we have a structure

```
type Metric struct {
  Name string
  Percent int
}

xload = Metric{Name: "somework", Percent: 25}
yload := Metric{}
```

#### In memory

* Pickle

```
pickle, err := ficklepickle.Pickle(xload)
if err != nil {
  panic("Pickle error for Metric")
}
// pickled data is 'pickle' as byte array
```

* UnPickle

```
if err := ficklepickle.Unpickle(pickle, &yload); err != nil {
  panic("Unpickle error for Metric")
}
fmt.Println(yload.Name) // shall output 'somework'
```


#### Persist using a mode (local file, database; remote TCP service)

```
// for normal local file
mode = ficklepickle.RwFile

// for github.com/abhishekkr/gol/golkeyval supported database
mode = ficklepickle.RwDb // database type default: leveldb, configurable

// for remote TCP server hosted pickles, can be used to share across services
// sample TCP service can be created similar to example code at: _tests_behavioral_/fickle_server.go
mode = ficklepickle.RwRpc // remote database type default: leveldb, configurable in Rpc Server
```

* Pickle

```
err := ficklepickle.Write(mode, "metric_xload", xload)
if err != nil {
  panic("Pickle error for Metric")
}
```

* UnPickle

```
if err := ficklepickle.Read(mode, "metric_xload", &yload); err != nil {
  panic("Unpickle error for Metric")
}
fmt.Println(yload.Name) // shall output 'somework'
```

* Delete Pickle

```
if err := ficklepickle.Delete(mode, "metric_xload"); err != nil {
  panic("Delete pickle error for Metric")
}

```

---
