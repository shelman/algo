def permutations(s):
    if len(s) == 1:
        return [s]
    res = []
    for idx, letter in enumerate(s):
        res.extend([letter + p for p in permutations(s[:idx] + s[idx+1:])])
    return set(res)


def main():
    print(permutations('abdia'))
    print(permutations('cat'))


if __name__ == '__main__':
    main()