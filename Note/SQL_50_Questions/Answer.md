# 数据库

### 学生表
`insert into student(学号,姓名,出生日期,性别)values('0001' , '猴子' , '1989-01-01' , '男');`

`insert into student(学号,姓名,出生日期,性别)values('0002' , '猴子' , '1990-12-21' , '女');`

`insert into student(学号,姓名,出生日期,性别)values('0003' , '马云' , '1991-12-21' , '男');`

`insert into student(学号,姓名,出生日期,性别)values('0004' , '王思聪' , '1990-05-20' , '男');`

### 成绩表
`insert into score(学号,课程号,成绩)values('0001' , '0001' , 80);`

`insert into score(学号,课程号,成绩)values('0001' , '0002' , 90);`

`insert into score(学号,课程号,成绩)values('0001' , '0003' , 99);`

`insert into score(学号,课程号,成绩)values('0002' , '0002' , 60);`

`insert into score(学号,课程号,成绩)values('0002' , '0003' , 80);`

`insert into score(学号,课程号,成绩)values('0003' , '0001' , 80);`

`insert into score(学号,课程号,成绩)values('0003' , '0002' , 80);`

`insert into score(学号,课程号,成绩)values('0003' , '0003' , 80);`

### 课程表
`insert into course(课程号,课程名称,教师号)values('0001' , '语文' , '0002');`

`insert into course(课程号,课程名称,教师号)values('0002' , '数学' , '0001');`

`insert into course(课程号,课程名称,教师号)values('0003' , '英语' , '0003');`

### 教师表
`insert into teacher(教师号,教师姓名)values('0001' , '孟扎扎');`

`insert into teacher(教师号,教师姓名)values('0002' , '马化腾');`

`insert into teacher(教师号,教师姓名)values('0003' , null);`

`insert into teacher(教师号,教师姓名)values('0004' , '');`

# 50道面试题
### 简单查询
**1.查询姓“猴“的学生名单**
       
SELECT * FROM student WHERE 姓名 LIKE '猴%';       
       
**2.查询姓名总最后一个字是‘猴’的学生**
       
SELECT * FROM student WHERE 姓名 LIKE '%猴';       
       
**3.查询姓名中带‘猴’的学生名单**
       
SELECT * FROM student WHERE 姓名 LIKE '%猴%';       
       
**4.查询姓“孟”老师的个数**
       
SELECT * FROM teacher WHERE 教师姓名 LIKE '孟%';       
       
### 汇总分析       
       
**5.查询课程编号为“0002”的总成绩**       
       
分析思路       
       
select 查询结果 [总成绩:汇总函数sum]             
from 从哪张表中查找数据[成绩表score]             
where 查询条件 [课程号是0002]       
       
SELECT SUM(成绩) FROM score WHERE 课程号 = '0002';       
       
**6.查询选了课程的学生人数**
         
分析思路       
          
这个题目翻译成大白话就是：查询有多少人选了课程       
select 学号，成绩表里学号有重复值需要去掉       
from 从课程表查找score;       
       
SELECT COUNT(DISTINCT 学号) as 学生人数 FROM score;       
       
### 分组

**7.查询各科成绩最高和最低的分**       
        
分析思路       
       
select 查询结果 [课程ID：是课程号的别名,最高分：max(成绩) ,最低分：min(成绩)]       
from 从哪张表中查找数据 [成绩表score]       
where 查询条件 [没有]       
group by 分组 [各科成绩：也就是每门课程的成绩，需要按课程号分组];       
       
SELECT 课程号,MAX(成绩) as 最高分,MIN(成绩) as 最低分 FROM score GROUP BY 课程号;       
       
**8.查询每门课程被选修的学生数**
       
分析思路       
       
select 查询结果 [课程号，选修该课程的学生数：汇总函数count]       
from 从哪张表中查找数据 [成绩表score]       
where 查询条件 [没有]       
group by 分组 [每门课程：按课程号分组];       
       
