import time
import sys

from pwn import *

words_list = ["vkactf", "security", "coding", "chill", "flag", "game", "process", "music", "create", "cookie", "python", "java", "computer", "science", "network", "lucky", "exploit", "attack", "defense", "pentest", "proxy", "random", "server", "virus", "firewall", "data", "information", "browser", "phone", "byte"]

def count_words(input_string: str):
    word_count = 0
    
    for word in words_list:
        index = input_string.find(word)
        while index != -1:
            word_count += 1
            index = input_string.find(word, index + 1)
    
    return word_count

r = remote({sys.argv[1]}, 10022)
r.recvuntil(b'How many words in the string? Send number of words!\n')

for i in range(50):
    data = r.recvline()
    print(data)
    if b"vka{" in data:
        print(data)
        break
    print(data)
    quan = str(count_words(data.decode()))
    print(quan)
    r.sendline(quan.encode())
    r.recvuntil(b'Good!\n')
    time.sleep(0.5)

r.interactive()
