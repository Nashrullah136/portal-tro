create table audits
(
    date_time   datetime default getdate() not null,
    id          bigint identity
        constraint audits_pk
        primary key,
    username    varchar(32)                not null,
    object      varchar(32),
    object_id   varchar(50),
    action      varchar(16),
    data_before nvarchar(512),
    data_after  nvarchar(512)
)
    go

create table roles
(
    id        bigint identity
        constraint roles_pk
        primary key,
    role_name varchar(32) not null
)
    go

create table users
(
    id         bigint identity
        constraint users_pk
        primary key,
    username   varchar(32)                not null
        constraint username_unique
            unique,
    password   varchar(255)               not null,
    role_id    bigint                     not null
        constraint role_id_users
            references roles,
    created_at datetime default getdate() not null,
    created_by varchar(32)                not null,
    updated_at datetime default getdate() not null,
    updated_by varchar(32)                not null
)
    go

