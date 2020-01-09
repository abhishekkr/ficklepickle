## ficklepickle

> it's a simple struct pickling library with compression and encryption

* [sample usage](_tests_behavioral_/pickler.go)

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


#### Persist in local file

* Pickle

```
err := ficklepickle.Write("vanilla-file", "metric_xload", xload)
if err != nil {
  panic("Pickle error for Metric")
}
```

* UnPickle

```
if err := ficklepickle.Read("vanilla-file", "metric_xload", &yload); err != nil {
  panic("Unpickle error for Metric")
}
fmt.Println(yload.Name) // shall output 'somework'
```

---

![FicklePickle](docs/ficklepickle.png) | [![Go Report Card](https://goreportcard.com/badge/github.com/abhishekkr/ficklepickle)](https://goreportcard.com/report/github.com/abhishekkr/ficklepickle)

---

### ToDo

* Read/Write to support RPC like server/client model

* Read/Write to support embedded KeyVal store

---
