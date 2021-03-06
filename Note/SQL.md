# 1. 常见的数据结构
### 栈(Stack)
先进后出，后进先出         
Go - 用数组实现栈       
//定义一个新的类型来表达栈，需要一个数组和一个指向最后一个元素的索引       
type stack struct {       
	i int       
	data [10]int       
}       
//注意，Go的数据传递中，是值传递,但是这样不能满足我们的预期。因为一个副本被创建并传递给函数，push和pop处理的都是副本。所以提供指针       
func (s *stack) push(k int) {       
	if s.i +1 > 10 {       
		fmt.Println("I'm full")       
		return       
	}       
	s.data[s.i] = k       
	s.i++       
}       
func (s *stack) pop() int{       
	if s.i -1 < 0 {       
		fmt.Println("I'm hungry")       
		return 0       
	}       
        
	s.i--       
	result := s.data[s.i]       
	s.data[s.i] = 0       
	return result       
}       
### 队列(Queue)
先进先出的一种特殊的线性链表     
### 数组(Array)
查询速度快： 通过地址值与索引可快速定位到数据      
删除效率低： 删除数据后要将每个数据前移     
添加效率极低： 添加位置后每个数据都后移，再添加数据     
### 链表
链表中的数据都是游离存储的，每个元素节点包含元素值与下一元素的地址     
查询速度慢，因为每次查询都要通过head指针依次查询     
添加，删除效率相对较高，因为只需要将指针重新指向新添加进来的元素，其他元素的位置不需要动     
### 二叉树
全名二叉搜索树。存入的数据以第一条数据为基准，小于放左，大于放右     
特点：     
1. 只能有一个根节点，每个节点最多支持两个直接子节点     
2. 节点的度： 节点拥有的子树的个数。节点的度数不大于2，如果度数为0，则称不叶子节点或者终端节点     
缺点：     
虽然二叉树容易提高一些效率，但是面对节点多时，或者树的深度很高时，还是会面临着查找速度慢的情况，而且还很容易出现退化链表的情况（存放数据是有序的时候）。    
### 平衡二叉树 AVL Tree
平衡二叉树是在满足二叉树的情况下，尽可能的让树的度数变低，以提高查询效率     
要求： 任意节点的两个左右子树高度差不超过1，任意节点的左右子树都是一个平衡二叉树     
在底层二叉树的基础上，对进行插入和删除操作时通过特定的操作（如**左旋转和右旋转等**）保持二叉查找树的平衡，从而获得较高查找性能。     
左旋转： 被旋转的节点从左侧上升到父节点     
右旋转： 被旋转的节点从右侧上升到父节点     
https://www.cs.usfca.edu/~galles/visualization/AVLtree.html     
缺点：     
1. 树的深度过高，还是查询慢     
2. 无法解决回旋查找（批量查找）问题     
3. 添加节点效率过低，因为节点旋转有可能会牵一发而动全身     
### 红黑树
自平衡二叉查找树     
每一个节点或者是红色的或者是黑色的，根节点必须是黑色的     
如果某一个节点是红色的，那么它的子节点必须是黑色的（不能出现两个红色项相连情况）     
对每一个节点，从该节点到其所有后代叶子节点的简单路径上，均包含相同数目的黑色节点     
如果一个节点没有子节点或者父节点，则该节点的相应指针属性值为Null,这些Null视为叶节点，叶节点是黑色     
添加节点：     
添加节点的颜色可以是红色也可以是黑色     
默认使用红色效率高     
**红黑树增删改查效率都比较好**
相对于要求严格的AVL树来说，旋转次数变少，所以对于搜索，插入删除操作多的请况下使用红黑树     
缺点：     
回旋查找：     
如找比六小的数，它需要先找到六的位置然后回旋回去找父节点7，然后7还要去找5，然后一级一级找出所有比六小的数据，这样就十分缓慢       
数据过大时效率慢     
### B树 Balance Tree
多路平衡树，在树的基础上对节点进行横向的拉伸     
特点：     
所有键值分布在整棵树中（索引值和具体的data都在每个节点里）；     
任何一个关键字出现且只出现在一个结点中；     
搜索有可能在非叶子节点结束（最好情况O(1)就能找到数据）；     
在关键字全集内做一次查找，性能逼近二分查找；     
规则     
每个结点最多有m课子树（m称为阶）     
### B+树     
也是一种多路搜索树。和B树类似，B+树是B树的变体，但对B树的基础上，做了一些改变     
**一颗m阶B+树主要有这些特点**
每个结点至多有m个子女；     
**非根节点关键值个数范围： ceiling[m/2]-1 <= k <= m-1**ceiling向上取整     
相邻叶子节点是通过指针连起来的，并且是关键字大小排序的     
**B+树和B树的主要区别如下：**     
B+树内部节点是保存数据的；而B+树内部节点是不保存数据的，只做**索引**作用，它的叶子节点才保存数据。     
B+树相邻的叶子节点之间是通过链表指针连起来的，B树却不是。     
查找过程中，B树在找到具体的数值以后就结束，而B+树则需要通过索引找到叶子节点中的数据才结束     
B树中任何一个关键字出现且只出现在一个结点中，而B+树可以出现多次。     
**B+树的插入**     
B+树插入都是在叶子结点进行的，就是插入前，需要先找到要插入的叶子结点。     
如果被插入关键字的叶子节点，当前含有的关键字数量是小于阶数m，则直接插入。     
如果插入关键字后，叶子节点当前含有的关键字数目等于阶数m，则插，该节点开始「分裂」为两个新的节点，一个节点包含⌊m/2⌋ 个关键字，另外一个关键字包含⌈m/2⌉个关键值。（⌊m/2⌋表示向下取整，⌈m/2⌉表示向上取整，如⌈3/2⌉=2）。     
分裂后，需要将第⌈m/2⌉的关键字上移到父结点。如果这时候父结点中包含的关键字个数小于m，则插入操作完成。     
分裂后，需要将⌈m/2⌉的关键字上移到父结点。如果父结点中包含的关键字个数等于m，则继续分裂父结点。     
# 索引
### 作用
提高查询速度     
### 定义     
将结构化数据中的一部分信息提取出来，重新组织使其变得有一定结构，我们将信息称之为索引     
## 1. 索引分类     
### 聚集索引     
聚集索引是一种索引，该索引中键值的逻辑顺序决定了表中相应的物理顺序。聚集索引也称为聚簇索引（Cluster Index），聚集索引事物理地址连续存放的索引     
特点：只能有一个，一般为主键（主键一定是聚集索引，聚集索引并不一定是主键）     
     
