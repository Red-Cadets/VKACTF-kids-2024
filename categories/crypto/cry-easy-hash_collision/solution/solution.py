import hashlib
for i in range(48,122):
    for j in range(48, 122):
        for k in range(48, 122):
            for n in range(48, 122):
                p = chr(i)+chr(j)+chr(k)+chr(n)
                stored_password_hash = hashlib.sha256(bytes(p.encode())).hexdigest()[0:6]
                s = hashlib.sha256(bytes(p.encode())).hexdigest()
                if stored_password_hash == '6adeb6':
                    print(p,s)