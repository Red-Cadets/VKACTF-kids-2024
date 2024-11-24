#!/bin/python
import requests
import time
import base64
import sys

URL = sys.argv[1] if len(sys.argv) > 1 else 'http://127.0.0.1:5000'

def fetch_generated_string(url):
    while True:
        try:
            response = requests.get(url)
            response.raise_for_status()  
            generated_string = response.text  


            try:
                decoded_string = base64.b64decode(generated_string).decode('utf-8')
                if 'vka' in decoded_string:
                    print(f"Найден флаг: {decoded_string}")


            except (base64.binascii.Error, UnicodeDecodeError) as e:
                print(f'Полученная строка: {base64.b64decode(generated_string)}')

        except requests.RequestException as e:
            print(f"Ошибка при запросе к серверу: {e}")
        
        time.sleep(0.5) 


if __name__ == '__main__':
    fetch_generated_string(f'{URL}/dozhd')