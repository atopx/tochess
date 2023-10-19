from PIL import Image
import os

def convert_to_png(src_path, png_path):
    image = Image.open(src_path)
    image.save(png_path)

# 目录下的所有文件
directory = '/Users/atopx/opensource/tochess/chessbox/assets/image'

# 对于目录下的每一个.gif文件
for filename in os.listdir(directory):
    if filename.endswith(".jpg"):
        gif_path = os.path.join(directory, filename)
        # 这将生成 png文件的路径，例如：/path/to/your/directory/myfile.png
        png_path = os.path.splitext(gif_path)[0] + '.png'
        convert_to_png(gif_path, png_path)