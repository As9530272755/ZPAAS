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

alter table webuser add column password varchar(1024) not null default '';

-- 默认密码 123456
pdate webuser set password='$2a$10$XEbCQ7LH4BYRWsDpOpA.k.TQzHnhPLpOOBw8ViofWCao8gVrZBbCG' where id=1;
