import ctypes

lib = ctypes.CDLL("./test.so")

def call_c(L):
    arr = (ctypes.c_char_p * (len(L) ))()
    arr[:] = L
    lib.printList(arr,len(arr))


def convertToBytes(arr):
    bts = []
    for i in arr:
        bts.append(bytes(i, 'utf-8'))
    return bts

call_c(convertToBytes(["abc", "def"]))

