-- app meta table for app manager

CREATE TABLE `app_meta` (
  -- app format chart.yaml
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` varchar(64) NOT NULL COMMENT 'uuid v4',
  `name` varchar(128) NOT NULL COMMENT 'app meta name',
  `version` varchar(32) NOT NULL COMMENT 'app meta version',
  `group` varchar(32) NOT NULL COMMENT 'app group',
  `description` varchar(128) DEFAULT '' COMMENT 'app description',
  `annotations` json COMMENT 'Contains app-suite, category etc.',
  `labels` json COMMENT 'extra infos for extends',
  `home` varchar(64) DEFAULT '' COMMENT 'app homepage',
  `sources` varchar(64) DEFAULT '' COMMENT 'source url',
  `dependencies` json COMMENT 'other apps that this app depends on',
  `maintainers` json COMMENT 'app developer',
  `icon` varchar(64) DEFAULT '' COMMENT 'app icons',
  `doc` varchar(64) DEFAULT '' COMMENT 'app docs',

  -- app format sage templates
  `sage` json COMMENT 'sage info: components, services, storage etc.',

  -- app format install.yaml
  `install` json COMMENT 'app install policy: policy, licenseCheck, updatePolicy, migration etc.',

  -- app format templates
  `templates` json COMMENT 'reserved keyword, connect with the community in the future',

  -- app format README.md
  `read_me` text COMMENT 'app README.md',

  -- attributes other than app format
  `status` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT 'app status, 0 frozen, 1 available',
  `installed` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT 'whether the application is installed, 0 uninstalled, 1 installed',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`),
  UNIQUE KEY `app_name_version` (`name`, `version`),
  INDEX `app_id` (`app_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT 'for sage app meta';
