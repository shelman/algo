class InvalidNumberException(Exception):
    pass


def atoi(s):
    as_bytes = bytearray(s, 'utf-8')
    as_int = 0
    for b in as_bytes:
        if b < 48 or b > 57:
            raise InvalidNumberException('%s cannot be converted' % s)
        as_int *= 10
        as_int += (b - 48)
    return as_int


def main():
    print(atoi('123'))
    print(atoi('38819'))
    print(atoi('392s'))


if __name__ == '__main__':
    main()