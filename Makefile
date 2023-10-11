server:
	go run cmd/main.go

proto.update:
	rm -rf pkg/proto && curl -L -O -H 'Cache-Control: no-cache, no-store' https://github.com/ymHan/fd-proto/archive/refs/heads/master.zip && unzip master.zip && mv fd-proto-master/proto pkg/ && rm -rf fd-proto-master master.zip

proto.user:
	protoc pkg/proto/user*.proto --go_out=./pkg/api_user/pb/ --go-grpc_out=./pkg/api_user/pb/

proto.sports:
	protoc pkg/proto/sports*.proto --go_out=./pkg/api_sports/pb/ --go-grpc_out=./pkg/api_sports/pb/

proto.ims:
	protoc pkg/proto/ims*.proto --go_out=./pkg/api_ims/pb/ --go-grpc_out=./pkg/api_ims/pb/

proto.cms:
	protoc pkg/proto/cms*.proto --go_out=./pkg/api_cms/pb/ --go-grpc_out=./pkg/api_cms/pb/

proto:
	make proto.user
	# make proto.user && make proto.sports && make proto.ims && make proto.cms
