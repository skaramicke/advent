#!/bin/bash

export YEAR=`date +%Y`
export MONTH=`date +%m`
export DAY=`date +%d`

mkdir "day$YEAR$MONTH$DAY"
envsubst < template > "./day$YEAR$MONTH$DAY/day$YEAR$MONTH$DAY.go"

import_placeholder="\t// new imports go here"
import_string="	\"github.com/skaramicke/advent/day$YEAR$MONTH$DAY\""
sed -i "" "s|$import_placeholder|$import_string\n$import_placeholder|" main.go

# Note that the option placeholder needs to stay above the newly inserted option.
option_placeholder="\t\t\t// new options go here"
option_string="			\"December $DAY, $YEAR\","
sed -i "" "s|$option_placeholder|$option_placeholder\n$option_string|" main.go

call_placeholder="\t\t// new calls go here"
call_string="	case \"December $DAY, $YEAR\":\n		day$YEAR$MONTH$DAY.Run()"
sed -i "" "s|$call_placeholder|$call_string\n$call_placeholder|" main.go
