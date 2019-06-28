export GO111MODULE=on
go mod init
go get k8s.io/client-go@v12.0.0+
go get k8s.io/client-go@v11.0.0
go get k8s.io/api@kubernetes-1.14.0
go get k8s.io/apimachinery@kubernetes-1.14.0