#!/usr/bin/env bash
mkdir -p fonts
wget https://github.com/life888888/cjk-fonts-ttf/releases/download/v0.1.0/NotoSerifCJK-SC.zip
wget https://github.com/life888888/cjk-fonts-ttf/releases/download/v0.1.0/NotoSansMonoCJK-SC.zip
wget https://github.com/life888888/cjk-fonts-ttf/releases/download/v0.1.0/NotoSansCJK-SC.zip

unzip NotoSerifCJK-SC.zip -d fonts
unzip NotoSansMonoCJK-SC.zip -d fonts
unzip NotoSansCJK-SC.zip -d fonts

# 正确展示 emoji
curl -O fonts/NotoSansCJK-SC.zip https://github.com/googlefonts/noto-emoji/raw/main/fonts/NotoColorEmoji.ttf
