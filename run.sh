#!/bin/bash
echo Building
go build
if [ $? -eq 0 ]
then
  echo Raytracing
  time ./go_raytrace > out.ppm
  echo Displaying
  spd-say "Rendering complete"
  display out.ppm
fi
