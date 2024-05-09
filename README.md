# “小胰宝 - 面向胰腺癌肿瘤患者的智能RAG平台”

### 【项目介绍】

小胰宝是由病友SamQin（本人）开发并上线的RAG应用，主旨为帮助每年**数十万**新增确诊病人以及幸存病人，克服专业医学和治疗的医患信息差。使用小胰宝，可7x24小时，帮助病患高效率，准确理解病情状态，治疗术语，规范治疗指南，以及综合治疗的复杂信息，克服慌乱情绪，选择科学和有效治疗路线，并最终获得更长的治疗收益，也即生命窗口。

小胰宝已经上线7个月，帮助了2000+患者稳定提供，并聚集了支持AI科技应用的专业医生，患者群，以及开源技术社区和社会公益人士。小胰宝的上线，为全癌种肿瘤患者AI智能服务，开辟和验证了新模式，期待开源社区能够推动为数千万新患者，提供更多类似应用，服务包括肺癌/乳腺癌/甲状腺/肝癌/肠癌等癌肿的患者，获得信息便利。


<img src="https://www.freeimg.cn/i/2024/05/08/663ae77288263.png" alt="图片描述文字" />


### 体验地址
https://aibot.samqin.asia 

![小胰宝新手宝典](https://github.com/samqin123/---RAG-/assets/103937568/4e1b2e11-38fe-4978-9660-d2761c2a0af1)

### 【项目开源原则和誓言】
使用小胰宝项目开源框架，包括基于本项目技术框架，产品框架，知识库数据全部或部分，默认自动接受如下原则和誓言：

1.  ### 永不向病患和家属收费，不以“智能助手”做任何商业化引流。
2.  ### 尊重知识贡献者所有权，不盗取知识库获利，公开知识库来源。
3.  ### 一切立足“规范治疗”，“循证”为知识库准入第一原则。

----

### 【小胰宝使用的产品框架介绍】

待更新

----
### 【小胰宝项目1.0版本使用开源组件】

小胰宝使用如下开源项目，也感谢这些项目大佬的付出，他们为患者创造了科技温情的价值。

1. Fastgpt: 
2. Chatgpt-on-wechat：
3. Dify：
4. LLM：零一万物 提供了LLM能力的接口及Token消耗赞助，让数千患者享受到了科技福利
5. 也感谢其它为小胰宝上线提供支持的开源社区人士 @v佬 @Richard 


---
### 【版本介绍和路线图】
V1.0 

#### 【功能介绍】
##### 1. 知识库构建：
###### 1.1 基于“规范治疗”原则，精心挑选和导入了胰腺癌规范治疗指南（包括中国CSCO/美国NCCN胰腺癌诊疗指南和患者指南），同时覆盖20+胰腺癌并发症和综合治疗相关指南（白细胞/血小板/梗阻/黄疸介入等）
###### 1.2 导入胰腺病友群积累的综合治疗经验宝典内容
###### 1.3 覆盖靶向治疗和基因治疗，和胰腺癌相关的基因定义，以及基因治疗中的注意要点知识
###### 1.4 导入营养支持规范和治疗
###### 1.5 精选和导入综合治疗的专题文章

##### 2. Agent构建
###### 2.1 Agent创建
###### 2.2 Prompt配置
###### 2.3 LLM配置和选择
###### 2.4 Tools工具匹配
###### 2.5 对话测试
###### 2.6 对话模版管理

##### 3. 服务管理
###### 3.1 API创建和管理
###### 3.2 外部使用链接创建和管理
###### 3.3 服务统计和看板

##### 4. 系统管理
###### 4.1 用户权限/角色管理
###### 4.2 日志管理


##### 5. 微信端管理
###### 5.1 微信机器人接入和自动化监测
###### 5.2 Bot对话配置管理
###### 5.3 其它应用管理


##### 6. 企业微信端管理
###### 6.1 企业微信应用接入
###### 6.2 Bot对话配置管理
###### 6.3 其它应用功能管理

##### 7. 运营体系是不可或缺的部分，对于知识库的运营，包括
###### 7.1. 知识体系设计
###### 7.2. 知识体系评价和筛选分级
###### 7.3. 知识库结构设计
###### 7.4. 知识内容导入和校验
###### 7.5. 知识库的优化，压缩和迭代


---

### V2.0版本

预计发布功能：
##### 2.1. 图像识别：能够实现监测报告上传识别内容，对于病情指标的定义和异常进行智能解读，便于病人掌握指标意义，及时和主治医生进行互动；
##### 2.2. 音视频知识库
##### 2.3. 语音互动：增加稳定的语音互动能力，便于年长者与小胰宝方便进行相关的沟通
##### 2.4  新版本知识库体系
  ##### 2.4.1 胰腺癌综合治疗药物
  ##### 2.4.2 胰腺癌综合治疗营养菜谱
##### 2.5 稳定的微信机器人：**急需开源社区的技术贡献者帮助**

---
### 【开源合作招募】
1. 微信bot：
很多患者群聚集在微信生态，因此需要稳定的微信机器人，目前使用的免费开源工具包括Chatgpt-on-wechat, 存在2-3天掉线问题，因此严重影响病患的使用和小胰宝及类似应用服务十万级用户体量的要求。开源社区如有支持者，请联系本人。 
2. 知识库迭代
3. LLM能力增强
4. 诊疗协同
5. 公益组织协同

### 【联络方式】
作者联系邮箱  qinxiaoqiang2002@hotmail.com,  微信 qinxiaoqiang 
小胰宝公益组织联系方式【待更新】
