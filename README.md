# 留影叙佳期-生日贺卡自动获取

> 支持自动续签api的token，支持获取生日贺图。

## 使用

### 下载可执行文件

前往[Releases](https://github.com/rroy233/genshinBirthdayHelper/releases)下载可执行文件

### 编辑配置文件

在可执行文件同目录下，新建`config.json`

```json
{
  "accounts": [
    {
      "server": "cn_gf01",
      "uid": "114514",
      "mys-cookie": ""
    }
  ]
}
```

如果有多个账号可以这样写

```json
{
  "accounts": [
    {
      "server": "cn_gf01",
      "uid": "114514",
      "mys-cookie": ""
    },
    {
      "server": "cn_gf01",
      "uid": "1145141",
      "mys-cookie": ""
    }
  ]
}
```

保存文件

#### 说明

server: 服务器代号（天空岛cn_gf01，世界树cn_qd01）

uid: 游戏内UID

mys-cookie: 米游社cookie（可从米游社网页版获取）



### 运行使用

自行设置定时任务，每天晚上0点过后执行一次可执行文件即可。

Mac or Linux

```shell
./文件名称
```

Windows

```
双击可执行文件
```




## 自行编译

### 环境依赖

* go 1.18

### 编译

```shell
go build
```

或

```shell
# 交叉编译
make
```

## License

MIT License.


