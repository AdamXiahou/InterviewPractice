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

namespace Webapplication1.Controllers

[Authorize]    
[ApiController]    
[Route("[controller]")]    
[RequiredScope(RequiredScopesConfigurationKey = "AzureAd:Scopes")]     

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
    
    [HttpGet(Name = "GetWeatherForecast")]     
    public IEnumerable<WeatherForecast> Get()     
    {        
        return Enumerable.Range(1, 5).Select(index => new WeatherForecast    
        {     
            Date = DateTime.Now.AddDays(index),    
            TemperatureC = Random.Shared.Next(-20, 55),    
            Summary = Summaries[Random.Shared.Next(Summaries.Length)]    
        })    
        .ToArray();    
    }    
}    

### 状态码总结

##### 1xx - 信息提示
100 - Continue 初始的请求已经接受，客户应当继续发送请求的其余部分。         
101 - Switching Protocols 服务器将遵从客户的请求转换到另外一种协议     

##### 2xx - 成功
这类状态代码表明服务器成功地接受了客户端请求。        
200 - OK 一切正常，对GET和POST请求的应答文档跟在后面。        
201 - Created 服务器已经创建了文档，Location头给出了它的URL。        
202 - Accepted 已经接受请求，但处理尚未完成。        
203 - Non-Authoritative Information 文档已经正常地返回，但一些应答头可能不正确，因为使用的是文档的拷贝，非权威性信息。        
204 - No Content 没有新文档，浏览器应该继续显示原来的文档。如果用户定期地刷新页面，而Servlet可以确定用户文档足够新，这个状态代码是很有用的。        
205 - Reset Content 没有新的内容，但浏览器应该重置它所显示的内容。用来强制浏览器清除表单输入内容。        
206 - Partial Content 客户发送了一个带有Range头的GET请求，服务器完成了它。        

##### 3xx - 重定向
客户端浏览器必须采取更多操作来实现请求。例如，浏览器可能不得不请求服务器上的不同的页面，或通过代理服务器重复该请求。        
300 - Multiple Choices 客户请求的文档可以在多个位置找到，这些位置已经在返回的文档内列出。如果服务器要提出优先选择，则应该在Location应答头指明。        
301 - Moved Permanently 客户请求的文档在其他地方，新的URL在Location头中给出，浏览器应该自动地访问新的URL。        
302 - Found 类似于301，但新的URL应该被视为临时性的替代，而不是永久性的。注意，在HTTP1.0中对应的状态信息是“Moved Temporatily”。出现该状态代码时，浏览器能够自动访问新的URL，因此它是一个很有用的状态代码。注意这个状态代码有时候可以和301替换使用。例如，如果浏览器错误地请求 http://host/~user （缺少了后面的斜杠），有的服务器返回301，有的则返回302。严格地说，我们只能假定只有当原来的请求是GET时浏览器才会自动重定向。请参见 307。        
303 - See Other 类似于301/302，不同之处在于，如果原来的请求是POST，Location头指定的重定向目标文档应该通过GET提取。        
304 - Not Modified 客户端有缓冲的文档并发出了一个条件性的请求（一般是提供If-Modified-Since头表示客户只想比指定日期更新的文档）。服务器告诉客户，原来缓冲的文档还可以继续使用。        
305 - Use Proxy 客户请求的文档应该通过Location头所指明的代理服务器提取。        

##### 4xx - 客户端错误
发生错误，客户端似乎有问题。例如，客户端请求不存在的页面，客户端未提供有效的身份验证信息。        
400 - Bad Request 请求出现语法错误。        
400的返回body里面经常有Invalid argument provided （提供无效的参数）表明header或parameter里面有无效的参数。        
401 - Unauthorized 访问被拒绝，客户试图未经授权访问受密码保护的页面。        
404 - Not Found 无法找到指定位置的资源。这也是一个常用的应答。        