SELECT 课程号,COUNT(DISTINCT 学号) as 学生数 FROM score GROUP BY 课程号;       
       
**9.查询男生，女生人数**
       
分析思路       
       
select 查询结果 [性别，对应性别的人数：汇总函数count]       
from 从哪张表中查找数据 [性别在学生表中，所以查找的是学生表student]       
where 查询条件 [没有]       
group by 分组 [男生、女生人数：按性别分组]       
having 对分组结果指定条件 [没有]       
order by 对查询结果排序[没有];       
       
SELECT 性别,COUNT(性别) as 学生数 FROM student GROUP BY 性别;       
       
SELECT 性别,COUNT(*) as 学生数 FROM student GROUP BY 性别;       
       
### 分组结果的条件
       
**10.查询平均成绩大于60分学生的学号和平均成绩**
       
分析思路       
       
平均成绩：展开来说就是计算每个学生的平均成绩       
这里涉及到“每个”就是要分组了       
平均成绩大于60分，就是对分组结果指定条件       
              
select 查询结果 [学号，平均成绩：汇总函数avg(成绩)]       
from 从哪张表中查找数据 [成绩在成绩表中，所以查找的是成绩表score]       
where 查询条件 [没有]       
group by 分组 [平均成绩：先按学号分组，再计算平均成绩]       
having 对分组结果指定条件 [平均成绩大于60分]       
       
SELECT 学号,AVG(成绩) as 平均成绩 FROM score GROUP BY 学号 HAVING AVG(成绩)>60;       
       
**11.查询至少选修两门课程的学生学号**
       
分析思路       
              
第1步，需要先计算出每个学生选修的课程数据，需要按学号分组       
第2步，至少选修两门课程：也就是每个学生选修课程数目>=2，对分组结果指定条件       
       
select 查询结果 [学号,每个学生选修课程数目：汇总函数count]       
from 从哪张表中查找数据 [课程的学生学号：课程表score]       
where 查询条件 [至少选修两门课程：需要先计算出每个学生选修了多少门课，需要用分组，所以这里没有where子句]       
group by 分组 [每个学生选修课程数目：按课程号分组，然后用汇总函数count计算出选修了多少门课]       
having 对分组结果指定条件 [至少选修两门课程：每个学生选修课程数目>=2]       
              
SELECT 学号,COUNT(课程号) as 选修课程数 FROM score GROUP BY 学号 HAVING COUNT(课程号)>=2;       
       
**12.查询同名同姓学生名单并统计同名人数**
       
分析思路         
         
1)查找出姓名相同的学生有谁，每个姓名相同学生的人数         
查询结果：姓名,人数         
条件：怎么算姓名相同？按姓名分组后人数大于等于2，因为同名的人数大于等于2         
         
select 查询结果 [姓名,人数：汇总函数count(*)]         
from 从哪张表中查找数据 [学生表student]         
where 查询条件 [没有]         
group by 分组 [姓名相同：按姓名分组]         
having 对分组结果指定条件 [姓名相同：count(*)>=2]         
order by 对查询结果排序[没有];         
         
SELECT 姓名,COUNT(姓名) as 同名人数 FROM student GROUP BY 姓名 HAVING COUNT(姓名)>=2;         
         
**13.查询不及格的课程并按课程号从大到小排列**
         
分析思路         
         
select 查询结果 [课程号]         
from 从哪张表中查找数据 [成绩表score]         
where 查询条件 [不及格：成绩 <60]         
group by 分组 [没有]         
having 对分组结果指定条件 [没有]         
order by 对查询结果排序[课程号从大到小排列：降序desc];         
         
SELECT 课程号 FROM score WHERE 成绩<60 ORDER BY 课程号 DESC;         
         
**14.查询每门课程的平均成绩，结果按平均成绩升序排序，平均成绩相同时，按课程号降序排列**
         
分析思路         
         
