#!/bin/sh
# Remove directory if exist
if [ ! -f /${1} ]; then
    rm -r ${1}
fi

# Create directory with file in it
mkdir ${1}
mkdir ${1}/about_me
mkdir -p ${1}/about_me/personal
touch ${1}/about_me/personal/facebook.txt
mkdir -p ${1}/about_me/professional 
touch ${1}/about_me/professional/linkedin.txt

mkdir ${1}/my_friends
touch ${1}/my_friends/list_of_my_friends.txt

mkdir ${1}/my_system_info
touch ${1}/my_system_info/about_this_laptop.txt
touch ${1}/my_system_info/internet_connection.txt

# File facebook.txt dan linkedin.txt berisikan url dengan username masing masing di argumen kedua dan ketiga
echo "https://facebook.com/${2}" > ${1}/about_me/personal/facebook.txt
echo "https://linkedin.com/in/${3}" > ${1}/about_me/professional/linkedin.txt

# File list_of_my_friends.txt berisikan daftar nama teman teman yang diambil menggunakan command curl dari file di bawah.
link="https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
curl -s $link >> ${1}/my_friends/list_of_my_friends.txt

# File about_this_laptop.txt berisikan nama user dan uname -a dengan format di bawah
echo "Masukkan username"
    read name
echo "My username : $name " >> ${1}/my_system_info/about_this_laptop.txt
echo -n "With host:" $(uname -a)  >> ${1}/my_system_info/about_this_laptop.txt

# File internet_connection.txt berisikan hasil keluaran ping ke google.com sebanyak 3 kali.
echo "Connection to google": > ${1}/my_system_info/internet_connection.txt
ping -c 3 google.com >> ${1}/my_system_info/internet_connection.txt