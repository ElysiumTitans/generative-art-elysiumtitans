#! /bin/bash

path=$1
# files=$(ls "$path" | grep .png)
# echo "List of files in folder:"
# echo "${files}"

# for FILE in $files; do

#     echo "Renaming file: ${FILE}"
#     newname=${FILE: -5}
#     echo "$newname"
#     mv "$path/$FILE" "$path/$newname"

# done

## STEP 1
# replace whitespace in filenames
pushd $path

for f in *\ *; do mv "$f" "${f// Headwear 2/ HeadwearSecond}"; done
for f in *\ *; do mv "$f" "${f// copy/}"; done
for f in *\ *; do mv "$f" "${f// copy(1)/}"; done
for f in *\ *; do mv "$f" "${f// /_}"; done
popd

# STEP 2
# path=$1
files=$(ls "$path" | grep .png)
echo "List of files in folder:"
echo "${files}"

for FILE in $files; do

    echo "Renaming file: ${FILE}"
    layer=${FILE:0:1}
    if [ "${layer}" == "_" ]
    then
        layer=${FILE:1:2}
    fi
    mkdir -p "$path/$layer"
    newname=${FILE: -6}
    if [ "${newname:0:1}" == "_" ]
    then
        shortname=${FILE: -5}
        mv "$path/$FILE" "$path/$layer/$shortname"
    fi
    mv "$path/$FILE" "$path/$layer/$newname"

done

## STEP 3
path=$1
cd "$path"
for ((i=1; i<=9; i++))
do
    pushd "$i"

    files=$(ls | grep .png)
    echo "List of files in folder:"
    echo "${files}"

    for FILE in $files; do
        echo "Renaming file: ${FILE}"
        layer=${FILE:0:1}
        size=${#FILE}
        if [[ $size -gt 5 ]]
        then
            newname=${FILE: -5}
            mkdir -p "$layer"
            mv "$FILE" "$layer/$newname"
            continue
        fi
    done
    popd

done
