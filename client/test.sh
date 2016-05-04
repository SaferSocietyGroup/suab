DIRNAME=`dirname $0`
WD=`readlink -nf $DIRNAME`
export GOPATH=$WD 

go test suab
go test config
go test submitters
go test shutupflags
