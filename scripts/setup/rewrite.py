import json
import sys

# コマンドライン引数からファイルパスを取得
json_input_path = sys.argv[1]
code_bytes_file_path = sys.argv[2]
json_output_path = sys.argv[3]

# JSONファイルを読み込む
with open(json_input_path, 'r') as f:
    data = json.load(f)


with open(code_bytes_file_path, 'r') as f:
    byte_str = f.read()

# print(byte_str)

data['app_state']['wasm']['codes'][0]['code_bytes'] = byte_str


with open(json_output_path, 'w') as f:
    json.dump(data, f)
