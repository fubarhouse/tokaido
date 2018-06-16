package drupaltmpl

// SettingsAppend - (Append) docroot/sites/default/settings.php
var SettingsAppend = []byte(`/*
* Generated by Tokaido
*/

if (file_exists($app_root . '/' . $site_path . '/settings.tok.php')) {
  include $app_root . '/' . $site_path . '/settings.tok.php';
}

/*
* END Generated by Tokaido
*/
?>
`)

// SettingsTok - docroot/sites/default/settings.tok.php
var SettingsTok = []byte(`<?php

/**
 * @file
 * Configuration file for Tokaido local dev environments. Add this to .gitignore
 *
 * Check out https://docs.tokaido.io for help managing your Tokaido environment
 *
 * Generated by Tokaido
 */

$databases['default']['default'] = [
  'host' => 'mysql',
  'database' => 'tokaido',
  'username' => 'tokaido',
  'password' => 'tokaido',
  'port' => 3306,
  'driver' => 'mysql',
  'namespace' => 'Drupal\\Core\\Database\\Driver\\mysql',
  'prefix' => '',
];

$settings['file_private_path'] = $app_root . '/../files-private/';

/*
* END Generated by Tokaido
*/

?>
`)