select 查询结果 [课程号,平均成绩：汇总函数avg(成绩)]         
from 从哪张表中查找数据 [成绩表score]         
where 查询条件 [没有]         
group by 分组 [每门课程：按课程号分组]         
having 对分组结果指定条件 [没有]         
order by 对查询结果排序[按平均成绩升序排序:asc，平均成绩相同时，按课程号降序排列:desc];         
         
SELECT 课程号,AVG(成绩) as 平均成绩 FROM score GROUP BY 课程号 ORDER BY 平均成绩 ASC,课程号 DESC;         
         
**15.检索课程编号为“0004”且分数小于60的学生学号，结果按按分数降序排列**         
         
分析思路         
         
select 查询结果 []         
from 从哪张表中查找数据 [成绩表score]         
where 查询条件 [课程编号为“04”且分数小于60]         
group by 分组 [没有]         
having 对分组结果指定条件 []         
order by 对查询结果排序[查询结果按按分数降序排列];         
         
SELECT 学号 FROM score WHERE 课程号='0004' AND 成绩<60 ORDER BY 成绩 DESC;         
         
**16.统计每门课程的学生选修人数(超过2人的课程才统计)，要求输出课程号和选修人数，查询结果按人数降序排序，若人数相同，按课程号升序排序**         
         
分析思路         
         
select 查询结果 [要求输出课程号和选修人数]         
from 从哪张表中查找数据 []         
where 查询条件 []         
group by 分组 [每门课程：按课程号分组]         
having 对分组结果指定条件 [学生选修人数(超过2人的课程才统计)：每门课程学生人数>2]         
order by 对查询结果排序[查询结果按人数降序排序，若人数相同，按课程号升序排序];         
         
SELECT 课程号,COUNT(学号) as 选修人数 FROM score GROUP BY 课程号 HAVING 选修人数>2 ORDER BY 选修人数 DESC,课程号 ASC;         
         
**17.查询两门以上不及格课程的同学的学号及其平均成绩**         
         
分析思路         
         
先分解题目：         
1)[两门以上][不及格课程]限制条件         
2)[同学的学号及其平均成绩]，也就是每个学生的平均成绩，显示学号，平均成绩                  
         
分析过程：         
第1步：得到每个学生的平均成绩，显示学号，平均成绩         
第2步：再加上限制条件：         
                    1)不及格课程         
                    2)两门以上[不及格课程]：课程数目>2         
         
第1步：得到每个学生的平均成绩，显示学号，平均成绩         
         
select 查询结果 [学号,平均成绩：汇总函数avg(成绩)]         
from 从哪张表中查找数据 [涉及到成绩：成绩表score]         
where 查询条件 [没有]         
group by 分组 [每个学生的平均：按学号分组]         
having 对分组结果指定条件 [没有]         
order by 对查询结果排序[没有];         
         
select 学号, avg(成绩) as 平均成绩         
from score         
group by 学号;         
         
第2步：再加上限制条件：         
         
1)不及格课程         
2)两门以上[不及格课程]         
         
select 查询结果 [学号,平均成绩：汇总函数avg(成绩)]         
from 从哪张表中查找数据 [涉及到成绩：成绩表score]         
where 查询条件 [限制条件：不及格课程，平均成绩<60]         
group by 分组 [每个学生的平均：按学号分组]         
having 对分组结果指定条件 [限制条件：课程数目>2,汇总函数count(课程号)>2]         
order by 对查询结果排序[没有];         
         
SELECT 学号,AVG(成绩) as 平均成绩 FROM score WHERE 成绩<60 GROUP BY 学号 HAVING COUNT(成绩)>2;         
         
### 汇总分析
         
**18.查询学生的总成绩并进行排名**         
         
【知识点】分组查询         
分析思路         
         
select 查询结果 [总成绩：sum(成绩), 学号]         
from 从哪张表中查找数据 [成绩表score]         
where 查询条件 [没有]         
group by 分组 [学生的总成绩：按照每个学生学号进行分组]         
order by 排序 [按照总成绩进行排序：sum(成绩)];         
         