什么情况下主键不是聚集索引呢？     
答：在建表的时候，并没有加主键，这个时候如果说建立了一个聚集索引，再建立主键，那么这个时候主键就不是聚集索引了     
### 非聚集索引     
非聚集索引是表中记录的物理顺序和逻辑顺序不同的索引（此外还有空间索引，筛选索引，XML索引）     
特点：可以有多个（999）     
### 索引说明     
每张表上最大的聚集索引为1     
每张表上最大的非聚集索引数为999     
每个索引最多包含的键列数为16     
索引键记录大小最多为900字节     
## 2. 索引数据结构
在SQL Server数据库中，索引的存储是以B+树（注意区分和二叉树的区别）     
结构来存储，又称索引树，其节点类型为如下两种：     
索引节点（Key）;     
叶子节点（Key + Value）；     
索引节点按照层级关系，有时又分为根节点和中间节点，其本质是一样的，都只包含下一层节点的入口值和入口指针；     
叶子节点就不同了，它包含数据，这个数据可能是表中真实的数据行，也有可能是索引列值和行书签，前者对应于聚集索引，后者对应于得聚集索引     

B+Tree     
数据页：数据库中保存数据的最小单位     
索引组织表： 一张表中有聚集索引就是索引组织表     
索引表： 一个索引对应一张索引表，索引表中每条数据都对一张数据页     

