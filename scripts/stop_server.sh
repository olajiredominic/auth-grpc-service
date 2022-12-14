#!/bin/bash
kill $(lsof -t -i:5000) > /dev/null 2> /dev/null < /dev/null &