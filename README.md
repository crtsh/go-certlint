# go-certlint
Go wrapper for [awslabs/certlint](https://github.com/awslabs/certlint)

# INSTALL
```
go get -d github.com/crtsh/go-certlint
cd $GOPATH/src/github.com/crtsh/go-certlint
sed -i "s/ruby-?.?/ruby-X.Y/g" go-certlint.go
make
go install
```
Change `ruby-X.Y` to your Ruby version (e.g., `ruby-2.3`).
