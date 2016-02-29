import hashlib

secret_key = "yzbqklnj"
number = 0


def get_md5_hash(secret_key, number):
    return hashlib.md5(secret_key + str(number)).hexdigest()


while True:
    md5_hash = get_md5_hash(secret_key, number)
    if md5_hash.startswith("00000"):
        print("number:" + str(number))
        break

    number += 1