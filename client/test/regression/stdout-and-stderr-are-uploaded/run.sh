#! /bin/bash
echo 'stdout'
(>&2 echo 'stderr')

