
.PHONY: grpc_inter
grpc_inter:
	@echo "Generating Go files for ml_server"
	cd grpc_inter && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

.PHONY: ml_data_collector
ml_data_collector:
	@echo "Generating Go files for ml_data_collector"
	cd ml-data-collector && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

.PHONY: ml_server
ml_server:
	@echo "Generating Go files for ml_data_collector"
	cd ml-server && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto


.PHONY: protoc
protoc: ml_data_collector ml_server


.PHONY: vupdate
vupdate:
	go mod tidy
	go mod vendor
