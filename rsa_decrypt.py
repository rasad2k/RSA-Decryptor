import sys

def decrypt(c, d, n):
    res = hex(pow(int(c, 16), int(d, 16), int(n, 16)))
    m = bytearray.fromhex(res[2:]).decode()
    print(m)

if __name__ == "__main__":

    if len(sys.argv) != 4:
        print("Usage: ./script.py <c number> <d number> <n number>")
    else:
        c = sys.argv[1]
        d = sys.argv[2]
        n = sys.argv[3]
        decrypt(c,d,n)