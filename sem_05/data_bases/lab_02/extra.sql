drop table if exists t1;
create table t1 (
    id int,
    var1 text,
    val_from date,
    val_to date
);

drop table if exists t2;
create table t2 (
    id int,
    var2 text,
    val_from date,
    val_to date
);

insert into t1 (id, var1, val_from, val_to)
values (1, 'A', '2018-09-01','2018-09-15'), (1, 'B', '2018-09-16','5999-12-31'),
       (2, 'C', '2018-09-01','5999-12-31');
insert into t2 (id, var2, val_from, val_to)
values (1, 'A', '2018-09-01','2018-09-18'), (1, 'B', '2018-09-19','5999-12-31');

-- select joined.id, joined.var1, joined.var2, joined.val_from, joined.val_to
-- from ((select * from t1) union (select * from t2)) as joined

select q1.id, var1, var2, q1.val_from, q2.val_to
from ((select id, var1, val_from, val_to, row_number() over (partition by id order by val_from) as q1_rn
      from (select *, row_number() over (partition by id, val_from) as rn
             from ((select id, var1, val_from, val_to from t1) union all
                   (select id, var2 as var1, val_from, val_to from t2)) as vars) as vars1
      where vars1.rn = 1) as q1
      join
      (select id, var2, val_from, val_to, row_number() over (partition by id order by val_to) as q2_rn
      from (select *, row_number() over (partition by id, val_to) as rn
      from ((select id, var1 as var2, val_from, val_to from t1) union all
            (select id, var2, val_from, val_to from t2)) as vars) as vars1
      where vars1.rn = 1) as q2 on q1.id = q2.id and q1.q1_rn = q2.q2_rn);

-- select id, var1, var2, val_from, val_to, row_number() over (order by val_from)
-- from (select *, row_number() over (partition by id, val_from) as rn
--       from ((select id, var1, var1 as var2, val_from, val_to from t1) union all
--             (select id, var2 as var1, var2, val_from, val_to from t2)) as vars) as vars1
-- where vars1.rn = 1;
--
-- select id, var1, var2, val_from, val_to, row_number() over (order by val_to)
-- from (select *, row_number() over (partition by id, val_to) as rn
--       from ((select id, var1, var1 as var2, val_from, val_to from t1) union all
--             (select id, var2 as var1, var2, val_from, val_to from t2)) as vars) as vars1
-- where vars1.rn = 1;
