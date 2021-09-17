-- 当 user 表不存在的时候创建
create database if not exists hellodb; default charset utf8mb4;

-- 创建 user 表，并且创建出所有字段所有的属性
create table if not exists webuser (
    id bigint primary key auto_increment,
    name varchar(32) not null default '',
    age varchar(20) not null default '',
    sex boolean not null default true,
    addr text not null default '',
    created_at datetime not null ,
    updated_at datetime not null ,
    deleted_at datetime
) engine=innodb default charset utf8mb4;

alter table WEBuser add column if not exists password varchar(1024) not null default '';
