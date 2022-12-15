#!/bin/bash
kill -9 $(lsof -t -i:5000) > /dev/null 2> /dev/null < /dev/null &