##### 5xx - 服务器错误
服务器由于遇到错误而不能完成该请求。        
500 - Internal Server Error 服务器遇到了意料不到的情况，不能完成客户的请求。        
501 - Not Implemented 服务器不支持实现请求所需要的功能，页眉值指定了未实现的配置。例如，客户发出了一个服务器不支持的PUT请求。        
502 - Bad Gateway 服务器作为网关或者代理时，为了完成请求访问下一个服务器，但该服务器返回了非法的应答。 亦说Web 服务器用作网关或代理服务器时收到了无效响应。        
503 - Service Unavailable 服务不可用，服务器由于维护或者负载过重未能应答。例如，Servlet可能在数据库连接池已满的情况下返回503。服务器返回503时可以提供一个 Retry-After头。这个错误代码为 IIS 6.0 所专用。        
504 - Gateway Timeout 网关超时，由作为代理或网关的服务器使用，表示不能及时地从远程服务器获得应答。        
505 - HTTP Version Not Supported 服务器不支持请求中所指明的HTTP版本。        

### 什么是中间件（Middleware）
中间件是独立的系统级软件，连接操作系统层和应用程序层，将不同操作系统提供应用的接口标准化，协议统一化，屏蔽具体操作的细节。  

##### 中间件功能
1. 通信支持         
中间件为其所支持的应用软件提供平台化的运行环境，该环境屏蔽底层通信之间的接口差异，实现互操作，所以通信支持是中间件一个最基本的功能。早期应用与分布式的中间件交互主要的通信方式为远程调用和消息两种方式。通信模块中，远程调用通过网络进行通信，通过**支持数据的转换和通信服务**，从而屏蔽不同的操作系统和网络协议。远程调用是提供给予过程的服务访问，为上层系统只**提供非常简单的编程接口或过程调用模型**。消息提供**异步交互的机制**
2. 应用支持    
中间件的目的就是**服务上层应用，提供应用层不同服务之间的互操作机制**。它为上层应用开发提供统一的平台和运行环境，并**封装不同操作系统提供API接口，向应用提供统一的标准接口**，使应用的开发和运行与操作系统无关，实现其独立性。中间件松耦合的结构，标准的封装服务和接口，有效的互操作机制，从而给应用结构化和开发方法提供有力的支持
3. 公共服务     
公共服务是对应用软件中共性功能或约束的提取。**将这些共性的功能或者约束分类实现，并支持复用，作为公共服务，提供给应用程序使用**。通过提供标准、统一的公共服务，可减少上层应用的开发工作量，缩短应用的开发时间，并有助于提高应用软件的质量

##### 中间件分类
1. 事务式中间件    
事务式中间件又称事务处理管理程序，是当前用的最广泛的中间件之一，其主要功能是提供**联机事务处理所需要的通信、并发访问控制、事务控制、资源管理、安全管理、负载平衡、故障恢复和其他必要的服务**。事务式中间件**支持**大量客户进程的**并发访问**，具有**极强的扩展性**。由于事务式中间件具有**可靠性高、极强的扩展性**等特点，主要应用于电信、金融、飞机订票系统、证券等拥有大量客户的领域。 
2. 过程式中间件    
过程式中间件又称**远程过程调用中间件**。过程中间件一般从逻辑上分为两部分：客户和服务器。客户和服务器是一个逻辑概念，既可以运行在同一计算机上，也可以运行在不同的计算机上，甚至客户和服务器底层的操作系统也可以不同。客户机和服务器之间的通信可以使用同步通信，也可以采用线程式异步调用。所以过程式中间件**有较好的异构支持能力，简单易用**，但由于客户和服务器之间采用访问连接，所以在易剪裁性和容错方面有一定的局限性。
3. 面向消息的中间件    
面向消息的中间件，简称为消息中间件，是一类**以消息为载体进行通信的中间件**，利用高效可靠的消息机制来实现不同应用间大量的数据交换。按其通信模型的不同，消息中间件的通信模型有两类：消息队列和消息传递。通过这两种消息模型，不同应用之间的**通信和网络的复杂性脱离，摆脱对不同通信协议的依赖，可以在复杂的网络环境中高可靠、高效率的实现安全的异步通信**。消息中间件的非直接连接，支持多种通信规程，达到多个系统之间的数据的共享和同步。面向消息中间件是一类常用的中间件。 
4. 面向对象中间件    
面向对象中间件又称分布对象中间件，是**分布式计算技术和面向对象技术发展的结合**，简称对象中间件。分布对象模型是面向对象模型在分布异构环境下的自然拓广。面向对象中间件给应用层提供各种不同形式的通信服务，通过这些服务，上层应用对事务处理、分布式数据访问，对象管理等处理更简单易行。OMG组织是分布对象技术标准化方面的国际组织，它制定出了CORBA等标准。  
5. Web应用服务器    
Web应用服务器是**Web服务器和应用服务器相结合的产物**。应用服务器中间件可以说是软件的基础设施，利用构件化技术将应用软件整合到一个确定的协同工作环境中，并提供多种通信机制，事务处理能力，及应用的开发管理功能。由于直接支持三层或多层应用系统的开发，应用服务器受到了广大用户的欢迎，是中间件市场上竞争的热点，J2EE架构是应用服务器方面的主流标准。 
6. 其他    
新的应用需求、新的技术创新、新的应用领域促成了新的中间件产品的出现。如，ASAAC在研究标准航空电子体系结构时提出的通用系统管理GSM，属于典型的**嵌入式航电系统的中间件，互联网云技术的发展云计算中间件、物流网的中间件**等随着应用市场的需求应运而生。 

