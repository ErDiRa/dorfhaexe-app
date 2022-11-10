#!/bin/bash

# simple script to resize images by using imageMagick

#source: https://www.smashingmagazine.com/2015/06/efficient-image-resizing-with-imagemagick/
smartresize() {
   mogrify -path $3 -filter Triangle -define filter:support=2 -thumbnail $2 -unsharp 0.25x0.08+8.3+0.045 -dither None -posterize 136 -quality 82 -define jpeg:fancy-upsampling=off -define png:compression-filter=5 -define png:compression-level=9 -define png:compression-strategy=1 -define png:exclude-chunk=all -interlace none -colorspace sRGB $1
}

simpleResise(){
	convert $1 -resize $2 $3
}

if [ -z "$1" ]
  then
    echo "Please provide an input path"
		exit 1
fi

if [ -z "$2" ]
  then
    echo "Please provide an output path"
		exit1
fi

INPUT_PATH=$1
OUTPUT_PATH=$2


## define image sizes
sizes=(780 1024 1440 2400)



for image in $(find . -type f -exec file --mime-type {} \+ | awk -F: '{if ($2 ~/image\//) print $1}');
do
  ## get image width
  width=`convert $image -ping -format "%w" info:`

	echo $image $width

  ## get image path and name
  dir=$(dirname "$image")
  filename=$(basename "$image")
	extension="${filename##*.}"
	filename="${filename%.*}"

  ## run through image sizes
  for size in ${sizes[@]}; do

    ## set new image name
    newname="$OUTPUT_PATH"/"$filename"-"$size""w"."$extension"

		echo $image -> $newname

    ## resize image with define widths
    simpleResise $image $size $newname

		#does not work
		# smartresize $OUTPUT_PATH $size $image

  done
done