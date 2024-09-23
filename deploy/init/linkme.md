## 表：casbin_rule

### 字段：
- `id`: 主键，自增。
- `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`: 用于存储权限规则的字段。

### 索引：
- `PRIMARY KEY (id)` 主键索引
- `UNIQUE KEY idx_casbin_rule (ptype, v0, v1, v2, v3, v4, v5)` 唯一索引

### 关系：
- 无外键约束。

## 表：checks

### 字段：
- `id`: 主键，自增。
- `post_id`: 关联到某个帖子的ID。
- `content`: 审核内容。
- `title`: 审核标题。
- `author_id`: 作者ID。
- `status`: 审核状态，默认值为'Pending'。
- `remark`: 备注。
- `created_at`, `updated_at`: 创建和更新时间。

### 索引：
- `PRIMARY KEY (id)` 主键索引
- `KEY idx_checks_author (author_id)` 普通索引
- `KEY idx_checks_updated_at (updated_at)` 普通索引

### 关系：
- 无外键约束。

## 表：comments

### 字段：
- `id`: 主键，自增。
- `user_id`: 用户ID。
- `biz`, `biz_id`: 业务类型及其ID。
- `content`: 评论内容。
- `post_id`: 关联到某个帖子的ID。
- `root_id`: 根评论ID。
- `pid`: 父评论ID。
- `created_at`, `updated_at`: 创建和更新时间。

### 索引：
- `PRIMARY KEY (id)` 主键索引
- `KEY idx_comments_p_id (pid)` 普通索引
- `KEY idx_user_id (user_id)` 普通索引
- `KEY idx_biz_type_id (biz, biz_id)` 普通索引
- `KEY idx_post_id (post_id)` 普通索引
- `KEY idx_comments_root_id (root_id)` 普通索引

### 关系：
- `FOREIGN KEY (pid) REFERENCES comments (id) ON DELETE CASCADE` 外键约束
- 外键 `fk_comments_parent_comment` 约束 `pid` 关联到 `comments` 表的 `id` 字段，级联删除。

## 表：interactives

### 字段：
- `id`: 主键，自增。
- `biz_id`, `biz_name`: 业务ID和名称。
- `read_count`, `like_count`, `collect_count`: 阅读、点赞和收藏数量。
- `updated_at`, `created_at`: 创建和更新时间。

### 关系：
- 无外键约束。

## 表：plates

### 字段：
- `id`: 主键，自增。
- `name`: 板块名称。
- `description`: 板块描述。
- `created_at`, `updated_at`, `deleted_at`: 创建、更新和删除时间。
- `deleted`: 是否删除标志。
- `uid`: 用户ID。

### 索引：
- `PRIMARY KEY (id)` 主键索引
- `UNIQUE KEY idx_plates_name (name)` 唯一索引
- `KEY idx_plates_uid (uid)` 普通索引

### 关系：
- 无外键约束。

## 表：posts

### 字段：
- `id`: 主键，自增。
- `created_at`, `updated_at`, `deleted_at`: 创建、更新和删除时间。
- `title`: 帖子标题。
- `content`: 帖子内容。
- `status`: 帖子状态。
- `author_id`: 作者ID。
- `slug`: 帖子别名。
- `category_id`, `plate_id`: 分类和板块ID。
- `tags`: 标签。
- `comment_count`: 评论数量。

### 索引：
- `PRIMARY KEY (id)` 主键索引
- `UNIQUE KEY idx_posts_slug (slug)` 唯一索引
- `KEY idx_posts_author_id (author_id)` 普通索引
- `KEY idx_posts_category_id (category_id)` 普通索引
- `KEY idx_posts_plate_id (plate_id)` 普通索引
- `KEY idx_posts_deleted_at (deleted_at)` 普通索引

### 关系：
- `FOREIGN KEY (plate_id) REFERENCES plates (id)` 外键约束
- 外键 `fk_plates_posts` 约束 `plate_id` 关联到 `plates` 表的 `id` 字段。

## 表：profiles

### 字段：
- `id`: 主键，自增。
- `user_id`: 用户ID。
- `nick_name`: 昵称。
- `avatar`, `about`: 头像和简介。
- `birthday`: 生日。

### 索引：
- `PRIMARY KEY (id)` 主键索引
- `KEY idx_profiles_user_id (user_id)` 普通索引

### 关系：
- `FOREIGN KEY (user_id) REFERENCES users (id)` 外键约束
- 外键 `fk_users_profile` 约束 `user_id` 关联到 `users` 表的 `id` 字段。

## 表：recent_activities

### 字段：
- `id`: 主键，自增。
- `user_id`: 用户ID。
- `description`: 活动描述。
- `time`: 活动时间。

### 关系：
- 无外键约束。

## 表：user_collection_bizs

### 字段：
- `id`: 主键，自增。
- `uid`: 用户ID。
- `biz_id`, `biz_name`: 业务ID和名称。
- `status`: 状态。
- `collection_id`: 收藏ID。
- `updated_at`, `created_at`: 创建和更新时间。
- `deleted`: 是否删除标志。

### 关系：
- 无外键约束。

## 表：user_like_bizs

### 字段：
- `id`: 主键，自增。
- `uid`: 用户ID。
- `biz_id`, `biz_name`: 业务ID和名称。
- `status`: 状态。
- `updated_at`, `created_at`: 创建和更新时间。
- `deleted`: 是否删除标志。

### 关系：
- 无外键约束。

## 表：users

### 字段：
- `id`: 主键，自增。
- `created_at`, `updated_at`, `deleted_at`: 创建、更新和删除时间。
- `nickname`: 昵称。
- `password_hash`: 密码哈希。
- `birthday`: 生日。
- `email`, `phone`: 邮箱和电话。
- `about`: 简介。
- `deleted`: 是否删除标志。

### 关系：
- 无外键约束。

## 表：v_code_sms_logs

### 字段：
- `id`: 主键，自增。
- `sms_id`: 短信ID。
- `sms_type`, `mobile`, `v_code`, `driver`, `status_code`: 短信相关信息。
- `status`: 状态。
- `created_at`, `updated_at`, `deleted_at`: 创建、更新和删除时间。

### 索引：
- `PRIMARY KEY (id)` 主键索引

### 关系：
- 无外键约束。