## 3. 索引为什么选择B+树
1、 B+树的磁盘读写代价更低：B+树的内部节点并没有指向关键字具体信息的指针，因此其内部节点相对B树更小，如果把所有同一内部节点的关键字存放在同一盘块中，那么盘块所能容纳的关键字数量也越多，一次性读入内存的需要查找的关键字也就越多，相对IO读写次数就降低了。     

2、B+树的查询效率更加稳定：由于非终结点并不是最终指向文件内容的结点，而只是叶子结点中关键字的索引。所以任何关键字的查找必须走一条从根结点到叶子结点的路。所有关键字查询的路径长度相同，导致每一个数据的查询效率相当。     

3、B+树更便于遍历：由于B+树的数据都存储在叶子结点中，分支结点均为索引，方便扫库，只需要扫一遍叶子结点即可，但是B树因为其分支结点同样存储着数据，我们要找到具体的数据，需要进行一次中序遍历按序来扫，所以B+树更加适合在区间查询的情况，所以通常B+树用于数据库索引。     

4、B+树更适合基于范围的查询：B树在提高了IO性能的同时并没有解决元素遍历的我效率低下的问题，正是为了解决这个问题，B+树应用而生。B+树只需要去遍历叶子节点就可以实现整棵树的遍历。而且在数据库中基于范围的查询是非常频繁的，而B树不支持这样的操作或者说效率太低。     

## 4. 索引设计原则
### 是不是索引越多越好？
肯定不行     
索引也是需要空间存储，索引太多意味着占用的空间越多。     
索引页也需要系统维护，在增，删，改数据时索引需要重新编排。就好像一本书某一页被撕掉了，对应的目录也需要进行编排。     
索引堆积久了，由于维护数据过程中会产生过多的索引碎片，反而不利于数据查询     
### 什么情况下可以建立索引（数据大于20万）     
1. 主键一定要建     
2. 外键一定要     
3. 经常查询的列     
4. 经常用作查询条件的列     
5. 经常用在order by， group by， distinct后面的列     
6. 重复值比较多的列不能建立索引     
7. 对于text， image， bit这些类型的字段不能建立索引     
8. 经常存取的列不要建立索引     
## 5. 使用索引
### 语法格式
Create [unique] [clustered / nonclustered]     
index index_name     
on table_name(column_name1, column_name2, …)     
unique: 唯一索引     
clustered: 聚集索引     
nonclustered: 非聚集索引     
index_name： 索引名称     
建立聚集索引     
create clustered index idx_userinfo_Id on UserInfo(Id);     
创建非聚集索引     
create nonclustered index idx_userinfo_account on UserInfo(Account);     
创建唯一非聚集索引     
create unique  nonclustered index idx_userinfo_Ipwd on UserInfo(pwd);     

**唯一特点： 索引字段必须唯一，但可以有一个值为Null**

查看索引     
exec sp_helpindex 'dbo.UserInfo'     
     
# 视图
## 视图的作用
提高安全性     
简化查询过程     
## 什么是视图
视图是用于简化查询过程，提高数据库安全性的数据库虚拟表对象,本质就是一堆封装好的SQL     
## 如何使用视图
## 创建视图
create view     
as select column_name from table_name[where 条件]     
     
例：     
create view v_studentscore     
as      
select  a.*,b.Degree,c.Cno,Cname,d.* from Student a     
inner join Score b on a,Sno=b,Sno     
inner join Course c on b.Cno=c.Cno     
inner join Teacher d on c.Tno-d.Tno     
     
使用视图     
select * from v_studentscore where nickname = '张旭'     
     
修改视图     
一定要记得把创建视图的代码保存起来，免得下次修改视图的时候，又要重新再写一遍     
alter view v_studentscore     
--with encryption --加密  去掉它再运行就是解密     
as     
select  a.*,b.Degree,c.Cno,Cname,d.* from Student a     
inner join Score b on a,Sno=b,Sno     
inner join Course c on b.Cno=c.Cno     
inner join Teacher d on c.Tno-d.Tno     
     
删除视图     
drop view v_studentscore     
     
