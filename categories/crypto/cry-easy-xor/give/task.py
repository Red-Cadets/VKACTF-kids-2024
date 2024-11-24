from Crypto.Util.number import bytes_to_long

flag = b'???????????'
parts = flag.split(b'_')
parts = [part.decode('utf-8') for part in parts]
print("Parts:", parts)

def encode(parts, x):
    ct = []
    for part, key in zip(parts, x):
        part_bytes = bytes(part, 'utf-8')
        part_long = bytes_to_long(part_bytes)
        key_int = int(key * 1000)  
        ct.append(part_long ^ key_int)
    return ct


x = [?, ?, ?, ?]  

ciphertext = encode(parts, x)
print("Ciphertext:", ciphertext)

'''
sqrt(2) * sin(x) * (sin(x) ** 2 - 1) + cos(x) ** 2 = 0
Ciphertext: [-27991888056766972,-1836402242,-1265523987,-1833007901]
pi = 3,14
Interval = [(-5*pi)/2 ; -pi]