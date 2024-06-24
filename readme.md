# 达梦一键导出导入工具

## 一、简介
由于国产数据库命令行参数不容易让人记住。为了让开发人员导出导入便利，故开发了一键导出导入工具

`使用条件:`

客户端与数据库服务器网络连通

## 二、安装
解压压缩包到任意目录即可
![img.png](/image/img.png)

## 三、一键导出
1、进入`dm-tool`目录编辑配置文件

使用文本编辑工具如`Notepad++`打开`settings.yaml`文件

在以下填入正确的信息，注意请保持yaml文件格式的缩进

![img_1.png](/image/img_1.png)


2、执行导出

点击`start_export.bat`一键导出

![img_2.png](/image/img_2.png)


输入yes确认导出

![img_3.png](/image/img_3.png)
![](media/17007300990756/17007305540038.jpg)

导出完毕

![img_4.png](/image/img_4.png)

导出的备份默认在当前执行程序所在目录

## 四、一键导入
1、进入`dm-tool`目录编辑配置文件

使用文本编辑工具如`Notepad++`打开`settings.yaml`文件

在以下填入正确的信息，注意请保持yaml文件格式的缩进

![img_1.png](/image/img_1.png)

2、执行导入

点击`start_import.bat`

![img_5.png](/image/img_5.png)!

在cmd窗口内根据提示，输入备份文件的绝对路径，例如`D:\test2-2023_11_23_16_59_49.dmp`

![img_7.png](/image/img_7.png)

注意:请核对数据库信息之后，输入`yes`开始导入

![img_9.png](/image/img_9.png)

导入完毕

![img_10.png](/image/img_10.png)
