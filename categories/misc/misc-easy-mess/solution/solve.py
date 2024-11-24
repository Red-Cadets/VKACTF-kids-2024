from pwn import *
import time

def connect_to_service():
    return remote('127.0.0.1', 11111)

def collect_data():
    received_data = {}    
    conn = connect_to_service()    
    while True:
        line = conn.recvline().decode().strip()        
        if not line:
            continue

        index_str, base64_part = line.split(" : ")
        print(f"Получена часть {index_str}")
        index = int(index_str)
        
        if index in received_data:
            print("Повторение строки")
            break

        received_data[index] = base64_part        
        # time.sleep(0.5)
    
    conn.close()

    return received_data

def reconstruct_data(received_data):
    sorted_parts = [received_data[i] for i in sorted(received_data)]    
    full_data = ''.join(sorted_parts)
    
    return full_data

def write_to_file(data):
    with open('solve.txt', 'w') as f:
        f.write(data)

def main():
    received_data = collect_data()
    print(f"Получено {len(received_data)} частей данных")    
    reconstructed_data = reconstruct_data(received_data)    
    write_to_file(reconstructed_data)    
    print("записано в solve.txt")

if __name__ == '__main__':
    main()
