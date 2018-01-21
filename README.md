# MeritMS

Demo(http://112.74.42.44:8080/)

成果、价值、进度管理系统
1. 本系统是一款在线成果（工作量）登记、成果校审流程、价值档案、价值评测等管理系统，理念是尽量减少占用技术人员时间去进行工作量登记，但又能符合管理者对大数据的需求；技术人员基于MeritMS在线动态维护自己的价值档案，充分展示自己的价值。本系统采用go语言开发，运行文件为编译后的二进制exe文件，所以无需像其他语言（php、nodejs、java等语言）编写的web应用那样，需要配置运行服务环境。 本系统既可以运行于工程师个人电脑，也可以放到服务器上运行，仅运行exe文件即可实现网络化管理项目和人员资料，免维护，轻量，开源，功能齐全，采用大量开源的先进插件，是工程设计管理者不可或缺的工具。

2. 数据库采用sqlite嵌入式数据库，所以也无需配置数据库服务环境。

3. 开箱即用，无需网络开发知识。


## 特性

1. 按照 编制（制图）、设计、校核、审查等角色和权限进行成果校审流程管理。特点：批量选择——批量填写（修改）校审人员（校审意见）——批量提交——系统根据用户名和他所在位置（校核位、审查位）自动走流程。比projectwise的校审流程方便很多，见下图对比表。
2. 后台定制组织架构：部门——科室——人员 或 部门——人员。
3. 整合了EngineerCMS，可以集中管理项目资料。
4. 技术人员无论是向MeritMS中上传成果或在个人EngineerCMS中上传 成果，都可以自动生成 成果清单，提交给MeritMS进行工作量统计。
5. 基于IP权限和登录权限管理，IP权限快捷，无需登录；登录权限用于远程访问。可定制系统管理员、部门管理者、科室管理者和技术人员等角色权限。
6. 基于科室、部门不同组织级别的成果、项目、贡献、排名等统计展示。
7. 分技术人员和系统管理员2种批量添加 成果列表，后者无需经过层层校审流程，直接进入统计。
8. 成果校审流程中提供附件链接和校审意见填写。
9. 后台定制组织结构里的价值分类和价值列表。
10. 技术人员在系统中动态维护自己的价值：注册、获奖、负责人、科研……。
11. 基于科室、部门不同组织级别的技术人员价值排名。
12. 为尽最大努力简化技术人员添加成果的工作量，提供了5种添加方式：1、手工添加；2、自己导入excel表格中成果清单；3、交给管理员导入excel中的成果清单，省去校审流程；4、想MeritMS中上传成果，自动生成成果清单，可批量进行修改，批量进行提交；5、在EngineerCMS中上传成果，自动生成成果清单，可批量进行修改，批量进行提交。

## 下载、安装

在release标签中下载二进制文件，直接运行exe文件。

## Quick Start

* 参见quickstart快速开始。

## Documentation

* [中文文档]——请查阅document文件夹

## 免费开源和问题反馈

* 开源地址[https://github.com/3xxx/merit/](https://github.com/3xxx/merit/)
* 问题反馈: [https://github.com/3xxx/merit/issues](https://github.com/3xxx/merit/issues)

## LICENSE

Merit source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).

1.成果在线登记
![default](https://user-images.githubusercontent.com/10678867/35192561-dd749b32-fecf-11e7-8b58-03a0e037e3c8.png)

2.与project wise校审流程对比
![](https://user-images.githubusercontent.com/10678867/33466124-e8c03198-d686-11e7-8359-b394064e6c39.png)

3.简洁的校审流程设计
![](https://user-images.githubusercontent.com/10678867/33466220-7401580e-d687-11e7-84f0-758e3c4be83d.png)

![](https://user-images.githubusercontent.com/10678867/33466218-73d4e562-d687-11e7-971c-47d6613d1699.png)

4.价值评测项目设置
![snap15](https://user-images.githubusercontent.com/10678867/35192570-1602be52-fed0-11e7-84a3-480c7c64e26a.png)

5.给部门选择价值评测项目
![snap16](https://user-images.githubusercontent.com/10678867/35192571-1c745502-fed0-11e7-87d0-0b4ef514ce21.png)

6.价值登记、审核
![snap18](https://user-images.githubusercontent.com/10678867/35192567-085a6282-fed0-11e7-9b31-009e2320013f.png)

7.价值统计
![snap19](https://user-images.githubusercontent.com/10678867/35192563-f509b2a0-fecf-11e7-9c8c-b706bc9f2713.png)

8.项目进度展示
![](https://cloud.githubusercontent.com/assets/10678867/25748417/8dc743b0-31dd-11e7-920f-8a54f7e5b23d.png)

带图标的select2设计

![](https://user-images.githubusercontent.com/10678867/31264191-fe325a56-aa99-11e7-9689-5cc1c130de85.gif)
