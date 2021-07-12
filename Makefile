go:
	protoc --go_out=plugins=grpc:. proto/msg.proto
python: 
	python -m grpc_tools.protoc -I./proto --python_out=./python/ --grpc_python_out=./python/ ./proto/msg.proto