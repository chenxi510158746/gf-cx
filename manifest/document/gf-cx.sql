# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.39)
# Database: gf-cx
# Generation Time: 2023-12-07 08:37:54 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table c_goods
# ------------------------------------------------------------

DROP TABLE IF EXISTS `c_goods`;

CREATE TABLE `c_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `title` varchar(191) NOT NULL DEFAULT '' COMMENT '商品名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '商品描述',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '商品价格（单位：分）',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '商品状态（0-下架；1-上架）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

LOCK TABLES `c_goods` WRITE;
/*!40000 ALTER TABLE `c_goods` DISABLE KEYS */;

INSERT INTO `c_goods` (`id`, `title`, `desc`, `price`, `status`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'商品测试1','商品描述测试1',5000,1,'2023-12-05 19:52:04','2023-12-05 19:52:04',NULL),
	(2,'商品测试2','商品描述测试2',8000,1,'2023-12-05 19:52:46','2023-12-05 19:52:46',NULL),
	(3,'商品测试3','商品描述测试3',10000,1,'2023-12-05 19:53:07','2023-12-05 19:53:07',NULL),
	(4,'商品测试4','商品描述测试4',12000,1,'2023-12-05 19:53:28','2023-12-05 19:53:28',NULL),
	(5,'商品测试5','商品描述测试5',14000,1,'2023-12-05 19:53:46','2023-12-05 19:53:46',NULL),
	(6,'商品测试6','商品描述测试6',16000,1,'2023-12-05 19:54:02','2023-12-05 19:54:02',NULL),
	(7,'商品测试7','商品描述测试7',18000,1,'2023-12-05 19:54:25','2023-12-05 19:54:25',NULL),
	(8,'商品测试8','商品描述测试8',20000,1,'2023-12-05 19:54:46','2023-12-05 19:54:46',NULL),
	(9,'商品测试10','商品描述测试10',22000,0,'2023-12-05 19:56:01','2023-12-05 19:58:32','2023-12-05 19:59:37'),
	(10,'收拾收拾收拾收拾','商品描述测试9',20000,1,'2023-12-06 20:44:43','2023-12-06 20:44:43',NULL),
	(11,'商品哈哈2','商品描述哈哈2',3250,1,'2023-12-07 14:34:04','2023-12-07 14:36:51','2023-12-07 14:39:48');

/*!40000 ALTER TABLE `c_goods` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table c_orders
# ------------------------------------------------------------

DROP TABLE IF EXISTS `c_orders`;

CREATE TABLE `c_orders` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `order_sn` varchar(191) NOT NULL DEFAULT '' COMMENT '订单编号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `goods_price` int(11) NOT NULL DEFAULT '0' COMMENT '商品价格（单位：分）',
  `pay_amount` int(11) NOT NULL DEFAULT '0' COMMENT '支付金额（单位：分）',
  `pay_status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '支付状态（1-待支付；2-已支付；3-已退款）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

LOCK TABLES `c_orders` WRITE;
/*!40000 ALTER TABLE `c_orders` DISABLE KEYS */;

INSERT INTO `c_orders` (`id`, `order_sn`, `user_id`, `goods_id`, `goods_price`, `pay_amount`, `pay_status`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'O1TTCIFQAJE0CXH3PJBKZNPC100FJ7Z3Z',1,1,5000,50,2,'2023-12-06 16:17:38','2023-12-06 16:17:38',NULL),
	(2,'O1TTCIFQAS90CXH3REWU4TUO100H49V0Y',2,2,8000,8000,2,'2023-12-06 16:20:06','2023-12-06 16:20:06',NULL),
	(3,'O1TTCIFQAS90CXH3S9XDFFPC200HZYMGR',3,3,10000,2000,2,'2023-12-06 16:21:13','2023-12-06 16:21:13',NULL),
	(4,'O1TTCIFQAS90CXH3SMQGDBBS300DUT8H7',3,3,10000,2000,2,'2023-12-06 16:21:41','2023-12-06 16:21:41',NULL),
	(5,'O1TTCIFQAS90CXH3SNTOZBXC400SPV187',3,3,10000,2000,2,'2023-12-06 16:21:43','2023-12-06 16:21:43',NULL),
	(6,'O1TTCIFQAS90CXH3SOKXZ8M05001DIX6J',3,3,10000,2000,2,'2023-12-06 16:21:45','2023-12-06 16:21:45',NULL),
	(7,'O1TTCIFQAS90CXH3SP3QJGOO600B1010M',3,3,10000,2000,2,'2023-12-06 16:21:46','2023-12-06 16:21:46',NULL),
	(8,'O1TTCIFQAS90CXH3SPJ7HGU8700MJT4VQ',3,3,10000,2000,2,'2023-12-06 16:21:47','2023-12-06 16:21:47',NULL),
	(9,'O1TTCIFQAS90CXH3SQ00JG2O800QIUKXU',3,3,10000,2000,2,'2023-12-06 16:21:48','2023-12-06 16:21:48',NULL),
	(10,'O1TTCIFQAS90CXH3SQDRXGRK900BWOUWK',3,3,10000,2000,2,'2023-12-06 16:21:49','2023-12-06 16:21:49',NULL),
	(11,'O1TTCIFQAS90CXH3SQS7PSA8A000S68XR',3,3,10000,2000,2,'2023-12-06 16:21:50','2023-12-06 16:21:50',NULL),
	(12,'O1TTCIFQAS90CXH3SR8BI4LSB005WLKGB',3,3,10000,2000,2,'2023-12-06 16:21:51','2023-12-06 16:21:51',NULL),
	(13,'O1TTCIFQ1DM3CXHXL6PGYB5C2002NJC3W',5,5,14000,14000,2,'2023-12-07 15:42:31','2023-12-07 15:42:31',NULL),
	(14,'O1TTCIFQ1DM3CXHXLR6LZKU0300YJYL5P',5,8,20000,20000,2,'2023-12-07 15:43:15','2023-12-07 15:43:15',NULL);

/*!40000 ALTER TABLE `c_orders` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table c_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `c_users`;

CREATE TABLE `c_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户姓名',
  `mobile` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `password` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `age` int(5) NOT NULL DEFAULT '0' COMMENT '年龄',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

LOCK TABLES `c_users` WRITE;
/*!40000 ALTER TABLE `c_users` DISABLE KEYS */;

INSERT INTO `c_users` (`id`, `user_name`, `mobile`, `password`, `age`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'ccx','15618959766','7e6614509f3948ed921e9e82589ee73b',18,'2023-12-04 19:54:44','2023-12-04 19:54:44',NULL),
	(2,'ccxx','15618959768','7e6614509f3948ed921e9e82589ee73b',18,'2023-12-04 19:56:08','2023-12-04 19:56:08',NULL),
	(3,'cx','15618959769','7e6614509f3948ed921e9e82589ee73b',18,'2023-12-05 09:40:55','2023-12-05 09:40:55',NULL),
	(4,'cn','15618959770','7e6614509f3948ed921e9e82589ee73b',18,'2023-12-05 14:54:43','2023-12-05 14:54:43',NULL),
	(5,'sssccc','15618959787','7e6614509f3948ed921e9e82589ee73b',21,'2023-12-07 14:17:02','2023-12-07 14:17:02',NULL);

/*!40000 ALTER TABLE `c_users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
