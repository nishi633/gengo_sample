
```
go get golang.org/x/tools/cmd/goyacc
go install golang.org/x/tools/cmd/goyacc
```

y.goというソースコードを生成する。  
```
goyacc parser.go
```

実行例
```
go run y.go "1 + 2"
```

## 参考
- https://news.mynavi.jp/article/gogogo-10/
- https://qiita.com/k0kubun/items/1b641dfd186fe46feb65
- http://loveruby.net/ja/rhg/book/yacc.html
- https://hake.hatenablog.com/entry/20150923/p1
