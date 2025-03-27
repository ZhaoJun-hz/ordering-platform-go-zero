CREATE DATABASE `ordering-platform` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `sys_menu`
(
    `menu_id`           bigint                                       NOT NULL AUTO_INCREMENT  COMMENT '主键编码',
    `parent_id`         bigint unsigned DEFAULT '0' NOT NULL COMMENT '父菜单ID',
    `sort`              int unsigned NOT NULL DEFAULT '1' COMMENT '排序编号',
    `menu_type`         varchar(1)                                   NOT NULL COMMENT '菜单类型 （菜单、目录、按钮）M 目录 C 菜单 F 按钮',
    `paths`             varchar(128)                                 NOT NULL COMMENT '菜单完整路径 /分割',
    `path`              varchar(255) COLLATE utf8mb4_bin DEFAULT ''  NOT NULL COMMENT '菜单路由路径',
    `component`         varchar(255) COLLATE utf8mb4_bin DEFAULT ''  NOT NULL COMMENT '组件路径',
    `permission`        varchar(255) COLLATE utf8mb4_bin DEFAULT ''  NOT NULL COMMENT '权限标识',
    `name`              varchar(255) COLLATE utf8mb4_bin             NOT NULL COMMENT '菜单名称',
    `title`             varchar(255) COLLATE utf8mb4_bin             NOT NULL COMMENT '菜单显示标题',
    `icon`              varchar(255) COLLATE utf8mb4_bin             NOT NULL COMMENT '菜单图标',
    `hide_in_menu`      tinyint(1) DEFAULT '0' NOT NULL COMMENT '是否隐藏菜单 0 不隐藏 1 隐藏',
    `ignore_keep_alive` tinyint(1) DEFAULT '0' NOT NULL COMMENT '取消页面缓存 0 不取消 1 取消',
    `link_flag`         tinyint(1) DEFAULT '0' NOT NULL COMMENT '是否外链 0 不是 1 是',
    `link`              varchar(255) COLLATE utf8mb4_bin DEFAULT ''  NOT NULL COMMENT '跳转路径 （外链）',
    `disabled`          tinyint(1) DEFAULT '0' NOT NULL COMMENT '是否停用 0 不停用 1 启用，对应 menuVisibleWithForbidden',
    `created_at`        timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
    `updated_at`        timestamp null on update CURRENT_TIMESTAMP COMMENT '最后更新时间',
    `deleted_at`        timestamp DEFAULT NULL COMMENT '删除时间',
    `create_by`         bigint                           DEFAULT '1' NOT NULL COMMENT '创建者',
    `update_by`         bigint                           DEFAULT '1' NOT NULL COMMENT '更新者',
    PRIMARY KEY (`menu_id`),
    UNIQUE KEY `menu_name` (`name`),
    UNIQUE KEY `menu_path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `sys_api`
(
    `id`         bigint                   NOT NULL AUTO_INCREMENT COMMENT '主键编码',
    `handle`     varchar(128) DEFAULT ''  NOT NULL COMMENT 'handle',
    `title`      varchar(128) DEFAULT ''  NOT NULL COMMENT '标题',
    `path`       varchar(128) DEFAULT ''  NOT NULL COMMENT '地址',
    `type`       varchar(16)  DEFAULT ''  NOT NULL COMMENT '接口类型',
    `action`     varchar(16)  DEFAULT ''  NOT NULL COMMENT '请求类型',
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
    `updated_at` timestamp null on update CURRENT_TIMESTAMP COMMENT '最后更新时间',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `create_by`  bigint       DEFAULT '1' NOT NULL COMMENT '创建者',
    `update_by`  bigint       DEFAULT '1' NOT NULL COMMENT '更新者',
    PRIMARY KEY (`id`),
    KEY          `idx_sys_api_deleted_at` (`deleted_at`),
    KEY          `idx_sys_api_create_by` (`create_by`),
    KEY          `idx_sys_api_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `sys_menu_api_rule`
(
    `id`               bigint NOT NULL AUTO_INCREMENT,
    `sys_menu_menu_id` bigint NOT NULL,
    `sys_api_id`       bigint NOT NULL COMMENT '主键编码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_sys_menu_api` (`sys_menu_menu_id`, `sys_api_id`),
    KEY                `fk_sys_menu_api_rule_sys_api` (`sys_api_id`),
    CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`),
    CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `sys_role`
(
    `role_id`        bigint                           NOT NULL AUTO_INCREMENT,
    `role_name`      varchar(128) DEFAULT ''          NOT NULL,
    `status`         varchar(4)   DEFAULT '' NOT NULL,
    `role_key`       varchar(128) DEFAULT '' NOT NULL COMMENT '角色码',
    `role_sort`      bigint       DEFAULT '1' NOT NULL,
    `remark`         varchar(255) DEFAULT '' NOT NULL,
    `admin`          tinyint(1) DEFAULT '0' NOT NULL COMMENT '是否是管理员 0 不是 1 是',
    `data_scope`     varchar(128) DEFAULT '0' NOT NULL COMMENT '数据范围',
    `default_router` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'dashboard' COMMENT '默认登录页面',
    `created_at`     timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
    `updated_at`     timestamp null on update CURRENT_TIMESTAMP COMMENT '最后更新时间',
    `deleted_at`     timestamp DEFAULT NULL COMMENT '删除时间',
    `create_by`      bigint       DEFAULT '1'         NOT NULL COMMENT '创建者',
    `update_by`      bigint       DEFAULT '1'         NOT NULL COMMENT '更新者',
    PRIMARY KEY (`role_id`),
    KEY              `idx_sys_role_deleted_at` (`deleted_at`),
    KEY              `idx_sys_role_create_by` (`create_by`),
    KEY              `idx_sys_role_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `ordering-platform`.sys_role
(role_id, role_name, status, role_key, role_sort, remark, admin, data_scope, default_router,create_by, update_by, deleted_at)
VALUES(1, '超级管理员', '2', 'admin', 1, '超级管理员', 1, '', '/dashboard', 1, 1, NULL);

CREATE TABLE `sys_role_menu`
(
    `id`      bigint NOT NULL AUTO_INCREMENT,
    `role_id` bigint NOT NULL,
    `menu_id` bigint NOT NULL,
    UNIQUE KEY `uniq_sys_role_menu` (`role_id`,`menu_id`),
    PRIMARY KEY (`id`),
    KEY       `fk_sys_role_menu_sys_menu` (`menu_id`),
    CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`),
    CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `sys_user`
(
    `user_id`     bigint                                       NOT NULL AUTO_INCREMENT COMMENT '编码',
    `status`      tinyint unsigned DEFAULT '1' COMMENT '状态 1 正常 2 禁用',
    `username`    varchar(255) COLLATE utf8mb4_bin             NOT NULL COMMENT '登录名',
    `password`    varchar(255) COLLATE utf8mb4_bin             NOT NULL COMMENT '密码',
    `nickname`    varchar(255) COLLATE utf8mb4_bin             NOT NULL COMMENT '昵称',
    `description` varchar(255) COLLATE utf8mb4_bin DEFAULT '' NOT NULL COMMENT '用户的描述信息',
    `mobile`      varchar(255) COLLATE utf8mb4_bin DEFAULT '' NOT NULL COMMENT '手机号',
    `email`       varchar(255) COLLATE utf8mb4_bin DEFAULT '' NOT NULL COMMENT '邮箱号',
    `avatar`      varchar(512) COLLATE utf8mb4_bin DEFAULT '' NOT NULL COMMENT '头像路径',
    `dept_id`     bigint unsigned DEFAULT '1' NOT NULL COMMENT '部门ID',
    `role_id`     bigint unsigned DEFAULT '1' NOT NULL COMMENT '角色ID',
    `created_at`  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
    `updated_at`  timestamp null on update CURRENT_TIMESTAMP COMMENT '最后更新时间',
    `deleted_at`  timestamp DEFAULT NULL COMMENT '删除时间',
    `create_by`   bigint                           DEFAULT '1' NOT NULL COMMENT '创建者',
    `update_by`   bigint                           DEFAULT '1' NOT NULL COMMENT '更新者',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `nickname` (`nickname`),
    UNIQUE KEY `user_username_email` (`username`,`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `ordering-platform`.sys_user
(user_id, status, username, password, nickname, description, mobile, email, avatar, dept_id, role_id, create_by, update_by, deleted_at)
VALUES(1, 1, 'admin', '$2a$10$JCKRxlTBJiUImgh2RSp/d.jH40JF0nrbBS7UQlknUS8wz.T1G8E22', 'admin', 'admin', '18237508888', 'zhaojuncodeing@163.com', 'qqqq', 1, 1, 1, 1, NULL);

CREATE TABLE `sys_casbin_rule`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v0`    varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v1`    varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v2`    varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v3`    varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v4`    varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v5`    varchar(100) COLLATE utf8mb4_bin DEFAULT '' NOT NULL,
    `v6`    varchar(25) COLLATE utf8mb4_bin  DEFAULT '' NOT NULL,
    `v7`    varchar(25) COLLATE utf8mb4_bin  DEFAULT '' NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sys_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `sys_dept`
(
    `dept_id`    bigint                   NOT NULL AUTO_INCREMENT COMMENT '主键编码',
    `parent_id`  bigint       DEFAULT '0' NOT NULL COMMENT '父级部门id',
    `dept_path`  varchar(255) DEFAULT '' NOT NULL COMMENT '部门路径 / 分割',
    `dept_name`  varchar(128) DEFAULT '' NOT NULL COMMENT '部门名字',
    `sort`       tinyint      DEFAULT '1' NOT NULL COMMENT '排序',
    `leader`     varchar(128) DEFAULT '' NOT NULL COMMENT '负责人',
    `phone`      varchar(11)  DEFAULT '' NOT NULL COMMENT '负责人手机号',
    `email`      varchar(64)  DEFAULT '' NOT NULL COMMENT '负责人邮箱',
    `status`     tinyint      DEFAULT '2' NOT NULL COMMENT '状态 1 停用 2 启用',
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  COMMENT '创建时间',
    `updated_at` timestamp null on update CURRENT_TIMESTAMP COMMENT '最后更新时间',
    `deleted_at` timestamp DEFAULT NULL COMMENT '删除时间',
    `create_by`  bigint       DEFAULT '1' NOT NULL COMMENT '创建者',
    `update_by`  bigint       DEFAULT '1' NOT NULL COMMENT '更新者',
    PRIMARY KEY (`dept_id`),
    KEY          `idx_sys_dept_create_by` (`create_by`),
    KEY          `idx_sys_dept_update_by` (`update_by`),
    KEY          `idx_sys_dept_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

