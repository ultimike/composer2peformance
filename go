#!/bin/bash

# Run this with "bash go" to run speed tests with local Composer. Be sure to update/rollback Composer first as needed.

# composer create-project 
time composer create-project drupal/recommended-project speed1 --profile 

time composer create-project drupal/recommended-project speed2 --profile

time composer create-project drupal/recommended-project speed3 --profile

# composer require
cd speed1
time composer require drupal/pathauto:1.7.0 --profile

cd ../speed2
time composer require drupal/pathauto:1.7.0 --profile

cd ../speed3
time composer require drupal/pathauto:1.7.0 --profile

# composer update
cd speed1
composer require drupal/pathauto:^1.7 --no-update
time composer update drupal/pathauto --profile

cd ../speed2
composer require drupal/pathauto:^1.7 --no-update
time composer update drupal/pathauto --profile

cd ../speed3
composer require drupal/pathauto:^1.7 --no-update
time composer update drupal/pathauto --profile
