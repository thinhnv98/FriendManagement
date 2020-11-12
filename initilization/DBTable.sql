create table if not exists public.useremails
(
    id int8 not null generated always as identity primary key,
    email varchar(100) not null
);

-- drop table public.useremails

create table if not exists public.friends
(
    id int8 not null generated always as identity primary key,
    firstid int8 not null,
    secondid int8 not null,
    constraint firstemail_fk foreign key (firstid) references public.useremails(id),
    constraint secondemail_fk foreign key (secondid) references public.useremails(id)
);

-- drop table public.friends

create table if not exists public.subscriptions
(
    id int8 not null generated always as identity primary key,
    requestorid int8 not null,
    targetid int8 not null,
    constraint requestid_fk foreign key (requestorid) references public.useremails(id),
    constraint targetid_fk foreign key (targetid) references public.useremails(id)
);

--drop table public.subscriptions

create table if not exists public.blocks
(
    id int8 not null generated always as identity primary key,
    requestorid int8 not null,
    targetid int8 not null,
    constraint requestid_fk foreign key (requestorid) references public.useremails(id),
    constraint targetid_fk foreign key (targetid) references public.useremails(id)
);





