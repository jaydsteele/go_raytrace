#!/bin/bash
echo Building
go build
if [ $? -eq 0 ]
then
  echo Raytracing
  time ./go_raytrace > out.ppm
  echo Displaying
  say "Rendering complete"
  open -a Preview out.ppm
fi
