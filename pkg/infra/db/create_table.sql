-- create campaign table--
create table if not exists campaign
(
    campaign_id    serial,
    user_id        integer,
    name           varchar(50) not null,
    "desc"         varchar,
    short_desc     varchar,
    perks          varchar,
    backer_count   integer,
    goal_amount    integer,
    current_amount integer,
    slug           varchar,
    primary key (campaign_id),
    foreign key (user_id) references "user"
);

alter table campaign
    owner to assignment;


-- create campaign image table--
create table if not exists campaign_image
(
    campaign_image_id integer default nextval('campaign_image_campaign_image_id_seq'::regclass) not null,
    campaign_id       integer,
    file_name         varchar,
    primary key (campaign_image_id),
    foreign key (campaign_id) references campaign
);

alter table campaign_image
    owner to assignment;


-- create campaign table--
create table if not exists transaction
(
    trx_id      integer default nextval('transaction_trx_id_seq'::regclass) not null,
    campaign_id integer,
    user_id     integer,
    amount      integer,
    status      varchar,
    code        varchar,
    payment_url varchar,
    primary key (trx_id),
    foreign key (campaign_id) references campaign,
    foreign key (user_id) references "user"
);

alter table transaction
    owner to assignment;


-- create user table --
create table if not exists "user"
(
    user_id          integer default nextval('user_user_id_seq'::regclass) not null,
    name             varchar(50),
    occupation       varchar(50),
    email            varchar(50),
    password_hash    varchar(50),
    avatar_file_name varchar(225),
    role             varchar(10),
    primary key (user_id)
);

alter table "user"
    owner to assignment;