SELECT 学号,SUM(成绩) as 总成绩 FROM score GROUP BY 学号 ORDER BY 总成绩 DESC;         
         
**19.查询平均成绩大于60分的学生的学号和平均成绩**         
         
【知识点】分组+条件         
分析思路         
         
select 查询结果 [学号, 平均成绩: avg(成绩)]         
from 从哪张表中查找数据 [成绩表score]         
where 查询条件 [没有]         
group by 分组 [学号]         
having 分组条件 [平均成绩大于60分：avg(成绩 ) >60]         
order by 排序 [没有];         
         
SELECT 学号,AVG(成绩) as 平均成绩 FROM score GROUP BY 学号 HAVING 平均成绩>60;         
         
### 复杂查询
         
**20.查询所有课程成绩小于60分学生的学号、姓名**         
         
【知识点】子查询         
分析思路         
         
1)查询结果：学生学号，姓名         
2)查询条件：所有课程成绩 < 60 的学生，需要从成绩表里查找，用到子查询         
         
第1步，写子查询(所有课程成绩 < 60 的学生)         
         
select 查询结果[学号]         
from 从哪张表中查找数据[成绩表：score]         
where 查询条件[成绩 < 60]         
group by 分组[没有]         
having 对分组结果指定条件[没有]         
order by 对查询结果排序[没有]         
limit 从查询结果中取出指定行[没有];         
         
select 学号         
from student         
where 成绩 < 60;         
         
第2步，查询结果：学生学号，姓名，条件是前面1步查到的学号         
         
select 查询结果[学号,姓名]         
from 从哪张表中查找数据[学生表:student]         
where 查询条件[用到运算符in]         
group by 分组[没有]         
having 对分组结果指定条件[没有]         
order by 对查询结果排序[没有]         
limit 从查询结果中取出指定行[没有];         
         
SELECT 学号,姓名 FROM student WHERE 学号 in (SELECT 学号 FROM score WHERE 成绩<60);         
         
**21.查询没有学全所有课的学生的学号、姓名**
       
【考察知识点】in，子查询       
分析思路       
       
查找出学号，条件：没有学全所有课，也就是该学生选修的课程数 < 总的课程数       
       
