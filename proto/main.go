package main

// -I --proto_path: 指定proto文件所在目录(--go_out第二个参数会使用该路径)
//go:generate /home/t04934/protobuf/protoc --proto_path=./../resource/proto --plugin=/home/t04934/go/bin/protoc-gen-go --go_out=./ proto_test.proto
func main() {
}
