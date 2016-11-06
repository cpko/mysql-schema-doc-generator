# doc generator for mysql schema

## Requirements
* Go 1.7 or higher

## Installation
go get -u github.com/cpko/mysql-schema-doc-generator

## Usage
mysql-schema-doc-generator [-host 127.0.0.1] [-port 3306] [-u root] [-p 123456 ] -d my_database

* -host host,default "127.0.0.1"
* -port port,default 3306
* -u username,default "root"
* -p password,default ""
* -d database name,required

then a file named ```schema_doc.md``` will be generated in your current path,
which contains structure description of all tables from your database,and sorted by table name,
like this:

## addresses 表

|  字段  |   说明   |   类型   |  长度  |  是否可空  |  是否索引  |   备注   |
|:-----:|:-------:|:--------:|:------:|:--------:|:--------:|:--------:|
| id | 主键 | bigint |  -  | NO | YES |  |
| pid | 父级ID | bigint |  -  | YES | NO |  |
| name | 名称 | varchar |  50  | YES | NO |  |
| level | 级别 | int |  -  | YES | NO |  |

## users 表

|  字段  |   说明   |   类型   |  长度  |  是否可空  |  是否索引  |   备注   |
|:-----:|:-------:|:--------:|:------:|:--------:|:--------:|:--------:|
| id | 主键 | bigint |  -  | NO | YES |  |
| name | 用户名 | varchar |  40  | YES | YES |  |
| email | 邮件 | varchar |  32  | YES | YES |  |
| mobile | 手机号码 | varchar |  16  | YES | YES |  |
| password | 登录密码 | varchar |  32  | YES | NO |  |
