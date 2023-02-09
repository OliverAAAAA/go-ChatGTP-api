create table if not exists `go-chatgpt`.ask_log
(
    id         int auto_increment
    primary key,
    user_id    int default 0 not null,
    method     varchar(30)   null,
    request_ip varchar(30)   null,
    address    varchar(30)   null,
    request    longtext      not null,
    data       longtext      not null,
    createTime datetime      not null
    );

