ASCII_ZERO = 48


def itoa(i, base=10):
    as_bytes = bytearray()
    while i > 0:
        rem = i % base
        as_bytes.insert(0, rem + ASCII_ZERO)
        i //= base
    return as_bytes.decode('utf-8')


def main():
    print(itoa(123))
    print(itoa(33553))


if __name__ == '__main__':
    main()