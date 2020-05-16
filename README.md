<h1 align="center">
  <br>
  <img src="https://olivia-ai.org/img/icons/olivia-with-text.png" alt="Olivia's character" width="300">
  <br>
</h1>

<h4 align="center">📦 The Olivia Kit for Golang</h4>

<p align="center">
  <a href="https://github.com/olivia-ai/olivia-kit-go/actions?query=workflow%3A%22Format+checker%22"><img src="https://github.com/olivia-ai/olivia-kit-go/workflows/Format%20checker/badge.svg"></a>
  <a href="https://pkg.go.dev/github.com/olivia-ai/olivia-kit-go"><img src="https://godoc.org/github.com/olivia-ai/olivia?status.svg" alt="GoDoc"></a>
</p>

## How to use it
Here is a code example to see how to use the Go kit.
```golang
var information map[string]interface{}
client, err := NewClient("localhost:8080", true, &information)
if err != nil {
	panic(err)
}

defer client.Close()

client.SendMessage("Hello Olivia!")
```

<p align="center">
  <img width="60" src="https://olivia-ai.org/img/icons/olivia.png">
<p>

<p align="center">
  Made with ❤️ by <a href="https://github.com/hugolgst">Hugo Lageneste</a>
</p>

![Olivia's wave](https://olivia-ai.org/img/background-olivia.png)
