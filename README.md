# gin-admin

## 安装
```
git clone https://github.com/qiuapeng921/gin-template-admin
```

### 必须

- Mysql
- Redis

### 运行
```
cd $GOPATH/src/gin-admin
cp .env.example .env

###修改.env配置
DB_HOST=127.0.0.1:3306
DB_DATABASE=test
DB_USERNAME=root
DB_PASSWORD=
DB_CHARSET=utf8mb4
DB_MAX_IDLE=20
DB_MAX_OPEN=100

REDIS_HOST=127.0.0.1:6379
REDIS_PASSWORD=
REDIS_MAX_IDLE=50
REDIS_MAX_ACTIVE=1200

go run main.go
```

## 特性
- Gin
- Xorm
- Redis
- Templates
