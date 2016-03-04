import hashlib

secret_key = "yzbqklnj"
number = 0


def get_md5_hash(secret_key, number):
    return hashlib.md5(secret_key + str(number)).hexdigest()


i = 0

while not get_md5_hash(secret_key, number).startswith("00000"):
    number += 1

print(number)
