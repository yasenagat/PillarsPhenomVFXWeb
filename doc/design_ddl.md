#<center>天工异彩项目设计</center>

##用户系统

CREATE DATABASE IF NOT EXISTS PillarsPhenomVFX DEFAULT CHARSET utf8;

### 4.1 用户管理--用户列表
CREATE TABLE `user` (
	`user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`password` CHAR(32) NOT NULL,#密码
	`display_name` CHAR(20) NOT NULL,#显示用户名
	`picture` MEDIUMTEXT NOT NULL,#头像照片的base64编码
	`email` CHAR(30) NOT NULL UNIQUE,#登陆账号
	`phone` CHAR(20) NOT NULL,#电话
	`user_authority` CHAR(20) NOT NULL,#用户权限（admin，制片，助理，分包商）
	`file_path` VARCHAR(2047) NOT NULL,#用户存储路径
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`user_id`),
	INDEX(`email`),
	INDEX(`user_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
##### 用户注销账号是只做逻辑删除，即标记status为1。
INSERT INTO user(user_code,password,display_name,picture,email,phone,user_authority,file_path,status) VALUES('119427f6aed6fbce53eaadaaa5519317','E10ADC3949BA59ABBE56E057F20F883E','管理员',"478e3dd1616187541b6dcc4e82865133","admin@mail.com","13512341234","admin","null",0);#插入一条管理员用户测试


### 4.2 项目管理--项目列表
CREATE TABLE `project` (
	`project_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`project_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`project_name` VARCHAR(255) NOT NULL,#项目名称
	`picture` MEDIUMTEXT NOT NULL,#项目缩略图的base64编码
	`project_leader` VARCHAR(100) NOT NULL,#项目负责人
	`project_type` VARCHAR(100) NOT NULL,#项目类型
	`start_datetime` TIMESTAMP,#项目开始时间
	`end_datetime` TIMESTAMP,#项目结束时间
	`project_detail` VARCHAR(1000) NOT NULL,#项目描述（备注）
	`user_code` CHAR(32) NOT NULL,
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`project_id`),
	INDEX(`project_name`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


### 4.3 素材管理--用户添加的库
CREATE TABLE `library`(
	`library_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`library_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`library_name` VARCHAR(255) NOT NULL,#库的名称
	`library_path` VARCHAR(1000) NOT NULL,#库的原始路径
	`dpx_path` VARCHAR(1000) NOT NULL,#DPX路径
	`jpg_path` VARCHAR(1000) NOT NULL,#Jpg路径
	`mov_path` VARCHAR(1000) NOT NULL,#Mov小样路径
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`library_id`),
	INDEX(`library_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.3 素材管理--素材信息
CREATE TABLE `material` (
	`material_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`library_code` CHAR(32) NOT NULL,#库的代码
	`material_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`material_name` VARCHAR(255) NOT NULL,#素材的名称
	`material_type` VARCHAR(20) NOT NULL,#文件格式（R3D）
	`material_path` VARCHAR(2047) NOT NULL,#素材的Base路径，例:library_path+Base+material_type
	`video_track_count` INT NOT NULL,#轨数
	`width` INT NOT NULL,#尺寸宽，格式为1920*1080
	`height` INT NOT NULL,#尺寸高，格式为1920*1080
	`video_audio_framerate` INT NOT NULL,#视频音频帧速率
	`timecode_framerate` INT NOT NULL,#时间线帧速率
	`video_frame_count` INT NOT NULL,#帧数
	`start_absolute_timecode` VARCHAR(20) NOT NULL,#绝对始码，格式为00:00:00:00
	`end_absolute_timecode` VARCHAR(20) NOT NULL,#绝对止码，格式为00:00:00:00
	`start_edge_timecode` VARCHAR(20) NOT NULL,#始码，格式为00:00:00:00
	`end_edge_timecode` VARCHAR(20) NOT NULL,#止码，格式为00:00:00:00
	`meta_data` text NOT NULL,#元数据，string JSON
	`picture` MEDIUMTEXT NOT NULL,#缩略图，Base64编码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`material_id`),
	INDEX(`material_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.3 素材管理--用户自定义素材组
CREATE TABLE `material_folder`(
	`folder_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`folder_name` VARCHAR(255) NOT NULL,#素材组名
	`father_code` CHAR(32) NOT NULL,#上级代码
	`leaf_flag` CHAR(1) NOT NULL,#子节点标志，0否，1是
	`folder_detail` VARCHAR(1000) NOT NULL,#描述
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`folder_id`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.3 素材管理--用户添加的素材组数据
CREATE TABLE `material_folder_data`(
	`data_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`data_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`folder_id` INT UNSIGNED NOT NULL,#素材组的ID
	`material_code` CHAR(32) NOT NULL,#素材的代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`data_id`),
	INDEX(`data_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


### 4.4 视效镜头管理(Post)--镜头,每个shot是material的一段
CREATE TABLE `shot` (
	`shot_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`shot_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`project_code` CHAR(32) NOT NULL,#当前的项目code
	`material_code` CHAR(32) NOT NULL,#soure_file对应的素材名称的素材code
	`library_code` CHAR(32) NOT NULL,#素材对应的库code
	`picture` MEDIUMTEXT NOT NULL,#镜头缩略图根据始码抓取素材中的一帧
	`shot_num` CHAR(11) NOT NULL,# EDL文件中数量号,作用未知，001,002
	`start_time` CHAR(11) NOT NULL,#该shot的入点，格式为00:00:00:00
	`end_time` CHAR(11) NOT NULL,#该shot的出点，格式为00:00:00:00
	`from_clip_name` VARCHAR(100) NOT NULL,#剪辑文件名称
	`source_file` VARCHAR(100) NOT NULL,#对应的素材名称
	`shot_type` VARCHAR(32) NOT NULL,#镜头类型
	`shot_name` VARCHAR(100) NOT NULL,# 可以自定义修改的镜头名字
	`shot_fps` SMALLINT NOT NULL,#帧速率
	`width` SMALLINT NOT NULL,#尺寸宽，格式为1920*1080
	`height` SMALLINT NOT NULL,#尺寸高，格式为1920*1080
	`shot_detail` VARCHAR(500) NOT NULL,#镜头的描述
	`shot_status` VARCHAR(100) NOT NULL,#镜头的状态
	`edl_code` CHAR(32) NOT NULL,#edl文件生成的code
	`edl_name` VARCHAR(100) NOT NULL,#edl文件名称
	`shot_flag` VARCHAR(10) NOT NULL,#edl(0)或者手动添加(1)的标识
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`shot_id`),
	INDEX(`shot_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--镜头的制作需求列表
CREATE TABLE `shot_demand` (
	`demand_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`demand_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`shot_code` CHAR(32) NOT NULL,#对应的镜头
	`project_code` CHAR(32) NOT NULL,#项目code
	`picture` MEDIUMTEXT NOT NULL,#Base64编码图片
	`demand_detail` VARCHAR(100) NOT NULL,#需求的内容描述
	`demand_level` TINYINT UNSIGNED NOT NULL,#需求的优先级
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`demand_id`),
	INDEX(`demand_code`),
	INDEX(`shot_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--镜头的参考素材
CREATE TABLE `shot_material` (
	`material_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`material_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`shot_code` CHAR(32) NOT NULL,#对应的镜头
	`project_code` CHAR(32) NOT NULL,#项目code
	`picture` MEDIUMTEXT NOT NULL,#Base64编码图片
	`material_name` VARCHAR(100) NOT NULL,#参考素材名称
	`material_type` VARCHAR(20) NOT NULL,#参考素材格式
	`material_detail` VARCHAR(500) NOT NULL,#参考素材的内容描述
	`material_path` VARCHAR(500) NOT NULL,#参考素材的保存路径
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`material_id`),
	INDEX(`material_code`),
	INDEX(`shot_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--镜头的相关沟通信息
CREATE TABLE `shot_note` (
	`note_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`note_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`shot_code` CHAR(32) NOT NULL,#对应的镜头
	`project_code` CHAR(32) NOT NULL,#项目code
	`picture` MEDIUMTEXT NOT NULL,#Base64编码图片
	`note_detail` VARCHAR(1000) NOT NULL,#讨论的内容
	`note_type` VARCHAR(10) NOT NULL,#分类
	`note_verson` VARCHAR(100) NOT NULL,#版本
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`note_id`),
	INDEX(`note_code`),
	INDEX(`shot_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--用户自定义镜头组
CREATE TABLE `shot_folder`(
	`folder_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`folder_name` VARCHAR(255) NOT NULL,#镜头目录名
	`father_code` CHAR(32) NOT NULL,#上级代码
	`leaf_flag` CHAR(1) NOT NULL,#子节点标志，0否，1是
	`folder_detail` VARCHAR(1000) NOT NULL,#目录描述
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`folder_id`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--用户自定义镜头组数据
CREATE TABLE `shot_folder_data`(
	`data_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`data_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`folder_id` INT UNSIGNED NOT NULL,#镜头目录的ID
	`shot_code` CHAR(32) NOT NULL,#镜头的代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`data_id`),
	INDEX(`data_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--镜头外包列表,分配镜头给外包商
CREATE TABLE `shot_vendor` (
	`vendor_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`vendor_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`project_code` CHAR(32) NOT NULL,#项目代码
	`vendor_user` CHAR(32) NOT NULL,#指定的外包商
	`vendor_name` VARCHAR(100) NOT NULL,#外包列表名
	`vendor_detail` VARCHAR(1000) NOT NULL,#描述
	`open_detail` TINYINT(1) DEFAULT 0,#可查看信息0否1是
	`open_demo` TINYINT(1) DEFAULT 0,#可查看小样
	`down_material` TINYINT(1) DEFAULT 0,#可下载素材
	`up_demo` TINYINT(1) DEFAULT 0,#可上传小样
	`up_product` TINYINT(1) DEFAULT 0,#可上传成品
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`vendor_id`),
	INDEX(`vendor_code`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.4 视效镜头管理(Post)--分配给外包商的镜头列表
CREATE TABLE `shot_vendor_data`(
	`data_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`data_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`vendor_code` CHAR(32) NOT NULL,#镜头外包列表code
	`vendor_user` CHAR(32) NOT NULL,#指定的外包商
	`shot_code` CHAR(32) NOT NULL,#镜头的代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`data_id`),
	INDEX(`data_code`),
	INDEX(`vendor_code`),
	INDEX(`vendor_user`),
	INDEX(`project_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 5 外包商--上传小样和成品(生成版本)
CREATE TABLE `shot_version`(
	`version_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`version_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`shot_code` CHAR(32) NOT NULL,#镜头代码
	`vendor_user` CHAR(32) NOT NULL,#外包商代码
	`version_num` TINYINT NOT NULL,#版本序号
	`picture` MEDIUMTEXT NOT NULL,#Base64编码图片
	`demo_name` VARCHAR(1024) NOT NULL,#小样名称
	`demo_type` VARCHAR(1024) NOT NULL,#小样类型
	`demo_path` VARCHAR(1024) NOT NULL,#小样存储路径
	`demo_detail` VARCHAR(1000) NOT NULL,#小样的内容描述
	`product_name` VARCHAR(1024) NOT NULL,#成品名称
	`product_type` VARCHAR(1024) NOT NULL,#成品类型
	`product_path` VARCHAR(1024) NOT NULL,#成品存储路径
	`product_detail` VARCHAR(1000) NOT NULL,#成品的内容描述
	`status` TINYINT UNSIGNED NOT NULL,#数据有效性(0代表正常，删除标记为1)
	`insert_datetime` TIMESTAMP NOT NULL,
	`update_datetime` TIMESTAMP NOT NULL,
	PRIMARY KEY (`version_id`),
	INDEX(`version_code`),
	INDEX(`shot_code`),
	INDEX(`version_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
