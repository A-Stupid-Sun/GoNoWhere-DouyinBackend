-- auto-generated definition
create table users
(
    id             bigint unsigned auto_increment comment '自增主键'
        primary key,
    create_at      timestamp default CURRENT_TIMESTAMP not null,
    deleted_at     timestamp                           null,
    name           varchar(34)                         not null,
    follow_count   bigint    default 0                 not null,
    follower_count bigint    default 0                 not null,
    user_id        bigint                              not null,
constraint idx_user_id
        unique (user_id),
        primary key (id)

);


