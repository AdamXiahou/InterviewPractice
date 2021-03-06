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
                 
**2.查询姓名总最后一个字是‘猴’的学生**
           
**3.查询姓名中带‘猴’的学生名单**
       
**4.查询姓“孟”老师的个数**
       
### 汇总分析       
       
**5.查询课程编号为“0002”的总成绩**       
    
**6.查询选了课程的学生人数**
   
### 分组

**7.查询各科成绩最高和最低的分**       
   
**8.查询每门课程被选修的学生数**
    
**9.查询男生，女生人数**
   
### 分组结果的条件
       
**10.查询平均成绩大于60分学生的学号和平均成绩**

**11.查询至少选修两门课程的学生学号**
 
**12.查询同名同姓学生名单并统计同名人数**
 
**13.查询不及格的课程并按课程号从大到小排列**
       
**14.查询每门课程的平均成绩，结果按平均成绩升序排序，平均成绩相同时，按课程号降序排列**
      
**15.检索课程编号为“0004”且分数小于60的学生学号，结果按按分数降序排列**         
     
**16.统计每门课程的学生选修人数(超过2人的课程才统计)，要求输出课程号和选修人数，查询结果按人数降序排序，若人数相同，按课程号升序排序**         
      
**17.查询两门以上不及格课程的同学的学号及其平均成绩**         
  
### 汇总分析
         
**18.查询学生的总成绩并进行排名**         
 
**19.查询平均成绩大于60分的学生的学号和平均成绩**         
 
### 复杂查询
         
**20.查询所有课程成绩小于60分学生的学号、姓名**         

**21.查询没有学全所有课的学生的学号、姓名**
   
**22.查询出只选修了两门课程的全部学生的学号和姓名**       
       
**23.1990年出生的学生名单**       

**24.查询各科成绩前两名的记录**       
     
**25.查询各学生的年龄(精确到月份)**       
      
**26.查询本月过生日的学生**       
       
### 多表查询       
       
**27.查询所有学生的学号、姓名、选课数、总成绩**       
      
**28.查询平均成绩大于85的所有学生的学号、姓名和平均成绩**       
       
**29.查询学生的选课情况：学号，姓名，课程号，课程名称**       
       
**30.查询出每门课程的及格人数和不及格人数**       

**31.使用分段[100-85],[85-70],[70-60],[<60]来统计各科成绩，分别统计：各分数段人数，课程号和课程名称**              

**32.查询课程编号为0003且课程成绩在80分以上的学生的学号和姓名**       
              
**33.行列变化，将score变成如下表形式:**
     
### 多表连接
     
**34.检索"0001"课程分数小于60，按分数降序排列的学生信息**     

**35.查询不同老师所教不同课程平均分从高到低显示**     
   
**36.查询课程名称为"数学"，且分数低于60的学生姓名和分数**
    
**37.查询任何一门课程成绩在70分以上的姓名、课程名称和分数**     

**38.查询两门及其以上不及格课程的同学的学号，姓名及其平均成绩**     

**39.查询不同课程成绩相同的学生的学生编号、课程编号、学生成绩**     
    
**40.查询课程编号为“0001”的课程比“0002”的课程成绩高的所有学生的学号**     
     
**41.查询学过编号为“0001”的课程并且也学过编号为“0002”的课程的学生的学号、姓名**
    
**42.查询学过“孟扎扎”老师所教的所有课的同学的学号、姓名**     
    
**43.查询没学过"孟扎扎"老师讲授的任一门课程的学生姓名**     
     
**44.查询选修“孟扎扎”老师所授课程的学生中成绩最高的学生姓名及其成绩**     
     
**45.查询至少有一门课与学号为“0001”的学生所学课程相同的学生的学号和姓名**     
          
**46.按平均成绩从高到低.显示所有学生的所有课程的成绩以及平均成绩，总成绩**          
    
### SQL高级功能：窗口函数
         
**47.查询学生平均成绩及其名次**         
         
**48.按各科成绩进行排序，并显示排名**         
         
**49.查询每门功成绩最好的前两名学生姓名**         
        
**50.查询所有课程的成绩第2名到第3名的学生信息及该课程成绩**         