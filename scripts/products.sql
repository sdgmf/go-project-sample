/*
 Navicat Premium Data Transfer

 Source Server         : db_local
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : 127.0.0.1:3306
 Source Schema         : products

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : 65001

 Date: 02/08/2019 16:32:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for details
-- ----------------------------
DROP TABLE IF EXISTS `details`;
CREATE TABLE `details` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `price` double DEFAULT NULL,
                         `created_time` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of details
-- ----------------------------
BEGIN;
INSERT INTO `details` VALUES (1, 'apple', 1, '2019-07-31 19:43:10');
INSERT INTO `details` VALUES (2, 'pear', 1, '2019-07-31 19:43:45');
INSERT INTO `details` VALUES (3, 'banana', 0.5, '2019-07-31 19:44:08');
COMMIT;

-- ----------------------------
-- Table structure for ratings
-- ----------------------------
DROP TABLE IF EXISTS `ratings`;
CREATE TABLE `ratings` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `product_id` bigint(20) unsigned DEFAULT NULL,
                         `score` int(10) unsigned DEFAULT NULL,
                         `updated_time` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of ratings
-- ----------------------------
BEGIN;
INSERT INTO `ratings` VALUES (1, 1, 5, '2019-07-31 19:48:04');
INSERT INTO `ratings` VALUES (2, 2, 4, '2019-07-31 19:48:21');
INSERT INTO `ratings` VALUES (3, 3, 5, '2019-07-31 19:48:34');
COMMIT;

-- ----------------------------
-- Table structure for reviews
-- ----------------------------
DROP TABLE IF EXISTS `reviews`;
CREATE TABLE `reviews` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `product_id` bigint(20) unsigned DEFAULT NULL,
                         `message` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                         `created_time` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of reviews
-- ----------------------------
BEGIN;
INSERT INTO `reviews` VALUES (1, 1, 'good', '2019-07-31 19:48:47');
INSERT INTO `reviews` VALUES (2, 2, 'bad', '2019-07-31 19:49:12');
INSERT INTO `reviews` VALUES (3, 3, 'good', '2019-07-31 20:15:56');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