### 什么是路由（Route）
路由是指路由器从一个接口上收到数据包，根据数据包的目的地址进行定向并转发到另一个接口的过程    
设置路由器的主要目的是找到数据包从源到目的地的最有效路径    
路由可以分为两类：静态路由和动态路由。在静态路由中，所有路由都是在一个路由器中手动设置的。因此，如果网络有任何变化，路由也不会有任何变化，除非有人手动更正它。    
在动态路由中，路由是由软件根据网络的当前状态来设置的。    
网络变化，如链路故障、流量变化等，将在每一个离散时间步更新。根据这些信息，将在每个时间步长确定新路线。动态路由优于静态路由，因为路由器会根据网络中的变化进行实时更新。    

# 概念名词

### 什么是JWT
JSON Web Token (JWT)是一个开放标准(RFC 7519)，它定义了一种紧凑的、自包含的方式，用于作为JSON对象在各方之间安全地传输信息。该信息可以被验证和信任，因为它是数字签名的。   
JWT多用于授权和信息交换   
https://zhuanlan.zhihu.com/p/86937325    

### 怎么判断用户权限
通过验证用户信息      
通过session验证     

### 如何保持用户登录状态
1. session 存储到redis
2. cookie 登陆成功之后创建登陆凭证加密后写道浏览器cookie,每次请求都发送cookie
3. 手动添加token,将token添加到http header或者做为参数添加到url    
https://www.cnblogs.com/-new/p/6960901.html   

### 创建branch的git命令
查看所有分支： git branch -a   
在本地新建一个分支： git branch branchName      
git checkout -b branchName //新创建分支并切换       
切换到你的新分支: git checkout branchName        
将新分支发布在github上： git push origin branchName     
在本地删除一个分支： git branch -d branchName     
在github远程端删除一个分支： git push origin :branchName (分支名前的冒号代表删除)      

### 什么是squash merage
把一个dev里的所有提交的commit先合并成一个再提交

### 如何回退checkin的代码
回退命令：    
git reset --hard HEAD^ 回退到上个版本     
git reset --hard HEAD~3 回退到前3次提交之前，以此类推，回退到n次提交之前   
git reset --hard commit_id 退到/进到，指定commit的哈希码（这次提交之前或之后的提交都会回滚）    

# 算法
有两个无限长的数字字符串A和B,返回二者相加的新数字字符串C    

思路：
字符串不是整数类型不能直接相加，模拟整数类型   
注意：     
1. 分别获取字符串A和B的长度，比较大小，找到最大长度      
2. 对较小长度的字符串进行补零使两字符串长度相同      
3. 模拟竖式加减思维，从最后一位依次往前递加，满位进一      
func addStrings(num1 string, num2 string) string {       
    add := 0    //进位        
    ans := ""    //新字符串      
    for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || add!= 0; i, j = i - 1, j - 1 {       
        var x, y int      
        if i >= 0 {      
            x = int(num1[i] - '0')     
        }      
        if j >= 0 {      
            y = int(num2[j] - '0')      
        }      
        result := x + y + add       
        ans = strconv.Itoa(result%10) + ans    //strconv.Itoa 把整型转成字符串并确保是如果有进位 新字符串各位数            
        add = result / 10     
    }       
    return ans       
}        