# 事务与锁
create database step2_unit13;
go
use step2_unit13;
go
--创建数据表
Create TABLE account (
	id INT PRIMARY KEY identity,
	NAME VARCHAR(10),
	balance decimal(10,2)
);
添加数据
INSERT INTO account (NAME, balance) VALUES ('张三',1000)，('李四',1000)，；
## 1. 应用场景说明
什么是事务： 在实际开发过程中，一个业务操作如：转账，往往是要多次访问数据库才能完成的。转账时一个用户扣钱，另一个用户加钱，如果其中有一条SQL语句出现异常，这条SQL就有可能执行失败。事务执行是一个整体，所有的SQL语句都必须执行成功。如果其中有一条SQL语句出现异常，则所有的SQL语句都要回滚，整个业务执行失败。

模拟张三给李四转500元钱，一个转账的业务操作最少要执行下面的两条语句：
张三账号-500
update account set balance = balance - 500 where name = '张三'；
李四账号+500
update account set balance = balance + 500 where name = '李四'；
假设当张三账号上的-500元，服务器崩溃了。李四账号并没有+500，数据就出现问题了。我们需要保证其中一条SQL语句出现问题，整个转账就算失败。只有两条SQL都成功了转账才算成功。这个时候就要用到事务
**案例演示**    
begin transaction --也可以写tran,开启事务
--持久化： 磁盘的变更，数据已经保存到磁盘上
begin try --异常
	--张三账号-500
	update account set balance = balance - 500 where name = '张三'；
	--李四账号+500
	update account set balance = balance + 500 where name = '李四'；
	commit;--提交
end try
begin catch
	rollback;--回滚
end catch

总结：如果事务中SQL语句没有问题，commit提交事务，会对数据库数据的数据进行改变。如果事务中SQL语句有问题，rollback回滚事务，会回退到开启事务时的状态。

## 2. 事务的原理
事务开启之后，所有的操作都会临时保存到事务日志中，事务日志只有在得到commit命令才会同步到数据表中，其他任何情况都会清空事务日志（rollback,断开连接）
**事务的执行步骤**
1. 客户端连接数据库服务器，创建连接时创建此用户临时日志文件
2. 开启事务以后，所有的操作都会先写入到临时日志文件中
3. 所有的查询操作从表中查询，但会经过日志文件加工后才返回
4. 如果事务提交则将日志文件中的数据写到表中，否则清空日志文件。
## 3. 四大特性
原子性，一致性，隔离性，持久性
A（Atomatic）：原子性：事务是最小的执行单位，不允许分割。事务的原子性确保动作要么全部完成，要么完全不起作用；
C（Consistency）：一致性：执行事务前后，数据保持一致，例如转账业务中，无论事务是否成功，转账者和收款人的总额应该是不变的；
I（Isolation）：隔离性：并发访问数据库时，一个用户的事务不被其他事务所干扰，各并发事务之间数据库是独立的；
D（Durability）：持久性： 一个事务被提交之后。它对数据库中数据的改变是持久的，即使数据库发生故障也不应该对其有任何影响。
### 1. 事物的隔离级别
事务在操作时的理想状态： 所有的事务之间保持隔离。互不影响。因为并发操作，
多个用户同时访问同一个数据。可能引发并发访问的问题：
|  并发访问问题     | 含义  |
|  ----  | ----  |
| 脏读  | 一个事务中读取到了另一个事务中尚未提交的数据 |
| 不可重复读  | 一个事务中两次读取的数据内容不一致，要求的是一个事务中多次读取的数据是一致的，这是事务update时引发的问题，针对update |
| 幻读 | 一个事务中两次读取的数据的数量不一致，要求在一个事务多次读取的数据的数量是一致，针对insert |

**数据库有四种隔离级别** 高频问题
SQL 标准定义了四个隔离级别：

