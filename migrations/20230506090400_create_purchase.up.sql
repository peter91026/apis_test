CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table purchase(
    purchase_id uuid default uuid_generate_v4() not null serial primary key,
    applicant varchar not null,
    company_name varchar not null,
    department varchar not null,
    product_name varchar not null,
    price int not null,
    purchase_quantity int not null,
    created_by varchar not  null ,
    created_at timestamp default now() not null
)

create unique index purchase_purchase_id_uindex
    on purchase (purchase_id);

create or replace function purchase_code()
returns text as
$$
declare
    old_id text :=(select  purchase_id  from purchase order by purchase_id desc limit 1);
    id_number char(4) :='0001';
    datetime char(4) :=substring(cast(now() as text),1,4);
    new_id text ;
    num integer;
begin
    if old_id is null then
        new_id:='P'||datetime||id_number;
        return new_id;
    end if;
    
    if datetime=substring(old_id,4,4) then
        num :=cast(right(old_id,4) as integer)+1;
        id_number:=
        case
            when num<10 then '000'||num
            when num<100 then '00'||num
            when num<1000 then '0'||num
            when num<10000 then cast(num as text)
        end;
    end if;
    
    new_id:='P'||datetime||id_number;
    return new_id;
end; 
$$
language plpgsql;