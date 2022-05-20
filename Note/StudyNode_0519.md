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

### Questions
1. 查询课程编号为“0002”的总成绩
`SELECT SUM(成绩) FROM score WHERE 课程号 = '0002';`

2. 查询各科成绩最高和最低的分
`SELECT 课程号,MAX(成绩) as 最高分,MIN(成绩) as 最低分 FROM score GROUP BY 课程号;`

3. 查询平均成绩大于60分学生的学号
`SELECT 学号 FROM score GROUP BY 学号 HAVING AVG(成绩)>60;`

### 要点：
1. Select 加查询条件+可以加聚合函数

2. 看清题目，查询时与插入时一样

3. HAVING 子句可以让我们筛选分组后的各组数据

# 微服务
``[ApiController]
[Route("[controller]")]
public class WeatherForecastController : ControllerBase
{
    private static readonly string[] Summaries = new[]
    {
        "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
    };

    private readonly ILogger<WeatherForecastController> _logger;

    public WeatherForecastController(ILogger<WeatherForecastController> logger)
    {
        _logger = logger;
    }

    [HttpGet]
    public IEnumerable<WeatherForecast> Get()
    {
        var rng = new Random();
        return Enumerable.Range(1, 5).Select(index => new WeatherForecast
        {
            Date = DateTime.Now.AddDays(index),
            TemperatureC = rng.Next(-20, 55),
            Summary = Summaries[rng.Next(Summaries.Length)]
        })
        .ToArray();
    }
}
``

<code>

            Date = DateTime.Now.AddDays(index),
            TemperatureC = rng.Next(-20, 55),
            Summary = Summaries[rng.Next(Summaries.Length)]
</code>

