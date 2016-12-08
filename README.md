

# swaggerドキュメント生成
go get github.com/yvasiyarov/swagger
swagger -apiPackage="cat" -mainApiFile="cat.go"

# テスト実行
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
