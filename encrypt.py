from Crypto.Cipher import AES
from Crypto import Random

class AESCipher:
    def __init__(self, key):
        self.key = key

    def encrypt(self, raw):
        iv = Random.new().read(AES.block_size)
        cipher = AES.new(self.key, AES.MODE_CFB, iv)
        encrypted = cipher.encrypt(raw)
        return iv + encrypted

    def decrypt(self, raw):
        iv = raw[:AES.block_size]
        cipher = AES.new(self.key, AES.MODE_CFB, iv)
        return cipher.decrypt(raw[AES.block_size:])


cipher = AESCipher("example key 1234")
m = cipher.encrypt(b"my small message")

with open("msg", "wb") as f:
    f.write(m)
