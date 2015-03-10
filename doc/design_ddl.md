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
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`user_id`),
	INDEX(`email`),
	INDEX(`user_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
##### 用户注销账号是只做逻辑删除，即标记status为1。
INSERT INTO user(user_code,password,display_name,picture,email,phone,user_authority,file_path,status) VALUES('119427f6aed6fbce53eaadaaa5519317','E10ADC3949BA59ABBE56E057F20F883E','管理员',"478e3dd1616187541b6dcc4e82865133","admin@mail.com","13512341234","admin","null",0);#插入一条管理员用户测试


### 4.2 项目管理--项目列表
CREATE TABLE `project` (
	`project_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_code` CHAR(32) NOT NULL,
	`project_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`project_name` VARCHAR(255) NOT NULL,#项目名称
	`picture` MEDIUMTEXT NOT NULL,#项目缩略图的base64编码
	`project_leader` VARCHAR(100) NOT NULL,#项目负责人
	`project_type` VARCHAR(100) NOT NULL,#项目类型
	`start_datetime` TIMESTAMP,#项目开始时间
	`end_datetime` TIMESTAMP,#项目结束时间
	`project_detail` VARCHAR(1000) NOT NULL,#项目描述（备注）
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
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
	`dpx_path` VARCHAR(1000),#DPX路径
	`jpg_path` VARCHAR(1000),#Jpg路径
	`mov_path` VARCHAR(1000),#Mov小样路径
	`user_code` CHAR(32) NOT NULL,#用户代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`library_id`),
	INDEX(`library_code`)
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
	`user_code` CHAR(32) NOT NULL,#用户代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`material_id`),
	INDEX(`material_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.3 素材管理--用户添加的素材组
CREATE TABLE `material_folder`(
	`folder_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`folder_name` VARCHAR(255) NOT NULL,#素材组名
	`father_code` CHAR(32) NOT NULL,#上级代码
	`leaf_flag` CHAR(1) NOT NULL,#子节点标志，0否，1是
	`folder_detail` VARCHAR(1000) NOT NULL,#描述
	`user_code` CHAR(32) NOT NULL,#用户代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`folder_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

### 4.3 素材管理--用户添加的素材组数据
CREATE TABLE `material_folder_data`(
	`data_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`data_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`folder_id` INT UNSIGNED NOT NULL,#素材组的ID
	`material_code` CHAR(32) NOT NULL,#素材的代码
	`user_code` CHAR(32) NOT NULL,#用户代码
	`project_code` CHAR(32) NOT NULL,#项目代码
	`status` TINYINT UNSIGNED NOT NULL,#状态0代表正常，1代表已注销
	`insert_datetime` TIMESTAMP,
	`update_datetime` TIMESTAMP,
	PRIMARY KEY (`data_id`),
	INDEX(`data_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


每个shot是material的一段，该表由EDL文件解析获得，EDL解析的程序需要自己写
CREATE TABLE `shot` (
    `shot_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `shot_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `project_code` CHAR(32) NOT NULL,#　当前的项目code
    `material_code` CHAR(32) NOT NULL ,#计算生成的唯一识别符
    `shot_num` CHAR(11) NOT NULL,# EDL文件中数量号,作用未知，001,002
    `start_time` CHAR(11) NOT NULL,#该shot的入点，格式为00:00:00:00
    `end_time` CHAR(11) NOT NULL,#该shot的出点，格式为00:00:00:00
    `clip_name` VARCHAR(32) NOT NULL,#
    `soure_file` VARCHAR(32) NOT NULL,#该shot 的源文件名字
    `shot_type` VARCHAR(32) NOT NULL,
　　　　`shot_name` VARCHAR(32) NOT NULL,# 可以自定义修改的镜头名字
　　　　`shot_fps` SMALLINT NOT NULL,
　　　　`width` SMALLINT NOT NULL,#尺寸宽，格式为1920*1080
　　　　`height` SMALLINT NOT NULL,#尺寸高，格式为1920*1080
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `descrube` VARCHAR(500) NOT NULL,#镜头的描述
    `requment` CHAR(32),#制作需求列表
    `reference` CHAR(32),#　素材参照表
    `produce` CHAR(32),#　制作素材
    `version` CHAR(32),#　版本
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`shot_id`),
	INDEX(`field_name`),
	INDEX(`project_code`)
	INDEX(`shot_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

存储shot（镜头）的缩略图，一个镜头可能有多个缩略图
CREATE TABLE `thumbnail` (
    `thumbnail_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `thumbnail_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `shot_code` CHAR(32) NOT NULL,#对应的镜头
    `thumbnail_image` BLOB NOT NULL,#最大64K
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`thumbnail_id`),
	INDEX(`shot_code`),
	INDEX(`thumbnail_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
存储shot（镜头）的左侧列表层级关系表
CREATE TABLE `relation` (
    `relation_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `relation_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `parent_code` CHAR(32) NOT NULL,#上一层级的code
   `child_code` CHAR(32) NOT NULL,#　当前的项目code
    `list_name` VARCHAR(32) NOT NULL,#　当前层级的名字
    `isShot` BOOLEAN DEFAULT false, # 辅助区分是镜头还是分组的标题
    `content` VARCHAR(100),#描述
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`relation_id`),
	INDEX(`parent_code`),
	INDEX(`relation_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

接下来讲需求发给若干个接包方
CREATE TABLE `vendor` (
    `vendor_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `vendor_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `project_code` CHAR(32) NOT NULL,
    `vendorName` CHAR(32) NOT NULL,#接包方用户名
 `open_detail` BOOLEAN DEFAULT false,
`open_demo` BOOLEAN DEFAULT false,
`down_material` BOOLEAN DEFAULT false,
`up_demo` BOOLEAN DEFAULT false,
`up_product` BOOLEAN DEFAULT false,
    `content` VARCHAR(2047) NOT NULL,#对于制作的具体说明
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`vendor_id`),
	INDEX(`shot_code`),
	INDEX(`vendor_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

接包方处理完之后上传到文件夹并在Daily里面进行记录
CREATE TABLE `note` (
    `note_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `note_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `shot_code` CHAR(32) NOT NULL,#对应的镜头
    `award_code` CHAR(32) ,#对应的哪个包
　`content` VARCHAR(100),#讨论的内容
`thumbnail`BLOB NOT NULL,#最大64K　　存储图片
    `target_path` VARCHAR(2047) NOT NULL,#目标路径
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`note_id`),
	INDEX(`shot_code`),
	INDEX(`note_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

权限管理

制作需求表
CREATE TABLE `requment`(
 `requment_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `requment_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
	`shot_code` CHAR(32) NOT NULL,#对应的镜头
	`important` INT,#优先级
	 `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
	`thumbnail`BLOB NOT NULL,#最大64K　　存储图片
 `content` VARCHAR(100),
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`requment_id`),
	INDEX(`shot_code`),
	INDEX(`requment_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
版本信息
CREATE TABLE `version`(
 `version_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `version_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `shot_code` CHAR(32) NOT NULL,#对应的镜头
    `vendor_code` CHAR(32) NOT NULL,#对应的哪个包
	`thumbnail`BLOB NOT NULL,#最大64K　　存储图片
	`demo` VARCHAR(5000) ,＃也应该存储路径吗？？？？
	`product` VARCHAR(200), #存储路径
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`version_id`),
	INDEX(`shot_code`),
	INDEX(`version_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `reference`(
 `reference_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `reference_code` CHAR(32) NOT NULL UNIQUE,#计算生成的唯一识别符
    `shot_code` CHAR(32) NOT NULL,#对应的镜头
`reference_type` CHAR(32) DEFAULT '参考素材'，
   `mat_name` CHAR(32) ,#素材名
`mat_type` CHAR(32) NOT NULL,＃格式
`mat_content` VARCHAR(200),＃描述
`mat_url` VARCHAR(1000) NOT NULL,＃路径
    `status` TINYINT UNSIGNED NOT NULL,#0代表正常，只做逻辑删除，即标记status为1
    `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`reference_id`),
	INDEX(`shot_code`),
	INDEX(`reference_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

