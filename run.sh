#!/bin/bash
echo Building
go build
if [ $? -eq 0 ]
then
  echo Raytracing
  ./go_raytrace > out.ppm
  echo Displaying
  display out.ppm
fi
