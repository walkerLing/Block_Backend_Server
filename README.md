## 本仓库代码是该项目智能合约的开发及部署文件。本智能合约的开发基于Hyperledger Fabric，一个企业级的区块链平台，他使用智能合约作为其核心组件之一。

#### 仓库主功能文件介绍
| 文件名  | 文件功能  |
|---|---|
| model/fruit.go  |  果实溯源智能合约结构体定义  |
| fruitcc.tar.gz  | 智能合约应用打包部署的压缩包  |
| main.go  | 项目主入口  |
| go.mod，go.sum  | go module命令生成，用于引入第三方库  |
| handler/base.go  | 封装智能合约底层调用方法  |
| handler/fruit.go  | 业务逻辑代码  |
| META-INF/..  | 定义索引，方便区块链底层的数据库建立索引  |
| 连接文件  | 用户加入通道时根据登记证书生成的连接文件  |


#### 智能合约开发流程

1.确定需求和设计智能合约：
    在开始Hyperledger Fabric智能合约开发之前，您需要与客户或利益相关者沟通，了解他们的需求，并根据这些需求设计智能合约。

2.编写智能合约代码：
    使用您选择的编程语言（如Go、JavaScript等），编写智能合约代码。智能合约的编写应遵循Solidity 编写智能合约的方式。

3.测试智能合约：
    在链码开发环境中，对编写的智能合约进行测试。检查智能合约是否符合预期要求，确保代码没有错误。

4.安装智能合约：
    将智能合约安装到Hyperledger Fabric节点上。

5.实例化智能合约：
    在链码开发环境中，实例化智能合约并部署到Fabric网络上。

6.调用智能合约：
    在调用智能合约之前，您需要创建一个交易，并指定要调用的智能合约和所需参数。最后提交交易。

7.验证智能合约：
    检查智能合约是否执行正确，并且交易已被正确处理。

#### 开发摘要

    - 怎么调用到哪条区块链，还有区块链上的哪个智能合约，智能合约的哪个方法？

    联盟fruitaff和联盟内的通道channel01都是最先搭建完成的，联盟内的通道就是一条区块链。

    借助fabric的框架创建联盟名为fruitaff的再创建通道channel01，然后再在通道上安装智能合约，这是发起调用的前提。

    先根据联盟名fruitaff，和通道名channel01,获取到果实溯源区块链的连接文件，再根据果实溯源的区块链的连接文件获取到果实溯源的fabric-client，再根据请求的接口，invoke或者query,找对应智能合约的方法。