READ-UNCOMMITTED(读取未提交)： 最低的隔离级别，允许读取尚未提交的数据变更，可能会导致脏读、幻读或不可重复读。
READ-COMMITTED(读取已提交)： 允许读取并发事务已经提交的数据，可以阻止脏读，但是幻读或不可重复读仍有可能发生。
REPEATABLE-READ(可重复读)： 对同一字段的多次读取结果都是一致的，除非数据是被本身事务自己所修改，可以阻止脏读和不可重复读，但幻读仍有可能发生。
SERIALIZABLE(可串行化)： 最高的隔离级别，完全服从 ACID 的隔离级别。所有的事务依次逐个执行，这样事务之间就完全不可能产生干扰，也就是说，该级别可以防止脏读、不可重复读以及幻读。    

|  级别   | 名称  | 隔离级别  | 脏读  | 不可重复读  | 幻读  |
|  ----  |  ----  |  ----  |  ----  |  ----  |  ----  |
| 1  |  读未提交 |  read uncommitted  | 是  | 是 | 是 |
| 2  |  读已提交 |  read committed | 否 | 是  | 是 |
| 3  |  可重复读 |  repeatable read |  否 | 否 | 是 |
| 4  |  串行化 | serializable | 否 | 否 | 否 |
| 2-3之间  |  快照 | SNAPSHOT  | 否 | 否 | 否 |
      
隔离级别越高，性能越差，安全性越高。      
### 2. 事务的快照
SNAPSHOT快照：SNAPSHOT和READ COMMITTED SNAPSHOT两种隔离（可以把事务已经提交的行的上一版本保存在**TEMPDB**数据库中）
SNAPSHOT隔离级别在逻辑上与SERIALIZABLE类似
READ COMMITTED SNAPSHOT隔离级别在逻辑上与 READ COMMITTED类似
不过在快照隔离级别下读操作不需要申请获得共享锁，所以即便是数据已经存在排他锁也不影响读操作。而且仍然可以得到和SERIALIZABLE与READ COMMITTED隔离级别类似的一致性；如果目前版本与预期的版本不一致，读操作可以从TEMPDB中获取预期的版本。

如果启用任何一种基于快照的隔离级别，DELETE和UPDATE语句在做出修改前都会把行的当前版本复制到TEMPDB中，而INSERT语句不需要在TEMPDB中进行版本控制，因为此时还没有行的旧数据

undo.log
|  事务ID   | 数据ID  | 数据版本  |
|  ----  | ----  | ----  |
| 1  | 1 | 1（快照） |
| 2  | 1 | 2（如果版本不一致，则事务会获取版本1那条数据，也就是快照） |
       
作用：
1. 读操作时不会陷入block和死锁的问题中，SNAPSHOT本身提高了数据库系统的事务处理的性能。
2. 避免了脏读，非一致性读，以及丢失更新，和不可重复读等多个问题
## 4. 隔离级别演示
查看当前的隔离级别
DBCC USEROPTIONS
### 1.脏读的演示
设置隔离级别为read uncommit;
--set transaction isolation leve1 <隔离级别>
先将两个账号的金额都恢复成1000元。
update account set balance=1000;
1. 打开A窗口登录，设置隔离级别为最低
set transaction isolation level1 read uncommitted
2. 打开B窗口，AB窗口都开启事务
set transaction isolation level1 read uncommitted
3. A窗口更新2个人的账户数据，未提交
--张三账号-500
update account set balance = balance - 500 where name = '张三'
--李四账号+500
update account set balance = balance + 500 where name = '李四'
4. B窗口查询账户
张三 500 李四 1500
5. A窗口回滚
6. 查询B窗口
张三 1000 李四 1000
脏读非常危险，比如张三向李四购买商品，张三开启事务，向李四账号转入500块，然后打电话给李四说钱已经转了。李四一查询钱到账了，发货给张三，张三抽到货之后回滚事务，不给钱了
**解决脏读问题**：将全局隔离级别进行提升
将数据进行恢复：
UPDATE account SET balance = 1000；
1. A窗口设置隔离级别为： repeatable read
set tran isolation level repeatable read;
2. B窗口开启事务，查询数据，先不提交
begin tran;
select * from account;
3. 在A窗口开启事务，并更新数据，此时你会发现，事务会一直处于等待状态，无法提交，直到B窗口中的事务提交完毕。
