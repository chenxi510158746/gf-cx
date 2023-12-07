# 1. gf-cx

A `API` project using `GoFrame`.

# 2. Functions

1. `User`  : `register, login, logout, info`
2. `order` : `create, info, list(page)`
3. `goods` : `create, update, delete, info, list(page)`

# 3. Quick start

1. Download the source code
```
git clone https://github.com/chenxi510158746/gf-cx.git
cd gf-cx
```

2. Update config manifest/config/config.yaml and edit database.default config
```
vim manifest/config/config.yaml
```

3. Import Db
   Import manifest/document/df-cx.sql to your Mysql

4. Run gf-cx And Enjoy
```
gf run main.go
```

then open http://127.0.0.1:8000/ and enjoy it.
```
Note: It is recommended to use Postman for testing API interfaces
```