SELECT 学号,姓名 FROM student WHERE 学号 IN ( SELECT 学号 FROM score GROUP BY 学号 HAVING COUNT(学号)       
       
**22.查询出只选修了两门课程的全部学生的学号和姓名**       
       
SELECT 学号,姓名 FROM student WHERE 学号 IN ( SELECT 学号 FROM score GROUP BY 学号 HAVING COUNT(课程号)=2);       
       
**23.1990年出生的学生名单**       
       
分析思路        
       
查找1990年出生的学生名单       
学生表中出生日期列的类型是datetime       
       
SELECT * FROM student WHERE 出生日期 LIKE "1990%";       
       
SELECT * FROM student WHERE YEAR(出生日期) = 1990;       
       
**24.查询各科成绩前两名的记录**       
       
分析思路       
       
第1步，查出有哪些组       
我们可以按课程号分组，查询出有哪些组，对应这个问题里就是有哪些课程号       
select 课程号,max(成绩) as 最大成绩       
       
from score       
group by 课程号;       
       
第2步：先使用order by子句按成绩降序排序(desc)，然后使用limt子句返回topN(对应这个问题返回的成绩前两名)       
-- 课程号'0001' 这一组里成绩前2名       
       
select *       
from score       
where 课程号 = '0001'       
order by 成绩 desc       
limit 2;       
       
同样的，可以写出其他组的(其他课程号)取出成绩前2名的sql       
       
第3步，使用union all 将每组选出的数据合并到一起       
-- 左右滑动可以可拿到全部sql       
       
(SELECT * FROM score WHERE 课程号 = '0001' ORDER BY 成绩 DESC LIMIT 2)       
UNION ALL       
(SELECT * FROM score WHERE 课程号 = '0002' ORDER BY 成绩 DESC LIMIT 2)       
UNION ALL       
(SELECT * FROM score WHERE 课程号 = '0003' ORDER BY 成绩 DESC LIMIT 2);       
       
**25.查询各学生的年龄(精确到月份)**       
       
下面说明了TIMESTAMPDIFF函数的语法。       
       
TIMESTAMPDIFF(unit,begin,end);       
TIMESTAMPDIFF函数返回begin-end的结果，其中begin和end是DATE或DATETIME表达式。       
TIMESTAMPDIFF函数允许其参数具有混合类型，例如，begin是DATE值，end可以是DATETIME值。 如果使用DATE值，则TIMESTAMPDIFF函数将其视为时间部分为“00:00:00”的DATETIME值。       
unit参数是确定(end-begin)的结果的单位，表示为整数。 以下是有效单位：       
       
MICROSECOND       
SECOND       
MINUTE       
HOUR       
DAY       
WEEK       
MONTH       
QUARTER       
YEAR       
NOW() 函数返回当前的日期和时间。       
       
select 学号 ,timestampdiff(month ,出生日期 ,now())/12 from student ;       
       
**26.查询本月过生日的学生**       
       
SELECT * FROM student WHERE MONTH(出生日期) = month(now());       
       
SELECT * FROM student WHERE MONTH(CURRENT_DATE) = MONTH(出生日期);       
       
### 多表查询       
       
**27.查询所有学生的学号、姓名、选课数、总成绩**       
       
SELECT a.学号,a.姓名,COUNT(b.课程号) as 选课数,SUM(b.成绩) as 总成绩 FROM student as a       
       
LEFT JOIN score as b on a.学号 = b.学号 GROUP BY a.学号;       
       
**28.查询平均成绩大于85的所有学生的学号、姓名和平均成绩**       
       
SELECT a.学号,a.姓名,avg(b.成绩) as 平均成绩 FROM student as a LEFT JOIN score as b on a.学号 = b.学号 GROUP BY a.学号 HAVING 平均成绩>85;       
       
**29.查询学生的选课情况：学号，姓名，课程号，课程名称**       
       
select a.学号, a.姓名, c.课程号,c.课程名称 from student a inner join score b on a.学号=b.学号 inner join course c on b.课程号=c.课程号;       
       
**30.查询出每门课程的及格人数和不及格人数**       
       
-- 考察case表达式       
       
select 课程号,       
sum(case when 成绩>=60 then 1       
else 0       
end) as 及格人数,       
sum(case when 成绩 < 60 then 1       
else 0       
end) as 不及格人数       
from score       
group by 课程号;       

**31.使用分段[100-85],[85-70],[70-60],[<60]来统计各科成绩，分别统计：各分数段人数，课程号和课程名称**       
       
-- 考察case表达式       
       
select a.课程号,b.课程名称,       
sum(case when 成绩 between 85 and 100       
then 1 else 0 end) as '[100-85]',       
sum(case when 成绩 >=70 and 成绩<85       
then 1 else 0 end) as '[85-70]',       
sum(case when 成绩>=60 and 成绩<70       
then 1 else 0 end) as '[70-60]',       
sum(case when 成绩<60 then 1 else 0 end) as '[<60]'       
from score as a right join course as b       
on a.课程号=b.课程号       
group by a.课程号,b.课程名称;       

**32.查询课程编号为0003且课程成绩在80分以上的学生的学号和姓名**       
       
SELECT a.学号,a.姓名 FROM student as a INNER JOIN score as b on a.学号 = b.学号 WHERE b.课程号='0003' AND b.成绩>80;       
       
**33.行列变化，将score变成如下表形式:**
     
第1步，使用常量列输出目标表的结构     
     
可以看到查询结果已经和目标表非常接近了     
     
select 学号,'课程号0001','课程号0002','课程号0003' from score;     
     
第2步，使用case表达式，替换常量列为对应的成绩     
     
select 学号,     
(case 课程号 when '0001' then 成绩 else 0 end) as '课程号0001',     
(case 课程号 when '0002' then 成绩 else 0 end) as '课程号0002',     
(case 课程号 when '0003' then 成绩 else 0 end) as '课程号0003'     
from score;     
     
第3关，分组     
分组，并使用最大值函数max取出上图每个方块里的最大值     
     
SELECT 学号,     
MAX(CASE 课程号 WHEN '0001' THEN 成绩 ELSE 0 END ) as '课程号0001',     
MAX(CASE 课程号 WHEN '0002' THEN 成绩 ELSE 0 END ) as '课程号0002',     
MAX(CASE 课程号 WHEN '0003' THEN 成绩 ELSE 0 END ) as '课程号0003'     
FROM score     
GROUP BY 学号;     
     
### 多表连接
     
**34.检索"0001"课程分数小于60，按分数降序排列的学生信息**     

SELECT a.*,b.成绩 FROM student as a INNER JOIN score as b on a.学号=b.学号 WHERE b.成绩<60 AND b.课程号='0001' ORDER BY b.成绩 DESC;

**35.查询不同老师所教不同课程平均分从高到低显示**     
     
-- 子查询+内连接     
     
SELECT a.教师号,a.教师姓名,AVG(b.成绩) as 平均成绩 FROM (SELECT a.教师号,a.教师姓名,b.课程号,b.课程名称 FROM teacher as a INNER JOIN course as b on a.教师号=b.教师号) as a INNER JOIN score as b on a.课程号=b.课程号 GROUP BY 教师号 ORDER BY 平均成绩 DESC;     
     
-- 简写，可以直接3个表内连接形成一个新表     
     
select a.教师号,a.教师姓名,avg(c.成绩) from teacher as a inner join course as b on a.教师号= b.教师号 inner join score c on b.课程号=c.课程号     
group by a.教师姓名     
order by avg(c.成绩) desc;     
     
**36.查询课程名称为"数学"，且分数低于60的学生姓名和分数**

SELECT a.姓名,b.成绩 as 分数 FROM student as a INNER JOIN score as b on a.学号=b.学号 INNER JOIN course as c on b.课程号=c.课程号 WHERE c.课程名称='数学' AND b.成绩<60;     
     
**37.查询任何一门课程成绩在70分以上的姓名、课程名称和分数**     

SELECT a.姓名,c.课程名称,b.成绩 as 分数 FROM student as a INNER JOIN score as b on a.学号=b.学号 INNER JOIN course as c on b.课程号=c.课程号 WHERE b.成绩>70;     

**38.查询两门及其以上不及格课程的同学的学号，姓名及其平均成绩**     

SELECT a.学号,a.姓名,AVG(b.成绩) as 平均成绩 FROM student as a INNER JOIN score as b on a.学号=b.学号 WHERE b.成绩<60 GROUP BY a.学号 HAVING COUNT(a.学号)>=2;     

**39.查询不同课程成绩相同的学生的学生编号、课程编号、学生成绩**     
     
SELECT DISTINCT a.学号,a.课程号,a.成绩 FROM score as a INNER JOIN score as b on a.学号=b.学号 WHERE b.成绩=a.成绩 AND a.课程号!=b.课程号;     
     
**40.查询课程编号为“0001”的课程比“0002”的课程成绩高的所有学生的学号**     
     
SELECT a.学号 FROM (SELECT 学号,成绩 FROM score WHERE 课程号='0001') as a INNER JOIN (SELECT 学号,成绩 FROM score WHERE 课程号='0002') as b on a.学号=b.学号 WHERE a.成绩>b.成绩;     
     
**41.查询学过编号为“0001”的课程并且也学过编号为“0002”的课程的学生的学号、姓名**

SELECT a.学号,c.姓名 FROM (SELECT 学号,成绩 FROM score WHERE 课程号='0001') as a INNER JOIN (SELECT 学号,成绩 FROM score WHERE 课程号='0002') as b on a.学号=b.学号 INNER JOIN student c on c.学号=a.学号;     
     
**42.查询学过“孟扎扎”老师所教的所有课的同学的学号、姓名**     
     
SELECT a.学号,a.姓名 FROM student as a INNER JOIN score as b on a.学号=b.学号 INNER JOIN course as c on c.课程号=b.课程号 INNER JOIN teacher as d on d.教师号=c.教师号 WHERE d.教师姓名='孟扎扎';     
     
**43.查询没学过"孟扎扎"老师讲授的任一门课程的学生姓名**     
     
SELECT 姓名,学号 FROM student WHERE 学号 NOT IN (SELECT a.学号 FROM student as a INNER JOIN score as b on a.学号=b.学号 INNER JOIN course as c on c.课程号=b.课程号 INNER JOIN teacher as d on d.教师号=c.教师号 WHERE d.教师姓名='孟扎扎');     
     
**44.查询选修“孟扎扎”老师所授课程的学生中成绩最高的学生姓名及其成绩**     
     
SELECT a.学号,a.姓名,b.成绩 FROM student as a INNER JOIN score as b on a.学号=b.学号 INNER JOIN course as c on c.课程号=b.课程号 INNER JOIN teacher as d on d.教师号=c.教师号 WHERE d.教师姓名='孟扎扎' ORDER BY b.成绩 DESC LIMIT 1;     
     
**45.查询至少有一门课与学号为“0001”的学生所学课程相同的学生的学号和姓名**     
     
SELECT 学号,姓名 FROM student WHERE 学号 IN (SELECT b.课程号 FROM student as a INNER JOIN score as b on a.学号=b.学号 WHERE b.学号='0001');     
          
**46.按平均成绩从高到低.显示所有学生的所有课程的成绩以及平均成绩，总成绩**          
     
SELECT a.学号,AVG(a.成绩) as 平均成绩,     
MAX(CASE WHEN b.课程名称='数学' THEN a.成绩 ELSE NULL END ) as '数学',     
MAX(CASE WHEN b.课程名称='语文' THEN a.成绩 ELSE NULL END ) as '语文',     
MAX(CASE WHEN b.课程名称='英语' THEN a.成绩 ELSE NULL END ) as '英语',     
SUM(a.成绩) as 总成绩     
FROM score as a INNER JOIN course as b on a.课程号=b.课程号 GROUP BY a.学号 ORDER BY 平均成绩 DESC;     
     
### SQL高级功能：窗口函数
         
**47.查询学生平均成绩及其名次**         
         
select s1.id,s1.name,c.av,c.rank from (select a.stu_id, a.av, (@row_num := @row_num + 1) as rank from (select stu_id, avg(grade) as av from score group by stu_id order by av desc) a join (select @row_num := 0) as b) as c join student as s1 on c.stu_id = s1.id order by c.rank;         
         
**48.按各科成绩进行排序，并显示排名**         
         
select stu_id,avg(grade) as av,(@rank := @rank + 1) as rank from score join (select @rank := 0) as b group by stu_id order by av desc;          
         
**49.查询每门功成绩最好的前两名学生姓名**         
         
select s1.name, s.* from student as s1 join (select * from score s1 where (select count(*) from score s2 where s2.course_id = s1.course_id and s2.grade >= s1.grade) < 3 order by s1.course_id asc) as s on s.stu_id = s1.id order by s.course_id desc,s.grade desc;         
         
**50.查询所有课程的成绩第2名到第3名的学生信息及该课程成绩**         
         
select s1.course_id,s1.grade from score as s1 join (select distinct course_id, grade from score) as s2 on s1.course_id = s2.course_id and
s1.grade <= s2.grade group by s1.course_id , s1.grade having count(*) < 3 order by s1.course_id desc,s1.grade desc;         
         
