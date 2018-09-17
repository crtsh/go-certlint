# go-certlint
Go wrapper for [https://github.com/awslabs/certlint](awslabs/certlint)

# INSTALL
```
go get -d github.com/crtsh/go-certlint
cd $GOPATH/src/github.com/crtsh/go-certlint
sed -i "s/ruby-?.?/ruby-X.Y/g" go-certlint.go
make
go install
```
...where X.Y is your Ruby version (e.g., 2.3).
