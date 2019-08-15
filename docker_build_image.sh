#! /bin/bash

echo "###### DEPLOY MODULE ######"
currentDate=`date +"%b-%d.%H.%M"`
echo "Build colegios-students docker image" $currentDate

version="0.1.0"
imagename="colegios-students"
filename=$imagename-v$version-$currentDate".tar"
echo "File name: "$filename


echo "Generate local docker image:"
docker build -t $imagename:$version . -f ./Dockerfile --rm=true || exit 1

docker rmi $(docker images -f dangling=true -q)

cd ..

echo "Save docker image in file *.tar:"
docker save "$imagename":"$version" > devops/images/v$version/$filename || exit 1
echo $filename

echo "Upload docker image to server:"
scp devops/images/v$version/$filename root@157.230.214.121:/colegios_api/devops/images || exit 1
echo "###### SUCCESS LOAD ######"

loadImage="docker load < colegios_api/devops/images/"$filename

echo "Enter server: "
ssh root@157.230.214.121<< EOF
    cd ..
    $loadImage
    cd colegios_api/devops
    docker-compose down
    docker-compose up 
EOF
echo "###### SUCCESS DEPLOY MODULE ######"

cd colegios-students

echo "Want you commit changes? [y/n]"
read commit

if [ $commit == 'y' ]
then 
    echo "Commit message: "
    read message
    git commit -am "$message"
    git push
    echo "###### SUCCESS SCRIPT ######"
else 
    echo "###### SUCCESS SCRIPT[NOT COMMIT] ######"
    exit 1
fi